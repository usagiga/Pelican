package path

import (
	"io/ioutil"
	"path"
	"path/filepath"
)

// Get all files recursive in specific dir.
// https://qiita.com/tanksuzuki/items/7866768c36e13f09eedb
func Dirwalk(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		if file.IsDir() {
			paths = append(paths, Dirwalk(filepath.Join(dir, file.Name()))...)
			continue
		}
		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths
}

func DirwalkWithExt(dir string, ext string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		fileName := file.Name()
		filePath := filepath.Join(dir, fileName)

		if file.IsDir() {
			paths = append(paths, DirwalkWithExt(filePath, ext)...)
			continue
		}
		if path.Ext(fileName) != ext {
			continue
		}

		paths = append(paths, filePath)
	}

	return paths
}
