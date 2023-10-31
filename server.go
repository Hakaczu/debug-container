package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "Hello There\n")
}

func home(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "M.A.R.I.A.N - DWSC\n")
    fmt.Fprintf(w, "Manualy Activated Recovery of Infrastructue, Applications and Networking - Debug and Web Server Container\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}

func main() {

	fmt.Printf("M.A.R.I.A.N - DWSC\n")
    fmt.Printf("Manualy Activated Recovery of Infrastructue, Applications and Networking - Debug and Web Server Container\n")
	http.HandleFunc("/", home)
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)

    http.ListenAndServe(":80", nil)
}