FROM registry.cn-qingdao.aliyuncs.com/gold-faas/gold-base-env-golang:1.0

# create work dir
WORKDIR /home

# pull the code
RUN git clone https://github.com/MarkLux/GOLD

# start the server
WORKDIR /home/GOLD/serving/restful
CMD go run main.go

# expose the port
EXPOSE 8080