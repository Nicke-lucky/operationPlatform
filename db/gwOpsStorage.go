package db

import (
	log "github.com/sirupsen/logrus"
	"operationPlatform/dto"
	"operationPlatform/types"
	"operationPlatform/utils"
)

//

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
	if err := db.Table("b_dm_wanggjcxx").Where("F_VC_WANGGBH=?", FVcWanggbh).Last(gwxx).Error; err != nil {
		log.Println("查询 网关基础信息表最新数据时 QueryTabledata error :", err)
		return err, nil
	}
	log.Println("查询网关基础信息表 数据，成功！数据结果:", gwxx)
	return nil, gwxx
}

//3、更新网关信息表 根据网关编号
func UpdateGatewaydata(Wanggbh string, gwdata *types.BDmWanggjcxx) error {
	db := utils.GormClient.Client
	//gwxx := new( types.BDmWanggjcxx)
	if err := db.Table("b_dm_wanggjcxx").Where("F_VC_WANGGBH=?", Wanggbh).Updates(gwdata).Error; err != nil {
		log.Println("更新网关基础信息表 error", err)
		return err
	}
	log.Println("更新网关基础信息表 ok !")
	return nil
}

//4、查询网关信息多条数据【所有】
func QueryGatewayALLdata(req *dto.QueryGatewayListQeqdata) (error, *[]types.BDmWanggjcxx) {
	db := utils.GormClient.Client
	gwxxs := make([]types.BDmWanggjcxx, 0)
	log.Println("req:", req)

	//1、校验参数 默认选择全部
	//GatewayNumber    //设备编号 网关编号 默认全部：0
	//ParkName           //停车场名称 默认全部：0
	//Status                //状态：2全部，1在线、0离线
	//Version                //软件版本
	//UpdateBeginTime          //起始时间
	//UpdateEndTime            //结束时间
	if err := db.Table("b_dm_wanggjcxx").Find(&gwxxs).Error; err != nil {
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

	//1、校验参数 默认选择全部
	//网关设备id

	if err := db.Table("b_dm_wanggjcxx").Where("F_VC_WANGGBH =?", req.TerminalId).Find(&gjs).Error; err != nil {
		log.Println("查询 告警信息表ALL数据时 error :", err)
		return err, nil
	}
	log.Println("查询告警表 数据，成功！数据结果:", "共", len(gjs), "个告警")
	return nil, &gjs
}

//查询重启信息
func QueryRestartALLdata(req *dto.QueryRestartMsgListQeq) (error, *[]types.BDmChongq) {
	db := utils.GormClient.Client
	gjs := make([]types.BDmChongq, 0)
	log.Println("req:", req)

	//1、校验参数 默认选择全部
	//	TerminalId string `json:"terminal_id"` // 设备ID，如CE4C37043A520C93	网关设备id
	//	BeginTime  string `json:"Begin_time"`  //重启列表请求起始时间
	//	EndTime    string `json:"end_time"`    //重启列表请求结束时间

	if err := db.Table("b_dm_chongq").Where("F_VC_WANGGBH =?", req.TerminalId).Find(&gjs).Error; err != nil {
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
	if err := db.Table("b_dm_chongq").Where("F_VC_WANGGBH =?", TerminalId).Last(&cq).Error; err != nil {
		log.Println("查询 重启信息表One数据时 error :", err)
		return err, nil
	}
	log.Println("查询重启信息表 数据，成功！数据结果:")
	return nil, cq
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
	log.Println("查询网关基础信息表 数据，成功！数据结果:")
	return nil, gwxx
}
