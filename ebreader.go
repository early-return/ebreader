package main

import (
	"ebreader/config"
	"ebreader/server"
	"ebreader/util/files"
	"ebreader/util/template"
	"flag"
	"fmt"
	"log"
	"os"
)

func init() {
	initConfig()
	config.Path += "/ebreader"
	files.Clean()
}

func main() {
	start()
}

func start() {
	err := files.Unepub()
	if err != nil {
		log.Fatalln(err)
	}

	err = template.Build(config.Path + "/toc.ncx")
	if err != nil {
		log.Fatalln(err)
	}

	server.Run()
}

//initConfig 解析命令行参数，初始化配置
func initConfig() {
	flag.IntVar(&config.Port, "p", 4444, "程序绑定的端口")
	flag.StringVar(&config.Path, "P", "/tmp", "Ebreader临时工作目录")
	flag.Parse()

	if flag.NArg() != 1 {
		showHelper()
	}
	config.File = flag.Arg(0)
}

//showHelper 显示帮助信息
func showHelper() {
	fmt.Println("用法：")
	fmt.Println("\t ebreader [-options=value] filename")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println()
	flag.PrintDefaults()
	os.Exit(2)
}
