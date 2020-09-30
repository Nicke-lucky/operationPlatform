package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"operationPlatform/dto"
	"operationPlatform/types"
	"operationPlatform/utils"
	"strconv"
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
//3定时任务 按分钟的
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
//4定时任务 按秒的
func HandleSecondTasks() {
	tiker := time.NewTicker(time.Second * 120) //每15秒执行一下
	for {
		log.Println("执行线程3，处理按分钟的定时任务44444444444444444444444444444444444444444444444444444444444444444444444444444444444444444444")
		//任务一
		//获取网关列表数据,并更新数据
		gwuperr := GatewayDataUpdate()
		if gwuperr != nil {
			log.Println("++++++++++++++++++++++++++【任务一 有错误 获取网关列表数据,并更新数据】+++++++++++++++++++++")
		}

		//任务二
		//获取指标列表数据,并更新数据
		gwMetricerr := GatewayMetricDataUpdate()
		if gwMetricerr != nil {
			log.Println("++++++++++++++++++++++++++【任务二 有错误 获取指标列表数据,并更新数据】+++++++++++++++++++++")
		}

		//任务三
		//获取告警数据,并更新数据
		Alarmerr := GatewayAlarmDataUpdate()
		if Alarmerr != nil {
			log.Println("++++++++++++++++++++++++++【 任务三 有错误 获取告警数据,并更新数据 】+++++++++++++++++++++")
		}

		//
		log.Println(utils.DateTimeFormat(<-tiker.C), "执行线程3，处理按分钟的定时任务【完成】44444444444444444444444444444444444444444444444444444444444444444444444444444")

	}

}

var Errormsg_address string
var Gwmsg_address string
var Metric_address string

