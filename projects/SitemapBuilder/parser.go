package main

import (
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Text string
	Href string
}

func ParseLinks(content string) []string {

	var urls []string

	doc, err := html.Parse(strings.NewReader(content))
	CheckError(err)
	nodeChannel := make(chan *html.Node)
	go getTargetNode(nodeChannel, doc)

	for node := range nodeChannel {
		urls = append(urls, extractHref(node))
	}
	return urls

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
