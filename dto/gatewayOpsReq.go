package dto

//网关运维监控请求对象

//添加网关设备请求信息
type GatewayDevicedata struct {
	GatewayNumber string `json:"gw_number"  example:"gw200abc"` //设备编号
	ParkName      string `json:"park_name"`                     //停车场名称
	Note          string `json:"note" example:"备注"`
}

//查询网关列表请求信息
type QueryGatewayListQeqdata struct {
	GatewayNumber   string `json:"gw_number"  example:"gw200abc"` //设备编号 网关编号
	ParkName        string `json:"park_name"`                     //停车场名称
	Status          int    `json:"park_name"`                     //状态：2全部，1在线、0离线
	Version         string `json:"version"`                       //软件版本
	UpdateBeginTime string `json:"updateBegin_time"`              //更新时间
	UpdateEndTime   string `json:"updateEnd_time"`
}
