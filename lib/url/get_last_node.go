package url

import (
	"net/url"
	"strings"
)

// GetLastNode returns Last Node of URL.
// Last Node is the very last fragment of URL divided by `/`.
// (ex: http://example.com/foo/bar.png 's Last Node is bar.png)
func GetLastNode(url *url.URL) (lastNode string) {
	path := url.Path
	pathFragments := strings.Split(path, "/")

	return pathFragments[len(pathFragments)-1]
}
