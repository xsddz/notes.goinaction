package mtypeembedding

import "fmt"

// notifier 是一个定义了通知类行为的接口
type notifier interface {
	notify()
}

// user 在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

// notify 实现了一个可以通过 user 类型值的指针调用的方法
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

/*
admin 代表一个拥有权限的管理员用户

1. 嵌入类型是将已有的类型直接声明在新的结构类型里。被嵌入的类型被称为新的外部类型的内部类型。
要嵌入一个类型，只需要声明这个类型的名字就可以了。
2. 通过嵌入类型，与内部类型相关的标识符会提升到外部类型上。这样外部类型就组合了内部类型包含的
所有属性，并且可以添加新的字段和方法。外部类型也可以通过声明与内部类型标识符同名的标识符来覆盖
内部标识符的字段或者方法。
3. 对外部类型来说，内部类型总是存在的。这就意味着，虽然没有指定内部类型对应的字段名，还是可以
使用内部类型的类型名，来访问到内部类型的值。
*/
type admin struct {
	user  // 嵌入类型
	level string
}

// notify 这里重新实现，通过 admin 类型值的指针调用的方法
// func (a *admin) notify() {
// 	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
// }

// Run type embedding 测试入口
func Run() {
	fmt.Printf("================================================================================\n")

	// 创建一个 admin 用户
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// 我们可以直接访问内部类型的方法
	ad.user.notify()

	/*
		内部类型的方法也被提升到外部类型

		1. 由于内部类型的提升，内部类型实现的接口会自动提升到外部类型。这意味着由于内部类型的实现，
		外部类型也同样实现了这个接口。
		2. 如果外部类型实现了 notify 方法，内部类型的实现就不会被提升。不过内部类型的值一直存在，
		因此还可以通过直接访问内部类型的值，来调用没有被提升的内部类型实现的方法。
	*/
	ad.notify()
	// 给 admin 用户发送一个通知，用于实现接口的内部类型的方法，被提升到外部类型
	sendNotification(&ad)

	fmt.Printf("================================================================================\n")
}

// sendNotification 接受一个实现了 notifier 接口的值，并发送通知
func sendNotification(n notifier) {
	n.notify()
}
