package main

import (
	"fmt"
	"github.com/pkwenda/PictureBedTransfor/upload"
	"github.com/pkwenda/PictureBedTransfor/util"
	"regexp"
	"strings"
)


var sourceImageUrlArr []string



func main() {

	origin := util.Ioutil("/Users/ZhuangXiaoDa/open-work/blog/source/_posts/compare-minio-with-seaweedfs.md")

	util.WriteWithIoutil("/Users/ZhuangXiaoDa/open-work/blog/source/_posts/compare-minio-with-seaweedfs-new.md",
		analysisPictureReplace(origin))
}



// target is directory path or file path
func analysisPictureReplace(origin string) string {
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
		filename, filePath := util.DownloadImg(url)
		uploadedPath := upload.MinioUpload(filename,filePath)
		origin = strings.Replace(origin,url,uploadedPath,-1)
	}
	return origin

}

