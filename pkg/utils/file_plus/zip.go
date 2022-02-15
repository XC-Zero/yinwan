package file_plus

import (
	"archive/zip"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Zip 将文件或文件夹打包成zip文件
func Zip(srcDir string, zipFileName string) error {

	// 不能为同一路径，会无限循环增长！
	if filepath.Dir(srcDir) == filepath.Dir(zipFileName) {
		return errors.New("The source file location cannot be the same as the destination file location! ")
	}
	// 预防：旧文件无法覆盖
	_ = os.RemoveAll(zipFileName)

	// 创建：zip文件
	zipFile, _ := os.Create(zipFileName)
	defer func(zipFile *os.File) {
		_ = zipFile.Close()
	}(zipFile)

	// 打开：zip文件
	archive := zip.NewWriter(zipFile)
	defer archive.Close()

	// 遍历路径信息
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, _ error) error {

		// 如果是源路径，提前进行下一个遍历
		if path == srcDir {
			return nil
		}

		// 获取：文件头信息
		header, _ := zip.FileInfoHeader(info)
		header.Name = strings.TrimPrefix(path, srcDir+`\`)

		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Deflate
		}

		// 创建：压缩包头部信息
		writer, err := archive.CreateHeader(header)
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
		return nil
	})
}
