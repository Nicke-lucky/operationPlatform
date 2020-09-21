package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"operationPlatform/dto"
	"operationPlatform/types"
	"operationPlatform/utils"
)

//运维网关监控

//1、网关列表查询
func Querygatewaylist(c *gin.Context) {
	//校验参数
	//网关编号
	//停车场名称
	//状态：全部，在线、离线
	//软件版本
	//更新时间
	//1.获取网关列表数据
	randStr := utils.GetRandStr(4)
	logrus.Println("随机获取验证码文字:", randStr)
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: randStr, Message: "随机获取验证码文字成功"})
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

//Addgatewaydevice
func Addgatewaydevice(c *gin.Context) {
	//校验参数
	//设备编号
	//部署停车场
	//备注：
	//1.获取网关列表数据
	randStr := utils.GetRandStr(4)
	logrus.Println("随机获取验证码文字:", randStr)
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: randStr, Message: "随机获取验证码文字成功"})
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
