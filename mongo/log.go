package mongo

import (
	"bufio"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"io"
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

	lineReader := bufio.NewReader(file)
	for {
		// 相同使用场景下可以采用的方法
		// func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)
		// func (b *Reader) ReadBytes(delim byte) (line []byte, err error)
		// func (b *Reader) ReadString(delim byte) (line string, err error)
		lineText, err := lineReader.ReadString('\n')
		if err == io.EOF {
			break
		}
		count++
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
							logger.Errorln(err)
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
						logger.Errorln(err)
					}
				}
			} else {
				doc := bson.M{}
				err = bson.UnmarshalExtJSON([]byte(data), true, &doc)
				if err == nil {
					InsertLogOne(db, collection, doc)
				} else {
					logger.Errorln(err)
				}
			}
		}
	}

	logger.Infof("总共导入mongo数据量：%d", count)
}
