// Package decode 这个示例程序展示如何解码 JSON 字符串
package decode

import (
	"encoding/json"
	"fmt"
	"log"
)

// Contact 结构代表我们的 JSON 字符串
type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Cell string `json:"cell"`
	} `json:"contact"`
}

// JSON 包含用于反序列化的演示字符串
var JSON = `{
	"name": "Gopher",
	"title": "programmer",
	"contact": {
		"home": "415.333.3333",
		"cell": "415.555.5555"
	}
}`

// RunStringDemo 示例入口
func RunStringDemo() {
	// 将 JSON 字符串反序列化到变量
	var cs Contact
	cserr := json.Unmarshal([]byte(JSON), &cs)
	if cserr != nil {
		log.Println("ERROR:", cserr)
		return
	}

	fmt.Println(cs)

	// 将 JSON 字符串反序列化到 map 变量
	var cm map[string]interface{}
	cmerr := json.Unmarshal([]byte(JSON), &cm)
	if cmerr != nil {
		log.Println("ERROR:", cmerr)
		return
	}

	fmt.Println("Name:", cm["name"])
	fmt.Println("title:", cm["title"])
	fmt.Println("Contact")
	fmt.Println("H:", cm["contact"].(map[string]interface{})["home"])
	fmt.Println("C:", cm["contact"].(map[string]interface{})["cell"])

}
