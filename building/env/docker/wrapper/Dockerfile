FROM registry.cn-qingdao.aliyuncs.com/gold-faas/gold-base-env-golang:1.0

# create work dir
WORKDIR /home

# pull the code
RUN git clone https://github.com/MarkLux/GOLD

# build the project
WORKDIR /home/GOLD/serving/wrapper
RUN go mod tidy

# expose the port
EXPOSE 8080