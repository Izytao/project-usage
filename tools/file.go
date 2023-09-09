package tools

import (
	"github.com/gobuffalo/packr/v2"
	"os"
)

// FileGetContent 获取文件内容，可以打包到二进制
func FileGetContent(file string, targetPath string) string {
	str := ""
	// 将打包箱命名为 "tmpl"， ../static 是要打包的文件或文件夹的路径。
	box := packr.New("tmpl", targetPath)
	content, err := box.FindString(file)
	if err != nil {
		return str
	}
	return content
}

// IsFileNotExist 判断文件文件夹不存在
func IsFileNotExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return true, nil
	}
	return false, err
}
