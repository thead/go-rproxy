/*
  Reverse proxy server
  Template from Eric Gravert demo code
*/
package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func main() {
	// Where the proxy talks to
	u, _ := url.Parse("http://localhost:1718")
	proxy := httputil.NewSingleHostReverseProxy(u)

	// I wrap the reverse proxy in another http.Handler so I can
	// do interesting things with the request
	recorder := &recorder{next: proxy}
	http.ListenAndServe(":6432", recorder)
}

type recorder struct {
	next http.Handler
}

// dump the request url and duration of the proxied call
func (rec *recorder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())
	start := time.Now()
	rec.next.ServeHTTP(w, r)
	fmt.Println(time.Since(start).String)
}
