package db

import (
	"log"
	"testing"
)

func TestNewTables(t *testing.T) {
	Newdb()
	//NewTables()
}
func TestQueryTable(t *testing.T) {
	//数据库生成的表名是结构体名称的复数形式
	s1 := []string{"b_jsjk_jiesjkptyhb",
		"b_jsjk_zhuanjssjjk",
		"b_jsjk_yuqsjtj",
		"b_jsjk_yicsjtj",
		"b_jsjk_yicsjtcctj",
		"b_jsjk_tingccjssjtj",
		"b_jsjk_shujtbjk",
		"b_jsjk_shujbjk",
		"b_jsjk_shengwtccjsqs",
		"b_jsjk_shengwqftj",
		"b_jsjk_shengwjszysjtj",
		"b_jsjk_shengwjssjfl",
		"b_jsjk_shengwjsqs",
		"b_jsjk_shengnyfssjtj",
		"b_jsjk_shengntccjsqs",
		"b_jsjk_shengnsssjjk",
		"b_jsjk_shengnqktj",
		"b_jsjk_shengnjssjfl",
		"b_jsjk_shengnjsqs",
		"b_jsjk_heimdjk",
		"b_jsjk_shengnjfsjtj",
		"b_jsjk_qingfhd",
		"b_jsjk_jiestj",
		"JieSuanWssjs",
	}

	tablenames := make([]string, 0)
	tablenames = append(tablenames, s1...)
	log.Println(len(tablenames), tablenames)
	Newdb()
	for i, tablename := range tablenames {
		log.Println(1+i, tablename)
		QueryTable(tablename)
	}
}

//插入数据
func TestInsertTabledata(t *testing.T) {
	Newdb()
	InsertTabledata(10000)
}

//
func TestQueryTabledata(t *testing.T) {
	Newdb()
	QueryTabledata(10000)
}

func TestUpdateTabledata(t *testing.T) {
	Newdb()
	//data := &types.BJsjkJiestj{FNbId: 19, FNbZongts: 9999, FNbZongje: 90000}
	//UpdateTabledata(data, 10000, 19)
}
