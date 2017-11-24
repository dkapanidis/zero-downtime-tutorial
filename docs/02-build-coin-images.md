# Build Container Image

## Coin

Coin is a web server that listens on port `:80`, faulty by design:

* On startup it flips a coin (50% change)
  * On success:
    * It always respond 200 on a GET request
  * On failure:
    * I always respond 500 on a GET request

### Build Coin

Build `spiddy/coin` Images with the [server.go](../images/coin/server.go) go app using this [Dockerfile](../images/coin/Dockerfile):

```shell
docker build -t spiddy/coin:v1.0 -t spiddy/coin:v2.0 images/coin/
```

> NOTE: The images are also prebuild and pushed on [Docker Hub](https://hub.docker.com/r/spiddy/coin/tags/)
>
> We've tagged two identical versions (`v1.0` and `v2.0`). We'll use them to simulate an upgrade later on.

## Run Coin

Run a coin app and expose it on port `:8080`

```shell
docker run --rm --name coin -d -p 8080:80 spiddy/coin:v1.0
curl -v localhost:8080
```

Clean up the running coin now, we'll restart it later using an orchestrator.

```shell
docker stop coin
```

> The container will be auto-removed when stopped, because we added the `--rm` flag set when started it.

Next: [Run Status Page](./03-run-status-page.md)