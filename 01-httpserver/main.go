package main

import (
	"fmt"
	"io"
	"net/http"
)
 
func IndexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello Cloud-Native")
	// fmt.Println("r.Header = ", r.Header)
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		fmt.Println(k,v)
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "200\n")
}
 
func main() {
    http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe("0.0.0.0:8000",nil)
	    if err != nil {
        fmt.Println("服务器错误")
    }
}