FROM registry.cn-qingdao.aliyuncs.com/gold-faas/gold-faas-wrapper:0.2.beta

# create work dir
WORKDIR /home

ARG REPO_URL
ARG REPO_NAME
ARG BRANCH
ARG COMMIT_ID

RUN git clone $REPO_URL \
    && cd $REPO_NAME \
    && git checkout $BRANCH \
    && git reset --hard $COMMIT_ID \
    && cp gold_biz.go /home/GOLD/serving/wrapper/gold/gold_biz.go

# build the project
WORKDIR /home/GOLD/serving/wrapper
RUN go build
CMD ./wrapper

# expose the port
EXPOSE 8080