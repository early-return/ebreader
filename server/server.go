package server

import (
	"ebreader/config"
	"fmt"
	"net/http"
)

//Run 监听端口函数
func Run() error {
	http.Handle("/", http.FileServer(http.Dir(config.Path)))
	address := fmt.Sprintf("0.0.0.0:%d", config.Port)
	fmt.Printf("正在开始监听%s 请用浏览器访问http://127.0.0.1:%d\n", address, config.Port)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}
	return nil
}
