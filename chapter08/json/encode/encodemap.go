package encode

import (
	"encoding/json"
	"fmt"
	"log"
)

// RunMapDemo 示例入口
func RunMapDemo() {
	// 创建一个保存键值对的映射
	c := make(map[string]interface{})

	c["name"] = "张三"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "北京海淀",
		"cell": "1024",
	}

	// 将这个映射序列化到 JSON 字符串
	// MarshalIndent 很像 Marshal，只是用缩进对输出进行格式化
	data, err := json.MarshalIndent(c, "", "    ")
	// data, err := json.Marshal(c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println(string(data))
}
