package domain

import (
	"github.com/usagiga/pelican/entity"
	"net/url"
)

type ImageMatchModel interface {
	GetAllImageURLs(doc *entity.Document) (imageUrls []url.URL, err error)
}
