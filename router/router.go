package router

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"operationPlatform/controller"

	"strings"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RouteInit(IpAddress string) {
	logrus.Print("服务端 IpAddress：", IpAddress)
	router := gin.New()
	router.Use(Cors()) //跨域资源共享

	url := ginSwagger.URL("http://127.0.0.1:8077/swagger/doc.json")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	apiV1 := router.Group("/operationplatform/api/v1")
	APIV1Init(apiV1)

	http.Handle("/", router)
	gin.SetMode(gin.ReleaseMode)

	runerr := router.Run(IpAddress)
	if runerr != nil {
		logrus.Print("Run error", runerr)
		return
	}
}
func APIV1Init(route *gin.RouterGroup) {
	AuthAPIInit(route)
}

func AuthAPIInit(route *gin.RouterGroup) {
	//用户注册
	//route.POST("/user/register", controller.Register)
	//用户登录
	route.GET("/user/imagecaptcha", controller.Imagecaptcha)
	route.POST("/user/login", controller.Login)

	//网关运维监控平台
	//1、网关列表查询
	route.POST("/gw/querygatewaylist", controller.Querygatewaylist)
	//2、告警列表查询
	route.POST("/gw/queryalarmlist", controller.QueryAlarmlist)
	//3、重启记录列表查询
	route.POST("/gw/restartrecordlist", controller.QueryRestartRecordlist)
	//4、天线列表查询
	route.POST("/gw/rsulist", controller.QueryRSURecordlist)
	// 5、网关设备详情查询
	route.POST("/gw/gatewaydevicedetails", controller.QueryGatewayDeviceDetails)

	//6、添加设备[1]
	route.POST("/gw/addgatewaydevice", controller.Addgatewaydevice)
	//7、远程连接  [暂时不做]

	//软件版本管理Version management
	//8、增加网关软件更新的版本【上传固件之增加版本记录到数据库】
	route.POST("/version/addversion", controller.AddNewVersion)
	//9、查询软件版本列表
	route.POST("/version/querygatewayversionlist", controller.QuerygatewayVersionlist)

	//11、删除软件版本【可以批量删除】
	route.POST("/version/deleteversion", controller.DeleteNewVersion)
	//12、上传固件之上传版本文件
	route.POST("/version/fileupload", controller.UploadVersionFile)

	//13、网关软件更新执行【？？？】
	route.POST("/version/update", controller.VersionUpdate)
	route.POST("/version/performupdate", controller.PerformVersionUpdate)

	//下拉框
	//软件版本下拉框
	route.GET("/version/queryversionlist", controller.QueryVersionlist)
	//网关设备下拉框
	route.GET("/version/querygatewaylist", controller.QueryGatewaylist)
	//停车场下拉框
	route.GET("/version/queryparknamelist", controller.QueryparkNamelist)
}

//以下为cors实现
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method               //请求方法
		origin := c.Request.Header.Get("Origin") //请求头部
		var headerKeys []string                  // 声明请求头keys
		for k := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Origin", "*") // 这是允许访问所有域

			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE") //服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
			//  header的类型
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			//              允许跨域设置                                                                                                      可以返回其他子段

			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析

			c.Header("Access-Control-Max-Age", "172800")          // 缓存请求信息 单位为秒
			c.Header("Access-Control-Allow-Credentials", "false") //  跨域请求是否需要带cookie信息 默认设置为true
			c.Set("content-type", "application/json")             // 设置返回格式是json
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		c.Next() //  处理请求
	}
}
