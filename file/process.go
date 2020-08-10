package file

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/sirupsen/logrus"
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
		fileSet, err := getFileSet(ctx.String("dictionary"))
		if err != nil {
			return err
		}



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
		return nil
	}
}
