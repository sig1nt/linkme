package main

import (
	"net/http"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		panic("Bad Arguments")
	}

	url, err := url.Parse(os.Args[1])
	if err != nil {
		panic(err)
	}

	url.Scheme = "https"

	http.ListenAndServe("0.0.0.0:8080", http.RedirectHandler(url.String(), 301))
}
