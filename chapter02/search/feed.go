package search

import (
	"encoding/json"
	"os"
)

/*
1. Go 编译器可以根据赋值运算符右边的值来推导类型，声明常量的时候不需要指定类型。
2. 这个常量的名称使用小写字母开头，表示它只能在 search 包内的代码里直接访问，而不暴露到包外面。
*/
const dataFile = "data/data.json"

// Feed 包含我们需要处理的数据源的信息。这个类型会对外暴露。
type Feed struct {
	/*
		每个字段的声明最后 ` 引号里的部分被称作标记（tag）。这个标记里描述了 JSON 解码的元数据，
		用于创建 Feed 类型值的切片。每个标记将结构类型里字段对应到 JSON 文档里指定名字的字段。
	*/
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

// RetrieveFeeds 读取并反序列化源数据文件
func RetrieveFeeds() ([]*Feed, error) {
	// 打开文件
	file, err := os.Open(dataFile)
	if err != nil {
		return nil, err
	}

	/*
		当函数返回时关闭文件。

		关键字 defer 会安排随后的函数调用在函数返回时才执行，且保证这个函数一定会被调用。
		哪怕函数意外崩溃终止，也能保证关键字 defer 安排调用的函数会被执行。
	*/
	defer file.Close()

	// 将文件解码到一个切片里，这个切片的每一项是一个指向一个 Feed 类型值的指针
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)

	// 这个函数不需要检查错误，调用者会做这件事
	return feeds, err
}
