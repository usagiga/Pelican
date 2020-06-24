package infra

import (
	"net/url"
)

type ImageInfra interface {
	Fetch(url *url.URL, destPath string) (err error)
}
