package search

import (
	"fmt"
	"log"
)

// Result 保存搜索的结果
type Result struct {
	Field   string
	Content string
}

/*
Matcher 定义了要实现的新搜索类型的行为

1. 命名接口的时候，也需要遵守 Go 语言的命名惯例。如果接口类型只包含一个方法，那么这个类型的名字
以 er 结尾。我们的例子里就是这么做的，所以这个接口的名字叫作 Matcher 。如果接口类型内部声明了
多个方法，其名字需要与其行为关联。
2. 一个接口的行为最终由在这个接口类型中声明的方法决定。
3. 如果要让一个用户定义的类型实现一个接口，这个用户定义的类型要实现接口类型里声明的所有方法。
*/
type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

// Match 函数，为每个数据源单独启动 goroutine 来执行这个，函数并发地执行搜索
func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// 对特定的匹配器执行搜索
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// 将结果写入通道
	for _, result := range searchResults {
		results <- result
	}
}

// Display 从每个单独的 goroutine 接收到结果后在终端窗口输出
func Display(results chan *Result) {
	// 通道会一直被阻塞，直到有结果写入
	// 一旦通道被关闭，for 循环就会终止
	for result := range results {
		fmt.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
