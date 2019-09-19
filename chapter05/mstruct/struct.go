package mstruct

import "fmt"

// user 在程序里定义一个用户类型
type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

// admin 需要一个 user 类型作为管理者，并附加权限
type admin struct {
	person user
	level  string
}

/*
Duration 基于 int64 声明一个新类型

虽然 int64 是基础类型，Go 并不认为 Duration 和 int64 是同一种类型。这两个类型是完全不同的有区别的类型。
*/
type Duration int64

// Run struct测试入口
func Run() {
	fmt.Printf("================================================================================\n")

	// 声明 user 类型的变量
	var u01 user
	fmt.Printf("u01:\t%T\t%v\n", u01, u01)

	// 声明 user 类型的变量，并初始化所有字段
	u02 := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}
	fmt.Printf("u02:\t%T\t%v\n", u02, u02)

	// 声明 user 类型的变量
	u03 := user{"Lisa", "lisa@email.com", 123, true}
	fmt.Printf("u03:\t%T\t%v\n", u03, u03)

	fmt.Printf("================================================================================\n")

	// 声明 admin 类型的变量
	a01 := admin{
		person: user{
			name:       "Lisa",
			email:      "lisa@email.com",
			ext:        123,
			privileged: true,
		},
		level: "super",
	}
	fmt.Printf("a01:\t%T\t%v\n", a01, a01)

	fmt.Printf("================================================================================\n")
}
