package main

import (
	"fmt"
	"net/http"
)

func main() {
	h1 := make(http.Header)
	var h2 http.Header
	h1.Add("a", "b")
	h1.Add("c", "f")
	h1.Add("d", "e")
	h2 = h1
	fmt.Println("h1:", h1)
	fmt.Println("h2:", h2)
}
