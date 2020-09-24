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
	//GatewayNumber    //设备编号 默认全部：0
	//ParkName         //停车场名称 默认全部：0
	//Status           //状态：2全部，1在线、0离线
	//Version          //软件版本
	//UpdateBeginTime  //起始时间
	//UpdateEndTime    //结束时间

	//全部
	if req.GatewayNumber == "0" && req.ParkName == "0" && req.Status == 2 && req.Version == "0" && req.UpdateEndTime == "0" && req.UpdateBeginTime == "0" {
		if err := db.Table("b_dm_wanggjcxx").Find(&gwxxs).Error; err != nil {
			log.Println("查询 网关基础信息表ALL数据时 error :", err)
			return err, nil
		}
		log.Println("查询网关基础信息表 数据，成功！数据结果:", "共", len(gwxxs), "个设备")
		return nil, &gwxxs
	} else {

		//1、查询设备编号
		if req.GatewayNumber != "0" && req.ParkName == "0" && req.Status == 2 && req.Version == "0" && req.UpdateEndTime == "0" && req.UpdateBeginTime == "0" {
			if err := db.Table("b_dm_wanggjcxx").Where("F_VC_WANGGBH = ?", req.GatewayNumber).Find(&gwxxs).Error; err != nil {
				log.Println("查询 网关基础信息表ALL数据时 error :", err)
				return err, nil
			}
			log.Println("查询网关基础信息表 数据，成功！数据结果:", "共", len(gwxxs), "个设备")
			return nil, &gwxxs
		} else {

			//2、查询停车场
			if req.GatewayNumber == "0" && req.ParkName != "0" && req.Status == 2 && req.Version == "0" && req.UpdateEndTime == "0" && req.UpdateBeginTime == "0" {
				if err := db.Table("b_dm_wanggjcxx").Where("FVcTingccbh = ?", req.ParkName).Find(&gwxxs).Error; err != nil {
					log.Println("查询 网关基础信息表ALL数据时 error :", err)
					return err, nil
				}
				log.Println("查询网关基础信息表 数据，成功！数据结果:", "共", len(gwxxs), "个设备")
				return nil, &gwxxs
			} else {
				//3、查询状态[]
				if req.GatewayNumber == "0" && req.ParkName == "0" && req.Status != 2 && req.Version == "0" && req.UpdateEndTime == "0" && req.UpdateBeginTime == "0" {
					if err := db.Table("b_dm_wanggjcxx").Where("F_NB_ZHUANGT = ?", req.Status).Find(&gwxxs).Error; err != nil {
						log.Println("查询 网关基础信息表ALL数据时 error :", err)
						return err, nil
					}
					log.Println("查询网关基础信息表 数据，成功！数据结果:", "共", len(gwxxs), "个设备")
					return nil, &gwxxs
				} else {
					//4、查询版本【】
					if req.GatewayNumber == "0" && req.ParkName == "0" && req.Status == 2 && req.Version != "0" && req.UpdateEndTime == "0" && req.UpdateBeginTime == "0" {
						if err := db.Table("b_dm_wanggjcxx").Where("F_VC_DANGQBBH = ?", req.Version).Find(&gwxxs).Error; err != nil {
							log.Println("查询 网关基础信息表ALL数据时 error :", err)
							return err, nil
						}
						log.Println("查询网关基础信息表 数据，成功！数据结果:", "共", len(gwxxs), "个设备")
						return nil, &gwxxs
					} else {
						//5、查询时间【】
						if req.GatewayNumber == "0" && req.ParkName == "0" && req.Status == 2 && req.Version == "0" && req.UpdateEndTime != "0" && req.UpdateBeginTime != "0" {
							if err := db.Table("b_dm_wanggjcxx").Where("F_DT_ZUIHGXSJ >=", req.UpdateBeginTime+" 00:00:00").Where("F_DT_ZUIHGXSJ <=", req.UpdateEndTime+" 23:59:59").Find(&gwxxs).Error; err != nil {
								log.Println("查询 网关基础信息表ALL数据时 error :", err)
								return err, nil
							}
							log.Println("查询网关基础信息表 数据，成功！数据结果:", "共", len(gwxxs), "个设备")
							return nil, &gwxxs
						} else {
							//6、停车场、状态
							if req.GatewayNumber == "0" && req.ParkName != "0" && req.Status != 2 && req.Version == "0" && req.UpdateEndTime == "0" && req.UpdateBeginTime == "0" {
								if err := db.Table("b_dm_wanggjcxx").Where("F_NB_ZHUANGT = ?", req.Status).Where("FVcTingccbh = ?", req.ParkName).Find(&gwxxs).Error; err != nil {
									log.Println("查询 网关基础信息表ALL数据时 error :", err)
									return err, nil
								}
								log.Println("查询网关基础信息表 数据，成功！数据结果:", "共", len(gwxxs), "个设备")
								return nil, &gwxxs
							} else {
								//7、全部
								if err := db.Table("b_dm_wanggjcxx").Find(&gwxxs).Error; err != nil {
									log.Println("查询 网关基础信息表ALL数据时 error :", err)
									return err, nil
								}
								log.Println("查询网关基础信息表 数据，成功！数据结果:", "共", len(gwxxs), "个设备")
								return nil, &gwxxs
							}
						}
					}
				}
			}
		}

	}
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
	if err := db.Table("b_dm_ruanjbb").Where("F_VC_RUANJBBH =?", banbh).First(v).Error; err != nil {
		log.Println("查询 软件版本表数据时 error :", err)
		return err, nil
	}
	log.Println("查询软件版本表 数据，成功！")
	return nil, v
}

//查询软件版本列表
func QueryVersionALLdata() (error, *[]types.BDmRuanjbb) {
	db := utils.GormClient.Client
	vs := make([]types.BDmRuanjbb, 0)
	//除去删除的软件版本
	if err := db.Table("b_dm_ruanjbb").Not("F_NB_ZHUANGT =?", 1).Find(&vs).Error; err != nil {
		log.Println("查询 软件版本表ALL数据时 error :", err)
		return err, nil
	}
	log.Println("查询软件版本表 数据，成功！数据结果:", "共", len(vs), "个版本")
	return nil, &vs
}

//查询软件版本更新次数
func QueryVersionNumdata(banbh string) (error, int) {
	db := utils.GormClient.Client
	vs := make([]types.BDmRuanjgxzx, 0)
	//除去删除的软件版本
	if err := db.Table("b_dm_ruanjgxzx").Where("F_VC_RUANJBBH =?", banbh).Where("F_NB_ZHUANGT=?", 1).Find(&vs).Error; err != nil {
		if fmt.Sprint(err) == "record not found" {
			log.Println("  err== `record not found`:", err)
			return nil, 0
		} else {
			log.Println("查询 软件版本表ALL数据时 error :", err)
			return err, 0
		}
	}
	log.Println("查询软件版本表 数据，成功！数据结果:", "共", len(vs), "次数")
	return nil, len(vs)
}

//删除软件版本
func DeleteVersionsdata(req *dto.DeleteVersionQeq) error {
	db := utils.GormClient.Client
	//删除软件版本
	for _, v := range req.Version {
		version := new(types.BDmRuanjbb)
		version.FNbZhuangt = 1
		if err := db.Table("b_dm_ruanjgxzx").Where("F_VC_RUANJBBH =?", v).Update(version).Error; err != nil {
			if fmt.Sprint(err) == "record not found" {
				log.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++err== `record not found`:", err)
				return nil
			} else {
				log.Println("删除 软件版本表 数据时 error :", err)
				return err
			}
		}
		log.Println("删除软件版本表 数据，成功")
	}
	return nil
}
