package db

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"operationPlatform/dto"
	"operationPlatform/types"
	"operationPlatform/utils"
	"time"
)

//
var FilePath string

//1、  1、新增网关列表
func GatewayInsert(gwxx *types.BDmWanggjcxx) error {
	db := utils.GormClient.Client
	if err := db.Table("b_dm_wanggjcxx").Create(gwxx).Error; err != nil {
		// 错误处理...
		log.Println("Insert b_dm_wanggjcxx error", err)
		return err
	}
	log.Println("新增 网关基础信息表 数据，插入成功！", "网关编号:=", gwxx.FVcWanggbh)
	return nil
}

//2、 Query查询网关信息 根据网关编号  b_dm_wanggjcxx
func QueryGatewaydata(FVcWanggbh string) (error, *types.BDmWanggjcxx) {
	db := utils.GormClient.Client
	gwxx := new(types.BDmWanggjcxx)
	//赋值
	if err := db.Table("b_dm_wanggjcxx").Where("F_VC_WANGGBH =?", FVcWanggbh).Last(gwxx).Error; err != nil {
		log.Println("查询 网关基础信息表最新数据时 QueryTabledata error :", err)
		return err, nil
	}
	log.Println("查询网关基础信息表 数据，成功！数据结果:", gwxx.FVcWanggbh)
	return nil, gwxx
}

//3、更新网关信息表 根据网关编号
func UpdateGatewaydata(Wanggbh string, gwdata *types.BDmWanggjcxx) error {
	db := utils.GormClient.Client
	if err := db.Table("b_dm_wanggjcxx").Where("F_VC_WANGGBH=?", Wanggbh).Updates(gwdata).Error; err != nil {
		log.Println("更新网关基础信息表 error", err)
		return err
	}
	log.Println("更新网关基础信息表 ok !")
	return nil
}

//所有的网关，用于判断是否离线
func QueryALlGatewaydata() (error, *[]types.BDmWanggjcxx) {
	db := utils.GormClient.Client
	gwxxs := make([]types.BDmWanggjcxx, 0)
	//赋值
	if err := db.Table("b_dm_wanggjcxx").Find(&gwxxs).Error; err != nil {
		log.Println("查询 网关基础信息表所有数据时QueryALlGatewaydata error :", err)
		return err, nil
	}
	log.Println("查询网关基础信息表 数据，成功！数据结果:")
	return nil, &gwxxs
}

//4、查询网关信息多条数据【所有】
func QueryGatewayALLdata(req *dto.QueryGatewayListQeqdata) (error, *[]types.BDmWanggjcxx) {
	db := utils.GormClient.Client
	gwxxs := make([]types.BDmWanggjcxx, 0)
	log.Println("req:", req)
	mytable := db.Table("b_dm_wanggjcxx")
	//全部设备
	if req.GatewayNumber != "" {
		//1、查询设备编号
		mytable = mytable.Where("F_VC_WANGGBH = ?", req.GatewayNumber)
	}
	if req.ParkName != "" {
		//2、查询停车场
		mytable = mytable.Where("F_VC_TINGCCBH = ?", req.ParkName)
	}
	if req.Status != 2 {
		//3、查询状态[]
		mytable = mytable.Where("F_NB_ZHUANGT = ?", req.Status)
	}
	if req.Version != "" {
		//4、查询版本【】
		mytable = mytable.Where("F_VC_DANGQBBH = ?", req.Version)
	}
	if req.UpdateEndTime != "" || req.UpdateBeginTime != "" {
		//5、查询时间【】
		mytable = mytable.Where("F_DT_ZUIHGXSJ>=?", req.UpdateBeginTime+" 00:00:00").Where("F_DT_ZUIHGXSJ<=?", req.UpdateEndTime+" 23:59:59")
	}

	if err := mytable.Find(&gwxxs).Error; err != nil {
		log.Println("查询 网关基础信息表ALL数据时 error :", err)
		return err, nil
	}
	log.Println("查询网关基础信息表 数据，成功！数据结果:", "共", len(gwxxs), "个设备")
	return nil, &gwxxs
}

