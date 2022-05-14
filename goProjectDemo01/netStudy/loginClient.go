package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

var client http.Client
var jar http.CookieJar
var localCookies []*http.Cookie
var index int

func init() {
	jar, _ = cookiejar.New(nil)
	client = http.Client{
		Jar: jar,
	}
	localCookies = make([]*http.Cookie, 10)
}

func login(username string, password string) {
	cookieUrl, _ := url.Parse("http://localhost:8000/login")
	fmt.Printf("client.Jar.Cookies(): %v\n", client.Jar.Cookies(cookieUrl))
	var form url.Values
	form = make(url.Values)
	form["username"] = []string{username}
	form["password"] = []string{password}
	resp, _ := client.PostForm("http://localhost:8000/login", form)
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(os.Stdout, "%s\n", string(b))
	resp.Body.Close()
	Cookies := resp.Cookies()
	client.Jar.SetCookies(cookieUrl, Cookies)
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		msgs := strings.Split(input.Text(), " ")
		login(msgs[0], msgs[1])
	}
}
