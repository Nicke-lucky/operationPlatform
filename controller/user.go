package controller

import (
	"fmt"
	"net/http"
	"operationPlatform/dto"
	"operationPlatform/service"
	"operationPlatform/types"
	"operationPlatform/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*  接口方法注释   */
//@Summary 注册api
//@Tags 注册
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /user/register [post]
/*func Register(c *gin.Context) {
	req := dto.ReqRegister{}
	respFailure := dto.ResponseFailure{}

	if err := c.Bind(&req); err != nil {
		logrus.Errorf("ReqRegister json unmarshal err: %v", err)
		respFailure.Code = -1
		respFailure.Message = fmt.Sprintf("json unmarshal err: %v", err)
		c.JSON(http.StatusOK, respFailure)
		return
	}
	//注册处理
	code, err := service.Register(req)
	if err != nil {
		logrus.Errorf("Register  err: %v", err)
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("Register err: %v", err)
		c.JSON(http.StatusOK, respFailure)
		return
	}

	//401 返回的一个状态
	if code == types.StatusRepeatedRegistration {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusRepeatedRegistration, Data: types.StatusText(types.StatusRepeatedRegistration), Message: "重复注册"})
	}
	if code == types.StatusRegisteredSuccessfully {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Message: "注册成功"})
	}
}*/

/*  接口方法注释   */
//@Summary 登录api
//@Tags 登录
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /user/login [post]
func Login(c *gin.Context) {
	req := dto.Reqlogin{}
	respFailure := dto.ResponseFailure{}

	if err := c.Bind(&req); err != nil {
		logrus.Println("Login json unmarshal err: %v", err)
		respFailure.Code = -1
		respFailure.Message = fmt.Sprintf("json unmarshal err: %v", err)
	}
	//登录处理
	code, err := service.Login(req)
	if err != nil {
		respFailure.Code = code
		respFailure.Message = fmt.Sprintf("登录时error:%v", err)
		logrus.Println("登录时error：", respFailure.Message)
		c.JSON(http.StatusOK, dto.Response{Code: respFailure.Code, Data: "登录时error", Message: "登录时error"})
		return
	}
	if code == types.StatusGETRedisError {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusPleaseRegister, Data: types.StatusText(types.StatusPleaseRegister), Message: "用户名输入错误,请重新输入"})
		return
	}

	if code == types.StatusPleaseRegister {
		logrus.Println("用户名输入错误,请重新输入")
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusPleaseRegister, Data: types.StatusText(types.StatusPleaseRegister), Message: "用户名输入错误,请重新输入"})
		return
	}

	if code == types.StatusPasswordError {
		logrus.Println("密码错误,请重新输入")
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusPasswordError, Data: types.StatusText(types.StatusPasswordError), Message: "密码错误,请重新输入"})
		return
	}
	if code == types.StatusNoVerificationcode {
		logrus.Println("验证码错误,请重新输入")
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusNoVerificationcode, Data: types.StatusText(types.StatusNoVerificationcode), Message: "验证码错误,请重新输入"})
		return
	}
	if code == types.StatusSuccessfully {
		logrus.Println("用户登录成功")
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: types.StatusText(types.StatusSuccessfully), Message: "用户登录成功"})
		return
	}
}

/*  接口方法注释   */
//@Summary 图片验证码api
//@Tags 图片验证码
//@version 1.0
//@Accept application/json
//@Param req body dto.Reqlogin true "请求参数"
//@Success 200 object dto.Response 成功后返回值
//@Failure 404 object dto.ResponseFailure 查询失败
//@Router /user/login [post]
func Imagecaptcha(c *gin.Context) {

	//1.随机获取验证码文字
	randStr := utils.GetRandStr(4)
	logrus.Println("随机获取验证码文字:", randStr)
	//2、验证码存redis
	conn := utils.Pool.Get()
	defer conn.Close()
	rseterr := utils.RedisSet(&conn, randStr, randStr)
	if rseterr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Message: "随机获取验证码文字 时，set redis error"})
		return
	}
	rserr := utils.RedisExpireSet(&conn, randStr, 300)
	if rserr != nil {
		c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Message: "设置过期时间 时，RedisExpireSet  error"})
		return
	}
	c.JSON(http.StatusOK, dto.Response{Code: types.StatusSuccessfully, Data: randStr, Message: "随机获取验证码文字成功"})
}
