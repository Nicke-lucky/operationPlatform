package main

import (
	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
	"operationPlatform/config"
	"operationPlatform/db"
	"operationPlatform/router"
	"operationPlatform/utils"

	"time"
)

// 项目介绍注释
// @title 结算数据监控平台
// @version 1.0
// @description Gin swagger 结算数据监控平台
// @host 127.0.0.1:8088
func main() {
	conf := config.ConfigInit() //初始化配置
	log.Println("配置文件信息：", *conf)
	utils.InitLogrus(conf.LogPath, conf.LogFileName, time.Duration(24*conf.LogMaxAge)*time.Hour, time.Duration(conf.LogRotationTime)*time.Hour)

	//结算监控数据库 "root:Microvideo_1@tcp(122.51.24.189:3307)/blacklist?charset=utf8&parseTime=true&loc=Local"
	mstr := conf.MUserName + ":" + conf.MPass + "@tcp(" + conf.MHostname + ":" + conf.MPort + ")/" + conf.Mdatabasename + "?charset=utf8&parseTime=true&loc=Local"
	db.DBInit(mstr) //初始化数据库
	utils.Pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //最大活跃连接数  0为没有限制
		IdleTimeout: 300, //空闲连接超时时间
		//连接方法
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", conf.RedisAddr)
		},
	}
	defer utils.Pool.Close()
	utils.Redisdatabasename = conf.Redisdatabasename

	IpAddress := conf.IpAddress

	//goroutine1
	go db.HandleDayTasks()
	//goroutine2
	go db.HandleHourTasks()
	//goroutine3
	go db.HandleMinutesTasks()

	//http处理
	router.RouteInit(IpAddress)

}