package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"strings"
)

/*
Given some HTML, parse it into a document and then count
how many words and images there are.

What happens if the html doc is empty?
*/

var raw = `
<!DOCTYPE html>
<html>
	<body>
		<h1>My First Heading</h1>
		<p>My first paragraph</p>
		<p>HTML images are defined with the img tag:</p>
		<img src="123.jpg" width="104" height="142"> 
	</body>
</html>
`

func visit(n *html.Node, words, pics *int) {

	if n.Type == html.TextNode {
		*words += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		*pics++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c, words, pics)
	}
}

func countWordsAndImages(doc *html.Node) (int, int) {
	var words, pics int
	visit(doc, &words, &pics)
	return words, pics
}

func main() {
	doc, err := html.Parse(bytes.NewReader([]byte(raw)))
	if err != nil {
		log.Panic(err)
	}

	words, pics := countWordsAndImages(doc)
	fmt.Println(words, pics)
}
