package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"operationPlatform/db"
	"operationPlatform/dto"
	"operationPlatform/service"
	"operationPlatform/types"
	"operationPlatform/utils"
	"os"
	"path"
	"strconv"
	"time"
)

//运维网关监控

//1、网关列表查询
func Querygatewaylist(c *gin.Context) {
	req := dto.QueryGatewayListQeqdata{}
	//获取请求数据
	if err := c.Bind(&req); err != nil {
		log.Println("查询网关列表,获取请求参数时 err:", err)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "查询网关列表,获取请求参数时 error"})
		return
	}
	log.Println("req.GatewayNumber:", req.GatewayNumber, "req.ParkName:", req.ParkName, "req.Status:", req.Status, "req.Version:", req.Version, "req.UpdateBeginTime:", req.UpdateBeginTime, "req.UpdateEndTime:", req.UpdateEndTime)

	qerr, wgxxs := db.QueryGatewayALLdata(&req)
	if qerr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询网关列表时 error"})
	}
	//数据赋值
	datas := make([]dto.QueryGatewayListResp, 0)
	for _, gwxx := range *wgxxs {
		data := new(dto.QueryGatewayListResp)
		data.TerminalId = gwxx.FVcWanggbh // 设备ID，如CE4C37043A520C93
		data.Parkid = gwxx.FVcTingccbh    // 停车场ID
		qpkerr, pm := db.QueryParkName(gwxx.FVcTingccbh)
		if qpkerr != nil {
			if fmt.Sprint(qpkerr) == "record not found" {
				log.Println("err:", qpkerr)

			} else {
				log.Println("++++++++++++++++++++++++++++++++查询停车场名称错误")
			}
		}
		if pm == nil {
			data.ParkName = gwxx.FVcTingccbh // 停车场名称
		} else {
			data.ParkName = pm.FVcMingc // 停车场名称
		}

		//data.CompanyId = gwxx               // 公司ID
		//data.CompanyName = gwxx             // 公司ID
		data.OnlineStatus = gwxx.FNbZhuangt //"	"status": "1"： 在线状态 0 :离线
		data.Gatewayip = gwxx.FVcIpdz       //   网关IP地址，多个地址则用”, ”分隔
		//CPU使用率
		data.CPU = gwxx.FNbCPUsyl
		MEMpercent, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", gwxx.FNbNeicsyl), 64)
		//内存使用率
		data.MEMpercent = MEMpercent
		Ysyncdx, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", gwxx.FNbYsyncdx), 64)
		//内存使用大小
		data.MEM = Ysyncdx
		//磁盘使用率
		Yingpsyl, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", gwxx.FNbYingpsyl), 64)
		data.DISKpercent = Yingpsyl
		//磁盘使用大小 Float64
		Yisyypdx, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", gwxx.FNbYisyypdx), 64)
		data.DISK = Yisyypdx
		//未处理告警数量
		data.UnprocessedErrors = gwxx.FNbWeiclgjs
		//告警总数
		data.Errors = gwxx.FNbGaojzs
		//重启次数
		data.Restarts = gwxx.FNbChongqcs
		//   场内网关当前版本号
		data.GetwayVersion = gwxx.FVcDangqbbh
		//   场内网关最后更新成功时间
		//查询重启记录表该网关最后更新时间
		restarterr, R := db.GatewayRestarNewTime(gwxx.FVcWanggbh)
		if restarterr != nil {
			if fmt.Sprint(restarterr) == "record not found" {
				log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++err:", restarterr)
				//没有重启记录

			} else {
				log.Println("++++++++++查询重启记录表该网关最后更新时间错误+++++++++")
			}

		} else {
			if R != nil {
				//网关最后重启时间
				data.LastversionUpdatedatetime = R.FDtChongqsj.Format("2006-01-02 15:04:05")
			}
		}

		//天线数量

		qerrRSUerr, errtxsl := db.QueryErrorRSUALLdata(gwxx.FVcWanggbh)
		if qerrRSUerr != nil {
			log.Println("++++++++++查询天线数量错误+++++++++")
		}
		data.RsuNum = errtxsl
		data.RsuALLNum = gwxx.FNbTianxsl
		//延迟
		data.Network = int64(gwxx.FNbWanglyc)
		data.Flag = false
		datas = append(datas, *data)
	}
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: datas, Message: "查询网关列表成功"})
}

