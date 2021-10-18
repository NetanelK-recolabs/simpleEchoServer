package main

import (
    "fmt"
    "net/http"
)

/* Simple Go(lang) echo server */

func main() {
    http.HandleFunc("/", HelloServer)
    http.ListenAndServe(":8080", nil)
}

// TODO: change this to echo on client msgs instead of URL path
func EchoServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