//任务一
//获取网关列表数据
func GatewayDataUpdate() error {

	//1、获取网关基本信息
	gwmsgs, err := GatewayDataPostWithJson()
	if err != nil {
		log.Println("获取网关基本信息失败", err)
		return err
	}
	//网关基本信息更新
	for _, gwmsg := range (*gwmsgs).Date {
		gwxx := new(types.BDmWanggjcxx)
		//数据赋值
		//如果网关编号存在就更新，如果不存在就插入
		qerr, gwd1 := QueryGatewaydata(gwmsg.MsgHead.TerminalId)
		if qerr != nil {
			//不存在就插入
			if fmt.Sprint(qerr) == "record not found" {
				log.Println("Queryerr== `record not found`:", err)
				log.Println("新网关", gwd1, "需要插入数据库")
				//
				gwxx1 := new(types.BDmWanggjcxx)
				gwxx1.FVcWanggbh = gwmsg.MsgHead.TerminalId //	 '网关编号',
				gwxx1.FVcGongsID = gwmsg.MsgHead.CompanyId  //	'公司ID',
				gwxx1.FVcTingccbh = gwmsg.MsgHead.Parkid    //'停车场编号',
				gwxx1.FDtZuijgxbbsj = time.Now()
				gwxx1.FDtChuangjsj = time.Now()
				gwxx1.FDtZuihgxsj = time.Now()
				gwxx1.FNbZhuangt = 1
				if !utils.StringExist(gwmsg.Gatewayip, ",") {
					gwxx.FVcIpdz = gwmsg.Gatewayip //	 'IP地址',
				} else {
					gwip := strings.Split(gwmsg.Gatewayip, ",")
					gwxx.FVcIpdz = gwip[0] //	 'IP地址',
				}

				gwxx.FVcDangqbbh = gwmsg.GetwayVersion                                    //'当前版本号',
				gwxx.FDtZuijgxbbsj = utils.StrTimeTotime(gwmsg.LastversionUpdatedatetime) // '最近更新版本时间',

				//插入新网关
				inerr := GatewayInsert(gwxx1)
				if inerr != nil {
					log.Println("++++++++++++++++++++++++++++++++++++++++插入新网关失败")
				}

				continue
			} else {
				log.Println("++++++++++++++++++++++++++++++++++++++++查询网关是否已经存在时，查询失败")
			}
		}

		//存在就更新记录
		gwxx.FVcWanggbh = gwmsg.MsgHead.TerminalId //'网关编号',
		gwxx.FVcGongsID = gwmsg.MsgHead.CompanyId  //'公司ID',
		gwxx.FVcTingccbh = gwmsg.MsgHead.Parkid    //'停车场编号',

		//判断是否在线，获取更新时间与现在的时间差大于5分钟就离线
		stamp1 := utils.StrTimeToTimestamp(gwmsg.UpdateTime) //
		stamp2 := utils.GetTimestamp()
		if (stamp2 - stamp1) > 300 {
			gwxx.FNbZhuangt = 0 //	'状态 0：离线、1：在线',[通过最新存储时间判断]
		} else {
			gwxx.FNbZhuangt = 1 //	 '状态 0：离线、1：在线',[通过最新存储时间判断]

		}

		if !utils.StringExist(gwmsg.Gatewayip, ",") {
			gwxx.FVcIpdz = gwmsg.Gatewayip //'IP地址',
		} else {
			gwip := strings.Split(gwmsg.Gatewayip, ",")
			gwxx.FVcIpdz = gwip[0] //'IP地址',
		}

		//gwxx.FNbYsyncdx = //	`F_NB_YISYNCDX` decimal(32, 10) DEFAULT NULL COMMENT '已使用内存大小 单位：MB',
		//gwxx.FNbZongncdx = //	`F_NB_ZONGNCDX` decimal(32, 10) DEFAULT NULL COMMENT '总内存大小 单位：MB',
		//gwxx.FNbYingpsyl = //	`F_NB_YINGPSYL` decimal(32, 10) DEFAULT NULL COMMENT '硬盘使用率 百分比',
		//gwxx.FNbYisyypdx = //	`F_NB_YISYYPDX` decimal(32, 10) DEFAULT NULL COMMENT '已使用硬盘大小 单位：GB',
		//gwxx.FNbZongypdx = //	`F_NB_ZONGYPDX` decimal(32, 10) DEFAULT NULL COMMENT '总硬盘大小 单位：GB',

		//告警总数是从告警列表中获取的
		errornum, qEerrorerr := QueryErrordata(gwmsg.MsgHead.TerminalId)
		if qEerrorerr != nil {
			log.Println("++++++++++++++++++++++++++++++++++++++++查询网关告警总数失败")
			gwxx.FNbGaojzs = int(errornum) // '告警总数',
		} else {
			gwxx.FNbGaojzs = int(errornum) // '告警总数',
		}

		//查询网关未处理告警总数
		unnum, qunerr := QueryUndisposedError(gwmsg.MsgHead.TerminalId)
		if qunerr != nil {
			log.Println("++++++++++++++++++++++++++++++++++++++++查询网关未处理告警总数失败")
			gwxx.FNbWeiclgjs = int(unnum) // '未处理告警数',
		} else {
			gwxx.FNbWeiclgjs = int(unnum) // '未处理告警数','
		}

		qRestarerr, RestartCount := QueryRestartCount(gwmsg.MsgHead.TerminalId)
		if qRestarerr != nil {
			log.Println("++++++++++++++++++++++++++++++++++++++++查询网关重启总数失败")
		}
		gwxx.FNbChongqcs = int(RestartCount)   //	`F_NB_CHONGQCS` int(11) NOT NULL DEFAULT '0' COMMENT '重启次数',
		gwxx.FVcDangqbbh = gwmsg.GetwayVersion //	`F_VC_DANGQBBH` varchar(512) DEFAULT NULL COMMENT '当前版本号',

		gwxx.FDtZuijgxbbsj = utils.StrTimeTotime(gwmsg.LastversionUpdatedatetime) // '最近更新版本时间',
		AntennaInfosNum := len(gwmsg.AntennaInfos)
		gwxx.FNbTianxsl = AntennaInfosNum //	'天线数量',

		yc, _ := strconv.Atoi(gwmsg.NetWorkDelay)
		gwxx.FNbWanglyc = yc          //'网络延迟 单位：ms',
		gwxx.FDtZuihgxsj = time.Now() //最后更新数据时间'
		//更新网关基本信息
		uperr := UpdateGatewaydata(gwmsg.MsgHead.TerminalId, gwxx)
		if uperr != nil {
			log.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++更新网关信息失败", time.Now())
		}
	}

	//2、获取网关使用信息
	//2.1告警信息
	//ErrorDataPostWithJson()
	return nil

}