//查询告警信息
func QueryErrorALLdata(req *dto.QueryErrorMsgListQeq) (error, *[]types.BDmGaoj) {
	db := utils.GormClient.Client
	gjs := make([]types.BDmGaoj, 0)
	log.Println("req:", req)
	mytable := db.Table("b_dm_gaoj")
	//全部设备
	if req.TerminalId != "" {
		//1、查询设备编号
		mytable = mytable.Where("F_VC_WANGGBH = ?", req.TerminalId)
	}
	if req.Status != 2 {
		//3、查询状态[]
		mytable = mytable.Where("F_NB_ZHUANGT = ?", req.Status)
	}
	if req.BeginTime != "" || req.EndTime != "" {
		//5、查询时间【】
		mytable = mytable.Where("F_DT_GAOJSJ >=?", req.BeginTime+" 00:00:00").Where("F_DT_GAOJSJ <=?", req.EndTime+" 23:59:59")
	}

	//1、校验参数 默认选择全部
	//网关设备id
	if err := mytable.Find(&gjs).Error; err != nil {
		log.Println("查询 告警信息表ALL数据时 error :", err)
		return err, nil
	}
	log.Println("查询告警表 数据，成功！数据结果:", "共", len(gjs), "个告警")
	return nil, &gjs
}

func QueryErrordata(gwid string) (int64, error) {
	db := utils.GormClient.Client
	//全部设备
	var Count int64
	if err := db.Table("b_dm_gaoj").Where("F_VC_WANGGBH = ?", gwid).Count(&Count).Error; err != nil {
		log.Println("查询告警表 数据，ALL数据时 error :", err)
		return 0, err
	}
	log.Println("++++++++++++++++++查询告警表 数据ok,count:", Count)
	return Count, nil
}

//未处理告警
func QueryUndisposedError(gwid string) (int64, error) {
	db := utils.GormClient.Client
	//全部设备
	var Count int64
	if err := db.Table("b_dm_gaoj").Where("F_VC_WANGGBH = ?", gwid).Where("F_NB_ZHUANGT=?", 0).Count(&Count).Error; err != nil {
		log.Println("查询 重启信息表ALL数据时 error :", err)
		return 0, err
	}
	log.Println("+++++++++++++++++查询未处理告警表 数据ok,count:", Count)
	return Count, nil
}

//查询重启信息
func QueryRestartALLdata(req *dto.QueryRestartMsgListQeq) (error, *[]types.BDmChongq) {
	db := utils.GormClient.Client
	gjs := make([]types.BDmChongq, 0)
	log.Println("req:", req)
	mytable := db.Table("b_dm_chongq")
	if req.TerminalId != "" {
		//1、查询设备编号
		mytable = mytable.Where("F_VC_WANGGBH = ?", req.TerminalId)
	}

	if req.BeginTime != "" || req.EndTime != "" {
		//5、查询时间【】
		mytable = mytable.Where("F_DT_CHONGQSJ >= ?", req.BeginTime+" 00:00:00").Where("F_DT_CHONGQSJ <= ?", req.EndTime+" 23:59:59")
	}

	if err := mytable.Find(&gjs).Error; err != nil {
		log.Println("查询 重启信息表ALL数据时 error :", err)
		return err, nil
	}
	log.Println("查询重启信息表 数据，成功！数据结果:", "共", len(gjs), "个重启")
	return nil, &gjs
}

//查询重启信息ByGWID
func QueryRestartOnedata(TerminalId string) (error, *types.BDmChongq) {
	db := utils.GormClient.Client
	cq := new(types.BDmChongq)
	log.Println("req.TerminalId:", TerminalId)

	if err := db.Table("b_dm_chongq").Where("F_VC_WANGGBH =?", TerminalId).Last(cq).Error; err != nil {
		log.Println("查询 重启信息表One数据时 error :", err)
		return err, nil
	}
	log.Println("查询重启信息表 数据，成功！数据结果:", cq.FNbChongqlxgzsc)
	return nil, cq
}

func QueryRestartCount(TerminalId string) (error, int64) {
	db := utils.GormClient.Client
	log.Println("req.TerminalId:", TerminalId)
	var Count int64
	if err := db.Table("b_dm_chongq").Where("F_VC_WANGGBH =?", TerminalId).Count(&Count).Error; err != nil {
		log.Println("查询 重启信息表One数据时 error :", err)
		return err, 0
	}
	log.Println("查询重启信息表 数据，成功！Count:", Count)
	return nil, Count
}

//查询天线信息
func QueryRSUALLdata(TerminalId string) (error, *[]types.BDmTianxxx) {
	db := utils.GormClient.Client
	gjs := make([]types.BDmTianxxx, 0)
	log.Println("req.TerminalId:", TerminalId)
	if err := db.Table("b_dm_tianxxx").Where("F_VC_WANGGBH =?", TerminalId).Find(&gjs).Error; err != nil {
		log.Println("查询 天线信息表ALL数据时 error :", err)
		return err, nil
	}
	log.Println("查询天线信息表 数据，成功！数据结果:", "共", len(gjs), "个天线")
	return nil, &gjs
}

