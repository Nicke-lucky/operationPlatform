package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"operationPlatform/db"
	"operationPlatform/dto"
	"operationPlatform/types"
	"operationPlatform/utils"
	"time"
)

//运维网关监控

//1、网关列表查询
func Querygatewaylist(c *gin.Context) {

	req := dto.QueryGatewayListQeqdata{}
	//获取请求数据
	if err := c.Bind(&req); err != nil {
		log.Println("查询网关列表,获取请求参数时 err: %v", err)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "查询网关列表,获取请求参数时 error"})
		return
	}
	//1、校验参数 默认选择全部
	//GatewayNumber    //设备编号 网关编号
	//ParkName                         //停车场名称
	//Status                //状态：2全部，1在线、0离线
	//Version                //软件版本
	//UpdateBeginTime          //更新时间
	//UpdateEndTime
	//网关编号
	//停车场名称
	//状态：全部，在线、离线
	//软件版本
	//更新时间

	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: "", Message: "随机获取验证码文字成功"})
}

/*
//未处理告警列表查询
func QueryUnprocessedAlarmlist(c *gin.Context) {
	//1.获取网关列表数据
	randStr := utils.GetRandStr(4)
	logrus.Println("随机获取验证码文字:", randStr)
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: randStr, Message: "随机获取验证码文字成功"})
}*/

//告警列表查询
func QueryAlarmlist(c *gin.Context) {
	//校验请求参数
	//网关id
	//起始时间，告警 时间
	//结束时间，告警 时间
	//处理状态 0：未处理的告警，1已处理的告警，2全部展示

	//1.获取网关列表数据
	randStr := utils.GetRandStr(4)
	logrus.Println("随机获取验证码文字:", randStr)
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: randStr, Message: "随机获取验证码文字成功"})
}

//
//重启记录列表查询
func QueryRestartRecordlist(c *gin.Context) {
	//校验请求参数
	//网关id
	//起始时间，重启时间2020-09-21 13:13:13
	//结束时间，重启时间
	//处理状态 0：未处理的告警，1已处理的告警，2全部展示
	//1.获取网关列表数据
	randStr := utils.GetRandStr(4)
	logrus.Println("随机获取验证码文字:", randStr)
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: randStr, Message: "随机获取验证码文字成功"})
}

//天线列表查询
func QueryRSURecordlist(c *gin.Context) {
	//校验请求参数
	//网关id
	//1.获取网关列表数据
	randStr := utils.GetRandStr(4)
	logrus.Println("随机获取验证码文字:", randStr)
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: randStr, Message: "随机获取验证码文字成功"})
}

// 网关设备详情
func QueryGatewayDeviceDetails(c *gin.Context) {
	//校验请求参数
	//网关id
	//1.获取网关列表数据
	randStr := utils.GetRandStr(4)
	logrus.Println("随机获取验证码文字:", randStr)
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: randStr, Message: "随机获取验证码文字成功"})
}

//添加设备 Addgatewaydevice
func Addgatewaydevice(c *gin.Context) {
	req := dto.GatewayDevicedata{}
	//1、获取请求数据
	if err := c.Bind(&req); err != nil {
		log.Println("添加设备 获取请求参数时 err: %v", err)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "添加设备，获取请求参数时 error"})
		return
	}
	//2、校验参数
	//设备编号
	if req.GatewayNumber == "" {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "添加设备，网关编号不能为空"})
		return
	}
	//部署停车场
	if req.ParkName == "" {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "添加设备，停车场名称不能为空"})
		return
	}
	//备注：

	//3、插入数据
	gwxx := new(types.BDmWanggjcxx)
	//3、数据赋值
	gwxx.FVcWanggbh = req.GatewayNumber //	`F_VC_WANGGBH` varchar(32) NOT NULL COMMENT '网关编号',
	//gwxx.FVcGongsID    = //	`F_VC_GONGSID` varchar(32) NOT NULL COMMENT '公司ID',

	gwxx.FVcTingccbh = req.ParkName //	`F_VC_TINGCCBH` varchar(32) NOT NULL COMMENT '停车场编号',
	gwxx.FNbZhuangt = 0             //	`F_NB_ZHUANGT` int(11) NOT NULL DEFAULT '0' COMMENT '状态 0：离线、1：在线',
	gwxx.FNbGaojzs = 0              //	`F_NB_GAOJZS` int(11) NOT NULL DEFAULT '0' COMMENT '告警总数',
	gwxx.FNbWeiclgjs = 0            //	`F_NB_WEICLGJS` int(11) NOT NULL DEFAULT '0' COMMENT '未处理告警数',
	gwxx.FNbChongqcs = 0            //	`F_NB_CHONGQCS` int(11) NOT NULL DEFAULT '0' COMMENT '重启次数',
	gwxx.FNbTianxsl = 0             //	`F_NB_TIANXSL` int(11) DEFAULT NULL COMMENT '天线数量',
	gwxx.FNbWanglyc = 0             //	`F_NB_WANGLYC` bigint(20) DEFAULT NULL COMMENT '网络延迟 单位：ms',
	gwxx.FDtChuangjsj = time.Now()  //	`F_DT_CHUANGJSJ` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	gwxx.FDtZuihgxsj = time.Now()
	//3、插入数据

	//插入前先校验数据
	qerr, gwdata := db.QueryGatewaydata(gwxx.FVcWanggbh)
	if qerr != nil {
		if fmt.Sprint(qerr) == "record not found" {
			log.Println("db.QueryGatewaydata err== `record not found`:", qerr)
		} else {
			c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "添加设备，先查询设备信息失败"})
			return
		}
	}
	if gwdata != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusDataISExist, Data: types.StatusText(types.StatusDataISExist), Message: "添加设备，网关设备信息id已经存在"})
		return
	}

	inerr := db.GatewayInsert(gwxx)
	if inerr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusINSERTDataError, Data: types.StatusText(types.StatusINSERTDataError), Message: "添加设备，新增设备信息失败，请检查设备信息"})
		return
	}
	//4、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: types.StatusText(types.StatusSuccessfully), Message: "添加设备成功"})
}

//网关软件更新
func AddNewgatewayVersion(c *gin.Context) {
	//校验参数
	//设备编号

	//1.获取网关列表数据
	randStr := utils.GetRandStr(4)
	logrus.Println("随机获取验证码文字:", randStr)
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: randStr, Message: "随机获取验证码文字成功"})
}

//查询软件版本列表
func QuerygatewayVersionlist(c *gin.Context) {
	//校验参数
	//设备编号

	//1.获取网关列表数据
	randStr := utils.GetRandStr(4)
	logrus.Println("随机获取验证码文字:", randStr)
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: randStr, Message: "随机获取验证码文字成功"})
}

//DeleteNewgatewayVersion
func DeleteNewgatewayVersion(c *gin.Context) {
	//校验参数
	//设备编号

	//1.获取网关列表数据
	randStr := utils.GetRandStr(4)
	logrus.Println("随机获取验证码文字:", randStr)
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: randStr, Message: "随机获取验证码文字成功"})
}
