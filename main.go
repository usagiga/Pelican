package main

import (
	"github.com/usagiga/pelican/application"
	"github.com/usagiga/pelican/domain"
	"github.com/usagiga/pelican/infra"
	"log"
)

func main() {
	// Infrastructure
	imageInfra := infra.NewImageInfra()

	// Domain
	imageMatchModel := domain.NewImageMatchModel()

	// Applications
	fetchImageApplication := application.NewFetchImageApplication(imageInfra, imageMatchModel)

	// Run
	err := fetchImageApplication.FetchAll()
	if err != nil {
		log.Fatalln("Can't fetch images. Details: ", err)
	}
}