//1.获取网关设备详情
func QueryOneGatewaydata(req *dto.QueryGatewayOneQeqdata) (error, *types.BDmWanggjcxx) {
	db := utils.GormClient.Client
	gwxx := new(types.BDmWanggjcxx)
	log.Println("req:", req)
	//校验请求参数
	//网关id
	if err := db.Table("b_dm_wanggjcxx").Where("F_VC_WANGGBH =?", req.TerminalId).Last(gwxx).Error; err != nil {
		log.Println("查询 网关基础信息表ALL数据时 error :", err)
		return err, nil
	}
	log.Println("查询网关基础信息表 数据，成功！数据结果:", gwxx.FVcWanggbh)
	return nil, gwxx
}

//增减软件版本
func AddVersion(req *dto.AddGatewayVersionQeq) error {
	db := utils.GormClient.Client
	version := new(types.BDmRuanjbb)

	version.FVcRuanjbbh = req.Version     //	`F_VC_RUANJBBH` varchar(512) NOT NULL COMMENT '软件版本号',
	version.FVcBanbgxnr = req.VersionNote //	`F_VC_BANBGXNR` varchar(1024) DEFAULT NULL COMMENT '版本更新内容',

	version.FDtShangcsj = time.Now() // utils.StrTimeTotime(req.Time) //	`F_DT_SHANGCSJ` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
	//version.FVcShangczid = //	`F_VC_SHANGCZID` varchar(32) NOT NULL COMMENT '上传者ID',
	version.FVcShangczxm = req.Name             //	`F_VC_SHANGCZXM` varchar(32) DEFAULT NULL COMMENT '上传者姓名',
	version.FVcWenjlj = FilePath + req.FileName // req.FilePath //	`F_VC_WENJLJ` varchar(512) DEFAULT NULL COMMENT '文件路径',
	version.FNbZhuangt = 0                      //	`F_NB_ZHUANGT` int(11) NOT NULL DEFAULT '0' COMMENT '状态 0：正常、1：已删除'
	if err := db.Table("b_dm_ruanjbb").Create(version).Error; err != nil {
		// 错误处理...
		log.Println("Insert b_dm_ruanjbb error", err)
		return err
	}
	log.Println("新增 软件版本表 数据，插入成功！", "软件版本号:=", version.FVcRuanjbbh)
	return nil
}
func QueryOneVersiondata(banbh string) (error, *types.BDmRuanjbb) {
	db := utils.GormClient.Client
	v := new(types.BDmRuanjbb)
	if err := db.Table("b_dm_ruanjbb").Where("F_VC_RUANJBBH =?", banbh).Last(v).Error; err != nil {
		log.Println("查询 软件版本表数据时 error :", err)
		return err, nil
	}
	log.Println("查询软件版本表 数据，成功！")
	return nil, v
}

//查询软件版本列表
func QueryVersionALLdata(req *dto.QueryVersionQeq) (error, *[]types.BDmRuanjbb) {
	db := utils.GormClient.Client
	vs := make([]types.BDmRuanjbb, 0)
	//全部
	mytable := db.Table("b_dm_ruanjbb")
	if req.Version != "" {
		//1、查询设备编号
		mytable = mytable.Where("F_VC_RUANJBBH =?", req.Version)
	}

	if req.BeginTime != "" || req.EndTime != "" {
		//5、查询上传时间
		mytable = mytable.Where("F_DT_SHANGCSJ = ?", req.BeginTime).Where("F_DT_SHANGCSJ = ?", req.EndTime)
	}
	//去除删除的
	if err := mytable.Not("F_NB_ZHUANGT = ?", 1).Find(&vs).Error; err != nil {
		log.Println("查询 软件版本表ALL数据时 error :", err)
		return err, nil
	}
	log.Println("查询软件版本表 数据，成功！数据结果:", "共", len(vs), "个版本")
	return nil, &vs
}

//查询软件版本设备使用数
func QueryVersionNumdata(banbh string) (error, int) {
	db := utils.GormClient.Client
	vs := make([]types.BDmRuanjgxzx, 0)
	//除去删除的软件版本
	//软件更新执行表  状态 0：未完成、1：已完成更新
	if err := db.Table("b_dm_ruanjgxzx").Where("F_VC_RUANJBBH = ?", banbh).Where("F_NB_ZHUANGT = ?", 1).Find(&vs).Error; err != nil {
		if fmt.Sprint(err) == "record not found" {
			log.Println("  err:", err)
			return nil, 0
		} else {
			log.Println("查询 软件更新执行表ALL数据时 error :", err)
			return err, 0
		}
	}
	log.Println("查询软件版本设备使用数，查询软件更新执行表 数据，成功！数据结果:", "共", len(vs), "次数")
	return nil, len(vs)
}

