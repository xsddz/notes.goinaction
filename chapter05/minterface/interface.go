package minterface

import (
	"fmt"
)

/*
notifier 是一个定义了通知类行为的接口

1. 接口是用来定义行为的类型。这些被定义的行为不由接口直接实现，而是通过方法由用户定义的类型实现。
如果用户定义的类型实现了某个接口类型声明的一组方法，那么这个用户定义的类型的值就可以赋给这个接口
类型的值。这个赋值会把用户定义的类型的值存入接口类型的值。
2. 接口值是一个两个字长度的数据结构，第一个字包含一个指向内部表的指针。这个内部表叫作 iTable，
包含了所存储的值的类型信息。iTable 包含了已存储的值的类型信息以及与这个值相关联的一组方法。
第二个字是一个指向所存储值的指针。
3. 对接口值方法的调用会执行接口值里存储的用户定义的类型的值对应的方法。因为任何用户定义的类型
都可以实现任何接口，所以对接口值方法的调用自然就是一种多态。在这个关系里，用户定义的类型通常叫作
实体类型，原因是如果离开内部存储的用户定义的类型的值的实现，接口值并没有具体的行为。
4. 如果一个类型实现了某个接口，所有使用这个接口的地方，都可以支持这种类型的值。
5. 方法集定义了一组关联到给定类型的值或者指针的方法。定义方法时使用的接收者的类型决定了这个方法
是关联到值，还是关联到指针，还是两个都关联。

规范里描述的方法集规则：

Values                 Methods Receivers
-----------------------------------------------
  T                      (t T)
 *T                      (t T) and (t *T)

可以看出，T 类型的值的方法集只包含值接收者声明的方法。而指向 T 类型的指针的方法集既包含
值接收者声明的方法，也包含指针接收者声明的方法。从接收者类型的角度来看，

Methods Receivers       Values
-----------------------------------------------
  (t T)                  T and *T
  (t *T)                 *T

如果使用指针接收者来实现一个接口，那么只有指向那个类型的指针才能够实现对应的接口。
如果使用值接收者来实现一个接口，那么那个类型的值和指针都能够实现对应的接口。
*/
type notifier interface {
	notify()
}

// user 在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

// notify 是使用指针接收者实现的方法
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

// admin 定义了程序里的管理员
type admin struct {
	name  string
	email string
}

// notify 使用值接收者实现了 notifier 接口
func (a admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.name, a.email)
}

// Run interface测试入口
func Run() {
	fmt.Printf("================================================================================\n")

	// 创建一个 user 值并传给 sendNotification
	bill := user{"Bill", "bill@email.com"}
	// 这里不能直接传入变量u，否则会报错说 user 类型并没有实现 notifier，因为 notify 方法是使用指针接收者声明的。
	// sendNotification(bill)
	sendNotification(&bill)

	// 创建一个 admin 值并传给 sendNotification
	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotification(lisa)
	sendNotification(&lisa)

	fmt.Printf("================================================================================\n")
}

// sendNotification 接受一个实现了 notifier 接口的值，并发送通知
func sendNotification(n notifier) {
	n.notify()
}
