package controllers

import (
	//"bytes"

	"fmt"
	"net/http"
	"strconv"

	//"net/url"

	"github.com/gorilla/mux"
	//"github.com/mjk712/kartinochki/cmd"
	"github.com/mjk712/kartinochki/pkg/cash"
	"github.com/mjk712/kartinochki/pkg/lib/e"
	"github.com/mjk712/kartinochki/pkg/models"
	"github.com/mjk712/kartinochki/pkg/utils"
)

func ImageShow(w http.ResponseWriter, r *http.Request) {

	c := cash.NewLru(1)

	var a int

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

	//Проверяю на наличие в кэшэ
	if img := c.Get(checkName); img != nil {
		buf, err := models.EncodeRawImage(img, checkName)
		if err != nil {
			fmt.Fprintln(w, "error while Encoding image")
		}
		utils.DeleteImg(imgName)

		w.Write(buf)
	}

	img := models.DecodeImage(imgPath, imgName)

	resImg := models.ResizeImage(uint(X), uint(Y), img)

	buf, err := models.EncodeImage(resImg, imgName, imageX, imageY)
	if err != nil {
		fmt.Fprintln(w, "error while Encoding image")
	}
	utils.DeleteImg(imgName)

	//прописать добавление картинки в кэш
	a++
	c.Set(imgName, resImg)
	utils.DeleteImg(checkName)

	w.Header().Set("Content-Type", "image/jpeg")

	w.Write(buf)

}
