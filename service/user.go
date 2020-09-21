package service

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"log"
	"operationPlatform/db"
	"operationPlatform/dto"
	"operationPlatform/types"
	"operationPlatform/utils"
)

//登录
func Login(req dto.Reqlogin) (int, error) {
	logrus.Print("登录请求参数：", req)
	//获取用户数据
	err, jg := db.QueryUserLoginmsg(req.UserName)
	//校验密码
	if "record not found" == fmt.Sprintf("%v", err) {
		logrus.Println("用户名不正确错误")
		return types.StatusPleaseRegister, nil
	}
	if err != nil {
		//查询用户是否被注册，查询失败
		return types.StatusPleaseRegister, err
	}
	psw := utils.GetMD5Encode(req.Password)
	logrus.Println("md5 pwd:", psw)
	logrus.Println("jg.FVcMim:", jg.FVcMim)

	if jg.FVcMim != psw {
		logrus.Println("密码错误")
		return types.StatusPasswordError, nil
	}

	conn := utils.Pool.Get()
	defer conn.Close()

	rgeterr, code := utils.RedisGet(&conn, req.Verificationcode)
	if rgeterr != nil {
		logrus.Println("RedisGet错误")
		return types.StatusGETRedisError, errors.New("RedisGet错误")
	}

	if code == nil {
		log.Println("验证码不存在")
		return types.StatusNoVerificationcode, nil
	}
	vstr := string(code.([]uint8))

	if vstr == req.Verificationcode {
		logrus.Println("验证码正确，登录成功")
		//删除验证码
		delerr := utils.RedisDelete(&conn, req.Verificationcode)
		if delerr != nil {
			logrus.Println("RedisDelete error")
			return types.StatusDELRedisError, delerr
		}
	} else {
		return types.StatusNoVerificationcode, errors.New("验证码不正确")
	}
	logrus.Println("密码正确")
	//返回数据
	return types.StatusSuccessfully, nil
}

/*
//用户注册
func Register(req dto.ReqRegister) (int, error) {
	logrus.Print("用户注册请求参数：", req)
	//获取请求数据
	err, jg := db.QueryUsermsg(req.UserName)
	if err != nil {
		//查询用户是否被注册，查询失败
		return types.StatusRegistError, err
	}
	//校验数据
	if jg > 0 {
		logrus.Println("重复注册")
		return types.StatusRepeatedRegistration, nil
	}
	//data := new(types.BJsjkJiesjkptyhb)
	//
	//data.FNbYonghid = req.UserName //   '用户id',//手机号 或者任意6位数字
	//data.FVcYonghmm = req.Password //   '用户密码',
	//data.FVcShoujh = req.Num       //   '手机号',
	//data.FVcYoux = req.Email       //   '邮箱',
	//data.FVcYonghnc = req.Name     // '用户昵称',
	inerr := db.UserInsert(data)
	if inerr != nil {
		logrus.Println("db.UserInsert error!")
		return types.StatusRegistError, inerr
	}
	logrus.Println("注册成功")
	//返回数据
	return types.StatusRegisteredSuccessfully, nil
}
*/
