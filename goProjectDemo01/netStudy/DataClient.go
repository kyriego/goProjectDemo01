package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	/* 	resp, err := http.Get("http://localhost:8000/list")
	   	if err != nil {
	   		log.Printf("no response for http get to %s:%v", "localhost:8000/list", err)
	   	}
	   	defer resp.Body.Close()
	   	b, _ := ioutil.ReadAll(resp.Body)
	   	fmt.Println(string(b)) */
	/* 	var values url.Values = make(url.Values)
	   	values.Set("item", "james")
	   	values.Set("price", "50.25")
	   	resp, _ := http.PostForm("http://localhost:8000/update", values)
	   	defer resp.Body.Close()
	   	b, _ := ioutil.ReadAll(resp.Body)
	   	fmt.Println(string(b)) */

	var values url.Values = make(url.Values)
	values.Add("item", "shoes")
	resp, _ := http.PostForm("http://localhost:8000/remove", values)
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))
}
