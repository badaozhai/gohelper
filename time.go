package gohelper

import (
	"time"
	"fmt"
)

// 今天日期 2016-08-09 格式 字符串
func Today() string {
	now := time.Now() // 获取当前时间(日期)
	today := now.Format("2006-01-02")
	return today
}
// 今天之前几天日期 正数表示往后推几天,负数表示往前推几天,比如 -1 就是表示昨天日期
func TodayAgo(ago int)string{
	nTime := time.Now()
	yesTime := nTime.AddDate(0,0,ago)
	logDay := yesTime.Format("2006-01-02")
	return logDay
}
// 今年 字符串
func ThisYear()string{
	nTime := time.Now()
	y := nTime.Format("2006")
	return y
}

// 2016-09-01 这种形式的 转成 Mon Sep 19 2016 00:00:00 GMT+0800 (中国标准时间) 这种形式的
func FmtDateToCn(today string )string{
	layout := "2006-01-02"
	t, _ :=time.Parse(layout, today)
	todaycn := t.Format(`Mon Jan 02 2006 00:00:00 GMT+0800 (中国标准时间)`)
	return todaycn
}

// unix时间戳 字符串
func GetTimeStamString() string {
	now := time.Now()
	nanos := now.UnixNano()
	millis := nanos / 1000000
	return fmt.Sprintf("%d", millis)
}