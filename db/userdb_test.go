package db

import (
	"github.com/sirupsen/logrus"
	"operationPlatform/types"
	"operationPlatform/utils"

	"testing"
)

//用户直接注册
func TestUserInsert(t *testing.T) {
	Newdb()
	user := types.BSysYongh{}
	user.FVcId = "admin5564512"
	user.FVcZhangh = "1324312admin55645"
	psw := utils.GetMD5Encode("123")
	user.FVcMim = psw
	user.FVcMingc = "1413运维管理平台1"
	user.FVcGongsid = "4123gandong155551"
	err := UserInsert(&user)
	if err != nil {
		logrus.Print("用户注册失败", err)
	} else {
		logrus.Print("用户注册ok")

	}

}

//
func TestQueryUsermsg(t *testing.T) {
	Newdb()
	err, resp := QueryUserLoginmsg("1324312admin55645")
	if err != nil {
		logrus.Print("查询用户能否被注册，失败", err)
	}
	logrus.Printf("查询用户已经被注册 %s,%s,%s,%s", resp.FVcGongsid, resp.FVcMingc, resp.FVcZhangh, resp.FVcMim)

}
