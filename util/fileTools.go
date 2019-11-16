package util

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Ioutil(name string) string {

	if contents,err := ioutil.ReadFile(name);err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		result := strings.Replace(string(contents),"\n","",1)
		return result
	}
	return ""
}


//使用ioutil.WriteFile方式写入文件,是将[]byte内容写入文件,如果content字符串中没有换行符的话，默认就不会有换行符
func WriteWithIoutil(name,content string) {
	data :=  []byte(content)
	if ioutil.WriteFile(name,data,0644) == nil {
		fmt.Println("写入文件成功:",content)
	}
}
