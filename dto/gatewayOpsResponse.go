package dto

//网关运维监控响应对象

type QueryResponse struct {
	Code    int `json:"code"  example:"200"` //3000
	CodeMsg string
	Data    interface{} `json:"data"`
	Message string      `json:"message" example:"响应成功信息"`
}

//网关运维成功响应
type GatewayOPSResponse struct {
	Code    int         `json:"code"  example:"200"`
	Data    interface{} `json:"data"`
	Message string      `json:"message" example:"响应成功信息"`
}

//网关运维失败响应
type GatewayOPSResponseFailure struct {
	Code    int         `json:"code"  example:"404"`
	Data    interface{} `json:"data"`
	Message string      `json:"message" example:"响应失败信息"`
}

//网关设备详情
type GatewayDeviceDetails struct {
	GatewayNumber int         `json:"gw_number"  example:"200"`
	Data          interface{} `json:"data"`
	Message       string      `json:"message" example:"响应成功信息"`
}
