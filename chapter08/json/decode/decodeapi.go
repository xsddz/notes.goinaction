// Package decode 这个示例程序展示如何使用 json 包和 NewDecoder 函数来解码 JSON 响应
package decode

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type book struct {
	// 到每个字段最后使用单引号声明了一个字符串。这些字符串被称作标签（tag），
	// 是提供每个字段的元信息的一种机制，将 JSON 文档和结构类型里的字段一一映射起来。
	// 如果不存在标签，编码和解码过程会试图以大小写无关的方式，直接使用字段的名字进行匹配。
	// 如果无法匹配，对应的结构类型里的字段就包含其零值。
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	Authors   string `json:"authors"`
	Publisher string `json:"publisher"`
	Language  string `json:"language"`
	Isbn10    string `json:"isbn10"`
	Isbn13    string `json:"isbn13"`
	Pages     string `json:"pages"`
	Year      string `json:"year"`
	Rating    string `json:"rating"`
	Desc      string `json:"desc"`
	Price     string `json:"price"`
	Image     string `json:"image"`
	URL       string `json:"url"`
}

// RunAPIDemo 示例入口
func RunAPIDemo() {
	uri := "https://api.itbook.store/1.0/books/9780134190440"

	resp, err := http.Get(uri)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	defer resp.Body.Close()

	// 将 JSON 响应解码到结构类型
	var b book
	err = json.NewDecoder(resp.Body).Decode(&b)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(b)
}
