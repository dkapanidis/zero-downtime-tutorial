# Deploy on Kubernetes

## Build Icon on Kubernetes Minikube

```shell
eval $(minikube docker-env)
docker build -t spiddy/coin images/coin/
```

## Deploy Coin & Status Page

The following components are created in Kubernetes

* A [Coin Deployment](../kubernetes/coin-deployment.yaml) that manages the Coin Pods
* A [Coin Service](../kubernetes/coin-service.yaml) that connects a Node Port to the Coin instances, load balancing them.
* A [Status Page Deployment](../kubernetes/statuspage-deployment.yaml) that manages the Status Page Pod
* A [Status Page Service](../kubernetes/statuspage-service.yaml) that connects a Node Port to the Status Page instance.

```shell
kubectl create -f kubernetes/
```

Get the Pods

```shell
$ kubectl get pods
sNAME                          READY     STATUS    RESTARTS   AGE
coin-7895dcd8b8-lxdbb         1/1       Running   0          19s
statuspage-7656d9bfd6-k7xj6   1/1       Running   0          19s
```

Get the Services:

```shell
$ kubectl get services
NAME         TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)        AGE
coin         NodePort    10.0.0.185   <none>        80:30367/TCP   40s
kubernetes   ClusterIP   10.0.0.1     <none>        443/TCP        16h
statuspage   NodePort    10.0.0.146   <none>        80:32669/TCP   40s
```

Get the Minikube IP

```shell
$ minikube ip
192.168.99.100
```

Open browser to `http://MINIKUBE_IP:STATUSPAGE_PORT`

## Scale Up Coins

```shell
$ kubectl scale deployment coin --replicas=10
deployment "coin" scaled
```

Check Pods

```shell
$ kubectl get pods

```

Check output of Coin

```shell
repeat 10 echo $(curl -s 192.168.99.100:30367)
```

## Activate Health Checks to Coins

We'll now reconfigure existing Coin deployment with [new configuration](../kubernetes-with-healthchecks/coin-deployment-with-healthchecks.yaml) that includes LivenessProbe and ReadinessProbe:

```shell
$ kubectl apply -f kubernetes-with-healthchecks/
deployment "coin" configured
```

Check Pods

```shell
$ watch kubectl get pod
NAME                          READY     STATUS    RESTARTS   AGE
coin-6bd789f884-dlzgx         1/1       Running   0          8m
statuspage-7656d9bfd6-qfwm9   1/1       Running   0          13m
```

Check output of Coin

```shell
repeat 10 echo $(curl -s 192.168.99.100:30367)
```
