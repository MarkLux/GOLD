# The custom basic envirnoment for golang
FROM centos
LABEL maintainer="MarkLux"
# install dependencies
RUN yun install -y \
    gcc g++ \
    git \
    wget \
    && yum clean all
# create work dir
RUN mkdir -p /home/go
WORKDIR /home/go
# install go from remote
RUN wget https://studygolang.com/dl/golang/go1.12.linux-amd64.tar.gz && \
    tar -xvzf go1.12.linux-amd64.tar.gz -C /usr/local && \
    rm -rf go1.12.linux-amd64.tar.gz
# set up common variables
ENV GOROOT /usr/local/go
ENV GOPATH /home/go
ENV PATH /usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/local/go/bin
