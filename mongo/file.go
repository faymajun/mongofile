package mongo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var MongoFile mongoFile

type mongoFile struct {
	logger *log.Logger
	f      *os.File
}

func InitMongoFile(fileName string) {
	MongoFile.openFile(fileName)
}

func (mf *mongoFile) openFile(fileName string) {
	abs, err := filepath.Abs("mongo")
	if err != nil {
		panic(fmt.Errorf("日志目录配置错误: Error=%s", err.Error()))
	}
	os.MkdirAll(abs, os.ModePerm)

	baseLogPath := filepath.Join(abs, fileName) + ".logs"
	//创建日志文件
	f, errOpenFile := os.OpenFile(baseLogPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if errOpenFile != nil {
		if errOpenFile != nil {
			logrus.Errorf("mongo local file system logger error. %v", errors.WithStack(err))
			panic("mongo local file system logger error")
		}
	}

	mf.f = f

	fileWriter := io.MultiWriter(f)
	mf.logger = log.New(fileWriter, "", log.Ldate|log.Ltime)
	// 写入日志内容
	mf.logger.Println("check to make sure it works")
}

func CloseMongoLog() {
	MongoFile.f.Close()
}

func WriteOne(document interface{}) {
	go func() {
		data, err := json.Marshal(document)
		if err == nil {
			MongoFile.logger.Println(string(data))
		}
	}()
}
