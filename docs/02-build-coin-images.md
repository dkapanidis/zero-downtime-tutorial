# Build Coin Images

Coin is a web server that listens on port `:80`, faulty by design:

* On startup it flips a coin (50% change)
  * On success:
    * It always respond 200 on a GET request
  * On failure:
    * I always respond 500 on a GET request

## Build Coin without Health check (coin:v1)

Build the `coin:v1` Image with the [server.go](../images/coin/server.go) go app using this [Dockerfile](../images/coin/Dockerfile):

```shell
docker build -t coin:v1 images/coin/
```

## Build Coin with Health Check (coin:v2)

Build the `coin:v2` Image using this [Dockerfile-v2](../images/coin/Dockerfile-v2):

```shell
docker build -f images/coin/Dockerfile-v2 -t coin:v2 images/coin/
```
