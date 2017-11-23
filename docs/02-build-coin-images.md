# Build Container Images

## Coin

Coin is a web server that listens on port `:80`, faulty by design:

* On startup it flips a coin (50% change)
  * On success:
    * It always respond 200 on a GET request
  * On failure:
    * I always respond 500 on a GET request

### Build Coin without Health check (spiddy/coin:v1.0)

Build the `spiddy/coin:v1.0` Image with the [server.go](../images/coin/server.go) go app using this [Dockerfile](../images/coin/Dockerfile):

```shell
docker build -t spiddy/coin:v1.0 images/coin/
```

### Build Coin with Health Check (spiddy/coin:v2.0)

Build the `spiddy/coin:v2.0` Image using this [Dockerfile-v2.0](../images/coin/Dockerfile-v2.0):

```shell
docker build -f images/coin/Dockerfile-v2.0 -t spiddy/coin:v2 images/coin/
```

## Run Coin

Run a coin app and expose it on port `:8080`

```shell
docker run --name coin -d -p 8080:80 spiddy/coin:v1.0
curl localhost:8080
```

Clean up

```shell
docker stop coin
docker rm coin
```