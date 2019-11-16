package util

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func DownloadImg(imgUrl string) (name string,path string) {
	imgPath := "/Users/ZhuangXiaoDa/open-work/image/blog/"

	fileName := CreateRandomString(10)
	writePath := imgPath + fileName + ".jpg"


	res, err := http.Get(imgUrl)
	if err != nil {
		fmt.Println(err)
		fmt.Println("A error occurred!")
		return
	}
	defer res.Body.Close()
	// 获得get请求响应的reader对象
	reader := bufio.NewReaderSize(res.Body, 32 * 1024 * 10)


	file, err := os.Create(writePath)
	if err != nil {
		panic(err)
	}
	// 获得文件的writer对象
	writer := bufio.NewWriter(file)

	written, _ := io.Copy(writer, reader)
	fmt.Printf("%s.jpg Total length: %d write finish ! \n",fileName, written)

	return fileName,writePath

}


