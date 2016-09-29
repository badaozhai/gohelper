package gohelper

import (
	"fmt"
	"testing"
)

func TestToday(t *testing.T) {
	fmt.Println(Today())
	// output:
	// 2016-09-29
}

func TestThisYear(t *testing.T) {
	fmt.Println(ThisYear())
	// output:
	// 2016-09
}

func TestFmtDateToCn(t *testing.T) {
	str := Today()
	fmt.Println(FmtDateToCn(str))
	//Thu Sep 29 2016 00:00:00 GMT+0800 (中国标准时间)
}