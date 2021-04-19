package command

import (
	"fmt"
	"gcli/util"
	"log"
	"os/exec"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

// Server 服务器信息结构体
type Server struct {
	Name       string
	Alias      string
	Cps        string
	Channel    int
	SubChannel int
	Desc       string
}

var servers = []Server{
	{
		Name:       "BiliBili",
		Alias:      "b",
		Cps:        "bilibili",
		Channel:    14,
		SubChannel: 0,
		Desc:       "哔哩哔哩",
	},
	{
		Name:       "Mihoyo",
		Alias:      "m",
		Cps:        "mihoyo",
		Channel:    1,
		SubChannel: 1,
		Desc:       "米哈游",
	},
}

// GetServer 根据输入情况 查找并返回找到的第一个服务器
func GetServer(nameOrDesc string) *Server {
	// 这里先简单遍历一下
	for _, server := range servers {
		// 比较时忽略大小写
		if strings.EqualFold(server.Name, nameOrDesc) || strings.EqualFold(server.Alias, nameOrDesc) || strings.EqualFold(server.Desc, nameOrDesc) {
			return &server
		}
	}
	return nil
}

// SwitchServer 切换服务器
func SwitchServer(server *Server) {
	// 获取配置文件路径
	_, configPath := GetGenshiConfigPath()

	// 覆写配置文件
	util.SetGenshinServerConfig(configPath, server.Cps, server.Channel, server.SubChannel)
	fmt.Printf("已经切换至：%+v\n", server)
}

// NewServer 实例化serverCommand
func NewServer() *cli.Command {

	return &cli.Command{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "选择游戏服务器，仅切换，不启动游戏",
		Action: func(c *cli.Context) error {
			if c.NArg() > 0 {
				return nil
			}
			serverDesc := ""
			prompt := &survey.Select{
				Message: "选择服务器 (使用方向键上下选择，回车键确定)",
				Options: []string{"哔哩哔哩", "米哈游"},
			}
			survey.AskOne(prompt, &serverDesc)
			server := GetServer(serverDesc)
			SwitchServer(server)
			return nil
		},
		Subcommands: []*cli.Command{
			{
				Name:  "ls",
				Usage: "列出可选服务器",
				Action: func(c *cli.Context) error {
					for _, server := range servers {
						fmt.Printf("%+v\n", server)
					}
					// 获取配置文件路径
					_, configPath := GetGenshiConfigPath()
					fmt.Println("当前：", GetServer(util.GetGenshinServerFromConfig(configPath).Cps).Desc)
					return nil
				},
			},
			{
				Name:    "use",
				Aliases: []string{"u"},
				Usage:   "选择服务器并启动游戏, m 官服， b b服",
				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return nil
					}
					// 查找server
					server := GetServer(c.Args().Get(0))
					if server == nil {
						fmt.Println("未找到输入的服务器", c.Args().Get(0), server)
						return nil
					}

					// 切换服务器
					SwitchServer(server)

					// 运行游戏
					cmd := exec.Command(GetGenshiLauncherPath())
					if err := cmd.Start(); err != nil {
						log.Fatalf("failed to call cmd.Run(): %v", err)
						return err
					}
					fmt.Println("启动完成")
					return nil
				},
			},
		},
	}
}
