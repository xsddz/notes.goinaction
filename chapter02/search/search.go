package search

import (
	"log"
	"sync"
)

/*
1. 在 Go 语言里，标识符要么从包里公开，要么不从包里公开。当代码导入了一个包时，程序可以
直接访问这个包中任意一个公开的标识符。这些标识符以大写字母开头。以小写字母开头的标识符是
不公开的，不能被其他包中的代码直接访问。
2. 这个变量没有定义在任何函数作用域内，所以会被当成包级变量。
*/
var matchers = make(map[string]Matcher)

// Run 执行搜索逻辑
func Run(searchTerm string) {
	/*
		创建一个无缓冲的通道，接收匹配后的结果

		1. 简化变量声明运算符（ := ）用于声明一个变量，同时给这个变量赋予初始值。
		2. 根据经验，如果需要声明初始值为零值的变量，应该使用 var 关键字声明变量；如果提供确切的
		非零值初始化变量或者使用函数返回值创建变量，应该使用简化变量声明运算符。
	*/
	results := make(chan *Result)

	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	/*
		构造一个 waitGroup，以便处理所有的数据源

		1. 在 Go 语言中，如果 main 函数返回，整个程序也就终止了。Go 程序终止时，还会关闭所有
		之前启动且还在运行的 goroutine。写并发程序的时候，最佳做法是，在 main 函数返回前，
		清理并终止所有之前启动的 goroutine。编写启动和终止时的状态都很清晰的程序，有助减少 bug，
		防止资源异常。
		2. WaitGroup 是一个计数信号量，我们可以利用它来统计所有的 goroutine 是不是都完成了工作。
	*/
	var waitGroup sync.WaitGroup
	// 设置需要等待处理每个数据源的 goroutine 的数量
	waitGroup.Add(len(feeds))

	/*
		为每个数据源启动一个 goroutine 来查找结果

		1. 关键字 range 可以用于迭代数组、字符串、切片、映射和通道。使用 for range 迭代切片时，
		每次迭代会返回两个值。第一个值是迭代的元素在切片里的索引位置，第二个值是元素值的一个副本。
		2. 下划线标识符的作用是占位符，占据了保存 range 调用返回的索引值的变量的位置。如果
		要调用的函数返回多个值，而又不需要其中的某个值，就可以使用下划线标识符将其忽略。
	*/
	for _, feed := range feeds {
		/*
			获取一个匹配器用于查找

			查找 map 里的键时，有两个选择：要么赋值给一个变量，要么为了精确查找，赋值给两个变量。
			赋值给两个变量时第一个值和赋值给一个变量时的值一样，是 map 查找的结果值。如果指定了
			第二个值，就会返回一个布尔标志，来表示查找的键是否存在于 map 里。如果这个键不存在，
			map 会返回其值类型的零值作为返回值，如果这个键存在，map 会返回键所对应值的副本。
		*/
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		/*
			使用关键字 go 启动一个 goroutine 来执行搜索

			1. 匿名函数是指没有明确声明名字的函数。
			2. 在 Go 语言中，所有的变量都以值的方式传递。因为指针变量的值是所指向的内存地址，
			在函数间传递指针变量，是在传递这个地址值，所以依旧被看作以值的方式在传递。
			3. 通过闭包，函数可以直接访问到那些没有作为参数传入的变量。匿名函数并没有拿到这些
			变量的副本，而是直接访问外层函数作用域中声明的这些变量本身。因为 matcher 和 feed
			变量每次调用时值不相同，所以并没有使用闭包的方式访问这两个变量。
		*/
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)

			// 每个 goroutine 完成其工作后，递减 WaitGroup 变量的计数值，当这个值递减到 0 时，
			// 我们就知道所有的工作都做完了。
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个 goroutine 来监控是否所有的工作都做完了
	go func() {
		/*
			等候所有任务完成

			调用 WaitGroup 的 Wait 方法。这个方法会导致 goroutine 阻塞，直到 WaitGroup 内部的计数到达 0。
		*/
		waitGroup.Wait()

		// 用关闭通道的方式，通知 Display 函数可以退出程序了
		close(results)
	}()

	log.Println("Display Result:")
	// 启动函数，显示返回的结果，并且在最后一个结果显示完后返回
	Display(results)
}

// Register 调用时，会注册一个匹配器，提供给后面的程序使用
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
