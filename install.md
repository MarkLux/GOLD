# GOLD Install Guide

## Preparation

To install & set up a GOLD FaaS Platform, 3 components is required:

- a k8s cluster (or minikube)

- a private docker registry

- a server to provide rest api

- [alternative] a web console (see [GOLD-console](https://github.com/MarkLux/GOLD-console))

## Set up k8s

1. clone the whole project into your k8s cluster master node and move into the k8s config directory:

    ```
    git clone https://github.com/MarkLux/GOLD && cd GOLD/building/env/k8s
    ``` 

2. create namespace

    ```
    $ kubectl create -f gold/namespace.yaml 
    namespace/gold created
    ```

3. deploy redis cluster (default: 6 pods, 3 master and 3 slaver)

    create config map:

    ```
    kubectl create configmap redis-conf --from-file=redis/redis.conf --namespace=gold 
    configmap/redis-conf created

    ```
    
    create headless service:
    
    ```
    kubectl create -f redis/headless_service.yaml 
    service/redis-service created
    ```
    
    create stateful set:
    
    ```
    kubectl create -f redis/stateful_set.yaml 
    statefulset.apps/redis-app created
    ```
    
    create entry service:
    
    ```
    kubectl create -f redis/gold_redis_service.yaml 
    service/gold-redis created
    ```
    
4. init redis cluster

    for redis version beyond 5.0, the redis-cli can be used to init the cluster.
    
    you can start a new pod in the cluster and install the redis-cli to init the cluster ([reference](http://marklux.cn/blog/108)).
    
5. deploy mongo replica sets (default: 4 pods)

    create a secret:
    
    ```
    TMPFILE=$(mktemp)
    /usr/bin/openssl rand -base64 741 > $TMPFILE
    kubectl create secret generic shared-bootstrap-data --from-file=internal-auth-mongodb-keyfile=$TMPFILE --namespace=gold
    rm $TMPFILE
    ```
    
    create pv (host path pv would be created defautly, you can change it into another storage)
    
    ```
    $ kubectl create -f mongo/volume.yaml 
    persistentvolume/gold-mongo-pv-0 created
    persistentvolume/gold-mongo-pv-1 created
    persistentvolume/gold-mongo-pv-2 created
    persistentvolume/gold-mongo-pv-3 created
    ```
    
    create headless service:
    
    ```
    $ kubectl create -f mongo/headless_service.yaml 
    service/mongo-service created
    ```
    
    create stateful set:
    
    ```
    $ kubectl create -f mongo/stateful_set.yaml 
    statefulset.apps/mongod created
    ```
    
6. set up mongo replica sets:

    there is a script to init the mongo, in `building/env/k8s/mongo/init.sh`
    
    you can run it like this:
    
    ```
    init.sh [root_pwd]
    ```
    
    then the mongo replica sets will be set up and the root user password will be set as `[root_pwd]`
    
 7. create http trigger:
 
    ```
    kubectl create -f gold/http-trigger.yaml
    ```
    
## Set up private docker registry

## Set up API server