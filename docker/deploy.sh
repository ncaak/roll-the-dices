#!/bin/bash

# version 1.0

### Auxiliar functions

# Generating database and app images
buildImages()
{
    echo -e "\n\e[93m Building images...\e[0m"
    docker build -t $image_prefix-db:$version ./db
    docker build -t $image_prefix-app:$version ./app
    echo -e "\e[92m Images done\e[0m\n"
}

# Creating bridge network
createNetwork()
{
    echo -e "\n\e[93m Creating brigde network named $image_prefix-nw...\e[0m"
	docker network create --driver bridge $image_prefix-nw
    echo -e "\e[92m Network created\e[0m\n"
}

#Runing containers
initContainers()
{
    DATABASE_HOSTNAME=database
    build_db_name=$(date +%s.%N | sha256sum | base64 | head -c 16 ; echo)
    build_db_user=$(date +%s.%N | sha256sum | base64 | head -c 16 ; echo)
    build_db_pass=$(date +%s.%N | sha256sum | base64 | head -c 32 ; echo)
    build_db_master=$(date +%s.%N | sha256sum | base64 | head -c 32 ; echo)

    echo -e "\n\e[93m Running containers...\e[0m"
    # Initialize database container and get the Id to reach it from App container
    db_container_id=$(docker run -td $1 --restart=always --log-opt max-size=10m \
        -e MYSQL_ROOT_PASSWORD=$build_db_master \
        -e MYSQL_DATABASE=$build_db_name \
        -e MYSQL_USER=$build_db_user \
        -e MYSQL_PASSWORD=$build_db_pass \
        $image_prefix-db:$version)

    db_container_ip=$(docker inspect \
        -f "{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}" \
        $db_container_id)

    # Initialize container linked to the database one through the hostname
    docker run -td $1 --restart=always --log-opt max-size=10m \
        --add-host=$DATABASE_HOSTNAME:$db_container_ip \
        -e "DATABASE_ACCESS=tcp($DATABASE_HOSTNAME)/$build_db_name" \
        -e "DATABASE_CREDENTIALS=$build_db_user:$build_db_pass" \
        $image_prefix-app:$version

    echo -e "\e[92m Container created\e[0m\n"
}

### Parameter management
# Default values for tagging images and running containers
image_prefix=dev
version=latest
# Deployment steps control
step_images=0
step_network=0
step_containers=0
# Initiation variables
build_prefix=dev
# Parameters control
while [ "$1" != "" ]; do
    case $1 in
        -n | --name )       image_prefix=$2
                            shift
                            ;;
        -v | --version )    version=$2
                            shift
                            ;;
        -i | --images )     step_images=1
                            ;;
        -n | --network )    step_network=1
                            ;;
        -c | --containers ) step_containers=1
                            ;;
        * )                 exit
    esac
    shift
done
### Main processing
# Isolated step for creating images
if [ "$step_images" == "1" ]; then
    echo -e "\e[93m :: Generating images ::\e[0m"
    buildImages
    echo -e "\e[93m :: Images generated successfully ::\e[0m"
    exit
fi
# Isolated step for creating the network
if [ "$step_network" == "1" ]; then
    echo -e "\e[93m :: Generating network ::\e[0m"
    createNetwork
    echo -e "\e[93m :: Network generated successfully ::\e[0m"
    exit
fi
# Isolated step for initiating the containers
if [ "$step_containers" == "1" ]; then
    echo -e "\e[93m :: Initiating containers ::\e[0m"
    initContainers
    echo -e "\e[93m :: Containers running successfully ::\e[0m"
    exit
fi
# If no step has been set, follow the complete process
echo -e "\e[93m :: Beginning complete installation ::\e[0m"
buildImages
createNetwork
initContainers --network=$image_prefix-nw
echo -e "\e[93m :: Installation ended successfully ::\e[0m"
exit

