package main

import (
    "fmt"
    "net/http"
    "os"
)

func Hello(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hello World.")
    fmt.Fprintf(w, "Hello World.\n")
    hostname, err := os.Hostname()
    if err == nil {
        fmt.Fprintf(w, "hostname: " + hostname + "\n")    
    }
}

func main() {
    http.HandleFunc("/", Hello)
    err := http.ListenAndServe("0.0.0.0:8080", nil)
    if err != nil {
        fmt.Println("http listen failed.")
    }
}
