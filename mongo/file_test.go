package mongo

import (
	"encoding/json"
	"testing"
	"time"
)

type ServerInfo struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	Time int64  `json:"time"`
}

var (
	TestDB         = "db"
	TestCollection = "collection"
)

func TestFile(t *testing.T) {
	InitMongoFile("test")
	testJson, err := json.Marshal(ServerInfo{Name: "testSever", Id: "1", Time: time.Now().Unix()})
	if err != nil {
		logger.Errorf("json marshal err: %v", err)
	}
	WriteLogOne(TestDB, TestCollection, string(testJson))
	time.Sleep(2 * time.Second)
	CloseMongoLog()
}

func TestFile1(t *testing.T) {
	InitMongoFile("test")
	for i := 0; i < 100; i++ {
		WriteLogOne(TestDB, TestCollection, "aaaaaaaaaaa")
	}
	time.Sleep(3 * time.Second)
	CloseMongoLog()
}

func TestFile2(t *testing.T) {
	InitMongoFile("test")
	for i := 0; i < 100; i++ {
		WriteLogOne(TestDB, TestCollection, "bbbbbbbbbb")
	}
	time.Sleep(3 * time.Second)
	CloseMongoLog()
}

func Test3(t *testing.T) {
	go TestFile1(t)
	go TestFile2(t)
	time.Sleep(4 * time.Second)
}
