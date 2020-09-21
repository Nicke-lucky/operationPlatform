package utils

import (
	"log"
	"time"
)

func DateTimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05") //后面的参数是固定的 否则将无输出
}

func DateFormatTimeToTime(data time.Time) time.Time {
	datestr := data.Format("2006-01-02 00:00:00")
	const Layout = "2006-01-02 15:04:05" //时间常量
	loc, _ := time.LoadLocation("Asia/Shanghai")
	time1, _ := time.ParseInLocation(Layout, datestr /*需要转换的时间类型字符串*/, loc)
	return time1
}

func DateFormatTimeTostrdate(data time.Time) string {
	datestr := data.Format("2006-01-02 00:00:00")
	b := []byte(datestr)
	return string(b[0:10])
}

//2006-01-02 15:04:05
func DateTimeNowFormat() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

//2006-01-02 15:04:05
func KuaizhaoTimeNowFormat() string {
	return time.Now().Format("2006-01-02 15:04:00")
}

func DateNowFormat() string {
	return time.Now().Format("2006-01-02")
}

//处理时间字符串转时间
func StrTimeTotime(strTime string) time.Time {
	const Layout = "2006-01-02 15:04:05" //时间常量
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tim, _ := time.ParseInLocation(Layout, strTime /*需要转换的时间类型字符串*/, loc)
	log.Println(tim)
	return tim
}

//处理时间字符串转时间
func StrTimeToNowtime() time.Time {
	strTime := time.Now().Format("2006-01-02 15:04:05")
	const Layout = "2006-01-02 15:04:05" //时间常量
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tim, _ := time.ParseInLocation(Layout, strTime /*需要转换的时间类型字符串*/, loc)
	log.Println(tim)
	return tim
}

//处理时间字符串转时间
func DateToNowdate() time.Time {
	strTime := time.Now().Format("2006-01-02")
	const Layout = "2006-01-02" //时间常量
	loc, _ := time.LoadLocation("Asia/Shanghai")

	tim, _ := time.ParseInLocation(Layout, strTime /*需要转换的时间类型字符串*/, loc)
	log.Println(tim)
	return tim
}

//处理时间字符串转时间
func StrTimeTodefaultdate() time.Time {
	strTime := "2020-01-01 00:00:00"
	const Layout = "2006-01-02 15:04:05" //时间常量
	loc, _ := time.LoadLocation("Asia/Shanghai")

	tim, _ := time.ParseInLocation(Layout, strTime /*需要转换的时间类型字符串*/, loc)
	log.Println(tim)
	return tim
}

//处理时间字符串转时间
func StrTimeTodefaultdatetimestr() string {
	return "2020-01-01 00:00:00"
}

//获取昨天的日期
func Yesterdaydate() string {
	nTime := time.Now()
	yesTime := nTime.AddDate(0, 0, -1)
	return yesTime.Format("2006-01-02")
}

//currentTime := time.Now()
//oldTime := currentTime.AddDate(0, 0, -2)
//获取前num天的日期
func OldData(num int) []string {
	days := num
	nTime := time.Now()
	//switch num {
	//case 1:
	//	days = -1
	//case 7:
	//	days = -7
	//case 14:
	//	days = -14
	//case 30:
	//	days = -7
	//default: //default case
	//	log.Println("  number  error")
	//	return nil
	//}
	daystrs := make([]string, 0)
	for i := 0; i < num; i++ {
		yesTime := nTime.AddDate(0, 0, -days)
		daystrs = append(daystrs, yesTime.Format("2006-01-02"))
		days--
	}
	//log.Println("daystrs:", daystrs)
	return daystrs
}
