package speeds

import (
    "net/url"
    "strings"
    "errors"
)

func GetFileName(url_path string) (string,error) {
    info, err :=url.Parse(url_path)
    if err != nil {
        return "", err
    }

    idx := strings.LastIndex(info.Path, "/")
    if idx == -1 {
        return "", errors.New("'/' not found in url path")
    }


    return info.Path[idx+1:], nil
}

func GetByRange(url_path string) {

}
