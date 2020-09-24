package db

import (
	"bytes"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"operationPlatform/dto"
	"operationPlatform/types"
	"operationPlatform/utils"
	"strings"

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
//3定时任务 按秒的
func HandleMinutesTasks() {
	tiker := time.NewTicker(time.Second * 10) //每15秒执行一下
	for {
		log.Println("执行线程3，处理按分钟的定时任务333333333333333333333333333333333333333333333333333333333333333333")
		//任务一
		//获取网关列表数据
		log.Println(utils.DateTimeFormat(<-tiker.C), "执行线程3，处理按分钟的定时任务【完成】333333333333333333333333333333333333333333333333333333333333333333")

	}

}

//goroutine3
//3定时任务 按秒的
func HandleSecondTasks() {
	tiker := time.NewTicker(time.Second * 10) //每15秒执行一下
	for {
		log.Println("执行线程3，处理按分钟的定时任务333333333333333333333333333333333333333333333333333333333333333333")
		//任务一
		//获取网关列表数据
		log.Println(utils.DateTimeFormat(<-tiker.C), "执行线程3，处理按分钟的定时任务【完成】333333333333333333333333333333333333333333333333333333333333333333")

	}

}

var Errormsg_address string
var Gwmsg_address string
var Metric_address string

//任务一
//获取网关列表数据
func GatewayDataUpdate() {

	//1、获取网关基本信息
	gwmsgs := GatewayDataPostWithJson()

	for _, gwmsg := range (*gwmsgs).Date {
		gwxx := new(types.BDmWanggjcxx)
		//数据赋值
		//gwxx.FVcWanggbh = gwmsg.MsgHead.TerminalId //	 '网关编号',
		gwxx.FVcGongsID = gwmsg.MsgHead.CompanyId //	`F_VC_GONGSID` varchar(32) NOT NULL COMMENT '公司ID',
		//gwxx.FVcTingccbh = gwmsg.MsgHead.Parkid   //	`F_VC_TINGCCBH` varchar(32) NOT NULL COMMENT '停车场编号',

		stamp1 := utils.StrTimeToTimestamp(gwmsg.UpdateTime)
		stamp2 := utils.GetTimestamp()
		if (stamp2 - stamp1) > 300 {
			gwxx.FNbZhuangt = 0 //	`F_NB_ZHUANGT` int(11) NOT NULL DEFAULT '0' COMMENT '状态 0：离线、1：在线',[通过最新存储时间判断]
		} else {
			gwxx.FNbZhuangt = 1 //	`F_NB_ZHUANGT` int(11) NOT NULL DEFAULT '0' COMMENT '状态 0：离线、1：在线',[通过最新存储时间判断]

		}

		if !utils.StringExist(gwmsg.Gatewayip, ",") {
			gwxx.FVcIpdz = gwmsg.Gatewayip //	`F_VC_IPDZ` varchar(32) DEFAULT NULL COMMENT 'IP地址',
		} else {
			gwip := strings.Split(gwmsg.Gatewayip, ",")
			gwxx.FVcIpdz = gwip[0] //	`F_VC_IPDZ` varchar(32) DEFAULT NULL COMMENT 'IP地址',
		}

		//gwxx.FNbCPUsyl = //	`F_NB_CPUSYL` decimal(32, 10) DEFAULT NULL COMMENT 'CPU使用率 百分比',
		//	gwxx.FNbNeicsyl = //	`F_NB_NEICSYL` decimal(32, 10) DEFAULT NULL COMMENT '内存使用率 百分比',
		//gwxx.FNbYsyncdx = //	`F_NB_YISYNCDX` decimal(32, 10) DEFAULT NULL COMMENT '已使用内存大小 单位：MB',
		//	gwxx.FNbZongncdx = //	`F_NB_ZONGNCDX` decimal(32, 10) DEFAULT NULL COMMENT '总内存大小 单位：MB',
		//gwxx.FNbYingpsyl = //	`F_NB_YINGPSYL` decimal(32, 10) DEFAULT NULL COMMENT '硬盘使用率 百分比',
		//	gwxx.FNbYisyypdx = //	`F_NB_YISYYPDX` decimal(32, 10) DEFAULT NULL COMMENT '已使用硬盘大小 单位：GB',
		//gwxx.FNbZongypdx = //	`F_NB_ZONGYPDX` decimal(32, 10) DEFAULT NULL COMMENT '总硬盘大小 单位：GB',

		//	gwxx.FNbGaojzs = //	`F_NB_GAOJZS` int(11) NOT NULL DEFAULT '0' COMMENT '告警总数',
		//gwxx.FNbWeiclgjs = //	`F_NB_WEICLGJS` int(11) NOT NULL DEFAULT '0' COMMENT '未处理告警数',

		//	gwxx.FNbChongqcs = //	`F_NB_CHONGQCS` int(11) NOT NULL DEFAULT '0' COMMENT '重启次数',
		gwxx.FVcDangqbbh = gwmsg.GetwayVersion //	`F_VC_DANGQBBH` varchar(512) DEFAULT NULL COMMENT '当前版本号',

		gwxx.FDtZuijgxbbsj = utils.StrTimeTotime(gwmsg.LastversionUpdatedatetime) //	`F_DT_ZUIJGXBBSJ` datetime DEFAULT NULL COMMENT '最近更新版本时间',
		AntennaInfosNum := len(gwmsg.AntennaInfos)
		gwxx.FNbTianxsl = AntennaInfosNum //	`F_NB_TIANXSL` int(11) DEFAULT NULL COMMENT '天线数量',

		//	gwxx.FNbWanglyc = //	`F_NB_WANGLYC` bigint(20) DEFAULT NULL COMMENT '网络延迟 单位：ms',

		gwxx.FDtZuihgxsj = time.Now()

		uperr := UpdateGatewaydata(gwmsg.MsgHead.TerminalId, gwxx)
		if uperr != nil {
			log.Println("更新网关信息error")
		}
	}

	//2、获取网关使用信息
	//ErrorDataPostWithJson()

}

//网关基础信息查询接口 一个dataserver的方法 ，再加 后面需要多个dataserver的
func GatewayDataPostWithJson() *dto.GatewayDeviceMsgResp {
	//post请求提交json数据
	gw := dto.GatewayDataReq{}
	ba, _ := json.Marshal(gw)
	//POST "http://172.18.70.22:8080/etcpark/dataserver/gateway/list"
	log.Println("Gwmsg_address:", Gwmsg_address)
	resp, _ := http.Post(Gwmsg_address, "application/json", bytes.NewBuffer([]byte(ba)))
	if resp.Body == nil {
		log.Println("resp.Body==nil")
	}

	body, _ := ioutil.ReadAll(resp.Body)

	Resp := new(dto.GatewayDeviceMsgResp)
	unmerr := json.Unmarshal(body, Resp)
	if unmerr != nil {
		log.Println("json.Unmarshal error")
	}
	log.Printf("Post request with json result:%v\n", Resp)
	return Resp
}

//告警信息查询接口
func ErrorDataPostWithJson(beginTime int64, endTime int64) *dto.ErrorMsgResp {
	//post请求提交json数据 时间戳
	errdatareq := dto.QueryErrorMsgQeq{BeginTime: beginTime, EndTime: endTime}
	ba, _ := json.Marshal(errdatareq)
	//POST
	log.Println("Errormsg_address:", Errormsg_address)
	resp, _ := http.Post(Errormsg_address, "application/json", bytes.NewBuffer([]byte(ba)))
	body, _ := ioutil.ReadAll(resp.Body)
	Resp := new(dto.ErrorMsgResp)
	unmerr := json.Unmarshal(body, Resp)
	if unmerr != nil {
		log.Println("json.Unmarshal error", unmerr)
	}
	log.Printf("Post request with json result:%v\n", Resp)
	return Resp
}

//

//指标信息查询接口
func MetricDataPostWithJson(metric string) *dto.MetricMsgResp {
	//post请求提交json数据 时间戳
	metricreq := dto.QueryMetricMsgQeq{Metric: metric}
	ba, _ := json.Marshal(metricreq)
	//POST
	log.Println("Metric_address:", Metric_address)
	resp, _ := http.Post(Metric_address, "application/json", bytes.NewBuffer([]byte(ba)))
	body, _ := ioutil.ReadAll(resp.Body)
	Resp := new(dto.MetricMsgResp)
	unmerr := json.Unmarshal(body, &Resp)
	if unmerr != nil {
		log.Println("json.Unmarshal error", unmerr)
	}
	log.Printf("Post request with json result:%v\n", Resp)
	return Resp
}
