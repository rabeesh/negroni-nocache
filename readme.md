# negroni middleware for nocache

Negroni middleware for no cache http headers

## Usage


~~~ go
package main

import (
    "fmt"
    "github.com/codegangsta/negroni"
    "github.com/rabeesh/negroni-nocache"
    "net/http"
)


func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
        fmt.Fprintf(rw, "Welcome to the home page!")
    })

    n := negroni.New()
    n.Use(nocache.New(true))
    n.Use(negroni.NewStatic(http.Dir("public")))
    n.UseHandler(mux)
    n.Run(":5000")
}
~~~
