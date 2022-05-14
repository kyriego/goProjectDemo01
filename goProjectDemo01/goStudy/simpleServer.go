package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var count int
var mu sync.RWMutex

func getRequestDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	getRequestDetails(w, r)
	fmt.Fprintf(w, "URL.Path = %s\n", r.URL.Path)
	mu.Lock()
	count++
	mu.Unlock()
}

func getCount(w http.ResponseWriter, r *http.Request) {
	getRequestDetails(w, r)
	mu.RLock()
	fmt.Fprintf(w, "count = %d\n", count)
	mu.RUnlock()
}
func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", getCount)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
