package util

import (
	"net/url"
	"strings"
)


func GetFileNameFromUrl(urlPath string) (string) {
	info, err := url.Parse(urlPath)
	if err != nil {
		return urlPath
	}

	fileName := info.Path

	// trim left part
	idx := strings.LastIndex(info.Path, "/")
	if idx != -1 {
		fileName = fileName[idx+1:]
	}

	// trim right part
	idx = strings.Index(fileName, "?")
	if idx != -1 {
		fileName = fileName[:idx]
	}

	return fileName
}

func GetByRange(url_path string) {

}
