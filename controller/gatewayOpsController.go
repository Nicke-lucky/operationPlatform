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
	"strconv"
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
	//查询网关列表
	qerr, wgxxs := db.QueryGatewayALLdata(&req)
	if qerr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询网关列表时 error"})
	}
	//数据赋值
	datas := make([]dto.QueryGatewayListResp, 0)
	for _, gwxx := range *wgxxs {
		data := new(dto.QueryGatewayListResp)
		data.TerminalId = gwxx.FVcWanggbh // 设备ID，如CE4C37043A520C93
		//data.Parkid = gwxx                  // 停车场ID
		data.ParkName = gwxx.FVcTingccbh // 停车场名称
		//data.CompanyId = gwxx               // 公司ID
		//data.CompanyName = gwxx             // 公司ID
		data.OnlineStatus = gwxx.FNbZhuangt //"	"status": "1"： 在线状态 0 :离线
		data.Gatewayip = gwxx.FVcIpdz       //   网关IP地址，多个地址则用”, ”分隔
		data.CPU = gwxx.FNbCPUsyl
		data.MEMpercent = gwxx.FNbNeicsyl
		data.MEM = gwxx.FNbYsyncdx
		data.DISKpercent = gwxx.FNbYingpsyl
		data.DISK = gwxx.FNbYisyypdx
		data.UnprocessedErrors = gwxx.FNbWeiclgjs
		data.Errors = gwxx.FNbGaojzs
		data.Restarts = gwxx.FNbChongqcs
		data.GetwayVersion = gwxx.FVcDangqbbh //   场内网关当前版本号

		data.LastversionUpdatedatetime = gwxx.FDtZuihgxsj.Format("2006-01-02 15:04:05") //   场内网关最后更新成功时间
		data.RsuNum = gwxx.FNbTianxsl
		data.Network = gwxx.FNbWanglyc
		datas = append(datas, *data)
	}

	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: datas, Message: "随机获取验证码文字成功"})
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
	req := dto.QueryErrorMsgListQeq{}
	//获取请求数据
	if err := c.Bind(&req); err != nil {
		log.Println("查询告警列表,获取请求参数时 err: %v", err)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "查询告警列表,获取请求参数时 error"})
		return
	}

	//1.获取告警列表数据
	qerr, gjs := db.QueryErrorALLdata(&req)
	if qerr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询告警列表时 error"})
	}
	datas := make([]dto.QueryErrorListResp, 0)
	for _, gaoj := range *gjs {
		data := new(dto.QueryErrorListResp)
		data.TerminalId = gaoj.FVcWanggbh
		data.ErrorTime = gaoj.FDtGaojsj.Format("2006-01-02 15:04:05") //
		data.ErrorDescribe = gaoj.FVcGaojms                           //
		data.ManId = gaoj.FVcChulrid
		data.ManName = gaoj.FVcChulrxm //
		data.Time = gaoj.FDtChulsj.Format("2006-01-02 15:04:05")
		datas = append(datas, *data)
	}
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: datas, Message: "查询告警列表成功"})
}

//
//重启记录列表查询
func QueryRestartRecordlist(c *gin.Context) {
	req := dto.QueryRestartMsgListQeq{}
	//获取请求数据
	if err := c.Bind(&req); err != nil {
		log.Println("查询重启列表,获取请求参数时 err: %v", err)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "查询重启记录列表,获取请求参数时 error"})
		return
	}

	//1.获取重启记录列表数据
	qerr, cqs := db.QueryRestartALLdata(&req)
	if qerr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询重启记录列表时 error"})
	}
	datas := make([]dto.QueryRestartListResp, 0)
	for _, chongq := range *cqs {
		data := new(dto.QueryRestartListResp)
		data.TerminalId = chongq.FVcWanggbh
		data.RestartTime = chongq.FDtChongqsj.Format("2006-01-02 15:04:05")
		data.WorkTime = strconv.Itoa(chongq.FNbChongqlxgzsc)

		datas = append(datas, *data)
	}
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: datas, Message: "查询重启记录列表成功"})
}

//天线列表查询
func QueryRSURecordlist(c *gin.Context) {
	req := dto.QueryRSUMsgListQeq{}
	//获取请求数据
	if err := c.Bind(&req); err != nil {
		log.Println("查询天线列表,获取请求参数时 err: %v", err)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "查询天线记录列表,获取请求参数时 error"})
		return
	}
	qerr, txs := db.QueryRSUALLdata(req.TerminalId)
	if qerr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询重启记录列表时 error"})
	}
	datas := make([]dto.QueryRSUMsgListResp, 0)
	for _, tx := range *txs {
		data := new(dto.QueryRSUMsgListResp)
		data.TerminalId = tx.FVcWanggbh
		data.RSUIP = tx.FVcIpdz                       // 天线ip
		data.Lane = tx.FVcChedwyid                    // 车道
		data.WorkTime = strconv.Itoa(tx.FNbLianxgzsc) //秒
		datas = append(datas, *data)
	}
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: datas, Message: "查询天线列表成功"})
}

// 网关设备详情
func QueryGatewayDeviceDetails(c *gin.Context) {
	req := dto.QueryGatewayOneQeqdata{}
	//获取请求数据
	if err := c.Bind(&req); err != nil {
		log.Println("查询网关列表,获取请求参数时 err: %v", err)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "查询网关列表,获取请求参数时 error"})
		return
	}
	//查询网关列表
	qerr, wgxx := db.QueryOneGatewaydata(&req)
	if qerr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询网关列表时 error"})
	}
	//数据赋值
	data := new(dto.QueryGatewayOneResp)
	data.TerminalId = wgxx.FVcWanggbh     // 设备ID，如CE4C37043A520C93
	data.ParkName = wgxx.FVcTingccbh      // 停车场名称
	data.Gatewayip = wgxx.FVcIpdz         //   网关IP地址，多个地址则用”, ”分隔
	data.GetwayVersion = wgxx.FVcDangqbbh //   场内网关版本号
	data.CPU = wgxx.FNbCPUsyl
	data.MEMpercent = wgxx.FNbNeicsyl
	data.MEM = wgxx.FNbZongncdx
	data.DISKpercent = wgxx.FNbYingpsyl
	data.DISK = wgxx.FNbZongypdx
	data.Network = wgxx.FNbWanglyc

	qerr, txs := db.QueryRSUALLdata(req.TerminalId)
	if qerr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询重启记录列表时 error"})
	}
	data.Restarts = len(*txs)

	qrerr, cq := db.QueryRestartOnedata(req.TerminalId)
	if qrerr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询重启记录列表时 error"})
	}
	data.WorkTime = cq.FNbChongqlxgzsc
	data.RestartTime = cq.FDtChongqsj.Format("2006-01-02 15:04:05")

	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: data, Message: "查询网关设备详情成功"})
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