//任务二
//获取指标列表数据,并更新数据
func GatewayMetricDataUpdate() error {

	//1、CPU使用率
	CPUmetric := "gateway.park.gateway.cpupercent"
	if cpudata, cpuerr := MetricDataPostWithJson(CPUmetric); cpuerr != nil {
		log.Println("查询CPU指标出错:", cpuerr)
	} else {
		//已获取数据
		if cpudata != nil {
			//把指标结果存数据库
			for _, cpu := range cpudata.MetricMsgDate.Date {
				//cpu.Time//采集时间
				//cpu.Endpoint//设备id
				//cpu.Value//指标值
				//
				log.Println("cpu.Time", cpu.Time, "cpu.Endpoint", cpu.Endpoint, "cpu.Value", cpu.Value)
				qgwerr, gwdata := QueryGatewaydata(cpu.Endpoint)
				if qgwerr != nil {
					if fmt.Sprint(qgwerr) == "record not found" {
						log.Println("  err== `record not found`:", qgwerr)
						//没有找到，说明该cpu信息还没有在数据库中有对应的设备记录
						log.Println(qgwerr, "没有找到，说明该cpu信息还没有在数据库中有对应的设备记录+++++++++++++++++")
						continue
					} else {
						log.Println("+++++++++++++++++++++++++++++=err==:", qgwerr)
						continue
					}
				}
				//更新cpu信息
				gwxx := new(types.BDmWanggjcxx)

				value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", cpu.Value), 64)
				gwxx.FNbCPUsyl = value //cpu使用率
				gwxx.FDtZuihgxsj = utils.StrTimeTotime(cpu.Time)
				upcpuerr := UpdateGatewaydata(gwdata.FVcWanggbh, gwxx)
				if upcpuerr != nil {
					log.Println(upcpuerr, "更新cpu信息失败+++++++++++++++++")
					continue
				}

			}

		} else {
			log.Println("查询CPU指标cpudata为空:", cpudata)
		}

	}

	//2、内存使用率
	MeMmetric := "gateway.park.gateway.mempercent"
	if MeMdata, MeMerr := MetricDataPostWithJson(MeMmetric); MeMerr != nil {
		log.Println("查询MeM指标出错:", MeMerr)
	} else {
		//已获取数据
		if MeMdata != nil {
			//把指标结果存数据库
			for _, MeM := range MeMdata.MetricMsgDate.Date {
				//MeM.Time//采集时间
				//MeM.Endpoint//设备id
				//MeM.Value//指标值
				//
				log.Println("MeM.Time", MeM.Time, "MeM.Endpoint", MeM.Endpoint, "MeM.Value", MeM.Value)
				qgwerr, memdata := QueryGatewaydata(MeM.Endpoint)
				if qgwerr != nil {
					if fmt.Sprint(qgwerr) == "record not found" {
						log.Println("err== `record not found`:", qgwerr)
						//没有找到，说明该MeM信息还没有在数据库中有对应的设备记录
						log.Println(qgwerr, "没有找到，说明该MeM信息还没有在数据库中有对应的设备记录+++++++++++++++++")
						continue
					} else {
						log.Println("+++++++++++++++++++++++++++++=err==:", qgwerr)
						continue
					}
				}
				//更新MeM信息
				gwxx := new(types.BDmWanggjcxx)
				//
				value, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", MeM.Value), 64)
				gwxx.FNbNeicsyl = value //内存使用率
				gwxx.FDtZuihgxsj = utils.StrTimeTotime(MeM.Time)
				upMeMerr := UpdateGatewaydata(memdata.FVcWanggbh, gwxx)
				if upMeMerr != nil {
					log.Println(upMeMerr, "更新MeM信息失败+++++++++++++++++")
					continue
				}

			}

		} else {
			log.Println("查询MeM指标MeMdata为空:", MeMdata)
		}
	}

	//3、内存已使用
	MeMYSYmetric := "mem.bytes.used"
	if MeMYSYdata, MeMYSYerr := MetricDataPostWithJson(MeMYSYmetric); MeMYSYerr != nil {
		log.Println("查询MeMYSY指标出错:", MeMYSYerr)
	} else {
		//已获取数据
		if MeMYSYdata != nil {
			//把指标结果存数据库
			for _, MeMYSY := range MeMYSYdata.MetricMsgDate.Date {
				//MeMYSY.Time//采集时间
				//MeMYSY.Endpoint//设备id
				//MeMYSY.Value//指标值
				//
				log.Println("MeMYSY.Time", MeMYSY.Time, "MeMYSY.Endpoint", MeMYSY.Endpoint, "MeMYSY.Value", MeMYSY.Value)
				qgwerr, MeMYSYdata := QueryGatewaydata(MeMYSY.Endpoint)
				if qgwerr != nil {
					if fmt.Sprint(qgwerr) == "record not found" {
						log.Println("err== `record not found`:", qgwerr)
						//没有找到，说明该MeMYSY信息还没有在数据库中有对应的设备记录
						log.Println(qgwerr, "没有找到，说明该MeMYSY信息还没有在数据库中有对应的设备记录+++++++++++++++++")
						continue
					} else {
						log.Println("+++++++++++++++++++++++++++++=err==:", qgwerr)
						continue
					}
				}
				//更新MeMYSY信息
				gwxx := new(types.BDmWanggjcxx)
				//内存已使用

				gwxx.FNbYsyncdx = utils.ByteToMB(MeMYSY.Value)
				gwxx.FDtZuihgxsj = utils.StrTimeTotime(MeMYSY.Time)
				upMeMYSYerr := UpdateGatewaydata(MeMYSYdata.FVcWanggbh, gwxx)
				if upMeMYSYerr != nil {
					log.Println(upMeMYSYerr, "更新MeMYSY信息失败+++++++++++++++++")
					continue
				}
			}
		} else {
			log.Println("查询MeMYSY指标MeMYSYdata为空:", MeMYSYdata)
		}
	}

	//4、内存总大小
	MeMZDXmetric := "mem.bytes.total"
	if MeMZDXdata, MeMZDXerr := MetricDataPostWithJson(MeMZDXmetric); MeMZDXerr != nil {
		log.Println("查询MeMYSY指标出错:", MeMZDXerr)
	} else {
		//已获取数据
		if MeMZDXdata != nil {
			//把指标结果存数据库
			for _, MeMZDX := range MeMZDXdata.MetricMsgDate.Date {
				//MeMZDX.Time//采集时间
				//MeMZDX.Endpoint//设备id
				//MeMZDX.Value//指标值
				//
				log.Println("MeMZDX.Time", MeMZDX.Time, "MeMZDX.Endpoint", MeMZDX.Endpoint, "MeMZDX.Value", MeMZDX.Value)
				qgwerr, MeMZDXdata := QueryGatewaydata(MeMZDX.Endpoint)
				if qgwerr != nil {
					if fmt.Sprint(qgwerr) == "record not found" {
						log.Println("err== `record not found`:", qgwerr)
						//没有找到，说明该MeMZDX信息还没有在数据库中有对应的设备记录
						log.Println(qgwerr, "没有找到，说明该MeMZDX信息还没有在数据库中有对应的设备记录+++++++++++++++++")
						continue
					} else {
						log.Println("+++++++++++++++++++++++++++++=err==:", qgwerr)
						continue
					}
				}
				//更新MeMZDX信息
				gwxx := new(types.BDmWanggjcxx)
				//
				gwxx.FNbZongncdx = utils.ByteToMB(MeMZDX.Value)
				gwxx.FDtZuihgxsj = utils.StrTimeTotime(MeMZDX.Time)
				upMeMZDXerr := UpdateGatewaydata(MeMZDXdata.FVcWanggbh, gwxx)
				if upMeMZDXerr != nil {
					log.Println(upMeMZDXerr, "更新MeMZDX信息失败+++++++++++++++++")
					continue
				}

			}

		} else {
			log.Println("查询MeMZDX指标MeMZDXdata为空:", MeMZDXdata)
		}
	}

	//5、磁盘使用率
	Diskmetric := "disk.cap.bytes.used.percent"
	if Diskdata, Diskerr := MetricDataPostWithJson(Diskmetric); Diskerr != nil {
		log.Println("查询Disk指标出错:", Diskerr)
	} else {
		//已获取数据
		if Diskdata != nil {
			//把指标结果存数据库
			for _, Disk := range Diskdata.MetricMsgDate.Date {
				//Disk.Time//采集时间
				//Disk.Endpoint//设备id
				//Disk.Value//指标值
				//
				log.Println("Disk.Time", Disk.Time, "Disk.Endpoint", Disk.Endpoint, "Disk.Value", Disk.Value)
				Diskerr, Diskdata := QueryGatewaydata(Disk.Endpoint)
				if Diskerr != nil {
					if fmt.Sprint(Diskerr) == "record not found" {
						log.Println("err== `record not found`:", Diskerr)
						//没有找到，说明该Disk信息还没有在数据库中有对应的设备记录
						log.Println(Diskerr, "没有找到，说明该Disk信息还没有在数据库中有对应的设备记录+++++++++++++++++")
						continue
					} else {
						log.Println("+++++++++++++++++++++++++++++=err==:", Diskerr)
						continue
					}
				}
				//更新Disk信息
				gwxx := new(types.BDmWanggjcxx)
				//utils.ByteToGB(Disk.Value)

				gwxx.FNbYingpsyl = Disk.Value
				gwxx.FDtZuihgxsj = utils.StrTimeTotime(Disk.Time)
				upDiskerr := UpdateGatewaydata(Diskdata.FVcWanggbh, gwxx)
				if upDiskerr != nil {
					log.Println(upDiskerr, "更新Disk信息失败+++++++++++++++++")
					continue
				}

			}

		} else {
			log.Println("查询Disk指标Diskdata为空:", Diskdata)
		}
	}

	//6、已使用磁盘大小
	DiskSYDXmetric := "disk.cap.bytes.used"
	if DiskSYDXdata, DiskSYDXerr := MetricDataPostWithJson(DiskSYDXmetric); DiskSYDXerr != nil {
		log.Println("查询DiskSYDX指标出错:", DiskSYDXerr)
	} else {
		//已获取数据
		if DiskSYDXdata != nil {
			//把指标结果存数据库
			for _, DiskSYDX := range DiskSYDXdata.MetricMsgDate.Date {
				//DiskSYDX.Time//采集时间
				//DiskSYDX.Endpoint//设备id
				//DiskSYDX.Value//指标值
				//
				log.Println("DiskSYDX.Time", DiskSYDX.Time, "DiskSYDX.Endpoint", DiskSYDX.Endpoint, "DiskSYDX.Value", DiskSYDX.Value)
				DiskSYDXerr, DiskSYDXdata := QueryGatewaydata(DiskSYDX.Endpoint)
				if DiskSYDXerr != nil {
					if fmt.Sprint(DiskSYDXerr) == "record not found" {
						log.Println("err== `record not found`:", DiskSYDXerr)
						//没有找到，说明该DiskSYDX信息还没有在数据库中有对应的设备记录
						log.Println(DiskSYDXerr, "没有找到，说明该DiskSYDX信息还没有在数据库中有对应的设备记录+++++++++++++++++")
						continue
					} else {
						log.Println("+++++++++++++++++++++++++++++=err==:", DiskSYDXerr)
						continue
					}
				}
				//更新DiskSYDX信息
				gwxx := new(types.BDmWanggjcxx)
				//

				gwxx.FNbYisyypdx = utils.ByteToGB(DiskSYDX.Value)
				gwxx.FDtZuihgxsj = utils.StrTimeTotime(DiskSYDX.Time)
				upDiskSYDXerr := UpdateGatewaydata(DiskSYDXdata.FVcWanggbh, gwxx)
				if upDiskSYDXerr != nil {
					log.Println(upDiskSYDXerr, "更新DiskSYDX信息失败+++++++++++++++++")
					continue
				}

			}

		} else {
			log.Println("查询DiskSYDX指标DiskSYDXdata为空:", DiskSYDXdata)
		}
	}

	//7、磁盘总大小
	DiskZDXmetric := "disk.cap.bytes.total"
	if DiskZDXdata, DiskZDXerr := MetricDataPostWithJson(DiskZDXmetric); DiskZDXerr != nil {
		log.Println("查询DiskZDX指标出错:", DiskZDXerr)
	} else {
		//已获取数据
		if DiskZDXdata != nil {
			//把指标结果存数据库
			for _, DiskZDX := range DiskZDXdata.MetricMsgDate.Date {
				//DiskZDX.Time//采集时间
				//DiskZDX.Endpoint//设备id
				//DiskZDX.Value//指标值
				//
				log.Println("DiskZDX.Time", DiskZDX.Time, "DiskZDX.Endpoint", DiskZDX.Endpoint, "DiskZDX.Value", DiskZDX.Value)
				DiskZDXerr, DiskZDXdata := QueryGatewaydata(DiskZDX.Endpoint)
				if DiskZDXerr != nil {
					if fmt.Sprint(DiskZDXerr) == "record not found" {
						log.Println("err== `record not found`:", DiskZDXerr)
						//没有找到，说明该DiskZDX信息还没有在数据库中有对应的设备记录
						log.Println(DiskZDXerr, "没有找到，说明该DiskZDX信息还没有在数据库中有对应的设备记录+++++++++++++++++")
						continue
					} else {
						log.Println("+++++++++++++++++++++++++++++=err==:", DiskZDXerr)
						continue
					}
				}
				//更新DiskZDX信息
				gwxx := new(types.BDmWanggjcxx)
				//
				gwxx.FNbZongypdx = utils.ByteToGB(DiskZDX.Value)
				gwxx.FDtZuihgxsj = utils.StrTimeTotime(DiskZDX.Time)
				upDiskZDXerr := UpdateGatewaydata(DiskZDXdata.FVcWanggbh, gwxx)
				if upDiskZDXerr != nil {
					log.Println(upDiskZDXerr, "更新DiskZDX信息失败+++++++++++++++++")
					continue
				}
			}

		} else {
			log.Println("查询DiskZDX指标DiskZDXdata为空:", DiskZDXdata)
		}
	}

	return nil
}

