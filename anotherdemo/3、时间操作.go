package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now() // 获取当前时间
	fmt.Printf("current time :%v\n", now)
	year := now.Year()
	mouth := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year, mouth, day, hour, minute, second)
	timestampDemo2(now.Unix())
}

func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)

	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}
