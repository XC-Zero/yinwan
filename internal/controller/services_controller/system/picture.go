package system

import (
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
)

var ImageFormatList = []string{
	".png",
	".jpg",
	".jpeg",
	".tiff",
	".tif",
	".svg",
}

const (
	IMAGE_FORMAT_ERROR = "图片格式不符合要求"
)

func SavePic(ctx *gin.Context) {
	file, err := ctx.FormFile("image_file")
	if err != nil {
		ctx.JSON(_const.REQUEST_PARM_ERROR, errs.CreateWebErrorMsg("上传图片失败!"))
		return
	}
	ext := strings.ToLower(path.Ext(file.Filename))
	if !imageFormat(ext) {
		ctx.JSON(_const.REQUEST_PARM_ERROR, errs.CreateWebErrorMsg(IMAGE_FORMAT_ERROR))
		return
	}
	object, err := client.MinioClient.PutObject()
	if err != nil {
		return
	}
}

func imageFormat(ext string) bool {
	for _, s := range ImageFormatList {
		if ext == s {
			return true
		}
	}
	return false
}
