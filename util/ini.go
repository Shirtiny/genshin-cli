package util

import (
	"strconv"

	"gopkg.in/ini.v1"
)

// SetGenshinServerConfig 覆写原神服务器配置 先简单的实现 写死
func SetGenshinServerConfig(path string, cps string, channel int, subChannel int) {
	// 取消空格美化
	ini.PrettyFormat = false
	cfg, _ := ini.Load(path)
	general := cfg.Section("General")
	general.Key("cps").SetValue(cps)
	general.Key("channel").SetValue(strconv.Itoa(channel))
	general.Key("sub_channel").SetValue(strconv.Itoa(subChannel))
	// 覆盖文件
	cfg.SaveTo(path)
}
