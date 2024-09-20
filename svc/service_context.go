package svc

import (
	"eftbot/config"

	LagCli "github.com/LagrangeDev/LagrangeGo/client"
	LagAuth "github.com/LagrangeDev/LagrangeGo/client/auth"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	C     config.Config
	QQCli *LagCli.QQClient

	// TODO: bot实例
}

func NewServiceContext(c config.Config) *ServiceContext {
	appInfo := LagAuth.AppList[c.AppInfo][c.AppInfoVer]
	qq_cli := LagCli.NewClient(c.QQUin, appInfo, c.SignUrl)

	// device info
	deviceInfo := getDeviceInfo(c.DeviceInfoSeed)
	qq_cli.UseDevice(deviceInfo)

	// sig
	// 登录信息
	sig := getSig()
	if sig != nil {
		qq_cli.UseSig(*sig)
	}

	// logger
	qq_cli.SetLogger(MyLogger{})

	// login
	// 观察原型得知该函数功能为阻塞并等待登录
	// 当c.sig内有登录信息时，会优先进行快速登录.密码为空则是扫码登录
	// TODO: 添加检测登录超时
	// log_ctx, _ := context.WithTimeout(context.Background(), 3*time.Minute)
	// threading.GoSafeCtx()
	err := qq_cli.Login(c.Password, qrCodePath)
	if err != nil {
		logx.Error("login fail: ", err.Error())
		return nil
	}

	return &ServiceContext{
		C:     c,
		QQCli: qq_cli,
	}
}
