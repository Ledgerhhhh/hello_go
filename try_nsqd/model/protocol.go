package model

type Protocol struct {
	Header  Header  `json:"Header"`
	Payload Payload `json:"Payload"`
}

// Header 协议头部分
type Header struct {
	// Namespace 请求类型(设备发现、设备控制、设备属性查询)
	Namespace string `json:"namespace"`
	// Name 请求具体目的,例：DiscoveryDevices 设备发现（获取设备列表）
	Name string `json:"name"`
	// Version Payload 的版本,目前版本为 1
	Version interface{} `json:"payLoadVersion"`
	// MessageID 方便请求追踪
	MessageID string `json:"messageId"`
}

// Payload 协议负载部分
type Payload struct {
	AccessToken string    `json:"accessToken,omitempty"`
	DeviceID    string    `json:"deviceId,omitempty"`
	DeviceIDs   []string  `json:"deviceIds,omitempty"`
	DeviceType  string    `json:"deviceType,omitempty"`
	Params      *Property `json:"params,omitempty"`
	// Attribute 设备控制指令作用的功能（如：电源、色光、风速增大）
	Attribute string `json:"attribute,omitempty"`
	// Value 对功能的操作内容（如：开、红光、风速增大到5挡）
	Value string `json:"value,omitempty"`
	// Devices 设备发现返回的设备列表
	Devices []TmDevice `json:"devices,omitempty"`
	// DeviceResponseList 设备控制返回结果
	DeviceResponseList []DeviceResponseList `json:"deviceResponseList,omitempty"`
	//** 报错时返回以下字段  **//
	ErrorCode string `json:"errorCode,omitempty"`
	ErrorMsg  string `json:"message,omitempty"`

	//** 小度分割线  **//
	// DiscoveredAppliances 场景/设备发现
	DiscoveredAppliances []DiscoveredAppliance `json:"discoveredAppliances,omitempty"`
	// DiscoveredGroups 设备发现分组
	DiscoveredGroups *[]DiscoveredGroup `json:"discoveredGroups,omitempty"`
	// Appliance 设备控制中的设备信息
	Appliance *DiscoveredAppliance `json:"appliance,omitempty"`
	// Function 表示打开设备的子功能
	Function string `json:"function,omitempty"`
	// ColorTemperatureInKelvin 色温
	ColorTemperatureInKelvin int `json:"colorTemperatureInKelvin,omitempty"`
	// DeltaPercentage 微调百分比
	DeltaPercentage *deltaPercentage `json:"deltaPercentage,omitempty"`
	// Brightness 亮度百分比
	Brightness *Brightness `json:"brightness,omitempty"`
	// Attributes 设备属性
	Attributes []Attribute `json:"attributes,omitempty"`
	// PreviousState 亮度修改前的对象
	PreviousState *PreviousState `json:"previousState,omitempty"`
	// DependentServiceName 依赖服务名称
	DependentServiceName string `json:"dependentServiceName,omitempty"`
	// DetalValue 设备属性设置
	DetalValue *detalValue `json:"detalValue"`
	DeltValue  *deltValue  `json:"deltValue"`
	// Color 颜色
	Color *Color `json:"color,omitempty"`
	// AchievedState 颜色更改后设备的状态
	AchievedState *AchievedState `json:"achievedState,omitempty"`
	// Mode 模式
	Mode *Mode `json:"mode,omitempty"`
	// TargetTemperature 设定温度
	TargetTemperature *TargetTemperature `json:"targetTemperature"`
	// FanSpeed 风速
	FanSpeed *FanSpeed `json:"fanSpeed"`
	// LockState 锁
	LockState string `json:"lockState"`
}

// TmDevice 天猫精灵设备
type TmDevice struct {
	DeviceID   string `json:"deviceId"`
	DeviceType string `json:"deviceType"`
	DeviceName string `json:"deviceName"`
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	Zone       string `json:"zone"`
	Icon       string `json:"icon"`
	//Properties []property `json:"properties"`
	Status Property `json:"status"`
	// Actions 设备支持的操作动作
	//Actions []string `json:"actions"`
	// Extensions 附加属性
	Extensions Extensions `json:"extensions"`
}

// Extensions 附加字段
type Extensions struct {
	Extensions1 string `json:"extensions1"`
}

