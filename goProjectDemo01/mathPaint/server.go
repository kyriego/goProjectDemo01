package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"Github.com/kyrieGo/goProjectDemo01/util"
)

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

func getmathpic(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	line := r.FormValue("line")
	ground := r.FormValue("ground")
	wid, _ := strconv.Atoi(r.FormValue("wid"))
	h, _ := strconv.Atoi(r.FormValue("h"))
	fmt.Printf("line = %s\n", line)
	fmt.Printf("ground = %s\n", ground)
	fmt.Printf("weight = %d\n", wid)
	fmt.Printf("height = %d\n", h)
	w.Header().Set("Content-Type", "text/html")
	w.Write(util.GetPicture(line, ground, wid, h))
}

func test(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", test)
	http.HandleFunc("/math", getmathpic)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
