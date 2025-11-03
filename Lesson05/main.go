package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//定义模版
	//解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err: %v\n", err)
		return
	}
	u1 := User{ //u1.Name
		Name:   "jackie",
		Gender: "男",
		Age:    20,
	}
	m1 := map[string]interface{}{
		"name":   "jackie",
		"gender": "男",
		"age":    18,
	}
	hobbylist := []string{"篮球", "足球", "双色球"}
	//渲染模版
	t.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbylist,
	})
}
func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTP server failed, err: %v\n", err)
		return
	}
}
