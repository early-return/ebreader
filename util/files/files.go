package files

import (
	"archive/zip"
	"ebreader/config"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

var ()

//Unepub 解压配置中的epub文件
func Unepub() error {
	reader, err := zip.OpenReader(config.File)
	if err != nil {
		return errors.New("文件不存在或打开文件出错！")
	}
	defer reader.Close()

	for _, file := range reader.File {
		lowerName := strings.ToLower(file.Name)
		if strings.HasPrefix(lowerName, "oebps") {
			newName := config.Path + lowerName[strings.Index(lowerName, "/"):]
			err := os.MkdirAll(getPath(newName), 0755)
			if err != nil {
				return err
			}

			writer, err := os.Create(newName)
			if err != nil {
				return err
			}
			defer writer.Close()

			reader, err := file.Open()
			if err != nil {
				return err
			}

			_, err = io.Copy(writer, reader)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//Clean 清理工作目录
func Clean() {
	if strings.HasSuffix(config.Path, "ebreader") {
		err := os.RemoveAll(config.Path)
		if err != nil {
			log.Fatalln(err)
		}
	}
}

//getPath 根据完整路径获得相应的目录
func getPath(file string) string {
	return file[0:strings.LastIndex(file, "/")]
}
