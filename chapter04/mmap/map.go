package mmap

import "fmt"

// Run 字典测试入口
func Run() {
	fmt.Printf("================================================================================\n")

	/*
		1. 映射是一个存储键值对的 无序 集合。
		2. 通过声明一个未初始化的映射来创建一个值为 nil 的映射 （称为 nil 映射 ）。映射不能用于存储键值对。
		3. 映射的增长没有容量或者任何限制。
		4. 映射的键可以是任何值。这个值的类型可以是内置的类型，也可以是结构类型，只要这个值可以
		使用 == 运算符做比较。切片、函数以及包含切片的结构类型这些类型由于具有引用语义，不能作为
		映射的键，使用这些类型会造成编译错误。
	*/

	// 创建一个映射，键的类型是 string，值的类型是 int
	dict01 := make(map[string]int)
	fmt.Printf("dict01:\t%T\t%v\n", dict01, dict01)

	// 创建一个映射，键和值的类型都是 string，使用两个键值对初始化映射
	dict02 := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}
	fmt.Printf("dict02:\t%T\t%v\n", dict02, dict02)

	// 创建一个空映射，用来存储颜色以及颜色对应的十六进制代码
	dict03 := map[string]string{}
	// 将 Red 的代码加入到映射
	dict03["Red"] = "#da1337"
	fmt.Printf("dict03:\t%T\t%v\n", dict03, dict03)

	// 获取键 Blue 对应的值
	value, exists := dict03["Blue"]
	// 这个键存在吗？
	if exists {
		fmt.Println(value)
	}

	fmt.Printf("================================================================================\n")

	// 创建一个映射，存储颜色以及颜色对应的十六进制代码
	dict04 := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}
	fmt.Printf("dict04:\t%T\t%v\n", dict04, dict04)
	// 显示映射里的所有颜色
	for key, value := range dict04 {
		fmt.Printf("Key: %-12s\tValue: %s\n", key, value)
	}
	// 删除键为 Coral 的键值对
	delete(dict04, "Coral")
	fmt.Printf("dict04:\t%T\t%v\n", dict04, dict04)

	fmt.Printf("================================================================================\n")
}
