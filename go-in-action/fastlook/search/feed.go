package search

import (
	"encoding/json"
	"os"
)

const dataFile = "fastlook/data/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	// 打开文件
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	// 函数执行完毕时关闭文件
	defer file.Close()

	// 文件解码并放到存放 Feed 指针的切片中
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	return feeds, err
}