//任务三
//获取告警数据,并更新数据
func GatewayAlarmDataUpdate() error {
	//查询的起始时间，查询的结束时间
	var beginTime, endTime int64

	//

	//调告警数据获取接口
	ErrorDataPostWithJson(beginTime, endTime)

	return nil
}

//网关基础信息查询接口 一个dataserver的方法 ，再加 后面需要多个dataserver的
func GatewayDataPostWithJson() (*dto.GatewayDeviceMsgResp, error) {
	//post请求提交json数据
	gw := dto.GatewayDataReq{}
	ba, _ := json.Marshal(gw)
	//POST "http://172.18.70.22:8080/etcpark/dataserver/gateway/list" [new]http://122.51.24.189:8080/etcpark/dataserver/gateway/list
	log.Println("网关基础信息查询接口 Gwmsg_address:", Gwmsg_address)
	resp, err := http.Post(Gwmsg_address, "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		log.Println("post请求网关基础信息查询接口失败:", err)
		return nil, err
	}

	if resp.Body == nil {
		log.Println("resp.Body==nil")
	}

	body, _ := ioutil.ReadAll(resp.Body)

	Resp := new(dto.GatewayDeviceMsgResp)
	//反序列化
	unmerr := json.Unmarshal(body, Resp)
	if unmerr != nil {
		log.Println("json.Unmarshal error", unmerr)
	}
	log.Println("网关基础信息查询接口 Post request with json result:", Resp)
	return Resp, nil
}

