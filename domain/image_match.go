package domain

import (
	"github.com/usagiga/pelican/entity"
	"net/url"
	"regexp"
)

type ImageMatchModelImpl struct {
	imageRule *regexp.Regexp
}

func NewImageMatchModel() (im ImageMatchModel) {
	imgRuleStr := "https?://.+/.+\\.(png|jpeg|jpg|jfif|gif|webp)"
	return &ImageMatchModelImpl{imageRule: regexp.MustCompile(imgRuleStr)}
}

func (m *ImageMatchModelImpl) GetAllImageURLs(doc *entity.Document) (imageUrls []url.URL, err error) {
	imageUrls = []url.URL{}

	contentStr, err := doc.ReadAllText()
	if err != nil {
		return nil, err
	}

	foundUrlStrs := m.imageRule.FindAllString(contentStr, -1)
	for _, foundUrlStr := range foundUrlStrs {
		foundUrl, err := url.Parse(foundUrlStr)
		if err != nil {
			return nil, err
		}

		imageUrls = append(imageUrls, *foundUrl)
	}

	return imageUrls, nil
}
