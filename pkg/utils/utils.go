package utils

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/mjk712/kartinochki/pkg/lib/e"
)

func DeleteImg(path string) error {
	return os.Remove(path)
}

func ParceUrl(u string) (string, error) {

	ur, err := url.ParseRequestURI(u)
	if err != nil {
		er := e.Wrap("Error while parce Url", err)
		return "", er
	}
	return ur.String(), nil
}

func GetImgCheckName(x, y, imgname string) string {
	imgName := x + "_" + y + imgname
	return imgName
}

func GetImgName(u string) (string, error) {
	r, err := http.NewRequest("GET", u, nil)

	if err != nil {
		er := e.Wrap("Error while Get Img Name", err)
		return "", er
	}

	return path.Base(r.URL.Path), nil
}

func DownloadFile(URL, filename string) error {
	const errMsg = "Error by Download"
	res, err := http.Get(URL)
	if err != nil {
		er := e.Wrap(errMsg, err)
		return er
	}
	defer res.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		er := e.Wrap(errMsg, err)
		return er
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		er := e.Wrap(errMsg, err)
		return er
	}
	return nil
}

func MoveFile(newPath, filename string) {
	err := os.Rename(filename, newPath)
	if err != nil {
		e.Wrap("Move error", err)
	}
}
