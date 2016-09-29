package gohelper

import (
	"strings"
	"regexp"
	"errors"
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