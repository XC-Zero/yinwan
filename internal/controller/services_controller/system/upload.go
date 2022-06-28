package system

import (
	"context"
	config2 "github.com/XC-Zero/yinwan/internal/config"
	"github.com/XC-Zero/yinwan/pkg/client"
	_const "github.com/XC-Zero/yinwan/pkg/const"
	"github.com/XC-Zero/yinwan/pkg/utils/errs"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"path"
	"strconv"
	"strings"
	"time"
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
	IMAGE_BUCKERT = "images"
)

const (
	IMAGE_FORMAT_ERROR = "图片格式不符合要求"
)

// SavePic 存储图片
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

	objectName := "PIC_" + strconv.FormatInt(time.Now().UnixNano(), 10) + ext

	open, err := file.Open()
	if err != nil {
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg("打开上传图片时发生错误! "))
		return
	}
	info, err := client.MinioClient.PutObject(
		context.Background(),
		IMAGE_BUCKERT, objectName, open, file.Size, minio.PutObjectOptions{})
	if err != nil {
		ctx.JSON(_const.INTERNAL_ERROR, errs.CreateWebErrorMsg("存储图片时发生错误! "))
		return
	}

	url := config2.CONFIG.StorageConfig.MinioConfig.EndPoint + "/images/" + info.Key
	if !strings.Contains(url, "http://") {
		url = "http://" + url
	}
	ctx.JSON(_const.OK, errs.CreateSuccessMsg("上传图片成功!!", map[string]interface{}{
		"url": url,
	}))
	return

}

func imageFormat(ext string) bool {
	for _, s := range ImageFormatList {
		if ext == s {
			return true
		}
	}
	return false
}

// SaveInvoice todo 存单据
func SaveInvoice(ctx *gin.Context) {

}

// SaveExcel todo 可能需要支持线上查看Excel

func SaveExcel(ctx *gin.Context) {
	//file, err := ctx.FormFile("excel")
	//if err != nil {
	//	ctx.JSON(_const.REQUEST_PARM_ERROR, errs.CreateWebErrorMsg("上传图片失败!"))
	//	return
	//}
}
