package htwml

import (
	"code.google.com/p/go-html-transform/h5"
	"encoding/xml"
	"io"
)

type Element struct {
	XMLName  xml.Name
	Value    string `xml:",chardata"`
	Children []Element
}

func (element *Element) Append(e Element) {
	element.Children = append(element.Children, e)
}

func Response() Element {
	return Element{XMLName: xml.Name{Local: "Response"}}
}

func Play(source string) Element {
	return Element{XMLName: xml.Name{Local: "Play"}, Value: source}
}

func Say(source string) Element {
	return Element{XMLName: xml.Name{Local: "Say"}, Value: source}
}

func Parse(r io.Reader) ([]byte, error) {
	p := h5.NewParser(r)
	err := p.Parse()

	if err != nil {
		return nil, err
	}

	tree := p.Tree()
	resp := Response()

	tree.Walk(func(n *h5.Node) {
		tag := n.Data()
		switch {
		case tag == "p":
			resp.Append(Say(n.Children[0].String()))
		case tag == "audio":
			for _, a := range n.Attr {
				if a.Name == "src" {
					resp.Append(Play(a.Value))
				}
			}
		}
	})

	return xml.MarshalIndent(resp, "", "  ")
}
