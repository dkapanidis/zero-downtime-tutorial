# Deploy on Docker Swarm

Docker Swarm is the orchestrator of Docker containers.

We've defined a stack on [docker-stack.yaml](../swarm/docker-stack.yaml) that deploys:

* Coin and exposes it on port `:8080` locally
* a Status Page that displays the Coin health and exposes it on port `:9000` locally

```shell
docker stack deploy -c swarm/docker-stack.yaml zerodowntime
```

Check the stacks

```shell
$ docker stack ls
NAME                SERVICES
zerodowntime        2
```

Check the services

```shell
$ docker service ls
ID                  NAME                      MODE                REPLICAS            IMAGE                      PORTS
z8unn9uzz8u3        zerodowntime_coin         replicated          3/3                 spiddy/coin:v1.0           *:8080->80/tcp
m39kmhoyww0i        zerodowntime_statuspage   replicated          1/1                 spiddy/statuspage:latest   *:9000->9000/tcp
```

Check the `coin` service

```shell
$ docker service ps zerodowntime_coin
ID                  NAME                  IMAGE               NODE                    DESIRED STATE       CURRENT STATE                ERROR               PORTS
mn5iboizxstr        zerodowntime_coin.1   spiddy/coin:v1.0    linuxkit-025000000001   Running             Running about a minute ago
s07in38bw4ec        zerodowntime_coin.2   spiddy/coin:v1.0    linuxkit-025000000001   Running             Running about a minute ago
su614xfba05h        zerodowntime_coin.3   spiddy/coin:v1.0    linuxkit-025000000001   Running             Running about a minute ago
```

Scale up the `coin` service

```shell
docker service update zerodowntime_coin --replicas=5
```

Scale up the `coin` service

```shell
docker service update zerodowntime_coin --replicas=20
```

Cleanup

```shell
docker stack rm zerodowntime
```

Next: [Deploy on Kubernetes](./04-deploy-on-kubernetes.md)