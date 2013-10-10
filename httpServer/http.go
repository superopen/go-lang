package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
)

func testHttp(w http.ResponseWriter,
				r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Fprintf(w, "Hello Http Client!")
	fmt.Println("path", r.URL.Path)
	fmt.Fprintf(w, "path: %s\n", r.URL.Path)

	fmt.Println("scheme", r.URL.Scheme)
	fmt.Fprintf(w, "scheme: %s\n", r.URL.Scheme)

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Fprintf(w, "key: %s\n", k)
		fmt.Println("val:", strings.Join(v, ""))
		fmt.Fprintf(w, "val: %s\n", strings.Join(v, ""))
	}
}

func main() {
	http.HandleFunc("/", testHttp)
	err := http.ListenAndServe(":4000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
