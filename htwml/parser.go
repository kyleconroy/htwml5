package htwml

import (
	"code.google.com/p/go-html-transform/h5"
	"encoding/xml"
	"io"
)

type Response struct {
		XMLName xml.Name `xml:"Response"`
		Plays []Play
}

type Play struct {
		XMLName xml.Name `xml:"Play"`
		Source string `xml:",innerxml"`
}

func Parse(r io.Reader) ([]byte, error) {
	p := h5.NewParser(r)
	err := p.Parse()

	if err != nil {
		return nil, err
	}

	tree := p.Tree()

	tree.Walk(func(n *h5.Node) {
	})

	resp := Response{Plays: make([]Play, 1)}
	resp.Plays[0] = Play{Source: "http://example.com/foo.mp3"}

	return xml.Marshal(resp)
}
