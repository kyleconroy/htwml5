package htwml

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	chk := func(err error) {
		if err != nil {
			t.Error(err)
		}
	}
	files, err := ioutil.ReadDir("testdata")
	chk(err)

	for _, fileInfo := range files {
		filename := fileInfo.Name()

		if strings.HasPrefix(filename, ".") {
			continue
		}

		if !strings.HasSuffix(filename, ".html") {
			continue
		}

		handle, err := os.Open(path.Join("testdata", filename))
		chk(err)

		output, err := Parse(handle)
		chk(err)

		twiml := strings.Replace(filename, ".html", ".xml", 1)

		xml, err := ioutil.ReadFile(path.Join("testdata", twiml))
		chk(err)

		if string(output)+"\n" != "\n"+string(xml) {
			fmt.Println(string(output))
			fmt.Println(string(xml))
			t.Error("Files are different")
		}
	}
}
