FROM golang:1.12-alpine

# create work dir
WORKDIR /home

# pull the code
RUN apk add git && git clone https://github.com/MarkLux/GOLD

# build the project
WORKDIR /home/GOLD/serving/wrapper
RUN go mod tidy

# expose the port
EXPOSE 8080