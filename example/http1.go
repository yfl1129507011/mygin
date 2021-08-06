package example

import (
	"fmt"
	"net/http"
)

func RunHttp1() {
	http.HandleFunc("/", indexHandle)
	http.HandleFunc("/hello", helloHandle)
	http.ListenAndServe(":9999", nil)
}

func indexHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

func helloHandle(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
