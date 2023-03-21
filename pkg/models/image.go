package models

import (
	"image"
	"image/jpeg"
	"io/ioutil"

	"os"

	"github.com/mjk712/kartinochki/pkg/lib/e"
	"github.com/mjk712/kartinochki/pkg/utils"
	"github.com/nfnt/resize"
)

func DecodeImage(url, filename string) image.Image {
	const errMsg = "Error By Decode Image"

	utils.DownloadFile(url, filename)
	file, err := os.Open(filename)

	if err != nil {
		e.Wrap(errMsg, err)
	}

	img, err := jpeg.Decode(file)
	if err != nil {
		e.Wrap(errMsg, err)
	}
	return img
}

func EncodeImage(img image.Image, filename, x, y string) ([]byte, error) {
	const errMsg = "Error By Encode Image"
	imgName := x + "_" + y + filename

	out, err := os.Create(imgName)
	if err != nil {
		e.Wrap(errMsg, err)
		return nil, err
	}
	jpeg.Encode(out, img, nil)
	buf, err := ioutil.ReadFile(imgName)

	if err != nil {
		e.Wrap(errMsg, err)
		return nil, err
	}
	return buf, nil
}

func ResizeImage(x, y uint, img image.Image) image.Image {

	competeImage := resize.Resize(x, y, img, resize.Lanczos3)

	return competeImage
}
