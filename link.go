package link

import (
	"strings"

	"golang.org/x/net/html"
)

//Link is the representation of an a tag in HTML
//document
type Link struct {
	address string
	text    string
}

//getLinkText extracts "a" tag text data ignoring all other html tags
func getLinkText(n *html.Node) string {
	var res string
	if n.Type == html.TextNode {
		res = n.Data
	}
	// searching for text in all nodes recursively
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res += getLinkText(c)
	}
	res = strings.Join(strings.Fields(res), " ")
	return res
}

//getLinkAddress extracts link href from an "a" tag
//by looping through given node's attributes
func getLinkAddress(n *html.Node) string {
	var res string
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			res = attr.Val
		}
	}
	return res
}

//GetLinks will extract all a tags from an HTML file and return them
//as a slice of links
func GetLinks(n *html.Node) []Link {
	var result []Link
	if n.Type == html.ElementNode && n.Data == "a" {
		link := new(Link)
		link.address = getLinkAddress(n)
		link.text = getLinkText(n)
		result = append(result, *link)
		return result
	}
	// traverse through all HTML nodes recursively
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result = append(result, GetLinks(c)...)
	}
	return result
}
