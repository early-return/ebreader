package main

import (
	"ebreader/config"
	"ebreader/server"
	"ebreader/util/files"
	"ebreader/util/template"
	"fmt"
	"log"
)

func main() {
	start()
}

func start() {
	files.Clean()

	fmt.Println("正在打开文件......")
	err := files.Unepub()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("正在解析文件......")
	err = template.Build(config.Path + "/toc.ncx")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("开始监听本地端口......")
	server.Run()
}