//删除软件版本
func DeleteVersionsdata(req *dto.DeleteVersionQeq) error {
	db := utils.GormClient.Client
	//删除软件版本
	for _, v := range req.Version {
		version := new(types.BDmRuanjbb)
		version.FNbZhuangt = 1 //1表示删除
		if err := db.Table("b_dm_ruanjbb").Where("F_VC_RUANJBBH = ?", v).Update(version).Error; err != nil {
			if fmt.Sprint(err) == "record not found" {
				log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++err:", err)
				continue
			} else {
				log.Println("删除 软件版本表 数据时 error :", err)
				return err
			}
		}
		log.Println("删除软件版本表 数据，成功")
	}
	log.Println("批量删除软件版本表数据，成功", len(req.Version))
	return nil
}

func VersionsUpdatedata(req *dto.VersionUpdateQeq) error {
	db := utils.GormClient.Client
	//要更新的设备软件版本
	for _, v := range req.Gwids {
		//1、查询这个设备，这个软件版本是否记录
		qverr, gxjl := QueryVersionISUpdate(v.Gwid, req.Version)
		if qverr != nil {
			if fmt.Sprint(qverr) == "record not found" {
				log.Println("+++++++++++++++++++++++++++++err:", qverr)
				//不存在，就要插入呀
				version := new(types.BDmRuanjgxzx)
				version.FVcWanggbh = v.Gwid                              //网关编号
				version.FVcRuanjbbh = req.Version                        //软件版本
				version.FNbJihgxcl = req.UpdateStatus                    //计划更新策略
				version.FDtJihgxsj = utils.StrTimeTotime(req.UpdateTime) //计划更新时间
				version.FDtGengxwcsj = utils.StrTimeTotime("2020-01-01 00:00:00")
				if err := db.Table("b_dm_ruanjgxzx").Create(version).Error; err != nil {
					log.Println("插入 软件版本更新表 数据时 error :", err)
					continue
				}
				log.Println("软件版本更新表数据更新插入成功")

			} else {
				log.Println("+++++++++++++++++++++++++++++++++++++[查询软件版本更新表数据失败]+++++++++++++++++++++err==:", qverr)
				continue
			}
		}

		//2、如果已经存在，说明要去查看版本更新是否执行到位
		if gxjl != nil && gxjl.FNbZhuangt == 0 {
			log.Println("++++++++++++++++[查询成功,要去查看版本更新是否执行到位,现在还没有更新完成]+++++++++++++++++++++gxjl.FNbZhuangt==0 ", gxjl.FNbZhuangt)
			continue
		}
	}
	log.Println("软件版本更新表数据更新完成++++++++")
	return nil
}

//查询该网关、该版本是否已经更新
func QueryVersionISUpdate(gwid, versionid string) (error, *types.BDmRuanjgxzx) {
	db := utils.GormClient.Client
	v := new(types.BDmRuanjgxzx)
	//查询是否存在如果已经存在，说明已经在更新中或者已经更新成功
	if err := db.Table("b_dm_ruanjgxzx").Where("F_VC_WANGGBH=?", gwid).Where("F_NB_BANBENID=?", versionid).Last(&v).Error; err != nil {
		log.Println("查询 软件更新执行 error :", err)
		return err, nil
	}
	log.Println("查询软件更新执行表 数据，成功！数据结果:", v.FVcWanggbh, v.FVcRuanjbbh, "v.FNbZhuangt(1:ok  0:ing):", v.FNbZhuangt)
	return nil, v
}

//查询软件版本列表下拉框
func QueryVersionALL() (error, *[]types.BDmRuanjbb) {
	db := utils.GormClient.Client
	vs := make([]types.BDmRuanjbb, 0)
	//全部
	//除去删除的软件版本
	//按时间排序
	if err := db.Table("b_dm_ruanjbb").Not("F_NB_ZHUANGT = ?", 1).Order("F_DT_SHANGCSJ desc").Find(&vs).Error; err != nil {
		log.Println("查询 软件版本表ALL数据时 error :", err)
		return err, nil
	}
	log.Println("查询软件版本表 数据，成功！数据结果:", "共", len(vs), "个版本")
	return nil, &vs
}

//查询设备列表下拉框
func QueryGatewayALL() (error, *[]types.BDmWanggjcxx) {
	db := utils.GormClient.Client
	gws := make([]types.BDmWanggjcxx, 0)
	//全部
	if err := db.Table("b_dm_wanggjcxx").Find(&gws).Error; err != nil {
		log.Println("查询 设备列表ALL数据时 error :", err)
		return err, nil
	}
	log.Println("查询设备列表 数据，成功！数据结果:", "共", len(gws), "个网关设备")
	return nil, &gws
}

