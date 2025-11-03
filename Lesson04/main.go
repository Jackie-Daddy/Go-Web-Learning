package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//2.解析模版
	t, err := template.ParseFiles("./hello.tmpl") //请勿刻舟求剑（使用go build）
	if err != nil {
		fmt.Printf("render template failed, err: %v\n", err)
		return
	}
	//3.渲染模版
	name := "jackie"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Printf("Execute template failed, err: %v\n", err)
		return
	}
}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTP server failed, err: %v\n", err)
		return
	}
}
