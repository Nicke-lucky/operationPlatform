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
	GatewayNumber   string `json:"gw_number"`         //设备编号 网关编号
	ParkName        string `json:"park_name"`         //停车场名称
	Status          int    `json:"status"`            //状态：2全部，1在线、0离线
	Version         string `json:"version"`           //软件版本
	UpdateBeginTime string `json:"update_begin_time"` //更新时间
	UpdateEndTime   string `json:"update_end_time"`
}

type GatewayDataReq struct {
}

//滴滴监控夜莺查询告警信息
type QueryErrorMsgQeq struct {
	BeginTime int64 `json:"startetime"` //查询告警的开始时间戳
	EndTime   int64 `json:"endetime"`   //查询告警的结束时间戳
}

//告警信息列表请求信息
type QueryErrorMsgListQeq struct {
	TerminalId string `json:"terminal_id"` // 设备ID，如CE4C37043A520C93	//网关id
	Status     int    `json:"status"`      //处理状态 0：未处理的告警，1已处理的告警，2全部展示
	BeginTime  string `json:"Begin_time"`  //起始时间，告警 时间
	EndTime    string `json:"end_time"`    //结束时间，告警 时间
}

//重启列表请求信息
type QueryRestartMsgListQeq struct {
	TerminalId string `json:"terminal_id"` // 设备ID，如CE4C37043A520C93	//网关id
	BeginTime  string `json:"Begin_time"`  //重启列表请求起始时间
	EndTime    string `json:"end_time"`    //重启列表请求结束时间
}

//天线列表请求信息
type QueryRSUMsgListQeq struct {
	TerminalId string `json:"terminal_id"` // 设备ID，如CE4C37043A520C93	//网关id
}

//网关设备详情请求信息
type QueryGatewayOneQeqdata struct {
	TerminalId string `json:"terminal_id"` // 设备ID，如CE4C37043A520C93	//网关id
}

//增减软件版本信息的请求信息
type AddGatewayVersionQeq struct {
	Version     string `json:"version"`      // 软件版本号
	VersionNote string `json:"version_note"` // 软件版本内容
	Name        string `json:"name"`
	FileName    string `json:"file_name"`
	//Time        string `json:"time"`
}

//
//增减软件版本信息的请求信息
type AddGatewayVersionFileQeq struct {
	FileName string `json:"file_name"`
	File     []byte `json:"file"`
}

//deleteVersion
type DeleteVersionQeq struct {
	Version []string `json:"version"` // 软件版本
}

/*
{
   "startetime" : 1600087800, 查询告警的开始时间戳,
   "endetime":  1600341810,查询告警的结束时间戳,
}
*/

//滴滴监控夜莺查询最近metric[指标值]
type QueryMetricMsgQeq struct {
	Metric string `json:"metric"` //指标
}

/*{【网关内存使用百分比】
    "metric":"gateway.park.gateway.mempercent"
}
gateway.park.gateway.mempercent 查询【网关内存使用百分比】指标
gateway.park.gateway.cpupercent 查询【网关CPU使用百分比】指标
*/
