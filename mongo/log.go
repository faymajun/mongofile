package mongo

import (
	"bufio"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"strings"
)

// 打开log文件-按行读取
func ReadLog(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(fmt.Sprintf("open log file err: %v", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineText := scanner.Text()
		fmt.Println(lineText)
		strArr := strings.Split(lineText, "#")
		if len(strArr) == 4 {
			db, collection, data := strArr[1], strArr[2], strArr[3]
			doc := bson.D{}
			err = bson.UnmarshalExtJSON([]byte(data), true, &doc)
			if err == nil {
				InsertLogOne(db, collection, doc)
			}
		}
	}
}

