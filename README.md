Hello Cloud Profiler
---

This is a sample of [Cloud Profiler](https://cloud.google.com/profiler/docs/profiling-go), a tool for profiling the application.

## Description

This program is a sample to run the [Cloud Profiler](https://cloud.google.com/profiler/docs/profiling-go) in kubernetes.
You can see the results of the profiling as follows.

![result](https://github.com/hyorimitsu/hello-cloud-profiler/blob/master/doc/img/result.png)

## Directory Structure

```
.
├── docker-compose-tools.yml  # => dev tools for local
├── k8s                       # => k8s definitions
└── src                       # => application source
    ├── Dockerfile
    ├── cmd
    ├── gcp
    │   └── profiler          # => Cloud Profiler wrapper
    │        ├── client.go
    │        ├── constants.go
    │        └── methods.go
    └── handlers

(some omitted)    
```

## Usage

### 1. Run the application

#### 1.1. in minikube

a. start minikube

```shell
minikube start --driver=virtualbox
```

b. build docker image in minikube

```shell
# change the destination to the docker in minikube
eval $(minikube docker-env)

# confirm docker context
docker context ls

# build image
docker build -t gcr.io/hello-cloud-profiler/app-api:1.0.0 ./src

# revert the destination to the local docker
eval $(minikube docker-env -u)

# confirm docker context
docker context ls
```

c. deploy to minikube

```shell
# confirm kubectl config
kubectl config get-contexts

# apply manifest
kubectl apply -f k8s/api.yml

# get service url
minikube service hello-cloud-profiler-api --url
```

**※ If run locally, Profiler will not start.**

#### 1.2. in GKE

a. create a project named `hello-cloud-profiler` in Google Cloud

b. enable `Kubernetes Engine API`

c. build docker image

```shell
docker build -t gcr.io/hello-cloud-profiler/app-api:1.0.0 ./src
```

d. setup gcloud

```shell
# reinitialize gcloud config
gcloud init

# confirm gcloud config
gcloud config list
```

e. push to GCR

```shell
# register gcloud as a Docker credential helper ( skip this if you have already set it up )
gcloud auth configure-docker

# push to GCR
docker push gcr.io/hello-cloud-profiler/app-api:1.0.0
```

f. create and connect cluster

```shell
# create cluster
gcloud container clusters create hello-cloud-profiler-api \
    --machine-type=n1-standard-1 \
    --num-nodes=1

# change the target of kubectl to the GKE cluster
gcloud container clusters get-credentials hello-cloud-profiler-api --zone asia-northeast1-a --project hello-cloud-profiler

# confirm the target of kubectl
kubectl config get-contexts
```

g. deploy to GKE

```shell
# apply manifest
kubectl apply -f k8s/api.yml

# get service
kubectl get service hello-cloud-profiler-api

# access to api
open http://{EXTERNAL-IP}:8081
```

h. confirm Profiler

![result](https://github.com/hyorimitsu/hello-cloud-profiler/blob/master/doc/img/result.png)

### 2. Stop the application

#### 2.1. in minikube

a. delete resources

```shell
# delete service
kubectl delete service hello-cloud-profiler-api

# delete deployment
kubectl delete deployment hello-cloud-profiler-api
```

b. stop minikube

```shell
minikube stop
```

c. delete minikube

```shell
minikube delete
```

#### 2.2. in GKE

a. delete a project named `hello-cloud-profiler` in Google Cloud.

## Dev Tools for Local

### 1. resolve dependencies

#### 1.1. golang

```shell
docker-compose -f docker-compose-tools.yml run go-mod
```