// DeviceResponseList 设备控制返回结果
type DeviceResponseList struct {
	DeviceID  string `json:"deviceId"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}

// Property 设备支持操控的属性
type Property struct {
	// Powerstate 电源开关，0：关，1：开
	Powerstate interface{} `json:"powerstate,omitempty"`
	// Onlinestate 在线状态，0：下线，1：在线
	Onlinestate interface{} `json:"onlinestate,omitempty"`
	// Brightness 灯光亮度
	Brightness interface{} `json:"brightness,omitempty"`
	// ColorTemperature 灯光色温
	ColorTemperature interface{} `json:"colorTemperature,omitempty"`
	// CurtainPosition 窗帘打开具体位置
	CurtainPosition interface{} `json:"curtainPosition,omitempty"`
	// CurtainPosition 窗帘执行动作
	CurtainConrtol interface{} `json:"curtainConrtol,omitempty"`
}

// DiscoveredAppliance 小度设备/场景发现
type DiscoveredAppliance struct {
	// ApplianceTypes 设备的品类
	ApplianceTypes []string `json:"applianceTypes"`
	// ApplianceId 设备ID
	ApplianceId string `json:"applianceId"`
	// ModelName 设备型号（对应的我们的设备中类）
	ModelName string `json:"modelName"`
	// Version 设备版本
	Version string `json:"version"`
	// FriendlyName 设备名称（取DeviceName）
	FriendlyName string `json:"friendlyName"`
	// FriendlyDescription 设备相关的描述，描述内容提需要提及设备厂商，使用场景及连接方式
	FriendlyDescription string `json:"friendlyDescription"`
	// IsReachable 设备当前是否能够到达，在我们的体系中等于是否在线
	IsReachable bool `json:"isReachable"`
	// Actions 设备支持的操作类型数组（turnOn、turnOff）
	Actions []string `json:"actions"`
	// AdditionalApplianceDetails 扩展信息，可以为空
	AdditionalApplianceDetails additionalApplianceDetails `json:"additionalApplianceDetails"`
	// ManufacturerName 设备厂商的名字
	ManufacturerName string `json:"manufacturerName"`
	// Attributes 设备属性
	Attributes []Attribute `json:"attributes"`
	// DeviceSubTypeNo 设备小类（非小度使用）
	DeviceSubTypeNo string `json:"subType"`
}

// additionalApplianceDetails 扩展信息
type additionalApplianceDetails struct {
	DeviceType      string `json:"deviceType"`
	DeviceSubTypeNo string `json:"deviceSubTypeNo"`
}

// Attribute 设备属性
type Attribute struct {
	Name       string      `json:"name"`
	Value      interface{} `json:"value"`
	Scale      string      `json:"scale"`
	LegalValue string      `json:"legalValue,omitempty"`
	// TimestampOfSample 属性值取样的时间戳，单位是秒。
	TimestampOfSample int64 `json:"timestampOfSample"`
	// UncertaintyInMilliseconds 属性值取样的时间误差,执行命令到上报属性的误差，毫秒
	UncertaintyInMilliseconds int `json:"uncertaintyInMilliseconds"`
}

// DiscoveredGroup 设备发现分组
type DiscoveredGroup struct {
	GroupName              string                 `json:"groupName"`
	ApplianceIds           []string               `json:"applianceIds"`
	GroupNotes             string                 `json:"groupNotes"`
	AdditionalGroupDetails additionalGroupDetails `json:"additionalGroupDetails"`
}

// additionalGroupDetails 分组扩展信息
type additionalGroupDetails struct {
}

// deltaPercentage 微调百分比
type deltaPercentage struct {
	Value int `json:"value"`
}

// Brightness 设置亮度
type Brightness struct {
	Value interface{} `json:"value"`
}

// PreviousState 亮度设置前的对象
type PreviousState struct {
	Brightness Brightness `json:"brightness"`
}

// detalValue 设置参数
type detalValue struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

// deltValue 设置参数
type deltValue struct {
	Value string `json:"value"`
	Scale string `json:"scale"`
}

// Color 颜色
type Color struct {
	Hue        float32 `json:"hue"`
	Saturation float32 `json:"saturation"`
	Brightness float32 `json:"brightness"`
}

// AchievedState 颜色更改后设备的状态
type AchievedState struct {
	Color *Color `json:"color"`
}

// Mode 模式
type Mode struct {
	DeviceType string `json:"deviceType"`
	Value      string `json:"value"`
}

// TargetTemperature 设定温度
type TargetTemperature struct {
	Value interface{} `json:"value"`
	Scale string      `json:"scale"`
}

// FanSpeed 风速
type FanSpeed struct {
	Value interface{} `json:"value"`
	Level string      `json:"level"`
}
