package main

import (
	"github.com/urfave/cli"
	"mongofile/file"
	"os"
)

// 1. 获取Logs文件路径
// 2. 打开Logs文件，读取注册表偏移位置
// 2. 定时获取最新的logs line
// 3. 把json数据结构存入mongo数据库
// 4. 保存当前的读取偏移位置
// 5. 关闭的时候写入注册表
func main() {
	app := cli.NewApp()
	app.Usage = "startup server"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "dictionary, d",
			Value: "mongo",
		},
		cli.StringFlag{
			Name:  "file, f",
			Value: "text.log",
		},
		cli.StringFlag{
			Name:  "mongoAddr, m",
			Value: "mongodb://192.168.1.224:27017",
		},
		cli.IntFlag{
			Name:  "scanFrequency, s",
			Value: 60,
		},
	}

	app.Action = file.Action()
	app.Run(os.Args)
}