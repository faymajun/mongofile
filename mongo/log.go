package mongo

import (
	"bufio"
	"fmt"
	"os"
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
	}
}