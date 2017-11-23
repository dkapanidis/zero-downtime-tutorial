# Build Coin Images

Coin is a web server that listens on port `:80`, faulty by design:

* On startup it flips a coin (50% change)
  * On success:
    * It always respond 200 on a GET request
  * On failure:
    * I always respond 500 on a GET request

## The code

The code is located at `images/coin` and is the following:

`server.go`

```go
package main

import (
    "fmt"
    "net/http"
    "log"
    "math/rand"
    "time"
)

var coin int

func serveHTTP(w http.ResponseWriter, r *http.Request) {
  if (coin == 1) {
    w.Write([]byte("200 - You are a lucky lucky person!"))
  } else {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("500 - Bad luck, try flipping the coin again!"))
  }
}

func main() {
    http.HandleFunc("/", serveHTTP) // set router
    rand.Seed( time.Now().UnixNano())
    coin = rand.Intn(2)
    fmt.Println(coin)

    err := http.ListenAndServe(":80", nil) // set listen port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
```

`Dockerfile`

```Dockerfile
FROM golang:1.8.0-alpine AS build
COPY server.go src/
RUN go build src/server.go

FROM alpine:3.5
RUN apk -U add curl
COPY --from=build /go/server /server
EXPOSE 80
CMD ["/server"]
```

To build the `coin:v1` (No health checks) image:

```shell
cd images/coin
docker build -t coin:v1 .
```

In the next version we add a health check on the `Dockerfile-v2`:

```Dockerfile
...
HEALTHCHECK --interval=30s --timeout=3s CMD curl --silent --fail http://localhost/ || exit 1
```

To build the `coin:v2` (health check included) image:

```shell
docker build -f Dockerfile-v2 -t coin:v2 .
```