//告警列表查询
func QueryAlarmlist(c *gin.Context) {
	req := dto.QueryErrorMsgListQeq{}
	//获取请求数据
	if err := c.Bind(&req); err != nil {
		log.Println("查询告警列表,获取请求参数时 err:", err)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "查询告警列表,获取请求参数时 error"})
		return
	}
	//
	if req.TerminalId == "" {
		//1、查询设备编号
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "查询告警列表,获取请求参数时 设备id 不能为空"})
		return
	}

	//1.获取告警列表数据
	qerr, gjs := db.QueryErrorALLdata(&req)
	if qerr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询告警列表时 error"})
		return
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
		log.Println("查询重启列表,获取请求参数时 err:", err)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "查询重启记录列表,获取请求参数时 error"})
		return
	}
	if req.TerminalId == "" {
		//1、查询设备编号
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "查询重启记录列表,获取请求参数时 设备id 不能为空"})
		return
	}

	//1.获取重启记录列表数据
	qerr, cqs := db.QueryRestartALLdata(&req)
	if qerr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询重启记录列表时 error"})
		return
	}
	datas := make([]dto.QueryRestartListResp, 0)
	for _, chongq := range *cqs {
		data := new(dto.QueryRestartListResp)
		data.TerminalId = chongq.FVcWanggbh
		data.RestartTime = chongq.FDtChongqsj.Format("2006-01-02 15:04:05")

		data.WorkTime = utils.SecondsToTime(chongq.FNbChongqlxgzsc)

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
		log.Println("查询天线列表,获取请求参数时 err:", err)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "查询天线记录列表,获取请求参数时 error"})
		return
	}
	qerr, txs := db.QueryRSUALLdata(req.TerminalId)
	if qerr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询重启记录列表时 error"})
		return
	}
	datas := make([]dto.QueryRSUMsgListResp, 0)
	for _, tx := range *txs {
		data := new(dto.QueryRSUMsgListResp)
		data.TerminalId = tx.FVcWanggbh
		data.RSUIP = tx.FVcIpdz    // 天线ip
		data.Lane = tx.FVcChedwyid // 车道
		data.Isregister = tx.FVcZhuczt
		data.AntennaStatus = tx.FVcTianxzt
		data.AntennaStatusUpdatetime = tx.FVcTianxztgxsj
		data.WorkTime = utils.SecondsToTime(tx.FNbLianxgzsc) //秒
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
		log.Println("查询网关列表,获取请求参数时 err:", err)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "查询网关列表,获取请求参数时 error"})
		return
	}

	//查询网关列表

	qerr, wgxx := db.QueryOneGatewaydata(&req)
	if qerr != nil {
		if fmt.Sprint(qerr) == "record not found" {

			c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询网关列表时 ，该设备不存在"})
			return
		}
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询网关列表时 error"})
		return
	}
	//数据赋值
	data := new(dto.QueryGatewayOneResp)
	data.TerminalId = wgxx.FVcWanggbh // 设备ID，如CE4C37043A520C93

	qpkerr, pm := db.QueryParkName(wgxx.FVcTingccbh)
	if qpkerr != nil {
		if fmt.Sprint(qpkerr) == "record not found" {
			log.Println("err:", qpkerr)

		} else {
			log.Println("++++++++++++++++++++++++++++++++查询停车场名称错误")
		}
	}
	if pm == nil {
		data.ParkName = wgxx.FVcTingccbh // 停车场名称
	} else {
		data.ParkName = pm.FVcMingc // 停车场名称
	}

	data.Gatewayip = wgxx.FVcIpdz         //   网关IP地址，多个地址则用”, ”分隔
	data.GetwayVersion = wgxx.FVcDangqbbh //   场内网关版本号
	data.CPU = wgxx.FNbCPUsyl
	data.MEMpercent = wgxx.FNbNeicsyl
	data.MEM = wgxx.FNbZongncdx
	data.DISKpercent = wgxx.FNbYingpsyl
	data.DISK = wgxx.FNbZongypdx
	data.Network = int64(wgxx.FNbWanglyc)
	data.WorkTime = utils.SecondsToTime(wgxx.FNbYunxsc) //工作时长

	//qerr, txs := db.QueryRSUALLdata(req.TerminalId)
	//if qerr != nil {
	//	if fmt.Sprint(qerr) == "record not found" {
	//	} else {
	//		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询天线记录列表,获取天线数量时 error"})
	//		return
	//	}
	//}
	//if txs != nil {
	//	data.Restarts = len(*txs)
	//} else {
	//
	//}
	data.Restarts = wgxx.FNbChongqcs
	qrerr, cq := db.QueryRestartOnedata(req.TerminalId)
	if qrerr != nil {
		if fmt.Sprint(qrerr) == "record not found" {
		} else {
			c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询重启记录列表,获取重启次数时 error"})
			return
		}
	}
	if cq != nil {
		data.RestartTime = cq.FDtChongqsj.Format("2006-01-02 15:04:05") //重启时间
	} else {
		data.RestartTime = ""
	}

	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: data, Message: "查询网关设备详情成功"})
}

