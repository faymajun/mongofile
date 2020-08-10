package mongo

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"golang.org/x/net/context"
)

func TestMongo(t *testing.T) {
	redisConf := MongoConfig{Addr: "mongodb://192.168.40.196:27017", Name: Local_Mongo}
	err := MongoMgr.Add(redisConf)
	if err != nil {
		t.Errorf("The error is not nil, but it should be, err=%v", err)
	}

	collection := LocalMongo().Database("battle").Collection("record")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	res, errInsert := collection.InsertOne(ctx, bson.M{"name": "pi", "value": 3.14159})
	if errInsert != nil {
		t.Errorf("The error is not nil, but it should be, err=%v", errInsert)
	}

	id := res.InsertedID.(primitive.ObjectID)

	if s, err := strconv.ParseInt(id.Hex()[0:8], 16, 32); err == nil {
		fmt.Printf("%+v|%d|\n", id, s)
	}

	fmt.Println(id.String())
	fmt.Println(id.Hex())
	t.Logf("the insert expected greeting is %v.\n", id)

}

func TestMongo3(t *testing.T) {
	redisConf := MongoConfig{Addr: "mongodb://192.168.40.196:27017", Name: Local_Mongo}
	err := MongoMgr.Add(redisConf)
	if err != nil {
		t.Errorf("The error is not nil, but it should be, err=%v", err)
	}

	db := LocalMongo().Database("test")
	err = db.RunCommand(
		context.Background(),
		bsonx.Doc{
			{"create", bsonx.String("collection")},
			{"capped", bsonx.Boolean(true)},
			{"size", bsonx.Int32(64 * 1024)},
			{"max", bsonx.Int32(1)},
		},
	).Err()

	if err != nil {
		//	t.Errorf("The error is not nil, but it should be, err=%v", err)
	}
	coll := db.Collection("collection")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, errInsert := coll.InsertOne(ctx, bson.M{"name": "pi", "value": 6666})
	if errInsert != nil {
		t.Errorf("The error is not nil, but it should be, err=%v", errInsert)
	}

}
