package server

import (
	"ebreader/config"
	"fmt"
	"log"
	"net/http"
)

//Run 监听端口函数
func Run() {
	http.Handle("/", http.FileServer(http.Dir(config.Path)))
	address := fmt.Sprintf("0.0.0.0:%d", config.Port)
	log.Printf("正在开始监听%s 请用浏览器访问http://127.0.0.1:%d", address, config.Port)
	http.ListenAndServe(address, nil)
}
