package main

import (
	"fmt"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {

}
func main() {
	http.HandleFunc("/", f1)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("HTTP server failed, err: %v\n", err)
		return
	}
}
