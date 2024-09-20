package svc

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

var DumpsPath string = "etc/dump"

// func init() {
// 	// Ensure DumpsPath exists
// }

type MyLogger struct{}

func (l MyLogger) Info(format string, args ...any) {
	logx.Infof(format, args...)
}
func (l MyLogger) Warning(format string, args ...any) {
	logx.Infof(format, args...)
}
func (l MyLogger) Error(format string, args ...any) {
	logx.Errorf(format, args...)
}
func (l MyLogger) Debug(format string, args ...any) {
	// 使用 logx.Debugf 进行格式化输出
	logx.Debugf(format, args...)
}
func (l MyLogger) Dump(dumped []byte, format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	if _, err := os.Stat(DumpsPath); err != nil {
		err = os.MkdirAll(DumpsPath, 0o755)
		if err != nil {
			logx.Errorf("出现错误 %v. 详细信息转储失败", message)
			return
		}
	}
	dumpFile := path.Join(DumpsPath, fmt.Sprintf("%v.dump", time.Now().Unix()))
	logx.Errorf("出现错误 %v. 详细信息已转储至文件 %v 请连同日志提交给开发者处理", message, dumpFile)
	_ = os.WriteFile(dumpFile, dumped, 0o644)
}
