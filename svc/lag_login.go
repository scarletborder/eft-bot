package svc

import (
	"os"

	LagCli "github.com/LagrangeDev/LagrangeGo/client"
	LagAuth "github.com/LagrangeDev/LagrangeGo/client/auth"
	"github.com/zeromicro/go-zero/core/logx"
)

var deviceInfoPath string = "etc/device.json"
var sigPath string = "etc/sig.bin"
var qrCodePath string = "etc/qrcode.png"

func getDeviceInfo(seed int) *LagAuth.DeviceInfo {
	deviceInfo, err := LagAuth.LoadOrSaveDevice(deviceInfoPath)
	if err != nil {
		logx.Infof("fail to read deviceInfo file in %s %s, initialize one first",
			deviceInfoPath, err.Error())

		deviceInfo = LagAuth.NewDeviceInfo(seed)
	}
	return deviceInfo
}

func getSig() *LagAuth.SigInfo {
	data, err := os.ReadFile(sigPath)
	if err != nil {
		logx.Info("read sig error:", err)
	} else {
		sig, err := LagAuth.UnmarshalSigInfo(data, true)
		if err != nil {
			logx.Info("load sig error:", err)
		} else {
			return &sig
		}
	}
	return nil
}

// 保存当前账号的所有登录信息
func SaveUinInfo(q *LagCli.QQClient) {
	defer q.Release()

	// device info
	err := q.Device().Save(deviceInfoPath)
	if err != nil {
		logx.Errorf("fail to save deviceInfo: %v", err)
	}

	// sig
	data, err := q.Sig().Marshal()
	if err != nil {
		logx.Errorf("fail to marshal sig: %v", err)
	} else {
		err = os.WriteFile(sigPath, data, 0644)
		if err != nil {
			logx.Errorf("fail to write sig to disk: %v", err)
		}
	}
}
