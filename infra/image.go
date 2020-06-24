package infra

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

type ImageInfraImpl struct {

}

func NewImageInfra() (ii ImageInfra) {
	return &ImageInfraImpl{}
}

func (i *ImageInfraImpl) Fetch(url *url.URL, destPath string) (err error) {
	// Get image
	resp, err := http.Get(url.String())
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create empty file
	w, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer w.Close()

	// Write image into file
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
