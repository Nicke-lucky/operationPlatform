package db

import (
	"github.com/sirupsen/logrus"

	"testing"
)

//用户直接注册
func TestUserInsert(t *testing.T) {
	Newdb()
	user := types.BJsjkJiesjkptyhb{}
	user.FVcYonghmm = "abc111"
	user.FVcShoujh = "123123"
	user.FVcYonghnc = "abc111"
	user.FNbYonghid = "10001"
	user.FVcYoux = "1233@123"
	err := UserInsert(&user)
	if err != nil {
		logrus.Print("用户注册失败", err)
	}
}

//
func TestQueryUsermsg(t *testing.T) {
	Newdb()
	err, resp := QueryUsermsg("10001")
	if err != nil {
		logrus.Print("查询用户能否被注册，失败", err)
	}
	logrus.Print("查询用户能否被注册 ", resp)

}
