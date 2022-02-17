package main

import (
	"fmt"
	"net/http"
)

func main()  {

	engine := New()
	engine.GET("/hzf", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	engine.Run(":8080")
}