//告警信息查询接口
func ErrorDataPostWithJson(beginTime int64, endTime int64) (*dto.ErrorMsgResp, error) {
	//post请求提交json数据 时间戳
	errdatareq := dto.QueryErrorMsgQeq{BeginTime: beginTime, EndTime: endTime}
	ba, _ := json.Marshal(errdatareq)
	//POST
	log.Println("Errormsg_address:", Errormsg_address)
	resp, err := http.Post(Errormsg_address, "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		log.Println("post请求告警信息查询接口失败:", err)
		return nil, err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	Resp := new(dto.ErrorMsgResp)
	unmerr := json.Unmarshal(body, Resp)
	if unmerr != nil {
		log.Println("json.Unmarshal error", unmerr)
	}
	log.Printf("Post request with json result:%v\n", Resp)
	return Resp, nil
}

//
//指标信息查询接口
func MetricDataPostWithJson(metric string) (*dto.MetricMsgResp, error) {
	//post请求提交json数据 时间戳
	metricreq := dto.QueryMetricMsgQeq{Metric: metric}
	ba, _ := json.Marshal(metricreq)
	//POST
	log.Println("Metric_address:", Metric_address)
	resp, err := http.Post(Metric_address, "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		log.Println("post请求指标信息查询接口失败:", err)
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	Resp := new(dto.MetricMsgResp)
	unmerr := json.Unmarshal(body, &Resp)
	if unmerr != nil {
		log.Println("json.Unmarshal error", unmerr)
	}
	log.Printf("Post request with json result:%v\n", Resp)
	return Resp, nil
}
