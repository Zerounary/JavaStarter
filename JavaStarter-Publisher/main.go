package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {
	var ftpServer string
	var ftpPort string
	var filePath string
	var remoteFile string

	flag.StringVar(&ftpServer, "ip", "127.0.0.1", "FTP Server ip")
	flag.StringVar(&ftpPort, "port", "2121", "FTP Server Port")
	flag.StringVar(&filePath, "lf", "", "local file")
	flag.StringVar(&remoteFile, "rf", "./", "remote file")

	flag.Parse()

	ftpDial := ftpServer + ":" + ftpPort

	c, err := ftp.Dial(ftpDial, ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	err = c.Login("admin", "123456")
	if err != nil {
		log.Fatal(err)
	}

	// Do something with the FTP conn

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	lastIndex := strings.LastIndex(remoteFile, "/")
	c.MakeDir(remoteFile[:lastIndex])
	err = c.Stor(remoteFile, bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}

	fmt.Println("上传成功")
}
