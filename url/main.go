package main

import (
	"fmt"
	"net/url"
)

func main() {
	val := make(url.Values)
	var myurl *url.URL

	myurl, _ = url.Parse("http://test.com")
	val.Add("per_page", "100")

	myurl.RawQuery = val.Encode()
	fmt.Println(myurl.String())
}