//添加设备 Addgatewaydevice
func Addgatewaydevice(c *gin.Context) {
	req := dto.GatewayDevicedata{}
	//1、获取请求数据
	if err := c.Bind(&req); err != nil {
		log.Println("添加设备 获取请求参数时 err:", err)
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
	gwxx.FNbZhuangt = 0             //	'状态 0：离线、1：在线',
	gwxx.FNbGaojzs = 0              //	`F_NB_GAOJZS` int(11) NOT NULL DEFAULT '0' COMMENT '告警总数',
	gwxx.FNbWeiclgjs = 0            //	`F_NB_WEICLGJS` int(11) NOT NULL DEFAULT '0' COMMENT '未处理告警数',
	gwxx.FNbChongqcs = 0            //	`F_NB_CHONGQCS` int(11) NOT NULL DEFAULT '0' COMMENT '重启次数',
	gwxx.FNbTianxsl = 0             //	`F_NB_TIANXSL` int(11) DEFAULT NULL COMMENT '天线数量',
	gwxx.FNbWanglyc = 0             //	`F_NB_WANGLYC` bigint(20) DEFAULT NULL COMMENT '网络延迟 单位：ms',
	gwxx.FDtChuangjsj = time.Now()  //	`F_DT_CHUANGJSJ` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
	gwxx.FDtZuihgxsj = time.Now()
	gwxx.FDtZuijgxbbsj = time.Now()
	//3、插入数据

	//插入前先校验数据
	qerr, gwdata := db.QueryGatewaydata(gwxx.FVcWanggbh)
	if qerr != nil {
		if fmt.Sprint(qerr) == "record not found" {
			log.Println("db.QueryGatewaydata err:", qerr)
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
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: types.StatusText(types.StatusADDSuccessfully), Message: "添加设备成功"})
}

//增加版本——软件更新版本
func AddNewVersion(c *gin.Context) {
	req := dto.AddGatewayVersionQeq{}
	//1、获取请求数据
	if err := c.Bind(&req); err != nil {
		log.Println("添加软件更新版本 获取请求参数时 err:", err)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "添加软件更新版本，获取请求参数时 error"})
		return
	}
	//2、校验参数
	if req.Version == "" {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "添加软件更新版本，版本号不能为空"})
		return
	}

	if req.VersionNote == "" {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "添加软件更新版本，更新内容不能为空"})
		return
	}

	if req.Name == "" {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "添加软件更新版本，上传者姓名不能为空"})
		return
	}
	if req.FileName == "" {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "添加软件更新版本，上传文件名不能为空"})
		return
	}

	qverr, data := db.QueryOneVersiondata(req.Version)
	if qverr != nil {
		if fmt.Sprint(qverr) == "record not found" {
			log.Println("db.QueryGatewaydata err:", qverr)
		} else {
			c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "添加软件更新版本，先查询软件版本是否已上传失败"})
			return
		}
	}
	if data != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusDataISExist, Data: types.StatusText(types.StatusDataISExist), Message: "添加软件版本，软件版本已经存在"})
		return
	}

	inerr := db.AddVersion(&req)
	if inerr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusINSERTDataError, Data: types.StatusText(types.StatusINSERTDataError), Message: "添加软件更新版本，新增软件更新版本失败，请检查上传信息"})
		return
	}
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: types.StatusText(types.StatusSuccessfully), Message: "添加软件更新版本，上传成功"})
}

