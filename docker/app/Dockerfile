### Application Image based on Golang image + Debian image
# version 1.0

###
### FIRST STAGE: Generate binary file from the repository using Golang official image
###
FROM golang:1.13.4 AS compiler

ENV APP roll-the-dices
ENV REPO github.com/ncaak/${APP}

WORKDIR /go/src/

# Get the code from repository (master branch) then build the binary saved in /go/bin
RUN go get -d -v ${REPO}
RUN CGO_ENABLE=0 GOOS=linux go install -i ./...


###
### SECOND STAGE: Get an slim debian machine to run the application
###
FROM debian:stable-slim

# Environment variables with data for the app
ENV API_URL https://localhost/
ENV DATABASE_ACCESS tcp(database)/database
ENV DATABASE_CREDENTIALS user:pass

WORKDIR /root/

# Retrieve binary file from previous stage
COPY --from=compiler /go/bin/${APP} .
COPY daemon.sh .

# Install needed applications to run Golang app
RUN apt-get update \
    && apt-get install -y ca-certificates

CMD ["./daemon.sh"]
