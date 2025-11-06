package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	// 解析模版
	t, err := template.New("index.tmpl").
		Delims("{[", "]}").
		ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Println("parse template failed,err:", err)
		return
	}
	// 渲染模版
	name := "jackie"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("execute template failed,err:", err)
		return
	}
}

func xss(w http.ResponseWriter, r *http.Request) {
	// 定义模版
	// 解析模版
	// 解析模版之前定义一个自定义的函数safe
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		fmt.Println("parse template failed,err:", err)
		return
	}
	// 渲染模版
	str1 := "<script>alert(123);</script>"
	str2 := "<a href='https://www.baidu.com'>百度</a>"
	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("http server failed,err:", err)
		return
	}
}
