package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	//定义模版
	//解析模版
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err: %v\n", err)
		return
	}
	msg := "jackie"
	//渲染模版
	t.Execute(w, msg)
}
func home(w http.ResponseWriter, r *http.Request) {
	//定义模版
	//解析模版
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err: %v\n", err)
		return
	}
	msg := "jackie"
	//渲染模版
	t.Execute(w, msg)
}
func index2(w http.ResponseWriter, r *http.Request) {
	//定义模版（模版继承的方式）
	//解析模版
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index2.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err: %v\n", err)
		return
	}
	//渲染模版
	name := "jackie"
	t.ExecuteTemplate(w, "index2.tmpl", name)
}
func home2(w http.ResponseWriter, r *http.Request) {
	//定义模版（模版继承的方式）
	//解析模版
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home2.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err: %v\n", err)
		return
	}
	//渲染模版
	name := "jackiedaddy"
	t.ExecuteTemplate(w, "home2.tmpl", name)
}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTP server failed, err: %v\n", err)
		return
	}
}
