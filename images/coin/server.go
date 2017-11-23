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
