package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an HTML
// document
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and will return a
// slice of links parsed from it
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	nodes := linkNodes(doc)
	//dfs(doc, "")
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		//fmt.Println(attr)
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}
	//ret.Text = strings.Join(strings.Fields(text(n)), " ")
	ret.Text = text(n)
	return ret
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		// Maybe building a string with a byte buffer would be more optimal
		// but let's stick with this right now
		ret += text(c)
	}
	return strings.Join(strings.Fields(ret), " ")
	//return ret
}

func linkNodes(n *html.Node) []*html.Node {

	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}
	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}

	return ret
}

// func dfs(n *html.Node, padding string) {
// 	if n.Type == html.ElementNode && n.Data == "a" {
// 		fmt.Println(padding, n.Data)
// 	}
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		dfs(c, padding+"  ")
// 	}

// }
