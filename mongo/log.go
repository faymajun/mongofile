package mongo

import (
	"bufio"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"strings"
)

// 打开log文件-按行读取
var count int64 = 0

func ReadLog(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("open log file err: %v", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		count++
		lineText := scanner.Text()
		strArr := strings.Split(lineText, "#")
		if len(strArr) == 4 {
			db, collection, data := strArr[1], strArr[2], strArr[3]
			if data[0] == '[' {
				size := len(data)
				start, end := 1, size-1
				ind := 2
				for ; ind < end; ind++ {
					if data[ind-2] == '}' && data[ind-1] == ',' && data[ind] == '{' {
						doc := bson.M{}
						err = bson.UnmarshalExtJSON([]byte(data[start:ind-1]), false, &doc)
						if err == nil {
							InsertLogOne(db, collection, doc)
						} else {
							logger.Println(err)
						}
						start = ind
					}
				}
				if start < end-1 {
					doc := bson.M{}
					err = bson.UnmarshalExtJSON([]byte(data[start:end]), false, &doc)
					if err == nil {
						InsertLogOne(db, collection, doc)
					} else {
						logger.Println(err)
					}
				}
			} else {
				doc := bson.M{}
				err = bson.UnmarshalExtJSON([]byte(data), true, &doc)
				if err == nil {
					InsertLogOne(db, collection, doc)
				} else {
					logger.Println(err)
				}
			}
		}
	}

	logger.Infof("总共导入mongo数据量：%d", count)
}
