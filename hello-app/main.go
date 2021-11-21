package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		io.WriteString(rw, "Hello Application")
	})

	fmt.Println("server 8080 started")
	http.ListenAndServe(":8080", nil)

}
