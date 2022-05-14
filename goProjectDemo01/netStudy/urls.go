package main

import "net/http"

func main() {
	cookie := http.Cookie{
		Name:   "token",
		Value:  "1212616g",
		MaxAge: 1800,
	}
}
