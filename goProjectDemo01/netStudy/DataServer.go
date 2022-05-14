package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	mux.Handle("/update", http.HandlerFunc(db.update))
	mux.Handle("/remove", http.HandlerFunc(db.remove))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	s := req.Method
	if s == "POST" {
		req.ParseForm()
		item := req.PostFormValue("item")
		price := req.PostFormValue("price")
		f, _ := strconv.ParseFloat(price, 64)
		db[item] = dollars(f)
		fmt.Fprintf(w, "数据更新成功!\n")
		fmt.Fprintf(w, "当前数据如下:\n")
		for key, value := range db {
			fmt.Fprintf(w, "%s:%s\n", key, value)
		}
	} else {
		w.WriteHeader(405)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func (db database) remove(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	if method == "POST" {
		req.ParseForm()
		item := req.PostFormValue("item")
		if _, ok := db[item]; !ok {
			fmt.Fprintf(w, "删除失败:不存在item为%s的数据条目\n", item)
		} else {
			delete(db, item)
			fmt.Fprintf(w, "%s已删除成功\n", item)
			fmt.Fprintf(w, "最新数据如下:\n")
			for key, value := range db {
				fmt.Fprintf(w, "%s:%s\n", key, value)
			}
		}
	} else {
		w.WriteHeader(405)
	}
}
