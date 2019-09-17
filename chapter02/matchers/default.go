package matchers

import (
	"notes.goinaction/chapter02/search"
)

/*
defaultMatcher 实现了默认匹配器

空结构在创建实例时，不会分配任何内存。
*/
type defaultMatcher struct{}

// init 函数将默认匹配器注册到程序里
func init() {
	var matcher defaultMatcher
	search.Register("default", matcher)
}

/*
Search 实现了默认匹配器的行为

1. 如果声明函数的时候带有接收者，则意味着声明了一个方法。这个方法会和指定的接收者的类型绑在一起。
2. 可以使用 defaultMatcher 类型的值或者指向这个类型值的指针来调用 Search 方法。无论我们是
使用接收者类型的值来调用这个方，还是使用接收者类型值的指针来调用这个方法，编译器都会正确地引用
或者解引用对应的值，作为接收者传递给 Search 方法。
3. 因为大部分方法在被调用后都需要维护接收者的值的状态，所以，一个最佳实践是，将方法的接收者声明为指针。
对于 defaultMatcher 类型来说，使用值作为接收者是因为创建一个 defaultMatcher 类型的值不需要分配内存。
由于 defaultMatcher 不需要维护状态，所以不需要指针形式的接收者。
4. 使用指针作为接收者声明的方法，只能在接口类型的值是一个指针的时候被调用。使用值作为接收者声明的方法，
在接口类型的值为值或者指针时，都可以被调用。
*/
func (m defaultMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	return nil, nil
}
