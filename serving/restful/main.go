package main

import (
    "fmt"
    "net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hello World.")
    fmt.Fprintf(w, "Hello World.\n")
}

func main() {
    http.HandleFunc("/", Hello)
    err := http.ListenAndServe("0.0.0.0:8080", nil)
    if err != nil {
        fmt.Println("http listen failed.")
    }
}
