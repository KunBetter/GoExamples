// web 服务器
package main

//包 http 通过任何实现了 http.Handler 的值来响应 HTTP 请求
/*
type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
*/

import (
	"fmt"
	"net/http"
)

type hello struct{}

func (h *hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func web() {
	var h hello
	http.ListenAndServe("localhost:4000", &h)
}

type String string

type Struct struct {
	Greeting, Punct, Who string
}

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

func (str *Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, str.Greeting)
}

func web_route() {
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	http.ListenAndServe("localhost:4000", nil)
}
