package nocache

import (
    "log"
    "os"
    "net/http"
    "github.com/codegangsta/negroni"
)

type NoCache struct {
    noEtag bool
    Logger  *log.Logger
}

func New(noEtag bool) *NoCache {
    return &NoCache{
        noEtag: noEtag,
        Logger: log.New(os.Stdout, "[negroni no cache enabled] ", 0),
    }
}

func (c *NoCache) ServeHTTP(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
    newRw := negroni.NewResponseWriter(rw)
    newRw.Before(func (rw negroni.ResponseWriter){
        c.setNoCacheHeader(rw)
    })

    next(newRw, req)
}

func (c *NoCache) setNoCacheHeader(rw http.ResponseWriter) {
    h := rw.Header()
    h.Set("Cache-Control", "private, no-store, no-cache, must-revalidate, proxy-revalidate")
    h.Set("Pragma", "no-cache")
    h.Set("Expires", "0")
    if (c.noEtag) {
      h.Del("ETag")
    }
}