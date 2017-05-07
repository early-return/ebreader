package config

import (
	"flag"
	"fmt"
	"os"
)

var (
	//File 要操作的文件名
	File string
	//Port 服务器绑定的端口
	Port int
	//Path 服务器监听路径（临时目录）
	Path string
)

func init() {
	initConfig()
	Path += "/ebreader"
}

//initConfig 解析命令行参数，初始化配置
func initConfig() {
	flag.IntVar(&Port, "p", 4444, "程序绑定的端口")
	flag.StringVar(&Path, "P", "/tmp", "Ebreader临时工作目录")
	flag.Parse()

	if flag.NArg() != 1 {
		showHelper()
	}
	File = flag.Arg(0)
}

//showHelper 显示帮助信息
func showHelper() {
	fmt.Println("EBReader用法：")
	fmt.Println("\t ebreader [-options=value] filename")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println()
	flag.PrintDefaults()
	os.Exit(2)
}
