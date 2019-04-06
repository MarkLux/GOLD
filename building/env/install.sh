#!/bin/bash

echo "creating gold namespace..."
kubectl create -f k8s/gold/namespace.yaml

echo "creating redis cluster..."
echo "[1] creating config-map..."

kubectl create configmap redis-conf --from-file=k8s/redis/redis.conf --namespace=gold

echo "[2] creating headless service..."
kubectl create -f k8s/redis/headless_service.yaml

echo "[3] creating stateful set..." 
kubectl create -f k8s/redis/stateful_set.yaml

echo "[4] creating gold redis service..."
kubectl create -f k8s/redis/gold_redis_service.yaml

echo "redis cluster deployed, you need to init it with redis-trib later."

echo "creating mongo replica sets:"
echo "[1] creating secret..."
TMPFILE=$(mktemp)
/usr/bin/openssl rand -base64 741 > $TMPFILE
kubectl create secret generic shared-bootstrap-data --from-file=internal-auth-mongodb-keyfile=$TMPFILE --namespace=gold
rm $TMPFILE

echo "[2] creating persist volumes..."
kubectl create -f k8s/mongo/volume.yaml

echo "[3] creating headless service..."
kubectl create -f k8s/mongo/headless_service.yaml

echo "[4] creating stateful set..."
kubectl create -f k8s/mongo/stateful_set.yaml

echo "mongo replica sets deployed, you need to run k8s/mongo/init.sh to set up."

echo "creating gold http trigger..."
kubectl create -f k8s/gold/http-trigger.yaml
