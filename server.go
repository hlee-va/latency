package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "fmt"
    "net"
    "time"
    "runtime"
)

func buildRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/", Index)
    return router
}

func Index(w http.ResponseWriter, r *http.Request) {
    start := time.Now()
    sum := 0
    for i := 0; i != 2000000000; i++ {
        sum += i
    }
    fmt.Fprintf(w, "%s:%d", time.Since(start), sum)
}

func main()  {
    runtime.GOMAXPROCS(4)
    listener, err := net.Listen("tcp", ":9114")
    if err != nil {
        fmt.Println("Err:", err.Error())
        return
    }
    server := &http.Server{
        Handler:buildRouter(),
    }
    err = server.Serve(listener)
    if err != nil {
        fmt.Println("Err:", err.Error())
    }
}