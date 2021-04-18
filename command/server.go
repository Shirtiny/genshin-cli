package command

import (
	"fmt"
	"gcli/util"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

// Server 服务器信息结构体
type Server struct {
	Name       string
	Cps        string
	Channel    int
	SubChannel int
	Desc       string
}

var servers = []Server{
	{
		Name:       "BiliBili",
		Cps:        "bilibili",
		Channel:    14,
		SubChannel: 0,
		Desc:       "哔哩哔哩",
	},
	{
		Name:       "Mihoyo",
		Cps:        "mihoyo",
		Channel:    1,
		SubChannel: 1,
		Desc:       "米哈游",
	},
}

// NewServer 实例化serverCommand
func NewServer() *cli.Command {
	serverMap := make(map[string]Server)

	for _, server := range servers {
		serverMap[server.Desc] = server
	}

	return &cli.Command{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "选择游戏服务器",
		Action: func(c *cli.Context) error {
			serverDesc := ""
			prompt := &survey.Select{
				Message: "选择服务器 (使用方向键上下选择，回车键确定)",
				Options: []string{"哔哩哔哩", "米哈游"},
			}
			survey.AskOne(prompt, &serverDesc)
			server := serverMap[serverDesc]
			fmt.Printf("已经选择：%+v", server)
			// 获取配置文件路径
			_, configPath := GetGenshiConfigPath()
			// 覆写配置文件
			util.SetGenshinServerConfig(configPath, server.Cps, server.Channel, server.SubChannel)
			return nil
		},
		Subcommands: []*cli.Command{
			{
				Name:  "ls",
				Usage: "列出可选服务器",
				Action: func(c *cli.Context) error {
					for _, server := range servers {
						fmt.Println(server.Name, "名称：", server.Desc)
					}
					return nil
				},
			},
		},
	}
}
