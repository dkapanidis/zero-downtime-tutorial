# Deploy on Kubernetes

## Build Icon on Kubernetes Minikube

```shell
eval $(minikube docker-env)
docker build -t spiddy/coin:v1.0 -t spiddy/coin:v2.0 images/coin/
```

## Deploy Coin & Status Page

The following components are created in Kubernetes

* A [Coin Deployment](../kubernetes/coin-deployment.yaml) that manages the Coin Pods
* A [Coin Service](../kubernetes/coin-service.yaml) that exposes Coin instances to Node Port `:31000`, load balancing them.
* A [Status Page Deployment](../kubernetes/statuspage-deployment.yaml) that manages the Status Page Pod
* A [Status Page Service](../kubernetes/statuspage-service.yaml) that exposes Status Page instance to Node Port `:32000`.

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
coin         NodePort    10.0.0.185   <none>        80:31000/TCP   40s
kubernetes   ClusterIP   10.0.0.1     <none>        443/TCP        16h
statuspage   NodePort    10.0.0.146   <none>        80:32000/TCP   40s
```

Get the Minikube IP

```shell
$ minikube ip
192.168.99.100
```

Get response from Coin at port `:31000`

```shell
$ curl 192.168.99.100:31000
500 - Bad luck, try flipping the coin again!
```

Open browser to [http://192.168.99.100:32000](http://192.168.99.100:32000)

## Scale Up Coins

```shell
$ kubectl scale deployment coin --replicas=10
deployment "coin" scaled
```

Check Pods

```shell
$ kubectl get pods
NAME                          READY     STATUS    RESTARTS   AGE
coin-7895dcd8b8-2swck         1/1       Running   0          14s
coin-7895dcd8b8-4gg56         1/1       Running   0          14s
coin-7895dcd8b8-b54gh         1/1       Running   0          14s
coin-7895dcd8b8-fmg6r         1/1       Running   0          14s
coin-7895dcd8b8-h2s4g         1/1       Running   0          14s
coin-7895dcd8b8-hmbzh         1/1       Running   0          14s
coin-7895dcd8b8-j7k5j         1/1       Running   0          14s
coin-7895dcd8b8-m5wmq         1/1       Running   0          14s
coin-7895dcd8b8-mrb9t         1/1       Running   0          5m
coin-7895dcd8b8-zk8b6         1/1       Running   0          14s
```

Check output of Coin

```shell
$ repeat 10 echo $(curl -s 192.168.99.100:31000)
500 - Bad luck, try flipping the coin again!
500 - Bad luck, try flipping the coin again!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
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
$ repeat 10 echo $(curl -s 192.168.99.100:31000)
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
```

## Upgrade Coin

Let's upgrade the Coin to `v2.0`:

```shell
$ kubectl set image deployment coin 'coin=spiddy/coin:v2.0'
deployment "coin" image updated
```

Check output of Coin

```shell
$ repeat 10 echo $(curl -s 192.168.99.100:31000)
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
200 - You are a lucky lucky person!
```

## Cleanup

```shell
kubectl delete -f kubernetes
```