//查询停车场下拉框
func QueryParkNameALL() (error, *[]types.BTccTingcc) {
	db := utils.GormClient.Client
	tccs := make([]types.BTccTingcc, 0)
	//全部
	if err := db.Table("b_tcc_tingcc").Find(&tccs).Error; err != nil {
		log.Println("查询 停车场列表ALL数据时 error :", err)
		return err, nil
	}
	log.Println("查询停车场列表 数据，成功！数据结果:", "共", len(tccs), "个停车场")
	return nil, &tccs
}

func QueryParkName(parkid string) (error, *types.BTccTingcc) {
	db := utils.GormClient.Client
	tcc := new(types.BTccTingcc)
	//全部
	if err := db.Table("b_tcc_tingcc").Where("F_VC_TINGCCBH=?", parkid).First(&tcc).Error; err != nil {
		log.Println("查询 停车场列表ALL数据时 error :", err)
		return err, nil
	}
	log.Println("查询停车场列表 数据，成功！数据结果:", tcc.FVcTingccbh, tcc.FVcMingc)
	return nil, tcc
}

//获取告警的最新告警时间
func QueryAlarm() (error, *types.BDmGaoj) {
	db := utils.GormClient.Client
	gjs := make([]types.BDmGaoj, 0)
	if err := db.Table("b_dm_gaoj").Order("F_DT_GAOJSJ desc").Limit(1).Find(&gjs).Error; err != nil {
		log.Println("查询 获取告警的最新告警时间 error :", err)
		return err, nil
	}
	log.Println("查询 告警信息表ALL数据")
	return nil, &gjs[0]
}

//告警信息新增插入
func GatewayErrorInsert(gaoj *types.BDmGaoj) error {
	db := utils.GormClient.Client
	if err := db.Table("b_dm_gaoj").Create(gaoj).Error; err != nil {
		// 错误处理...
		log.Println("Insert b_dm_gaoj error:", err)
		return err
	}
	log.Println("新增 告警信息表 数据，插入成功！", "网关编号:=", gaoj.FVcWanggbh)
	return nil
}

//告警信息查询是否已经插入过了
func QueryGatewayError(gwid, time string) error {
	db := utils.GormClient.Client
	gaoj := new(types.BDmGaoj)
	if err := db.Table("b_dm_gaoj").Where("F_VC_WANGGBH=?", gwid).Where("F_DT_GAOJSJ=?", time).Last(gaoj).Error; err != nil {
		// 错误处理...
		log.Println("query b_dm_gaoj error:", err)
		return err
	}
	log.Println("查询 告警信息表 数据，插入成功！", "网关编号:=", gaoj.FVcWanggbh)
	return nil
}

//重启信息新增插入
func GatewayRestarInsert(Restar *types.BDmChongq) error {
	db := utils.GormClient.Client
	if err := db.Table("b_dm_chongq").Create(Restar).Error; err != nil {
		// 错误处理...
		log.Println("Insert b_dm_chongq error:", err)
		return err
	}
	log.Println("新增 重启信息表 数据，插入成功！", "网关编号:=", Restar.FVcWanggbh)
	return nil
}

//重启信息查询是否已经插入过了
func QueryGatewayRestar(gwid, time string) error {
	db := utils.GormClient.Client
	Restar := new(types.BDmChongq)
	if err := db.Table("b_dm_chongq").Where("F_VC_WANGGBH=?", gwid).Where("F_DT_CHONGQSJ=?", time).Last(Restar).Error; err != nil {
		// 错误处理...
		log.Println("query b_dm_chongq error:", err)
		return err
	}
	log.Println("查询 重启信息表 数据，插入成功！", "网关编号:=", Restar.FVcWanggbh)
	return nil
}

//网关设备重启信息查询最新时间
func GatewayRestarNewTime(gwid string) (error, *types.BDmChongq) {
	db := utils.GormClient.Client
	cqs := make([]types.BDmChongq, 0)
	if err := db.Table("b_dm_chongq").Where("F_VC_WANGGBH=?", gwid).Order("F_DT_CHONGQSJ desc").Limit(1).Find(&cqs).Error; err != nil {
		// 错误处理...
		log.Println(" query b_dm_chongq error:", err)
		return err, nil
	}
	log.Println("查询 重启信息表 数据，插入成功！", "网关编号:=", cqs[0].FVcWanggbh)
	return nil, &cqs[0]
}
