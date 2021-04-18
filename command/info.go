package command

import (
	"fmt"
	"gcli/util"

	"github.com/urfave/cli/v2"
)

// CONFIG_RELATIVE_PATH 配置文件路径（相对于安装路径） 暂时简单用string写死
const CONFIG_RELATIVE_PATH = `\Genshin Impact Game\config.ini`

// GetGenshiConfigPath 获取原神配置文件路径
func GetGenshiConfigPath() (string, string) {
	installPath := util.GetGenshinInstallPathByReg()
	configPath := installPath + CONFIG_RELATIVE_PATH
	return installPath, configPath
}

// NewInfo 实例化info命令
func NewInfo() *cli.Command {
	return &cli.Command{
		Name:    "info",
		Aliases: []string{"i"},
		Usage:   "显示查询到的游戏相关信息",
		Action: func(c *cli.Context) error {
			installPath, configPath := GetGenshiConfigPath()
			fmt.Println("查询到的原神安装路径：", installPath)
			fmt.Println("推测的配置文件路径：", configPath)
			return nil
		},
	}
}
