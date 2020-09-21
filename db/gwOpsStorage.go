package db

import (
	log "github.com/sirupsen/logrus"
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

//2、 Query网关信息 根据网关编号  b_dm_wanggjcxx
func QueryGatewaydata(FVcWanggbh string) (error, *types.BDmWanggjcxx) {
	db := utils.GormClient.Client
	gwxx := new(types.BDmWanggjcxx)
	//赋值
	if err := db.Table("b_dm_wanggjcxx").Where("F_VC_WANGGBH=?", FVcWanggbh).Last(gwxx).Error; err != nil {
		log.Println("查询 结算监控统计表最新数据时 QueryTabledata error :", err)
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
