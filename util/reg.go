package util

import (
	"log"

	"golang.org/x/sys/windows/registry"
)

// GetGenshinInstallPath 获取原神安装路径 暂时先这么写死
func GetGenshinInstallPathByReg() string {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows\CurrentVersion\Uninstall\原神`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	path, _, err := k.GetStringValue("installPath")
	if err != nil {
		log.Fatal(err)
	}
	return path
}
