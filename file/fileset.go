package file

import (
	"fmt"
	"io/ioutil"
)

func getFileSet(modulePath string) ([]string, error) {
	fmt.Println(modulePath)
	fileInfos, err := ioutil.ReadDir(modulePath)
	if err != nil {
		return []string{}, err
	}
	fileset := []string{}
	for _, fi := range fileInfos {
		if !fi.IsDir() {
			fileset = append(fileset, fi.Name())
		}
	}
	return fileset, nil
}
