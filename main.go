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

func GetInternalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", errors.New("internal IP fetch failed, detail:" + err.Error())
	}
	defer conn.Close()
	res := conn.LocalAddr().String()
	res = strings.Split(res, ":")[0]
	return res, nil
}

func runService(dir string) {
	localIpAddr, err := GetInternalIP()
	if err != nil {
		log.Println(err)
		log.Println("并没有获取到本机的ip地址呢，请手动查询～")
	} else {
		log.Println("本机ip：" + localIpAddr)
	}
	port := 8000

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
