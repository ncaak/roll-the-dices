### Database Image based on mariaDB image
# version 1.0

# Load mariadb image service
FROM mariadb:latest

# Environment variables for the image following offical mariaDB documentation
# @link:(https://hub.docker.com/_/mariadb)
# @link:(https://mariadb.com/kb/en/library/installing-and-using-mariadb-via-docker/)
ENV MYSQL_ROOT_PASSWORD testpassword
ENV MYSQL_DATABASE database
ENV MYSQL_USER default_user
ENV MYSQL_PASSWORD default_pass

# SQL and SH files on this folder will be triggered on initialization on alphabetic order
WORKDIR /docker-entrypoint-initdb.d/

# Default database file
COPY db-dump.sql ./

# No entrypoint script, it causes the service to not initiate
# as it interfere with the default entry script

