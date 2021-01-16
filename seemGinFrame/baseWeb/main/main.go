package main

import (
	"fmt"
	"net/http"

	"github.com/mygin"
)

func main() {
	r := mygin.New()
	r.GET("/diwang", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "<b> hello %s </b>  This is the end of mygin", req.URL.Path[1:len(req.URL.Path)-1])
	})
	r.Run(":9999")
}
