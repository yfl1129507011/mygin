package example

import (
	"fgin"
	"net/http"
)

func RunHttp3() {
	r := fgin.New()
	r.GET("/", func(c *fgin.Context) {
		c.HTML(http.StatusOK, "<h1>fenlon, hello</h1>")
	})
	r.Run(":9990")
}
