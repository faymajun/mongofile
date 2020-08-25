package file

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"mongofile/mongo"
	"mongofile/mongo/routine"
	"time"
)

var (
	status struct {
		started int32
		stoped  int32
	}

	logger = logrus.WithField("component", "process")
)

func Action() func(*cli.Context) error {
	return func(ctx *cli.Context) error {
		defaultConfig.scanFrequency = time.Duration(ctx.Int("scanFrequency")) * time.Second
		//fileSet, err := getFileSet(ctx.String("dictionary"))
		//if err != nil {
		//	return err
		//}
		// 连接Mongo数据库
		mongo.InitDefaultConfig(ctx.String("mongoAddr"))
		defer mongo.CloseDefaultMongo()

		fileName := ctx.String("file")
		mongo.ReadLog(fileName)
		routine.Pool.Stop() // 协程池任务关闭

		return nil
	}
}
