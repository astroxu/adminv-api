// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginResp struct {
	Code             string `json:"code"`
	Message          string `json:"message"`
	Status           string `json:"status"`
	CurrentAuthority string `json:"currentAuthority"`
	Id               int64  `json:"id"`
	UserName         string `json:"userName"`
	AccessToken      string `json:"token"`
	AccessExpire     int64  `json:"accessExpire"`
	RefreshAfter     int64  `json:"refreshAfter"`
}

type ListSysLogReq struct {
	Current  int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}

type ListSysLogData struct {
	Id             int64  `json:"id"`             // 编号
	UserName       string `json:"userName"`       // 用户名
	Operation      string `json:"operation"`      // 用户操作
	Method         string `json:"method"`         // 请求方法
	Params         string `json:"params"`         // 请求参数
	Time           int64  `json:"time"`           // 执行时长(毫秒)
	Ip             string `json:"ip"`             // IP地址
	CreateBy       string `json:"createBy"`       // 创建人
	CreateTime     string `json:"createTime"`     // 创建时间
	LastUpdateBy   string `json:"lastUpdateBy"`   // 更新人
	LastUpdateTime string `json:"lastUpdateTime"` // 更新时间
}

type ListSysLogResp struct {
	Code     string            `json:"code"`
	Message  string            `json:"message"`
	Current  int64             `json:"current,default=1"`
	Data     []*ListSysLogData `json:"data"`
	PageSize int64             `json:"pageSize,default=20"`
	Success  bool              `json:"success"`
	Total    int64             `json:"total"`
}

type DeleteSysLogReq struct {
	Id int64 `path:"id"`
}

type DeleteSysLogResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ListLoginLogReq struct {
	Current  int64 `json:"current,default=1"`
	PageSize int64 `json:"pageSize,default=20"`
}

type ListLoginLogData struct {
	Id             int64  `json:"id"`             // 编号
	UserName       string `json:"userName"`       // 用户名
	Status         string `json:"status"`         // 登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）
	Ip             string `json:"ip"`             // IP地址
	CreateBy       string `json:"createBy"`       // 创建人
	CreateTime     string `json:"createTime"`     // 创建时间
	LastUpdateBy   string `json:"lastUpdateBy"`   // 更新人
	LastUpdateTime string `json:"lastUpdateTime"` // 更新时间
}

type ListLoginLogResp struct {
	Code     string              `json:"code"`
	Message  string              `json:"message"`
	Current  int64               `json:"current,default=1"`
	Data     []*ListLoginLogData `json:"data"`
	PageSize int64               `json:"pageSize,default=20"`
	Success  bool                `json:"success"`
	Total    int64               `json:"total"`
}

type DeleteLoginLogReq struct {
	Id int64 `path:"id"`
}

type DeleteLoginLogResp struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
