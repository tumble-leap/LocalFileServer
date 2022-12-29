package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func runService(dir string) {
	localIpAddr := ""
	port := 8000
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				localIpAddr = ipnet.IP.String()
			}
		}
	}
	if localIpAddr == "" {
		log.Println("并没有获取到本机的ip地址呢，请手动查询～")
	} else {
		log.Println("本机ip：" + localIpAddr)
	}
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.StaticFS("/file", http.Dir(dir))
	t := time.NewTimer(time.Second * 2)
	go func() {
		for {
			err = r.Run(":" + strconv.Itoa(port))
			if err != nil {
				t.Stop()
				log.Println(err)
				fmt.Print("端口可能被占用，请输入新的端口号：")
				fmt.Scanf("%d", &port)
				t.Reset(time.Second)
			} else {
				t.Reset(time.Second)
				break
			}
		}
	}()

	<-t.C
	log.Println("服务启动成功")
	log.Println("请使用浏览器打开：http://" + localIpAddr + ":" + strconv.Itoa(port) + "/file/")
	select {}
}

func main() {
	var dir string
	currentDir, _ := os.Getwd()

	flag.StringVar(&dir, "t", currentDir, "The file server root directory, the default is the current directory")
	flag.Parse()
	_, err := os.ReadDir(dir)
	if err != nil {
		log.Println("文件夹不存在！")
		return
	}
	log.Println("当前文件服务器根目录为: " + dir)
	runService(dir)
}
