FROM registry.cn-qingdao.aliyuncs.com/gold-faas/gold-base-env-golang:1.0

# create work dir
WORKDIR /home

# install dependency

# pull the code
RUN git clone https://github.com/MarkLux/GOLD

# build the project
WORKDIR /home/GOLD/serving/wrapper
RUN go build
CMD ./wrapper

# expose the port
EXPOSE 8080