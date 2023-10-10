package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	link := "http://localhost/b9c440b7f27e73af319e12c58325eacd.zip"
	response, err := http.Get(link)
	if err != nil {
		panic(err)
	}

	_, err = ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

}
