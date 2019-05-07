FROM golang:1.12-alpine

# create work dir
WORKDIR /home

# install dependency

# pull the code
RUN apk add git && git clone https://github.com/MarkLux/GOLD

# build the project
WORKDIR /home/GOLD/eventing/http
RUN go build
CMD ./http

# expose the port
EXPOSE 8080