//
//上传版本文件
func UploadVersionFile(c *gin.Context) {
	//接收文件
	file, _ := c.FormFile("file")
	log.Println("FileName:", file.Filename, "file.Header", file.Header)
	dst := path.Join("./version/", file.Filename)
	//保存文件
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "添加软件更新版本，保存文件失败"})
		return
	}

	log.Println("要发送上传文件的path:=", dst)
	//读取文件
	f, oserr := os.Open("./version/" + dst)
	if oserr != nil {
		log.Println("os.Open error:", oserr)
		return
	}
	data, rerr := ioutil.ReadAll(f)
	if rerr != nil {
		return
	}

	defer func() {
		_ = f.Close()
	}()

	//3、把文件上传到OSS对象服务器上
	//log.Println("req.FileName:", req.FileName)
	service.FileUpload(data, file.Filename)

	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: types.StatusText(types.StatusSuccessfully), Message: "添加软件更新版本，上传成功"})
}

//查询软件版本列表
func QuerygatewayVersionlist(c *gin.Context) {
	req := dto.QueryVersionQeq{}
	//1、获取请求数据
	if err := c.Bind(&req); err != nil {
		log.Println("查询软件版本列表，获取请求参数时 err:", err)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "查询软件版本列表，获取请求参数时 error"})
		return
	}
	//1.查询软件版本列表数据
	qerr, vs := db.QueryVersionALLdata(&req)
	if qerr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询软件版本列表失败"})
		return
	}
	datas := make([]dto.QueryVersionListResp, 0)
	for _, v := range *vs {
		resq := new(dto.QueryVersionListResp)
		resq.Version = v.FVcRuanjbbh                            //版本号
		resq.VersionNote = v.FVcBanbgxnr                        //版本描述
		resq.Time = v.FDtShangcsj.Format("2006-01-02 15:04:05") //版本上传时间
		err, num := db.QueryVersionNumdata(v.FVcRuanjbbh)       //版本使用设备数
		if err != nil {
			c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询软件版本使用次数失败"})
			return
		}
		resq.Num = num //版本使用设备数
		resq.Name = v.FVcShangczxm
		datas = append(datas, *resq)

	}
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: datas, Message: "查询软件版本列表成功"})
}

//DeleteNewgatewayVersion
func DeleteNewVersion(c *gin.Context) {
	req := dto.DeleteVersionQeq{}
	//1、获取请求数据
	if err := c.Bind(&req); err != nil {
		log.Println("添加软件更新版本 获取请求参数时 err:", err)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "添加软件更新版本，获取请求参数时 error"})
		return
	}
	//2、校验参数
	if len(req.Version) == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "请选择要删除的软件版本"})
		return
	}
	derr := db.DeleteVersionsdata(&req)
	if derr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusDeleteDataError, Data: types.StatusText(types.StatusDeleteDataError), Message: "删除软件版本时错误"})
		return
	}
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: types.StatusText(types.StatusSuccessfully), Message: "删除软件版本成功"})
}

//QueryVersionlist
//软件版本下拉框
func QueryVersionlist(c *gin.Context) {
	qerr, datas := db.QueryVersionALL()
	if qerr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询软件版本列表失败"})
		return
	}
	var resp dto.QueryVersionsResp
	for _, data := range *datas {
		var v dto.VersionMsg
		v.Version = data.FVcRuanjbbh
		v.VersionNote = data.FVcBanbgxnr
		resp.Versions = append((resp.Versions), v)
	}
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: resp, Message: "获取软件版本下拉框成功"})
}

//QueryGatewaylist 网关设备下拉框
func QueryGatewaylist(c *gin.Context) {
	qerr, datas := db.QueryGatewayALL()
	if qerr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询设备列表失败"})
		return
	}
	var resp dto.QueryGatewaysResp
	for _, data := range *datas {
		resp.TerminalId = append((resp.TerminalId), data.FVcWanggbh)
	}
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: resp, Message: "获取设备下拉框成功"})
}

