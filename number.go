package gohelper

import (
	"strconv"
	"math/rand"
	"time"
)

// 获取随机小数 并以字符串方式返回
func GetRandFloat()string{
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fl := r.Float64()
	return strconv.FormatFloat(fl, 'g', -1, 64)
}

func RandomInt(end int ) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	in := r.Intn(end)
	return in
}