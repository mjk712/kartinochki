package controllers

import (
	// "bytes"

	"fmt"
	"image/jpeg"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mjk712/kartinochki/cash"
	"github.com/mjk712/kartinochki/config"
	"github.com/mjk712/kartinochki/lib/e"
	"github.com/mjk712/kartinochki/pkg/models"
	"github.com/mjk712/kartinochki/utils"
)

type Controller struct {
	cache *cash.LRU
}

func NewController(cache *cash.LRU) Controller {
	return Controller{
		cache: cache,
	}
}

func (c *Controller) ImageShow(w http.ResponseWriter, r *http.Request) {
	db := config.DbPath()
	vars := mux.Vars(r)
	imageX := vars["imageX"]
	imageY := vars["imageY"]
	rawUrl := vars["imgUrl"]
	imgUrl := "https://" + rawUrl

	Y, err := strconv.ParseInt(imageY, 0, 0)
	if err != nil {
		e.Wrap("error while parsing Y value", err)
	}
	X, err := strconv.ParseInt(imageX, 0, 0)
	if err != nil {
		e.Wrap("error while parsing X value", err)
	}

	imgPath, err := utils.ParceUrl(imgUrl)
	if err != nil {
		fmt.Print(err)
	}
	imgName, err := utils.GetImgName(imgPath)
	if err != nil {
		fmt.Print(err)
	}
	checkName := utils.GetImgCheckName(imageX, imageY, imgName)
	fmt.Println(checkName)
	// Проверяю на наличие в кэшэ

	if img, ok := c.cache.Get(checkName); ok == true {

		buf, err := models.EncodeRawImage(img, checkName)
		if err != nil {
			fmt.Fprintln(w, "error while Encoding image")
		}
		utils.DeleteImg(checkName)
		fmt.Println("found in cache")
		w.Write(buf)
		//проверяю на наличие в db

	} else if utils.FileExists(checkName) {

		//file, err := os.Open("/home/greg/Рабочий стол/kartinochki/kartinochki/cmd/db/" + checkName)
		file, err := os.Open(db + checkName)
		if err != nil {
			fmt.Println("err db img")
		}

		dbimg, err := jpeg.Decode(file)
		if err != nil {
			fmt.Println("err db img")
		}

		buf, err := models.EncodeRawImage(dbimg, checkName)
		if err != nil {
			fmt.Fprintln(w, "error while Encoding image")
		}
		utils.DeleteImg(checkName)
		fmt.Println("found in db")
		w.Write(buf)

	} else {
		img := models.DecodeImage(imgPath, imgName)

		resImg := models.ResizeImage(uint(X), uint(Y), img)

		buf, newImgName, err := models.EncodeImage(resImg, imgName, imageX, imageY)
		if err != nil {
			fmt.Fprintln(w, "error while Encoding image")
		}
		utils.DeleteImg(newImgName)

		c.cache.Set(newImgName, resImg)
		fmt.Println(newImgName + "gay")
		utils.DeleteImg(imgName)

		w.Header().Set("Content-Type", "image/jpeg")

		w.Write(buf)
	}
}
