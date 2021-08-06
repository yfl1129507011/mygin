package example

import (
	"fgin"
	"fmt"
	"net/http"
)

func RunHttp3() {
	r := fgin.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL=%q\n", req.URL)
	})

	r.Run(":9998")
}
