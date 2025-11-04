package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//定义一个函数kua
	//要么只有一个返回值，要么有两个返回值，第二个返回值必须是一个error类型
	k := func(name string) (string, error) {
		return name + "年轻又帅气", nil
	}
	//定义模版
	t := template.New("f.tmpl") //创建一个名字是f的模版对象，名字一定要与模版的名字能对应上
	//告诉模版引擎，我现在多了一个自定义的函数
	t.Funcs(template.FuncMap{
		"kua99": k,
	})
	//解析模版
	_, err := t.ParseFiles("./f.tmpl")
	if err != nil {
		fmt.Printf("parse template failed, err: %v\n", err)
		return
	}
	name := "jackie"
	//渲染模版
	t.Execute(w, name)
}
func demo1(w http.ResponseWriter, r *http.Request) {
	//定义模版
	//解析模版
	t, err := template.ParseFiles("./t.tmpl", "./ul.tmpl") //被包含的写在后面
	if err != nil {
		fmt.Printf("parse template failed, err: %v\n", err)
		return
	}
	name := "小王子"
	//渲染模版
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo1)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTP server failed, err: %v\n", err)
		return
	}
}
