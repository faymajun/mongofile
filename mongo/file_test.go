package mongo

import (
	"encoding/json"
	"testing"
	"time"
)

type ServerInfo struct {
	Name string `json:"name"`
	Id string `json:"id"`
	Time int64 `json:"time"`
}

func TestFile(t *testing.T) {
	InitMongoFile("test")
	testJson, err := json.Marshal(ServerInfo{Name: "testSever", Id:"1", Time: time.Now().Unix()})
	if err != nil {
		logger.Errorf("json marshal err: %v", err)
	}
	WriteLogOne("db", "collection", string(testJson))
	time.Sleep(2 * time.Second)
	CloseMongoLog()
}
