package mongo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
	"time"
)

func TestReadLog(t *testing.T) {
	InitDefaultConfig("mongodb://192.168.1.224:27017")
	ReadLog("C:\\Personally\\mongofile\\mongo\\mongo\\test.log")
	time.Sleep(2 * time.Second)
}

func TestBson(t *testing.T) {
	doc := bson.D{}
	err := bson.UnmarshalExtJSON([]byte("{\"name\":\"testSever\",\"id\":\"1\",\"time\":1597200094}"), true, &doc)
	fmt.Println(doc, err)
}