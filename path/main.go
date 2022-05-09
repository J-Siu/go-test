package main

import (
	"fmt"
	"net/url"
	"path"
)

type Self struct {
	User string
	Repo string
}

func main() {
	var p string
	var u *url.URL
	var self Self
	self.User = "john"
	self.Repo = "test"

	fmt.Println("hello")
	// self.In.Endpoint = "/repos/" + self.User + "/" + self.Repo
	p = path.Join("/repos", self.User, self.Repo)
	fmt.Println("p:" + p)

	p = path.Join("http://git.com", "repos", self.User, self.Repo)
	fmt.Println("p:" + p)

	u, _ = url.Parse("http:/test.com")
	fmt.Println("u:" + u.String())
	u.Path = path.Join(u.Path, "repo\\"+self.User, "/", self.Repo)
	fmt.Println("u:" + u.Path)
	fmt.Println("u:" + u.String())

	fmt.Println(`Path.Base(""):` + path.Base(""))
	fmt.Println(`Path.Base("/"):` + path.Base("/"))
	fmt.Println(`Path.Base("test"):` + path.Base("test"))
	fmt.Println(`Path.Base("test/test"):` + path.Base("test/test"))
}
