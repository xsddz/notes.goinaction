package mmethod

import (
	"fmt"
	"reflect"
)

// user 在程序里定义一个用户类型
type user struct {
	name  string
	email string
}

/*
notify 使用值接收者实现了一个方法

1. 方法能给用户定义的类型添加新的行为。方法实际上也是函数，只是在声明时，在关键字 func 和方法名
之间增加了一个参数。
2. 关键字 func 和函数名之间的参数被称作接收者，将函数与接收者的类型绑在一起。如果一个函数有接收者，
这个函数就被称为方法。
3. Go 语言里有两种类型的接收者：值接收者和指针接收者。如果使用值接收者声明方法，调用时会使用
这个值的一个副本来执行。
4. 在声明一个新类型之后，声明一个该类型的方法之前，如果给这个类型增加或者删除某个值，是要创建一个
新值，还是要更改当前的值？如果是要创建一个新值，该类型的方法就使用值接收者。如果是要修改当前值，
就使用指针接收者。这个答案也会影响程序内部传递这个类型的值的方式：是按值做传递，还是按指针做传递。
*/
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

// changeEmail 使用指针接收者实现了一个方法
func (u *user) changeEmail(email string) {
	u.email = email
}

// changeName 使用值接收者实现了一个方法
func (u user) changeName(name string) {
	u.name = "u is a copy"
}

// Run method测试入口
func Run() {
	fmt.Printf("================================================================================\n")

	// user 类型的值
	bill := user{"Bill", "bill@email.com"}
	fmt.Printf("bill:\t%T\t%v\n", bill, bill)

	// 指向 user 类型值的指针
	lisa := &user{"Lisa", "lisa@email.com"}
	fmt.Printf("lisa:\t%T\t%v\n", lisa, lisa)

	fmt.Printf("\n")
	nice(bill)

	fmt.Printf("================================================================================\n")

	fmt.Printf("user 类型的值用来调用使用值接收者声明的方法:\n")
	fmt.Printf("before change name:\t%T\t%v\n", bill, bill)
	bill.changeName("billNewName")
	fmt.Printf("after change name:\t%T\t%v\n", bill, bill)

	fmt.Printf("\n")
	bill.notify()
	fmt.Printf("\n")

	fmt.Printf("指向 user 类型值的指针用来调用使用值接收者声明的方法:\n")
	fmt.Printf("before change name:\t%T\t%v\n", lisa, lisa)
	lisa.changeName("lisaNewName")
	fmt.Printf("after change name:\t%T\t%v\n", lisa, lisa)

	fmt.Printf("\n")
	lisa.notify()
	fmt.Printf("\n")

	fmt.Printf("user 类型的值用来调用使用指针接收者声明的方法:\n")
	fmt.Printf("before change email:\t%T\t%v\n", bill, bill)
	bill.changeEmail("billNewEmail")
	fmt.Printf("after change email:\t%T\t%v\n", bill, bill)

	fmt.Printf("\n")

	fmt.Printf("指向 user 类型值的指针用来调用使用指针接收者声明的方法:\n")
	fmt.Printf("before change email:\t%T\t%v\n", lisa, lisa)
	lisa.changeEmail("lisaNewEmail")
	fmt.Printf("after change email:\t%T\t%v\n", lisa, lisa)

	fmt.Printf("================================================================================\n")
}

func nice(i interface{}) {
	t := reflect.TypeOf(i).NumMethod()
	fmt.Printf("===%T\t%v\n", t, t)

}
