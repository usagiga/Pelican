package application

import (
	"github.com/usagiga/pelican/domain"
	"github.com/usagiga/pelican/entity"
	"github.com/usagiga/pelican/infra"
	libpath "github.com/usagiga/pelican/lib/path"
	liburl "github.com/usagiga/pelican/lib/url"
	"path"
)

type FetchImageApplicationImpl struct {
	imageInfra      infra.ImageInfra
	imageMatchModel domain.ImageMatchModel
}

func NewFetchImageApplication(
	imageInfra infra.ImageInfra,
	imageMatchModel domain.ImageMatchModel,
) (fia FetchImageApplication) {
	return &FetchImageApplicationImpl{
		imageInfra:      imageInfra,
		imageMatchModel: imageMatchModel,
	}
}

func (a *FetchImageApplicationImpl) FetchAll() (err error) {
	docs := a.getAllDocs(".")

	for _, doc := range docs {
		// Get all images from doc
		imageUrls, err := a.imageMatchModel.GetAllImageURLs(&doc)
		if err != nil {
			return err
		}

		// Download all images
		for _, imgUrl := range imageUrls {
			fileName := liburl.GetLastNode(&imgUrl)
			imgDestPath := path.Join("./", fileName)
			err = a.imageInfra.Fetch(&imgUrl, imgDestPath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (a *FetchImageApplicationImpl) getAllDocs(srcDir string) (docs []entity.Document) {
	files := libpath.DirwalkWithExt(srcDir, ".md")

	docs = []entity.Document{}
	for _, file := range files {
		doc := entity.Document{
			FilePath: file,
		}

		docs = append(docs, doc)
	}

	return docs
}
