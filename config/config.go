package config

type Config struct {
	SubscribedGroups []uint32 // 订阅群组

	BotPrefix string `json:",optional"`

	QQUin    uint32
	Password string `json:",optional"`

	AppInfo    string `json:",optional"`
	AppInfoVer string `json:",optional"`
	SignUrl    string `json:",optional"`
	// 如果没有指定路径的DeviceInfo,使用seed初始化
	DeviceInfoSeed int `json:",optional"`
}

func NewConfig() Config {
	return Config{
		AppInfo:        "linux",
		AppInfoVer:     "3.1.2-13107",
		SignUrl:        "https://sign.lagrangecore.org/api/sign/25765",
		DeviceInfoSeed: 898989,
		BotPrefix:      "eft",
	}
}
