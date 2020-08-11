package mongo

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/net/context"
	"sync"
	"time"
)

var (
	MongoMgr = &MongoManager{stop: 0}
	logger   = logrus.WithField("component", "mongo")
)

type MongoManager struct {
	dbs  sync.Map
	stop int32
}

type Mongo struct {
	*mongo.Client
	Conf MongoConfig
}

type MongoConfig struct {
	Addr string
	Name string
}

func (mgr *MongoManager) GetMongo(name string) *Mongo {
	db, ok := mgr.dbs.Load(name)
	if !ok {
		logger.Errorf("mongo get failed:%s", name)
		return nil
	}
	rds, ok := db.(*Mongo)
	if !ok {
		logger.Errorf("mongo get type failed:%s", name)
		return nil
	}
	return rds
}

func (mgr *MongoManager) Exist(name string) bool {
	_, ok := mgr.dbs.Load(name)
	return ok
}

func (mgr *MongoManager) Add(conf MongoConfig) error {
	if mgr.Exist(conf.Name) {
		return fmt.Errorf("redis already have rname:%s %s", conf.Name, conf.Addr)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.Addr))
	if err != nil {
		return err
	}

	mgo := &Mongo{
		Client: client,
		Conf:   conf,
	}
	mgr.dbs.Store(conf.Name, mgo)
	logger.Infof("connect to mongo:%s %s", conf.Name, conf.Addr)
	return nil
}

func (mgr *MongoManager) Del(name string) bool {
	rds := mgr.GetMongo(name)
	if rds == nil {
		return false
	}
	mgr.dbs.Delete(name)
	logger.Errorf("del mongo from mgr:%s", name)
	return true
}