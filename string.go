package gohelper

import (
	"strings"
	"regexp"
	"errors"
	"os"
	"fmt"
	"bufio"
	"encoding/base64"
)
// 从目标字符串中截取字符串
// 用法 见 test
func GetValueBySplit(str ,start ,end string)(string,error){
	arr1 := strings.Split(str,start)
	if len(arr1)>1{
		str2 := arr1[1]
		arr2 := strings.Split(str2,end);
		if len(arr2)>1{
			return arr2[0],nil
		}else{
			return "",errors.New("字符串:"+str+"中没有"+end)
		}
	}else {
		return "",errors.New("字符串:"+str+"中没有"+start)
	}

}
func GetValueByRegexp(str ,start ,end string)(string){
	reg, err := regexp.Compile(start+`.*?`+end)
	if err != nil {
		return ""
	} else {
		src := reg.FindString(str)
		if src !=""{
			src = strings.Replace(src,start,"",-1)
			src = strings.Replace(src,end,"",-1)
			return src
		}
		return src
	}
}
//截取字符串 start 起点下标 length 需要截取的长度
func Substr(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

// 图片文件转 base65string 字符串
func Image2Base64String(filePath string) string {
	imgFile, err := os.Open(filePath) // a QR code image
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer imgFile.Close()
	// create a new buffer base on file size
	fInfo, _ := imgFile.Stat()
	var size int64 = fInfo.Size()
	buf := make([]byte, size)
	// read file content into buffer
	fReader := bufio.NewReader(imgFile)
	fReader.Read(buf)
	// if you create a new image instead of loading from file, encode the image to buffer instead with png.Encode()
	// png.Encode(&buf, image)
	// convert the buffer bytes to base64 string - use buf.Bytes() for new image
	imgBase64Str := base64.StdEncoding.EncodeToString(buf)
	return imgBase64Str
}