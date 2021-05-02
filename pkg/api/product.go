package api

import (
	"github.com/gin-gonic/gin"
	"github.com/maulanay85/go-micro-product/pkg/config"
	"github.com/maulanay85/go-micro-product/pkg/util"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func UploadProductImage(c *gin.Context) {
	allowedExtension := []string{".JPG", ".JPEG", ".PNG"}
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, config.ResponseBadRequest(""))
		return
	}

	extension := strings.ToUpper(filepath.Ext(file.Filename))

	if found := util.FindElementExist(extension, allowedExtension); found == false {
		c.JSON(http.StatusBadRequest, config.ResponseBadRequest("Extension Not Allowed"))
		return
	}

	folderImage := config.GetString("path.product_image")
	folderExist := util.CheckDirectoryExist(folderImage + "/image/product")
	if folderExist == false {
		os.MkdirAll(folderImage+"/image/product", os.ModePerm)
	}

	if err := c.SaveUploadedFile(file, folderImage+"/image/product/product"+util.GeneratedFileName()+extension); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, config.ResponseServerError(""))
		return
	}

	c.JSON(http.StatusOK, config.ResponseSuccess(util.GeneratedFileName()+extension))

}
