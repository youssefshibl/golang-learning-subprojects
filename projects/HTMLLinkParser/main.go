package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Text string
	Href string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	fileName := flag.String("file", "ex1.html", "Html file name")
	flag.Parse()
	fmt.Println(*fileName)

	r, err := os.Open(*fileName)
	check(err)
	defer r.Close()
	doc, err := html.Parse(r)
	check(err)
	nodeChannel := make(chan *html.Node)
	go getTargetNode(nodeChannel, doc)

	for node := range nodeChannel {

		fmt.Println(Link{
			Text: extractText(node),
			Href: extractHref(node),
		})
	}
}

func getTargetNode(nc chan *html.Node, node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		nc <- node
		return
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		getTargetNode(nc, c)
	}

	if node.Parent == nil {
		close(nc)
	}

}

func extractHref(a *html.Node) string {
	for _, att := range a.Attr {
		if att.Key == "href" {
			return att.Val
		}
	}
	return ""
}

func extractText(a *html.Node) string {
	var res string

	for c := a.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			res += c.Data
			continue
		}
		res += extractText(c)
	}
	return strings.TrimSpace(res)
}
