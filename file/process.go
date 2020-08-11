package file

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"mongofile/mongo"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
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
		fileName := ctx.String("file")
		mongo.ReadLog(fileName)

		// 连接Mongo数据库
		mongo.InitDefaultConfig(ctx.String("mongoAddr"))

		// 等待退出信号
		stopChan := make(chan os.Signal, 1)
		signal.Notify(stopChan, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM, syscall.SIGINT)
		select {
		case sig, _ := <-stopChan:
			logger.Infof("<<<==================>>>")
			logger.Infof("<<<stop process by:%v>>>", sig)
			logger.Infof("<<<==================>>>")
			break
		}
		if atomic.LoadInt32(&status.started) == 0 || atomic.AddInt32(&status.stoped, 1) != 1 {
			return fmt.Errorf("Server stop duplication")
		}

		// 关闭Mongo数据库
		mongo.CloseDefaultMongo()
		return nil
	}
}
