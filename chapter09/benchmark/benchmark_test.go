/*
用来检测要将整数值转为字符串，使用哪个函数会更好的基准测试示例。先使用 fmt.Sprintf 函数，
然后使用strconv.FormatInt 函数，最后使用 strconv.Itoa

1. 和单元测试文件一样， 基准测试的文件名也必须以 _test.go 结尾。
*/
package benchmark_test

import (
	"fmt"
	"strconv"
	"testing"
)

/*
BenchmarkSprintf 对 fmt.Sprintf 函数进行基准测试

eg:
	go test -v -run="none" -bench="BenchmarkSprintf" -benchtime="3s"
	go test -v -run="none" -bench=. -benchtime="3s"

	-bench 指定基准测试函数（可以接受正则表达式，来决定需要运行哪些测试）。
	-run 指定单元测试函数（可以接受正则表达式，来决定需要运行哪些测试）。这里字符串 "none" ，来保证在运行制订的基准测试函数之前没有单元测试会被运行。由于例子里没有单元测试函数的名字中有 none ，所以使用 none 可以排除所有的单元测试。
	-benchtime 用来更改测试执行的最短时间。
	-benchmem 提供每次操作分配内存的次数，以及总共分配内存的字节数。单位为 allocs/op 的值表示每次操作从堆上分配内存的次数。单位为 B/op 的值表示每次操作分配的字节数。

1. 基准测试函数必须以 Benchmark 开头，接受一个指向 testing.B 类型的指针作为唯一参数。
*/
func BenchmarkSprintf(b *testing.B) {
	number := 10

	// 这个方法用来重置计时器，保证测试代码执行前的初始化代码，不会干扰计时器的结果。
	// 为了保证得到的测试结果尽量精确，需要使用这个函数来跳过初始化代码的执行时间。
	b.ResetTimer()

	// 基准测试框架默认会在持续 1 秒的时间内，反复调用需要测试的函数。测试框架每次调用测试函数时，
	// 都会增加 b.N 的值。第一次调用时，b.N 的值为 1 。需要注意，一定要将所有要进行基准测试的
	// 代码都放到循环里，并且循环要使用 b.N 的值。否则，测试的结果是不可靠的。
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", number)
	}
}

// BenchmarkFormat 对 strconv.FormatInt 函数进行基准测试
func BenchmarkFormat(b *testing.B) {
	number := int64(10)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.FormatInt(number, 10)
	}
}

// BenchmarkItoa 对 strconv.Itoa 函数进行基准测试
func BenchmarkItoa(b *testing.B) {
	number := 10

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		strconv.Itoa(number)
	}
}
