# Prerequisites

This tutorial uses Docker to build container images and Docker Swarm or Kubernetes to orchestrate the containers.

## Install Docker

To install Docker in your machine follow the instructions [here](https://docs.docker.com/engine/installation/) depending your OS.

Once Docker is installed you can test the installation with the following command:

```shell
$ docker run --rm hello-world

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://cloud.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/engine/userguide/
```

## Activate Docker Swarm mode

Docker is able to run containers but if we want to orchestrate them and take advantage of Health checks we need to use Docker Swarm mode.

If you want to see if Docker Swarm mode is activated in your machine, ask for the list of nodes, if the node is not a swarm manager

```shell
$ docker swarm init
Swarm initialized: current node (XXXXXXXXXXXXXXXX) is now a manager.

To add a worker to this swarm, run the following command:

    docker swarm join --token XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX XXX.XXX.XXX.XXX:2377

To add a manager to this swarm, run 'docker swarm join-token manager' and follow the instructions.
```

## Install Kubernetes

To install Kubernetes you can use [Minikube](https://github.com/kubernetes/minikube).

Once installed you can start a Kubernetes instance with:

```shell
$ minikube start
Starting local Kubernetes v1.8.0 cluster...
Starting VM...

Getting VM IP address...
Moving files into cluster...
Setting up certs...
Connecting to cluster...
Setting up kubeconfig...
Starting cluster components...
Kubectl is now configured to use the cluster.
```

You'll also need to install [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/) to communicate with the cluster.

Once installed you can check the connectivity with:

```shell
$ kubectl get nodes
NAME       STATUS    ROLES     AGE       VERSION
minikube   Ready     <none>    1m        v1.8.0
```

## Download Code

Download this repository to be able to run the commads locally:

```shell
git clone https://github.com/spiddy/zero-downtime-tutorial.git
cd zero-downtime-tutorial
```

Next steps assume you're on the root directory of this repo.

Next: [Build Coin Images](./02-build-coin-images.md)