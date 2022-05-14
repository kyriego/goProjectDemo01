package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const (
	sysUsername1 string = "kyriego"
	sysPassword1 string = "787148928"

	sysUsername2 string = "blackops"
	sysPassword2 string = "123456"
)

var userCookies map[string]string = make(map[string]string)

func test(w http.ResponseWriter, r *http.Request) {

}

func math(w http.ResponseWriter, r *http.Request) {

}

func loginn(w http.ResponseWriter, r *http.Request) {
	fmt.Println(userCookies)
	r.ParseForm()
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	Cookies := r.Cookies()
	for _, cookie := range Cookies {
		if cookie.Name == "loginkey" {
			if cookie.Value == userCookies[username] {
				msg := fmt.Sprintf("免密登录成功")
				w.Write([]byte(msg))
				return
			}
		}
	}

	if (username == sysUsername1 && password == sysPassword1) ||
		(username == sysUsername2 && password == sysPassword2) {
		rand.Seed(time.Now().Unix())
		uustr := strconv.Itoa(rand.Intn(10000))
		cookie := &http.Cookie{
			Name:  "loginkey",
			Value: uustr,
		}
		userCookies[username] = uustr
		http.SetCookie(w, cookie)
		msg := "登录成功"
		w.Write([]byte(msg))
		return
	} else {
		msg := "用户名或密码输入错误"
		w.Write([]byte(msg))
		return
	}
}

func main() {
	/* 	http.HandleFunc("/test", test)
	   	http.HandleFunc("/math", math)
	   	http.HandleFunc("/login", login)
	   	http.ListenAndServe("localhost:8000", nil) */
	sm := http.NewServeMux()
	sm.HandleFunc("/test", test)
	sm.HandleFunc("/math", math)
	sm.HandleFunc("/login", loginn)
	server := &http.Server{
		Addr:    "localhost:8000",
		Handler: sm,
	}
	server.ListenAndServe()

}
