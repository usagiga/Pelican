package entity

import (
	"io/ioutil"
)

type Document struct {
	FilePath string
}

func (d *Document) ReadAllText() (contents string, err error) {
	contentBytes, err := ioutil.ReadFile(d.FilePath)
	if err != nil {
		return "", err
	}

	contents = string(contentBytes)

	return contents, err
}