//QueryparkNamelist停车场下拉框
func QueryparkNamelist(c *gin.Context) {
	qerr, datas := db.QueryParkNameALL()
	if qerr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusQueryDataError, Data: types.StatusText(types.StatusQueryDataError), Message: "查询停车场列表失败"})
		return
	}
	var resp dto.QueryParkNamesResp
	for _, data := range *datas {
		p := new(dto.ParkMSG)
		p.ParkNum = data.FVcTingccbh
		p.ParkName = data.FVcMingc
		resp.Parkmsg = append((resp.Parkmsg), *p)
	}
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: resp, Message: "获取停车场下拉框成功"})
}

//VersionUpdate
func VersionUpdate(c *gin.Context) {
	req := dto.VersionUpdateQeq{}
	//1、获取请求数据
	if err := c.Bind(&req); err != nil {
		log.Println("添加软件更新版本 获取请求参数时 err:", err)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "添加软件更新版本，获取请求参数时 error"})
		return
	}
	//2、校验参数
	if len(req.Gwids) == 0 {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "请选择要更新的设备网关"})
		return
	}
	if req.Version == "" {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "请选择要更新的软件版本"})
		return
	}

	if req.UpdateStatus == 0 && req.UpdateTime == "" {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "请选择要立即更新的时间"})
		return
	}

	if req.UpdateStatus == 1 && req.UpdateTime == "" {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "请选择要定时更新的时间"})
		return
	}

	derr := db.VersionsUpdatedata(&req)
	if derr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusDeleteDataError, Data: types.StatusText(types.StatusDeleteDataError), Message: "更新软件版本时错误"})
		return
	}
	//2、返回数据
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: types.StatusText(types.StatusSuccessfully), Message: "更新软件版本成功"})

	//执行版本更新,把要更新的网关设备已经版本信息，以及版本文件的地址 反馈给dataserver

	//1、查询要更新的设备与版本

}

//PerformVersionUpdate
func PerformVersionUpdate(c *gin.Context) {
	req := dto.PerformVersionUpdateQeq{}
	//1、获取请求数据
	if err := c.Bind(&req); err != nil {
		log.Println("添加软件更新版本 获取请求参数时 err:", err)
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "添加软件更新版本，获取请求参数时 error"})
		return
	}
	//2、校验参数
	if req.TerminalId == "" {
		log.Println("要更新的设备网关为空")
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "要更新的设备网关不能为空"})
		return
	}
	//if req.OSVersion == "" {
	//	log.Println("要更新的设备网关操作系版本为空" )
	//	c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "要更新的设备网关操作系版本不能为空"})
	//	return
	//}
	//
	//if req.OSArch == "" {
	//	log.Println("要更新的设备网关的操作系处理器架构为空" )
	//	c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "要更新的设备网关的操作系处理器架构不能为空"})
	//	return
	//}

	if req.GatewayVersion == "" {
		log.Println("要更新的设备网关的场内网关版本号为空")
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "要更新的设备网关的场内网关版本号不能为空"})
		return
	}

	//if req.CurrversionMd5 == "" {
	//	log.Println("要更新的设备网关的场内网关gateway文件MD5值为空" )
	//	c.JSON(http.StatusOK, dto.Response{Code: types.StatusGetReqError, Data: types.StatusText(types.StatusGetReqError), Message: "要更新的设备网关的场内网关gateway文件MD5值不能为空"})
	//	return
	//}

	derr, banbgx := db.PerformVersionsUpdatedata(&req)
	if derr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusDeleteDataError, Data: types.StatusText(types.StatusDeleteDataError), Message: "更新软件版本时错误"})
		return
	}
	log.Println(banbgx)
	//2、返回数据
	resp := new(dto.PerformVersionUpdateResp)
	resp.TerminalId = "CE4C37043A520C93" //设备ID，如CE4C37043A520C93
	resp.GatewayVersion = "version"      //网关新版本号
	resp.Download_url = "http://abc.com" //网关下载的URL
	resp.Upgrade = "0"                   //是否需要升级:0不需要 1需要升级
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: resp, Message: "查询设备网关软件版本是否更新成功"})
}
