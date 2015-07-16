package nocache

import (
    "net/http"
    "net/http/httptest"
    "fmt"
    "testing"
)

func ncHandleFunc (rw http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(rw, "")
}

func Test_NoCache(t *testing.T) {

    nc := New(true)

    rw := httptest.NewRecorder()
    req, err := http.NewRequest("GET", "http://localhost/foobar", nil)

    if err != nil {
        t.Fatal(err)
    }

    req.Header.Set("ETag", "Blabla")

    nc.ServeHTTP(rw, req, ncHandleFunc)
    h := rw.Header()

    if "private, no-store, no-cache, must-revalidate, proxy-revalidate" != h.Get("Cache-Control") {
        t.Error("Expected a valid cache-Control header")
    }

    if "no-cache" != h.Get("Pragma") {
        t.Error("Expected a valid pragma header")
    }

    if "0" != h.Get("Expires") {
        t.Error("Expected a valid expires header")
    }

    if "" != h.Get("ETag") {
        t.Error("Expected a empty ETag header")
    }

}
