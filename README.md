# 原神命令行工具

简单实现，目前可方便的在 b 服和官服之间切换

## 使用(管理员身份下)

查看游戏信息

> 目前还在测试，如果没有列出游戏安装路径，请不要继续使用

```shell
gcli info
```

切换服务器

> 目前只支持国内 b 服和官服

```shell
gcli server
```

切换并启动指定服务器

```shell
# 切换到b服
gcli server use b
#切换到官服
gcli server use m
```

查看支持的服务器列表和当前服务器

```shell
gcli server ls
```

帮助

```shell
gcli -h
```

原神2.1 B服需要将[PCGameSDK.dll](https://github.com/Shirtiny/genshin-cli/blob/main/PCGameSDK.dll)拷贝至`Genshin Impact Game`文件夹下才可切换
