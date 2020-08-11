package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"time"

	"golang.org/x/net/context"
)

const (
	Local_Mongo = "DefaultMongo"
)

func InitDefaultConfig(addr string) {
	redisConf := MongoConfig{Addr: addr, Name: Local_Mongo}
	MongoMgr.Add(redisConf)
}

func CloseDefaultMongo() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	DefaultMongo().Disconnect(ctx)
}

func DefaultMongo() *Mongo {
	return MongoMgr.GetMongo(Local_Mongo)
}

func InsertOne(c *mongo.Collection, document interface{}) {
	go func() {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		if _, error := c.InsertOne(ctx, document); error != nil {
			logger.Errorf("BattleInsert error, collection=%s, document=%v, InsertOne error=%s", c.Name(), document, error)
		}
	}()
}

func InsertMany(c *mongo.Collection, documents []interface{}) {
	go func() {
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		if _, error := c.InsertMany(ctx, documents); error != nil {
			logger.Errorf("BattleInsert error, collection=%s, document=%v, InsertMany error=%s", c.Name(), documents, error)
		}
	}()
}
