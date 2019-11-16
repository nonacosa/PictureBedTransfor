package main

import (
	"../util"
	"fmt"
	"github.com/minio/minio-go/v6"
	"log"
	"regexp"
	"strings"
)

//type Callback func(x, y int) int

var sourceImageUrlArr []string

//func upload(bedType string){
//	switch bedType {
//	case "minio":
//		minio()
//	}
//}
//
//
//func minio()  {
//
//}

func main() {

	endpoint := "www.gitrue.com:9000"
	accessKeyID := "pkwenda"
	secretAccessKey := "886pkxiaojiba"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now setup

	//
	//origin := util.Ioutil("/Users/ZhuangXiaoDa/open-work/blog/source/_posts/website-architecture-separate.md")
	//analysisPictureReplace(origin)


	//util.WriteWithIoutil("testtest",origin)
}



// target is directory path or file path
func analysisPictureReplace(origin string)  {
	fmt.Print(origin)
	//if ok, _ := regexp.Match("^[0-9a-zA-Z_]+$", []byte("hello")); ok {
	//	fmt.Println("ok");
	//}

	re3, _ := regexp.Compile(`https?://([\w-]+\.)+[\w-]+(/[\w-./?%&=]*)?`)

	for _,url := range re3.FindAllString(origin, -1) {
		if util.UrlIsIamge(url) {
			url = strings.Replace(strings.Replace(url, "[", "", 1 ),"]","",1)
			sourceImageUrlArr = append(sourceImageUrlArr, url)

		}
	}

	for _,url := range sourceImageUrlArr {
		util.DownloadImg(url)
	}
	rep := re3.ReplaceAllStringFunc(origin, xxx)
	fmt.Println(rep)

}

func xxx(ss string) string {
	return "dsadasdasdsada"
}
