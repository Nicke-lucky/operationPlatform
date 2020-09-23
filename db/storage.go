package db

import (
	log "github.com/sirupsen/logrus"
	"operationPlatform/config"
	"operationPlatform/types"
	"operationPlatform/utils"

	"time"
)

//结算监控平台数据层：数据的增删改查
func Newdb() {
	conf := config.ConfigInit() //初始化配置
	utils.InitLogrus(conf.LogPath, conf.LogFileName, time.Duration(24*conf.LogMaxAge)*time.Hour, time.Duration(conf.LogRotationTime)*time.Hour)
	mstr := conf.MUserName + ":" + conf.MPass + "@tcp(" + conf.MHostname + ":" + conf.MPort + ")/" + conf.Mdatabasename + "?charset=utf8&parseTime=true&loc=Local"
	DBInit(mstr) //初始化数据库
	Errormsg_address = conf.Errormsg_address
	Gwmsg_address = conf.Gwmsg_address
	Metric_address = conf.Metric_address
}

//1、查询表是否存在
func QueryTable(tablename string) {
	db := utils.GormClient.Client
	is := db.HasTable(tablename)
	if is == false {
		log.Println("不存在", tablename)
		return
	}
	log.Println("表存在：", tablename, is)
}

//1、Insert b_jsjk_jiestj 新增结算统计
func InsertTabledata(lx int) error {
	db := utils.GormClient.Client
	Jiestj := new(types.BJsJiessj)
	//赋值
	Jiestj.FNbJusbj = lx //统计类型 10000 ：省外
	if err := db.Table("b_js_jiessj").Create(&Jiestj).Error; err != nil {
		// 错误处理...
		log.Println("Insert b_js_jiessj error", err)
		return err
	}
	log.Println("省外-结算统计表插入成功！", "开始统计时间:=", Jiestj)
	return nil
}

//2、 Query b_jsjk_jiestj
func QueryTabledata(lx int) (error, *types.BJsJiessj) {
	db := utils.GormClient.Client
	//Jiestjs := make([]types.BJsjkJiestj, 0)
	Jiestjs := new(types.BJsJiessj)
	//赋值
	if err := db.Table("b_js_jiessj").Where("F_NB_KAWLH=?", lx).Last(&Jiestjs).Error; err != nil {
		log.Println("查询 结算监控统计表最新数据时 QueryTabledata error :", err)
		return err, nil
	}
	log.Println("查询结算监控统计表最新数据结果:", Jiestjs)
	return nil, Jiestjs
}

//3、更新结算统计表 update b_jsjk_jiestj
func UpdateTabledata(data *types.BJsJiessj, lx int, id int) error {
	db := utils.GormClient.Client
	Jiestj := new(types.BJsJiessj)

	if err := db.Table("b_js_jiessj").Where("F_NB_ID=?", id).Where("F_NB_KAWLH=?", lx).Updates(&Jiestj).Error; err != nil {
		log.Println("更新结算统计表 error", err)
		return err
	}
	return nil
}
