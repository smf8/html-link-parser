package main

import (
	"flag"
	"fmt"
	"os"

	link "github.com/smf8/html-link-parser"
	"golang.org/x/net/html"
)

func main() {
	filename := flag.String("file", "files/ex2.html", "HTML file address")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}
	doc, err := html.Parse(file)
	if err != nil {
		panic(err)
	}
	links := link.GetLinks(doc)

	for _, v := range links {
		fmt.Printf("%+v\n", v)
	}
}
