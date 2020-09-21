package db

import (
	log "github.com/sirupsen/logrus"

	"operationPlatform/utils"

	"time"
)

//goroutine1
//1定时任务 一天一次的
func HandleDayTasks() {
	for {
		now := time.Now()               //获取当前时间，放到now里面，要给next用
		next := now.Add(time.Hour * 24) //通过now偏移24小时

		next = time.Date(next.Year(), next.Month(), next.Day(), 3, 0, 0, 0, next.Location()) //获取下一个凌晨的日期

		t := time.NewTimer(next.Sub(now)) //计算当前时间到凌晨的时间间隔，设置一个定时器
		<-t.C
		log.Println("执行线程1，处理一天一次的定时任务【完成】11111111111111111111111111111111111111111111111111111111111111111")
	}
}

//goroutine2
//2定时任务 按小时的
func HandleHourTasks() {
	tiker := time.NewTicker(time.Minute * 60) //每15秒执行一下
	for {
		log.Println("执行线程2，处理按小时的定时任务222222222222222222222222222222222222222222222222")
		//任务一

		log.Println(utils.DateTimeFormat(<-tiker.C), "执行线程2，处理按小时的定时任务【完成】222222222222222222222222222222222222222222222222")

	}

}

//goroutine3
//3定时任务 按分钟的
func HandleMinutesTasks() {
	tiker := time.NewTicker(time.Second * 10) //每15秒执行一下
	for {
		log.Println("执行线程3，处理按分钟的定时任务333333333333333333333333333333333333333333333333333333333333333333")

		log.Println(utils.DateTimeFormat(<-tiker.C), "执行线程3，处理按分钟的定时任务【完成】333333333333333333333333333333333333333333333333333333333333333333")

	}

}
