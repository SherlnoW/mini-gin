package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
*
定义空的结构体 Engine，实现方法ServeHTTP
*/
type Engine struct{}

/*
ResponseWriter: 构造针对该请求的响应
Request: 包含该HTTP请求的所有的信息
*/
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
