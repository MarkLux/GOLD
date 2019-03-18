FROM registry.cn-qingdao.aliyuncs.com/gold-faas/gold-base-env-golang:1.0

# create work dir
WORKDIR /home

# install dependency

# pull the code
RUN git clone https://github.com/MarkLux/GOLD

# copy biz code
COPY gold_biz.go /home/GOLD/serving/wrapper/gold/gold_biz.go

# build the project
WORKDIR /home/GOLD/serving/wrapper
RUN go build
CMD ./wrapper

# expose the port
EXPOSE 8080