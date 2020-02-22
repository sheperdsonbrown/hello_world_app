package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

func main() {

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World! %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Alexander Brown")
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
