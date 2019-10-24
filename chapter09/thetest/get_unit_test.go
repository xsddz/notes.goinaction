/*
这个示例程序展示如何写基础单元测试

1. Go 语言的测试工具只会认为以_test.go 结尾的文件是测试文件。如果没有遵从这个约定，
在包里运行 go test 的时候就可能会报告没有测试文件。
2. 如果执行 go test 的时候没有加入冗余选项（ -v ），除非测试失败，否则我们是看不到任何测试输出的。
3. 如果测试函数执行时没有调用过 t.Fatal 或者 t.Error 方法，就会认为测试通过了。
*/
package thetest

import (
	"net/http"
	"testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

/*
 TestDownloadUnit 确认 http 包的 Get 函数可以下载内容

1. 一个测试函数必须是公开的函数，并且以 Test 单词开头。不但函数名字要以 Test 开头，
而且函数的签名必须接收一个指向 testing.T 类型的指针，并且不返回任何值。如果没有遵守这些约定，
测试框架就不会认为这个函数是一个测试函数，也不会让测试工具去执行它。
*/
func TestDownloadUnit(t *testing.T) {
	url := "http://www.goinggo.net/feeds/posts/default?alt=rss"
	statusCode := 200

	t.Log("Given the need to test downloading content.")
	{
		// 每个测试函数都应该通过解释这个测试的给定要求（given need），来说明为什么应该存在这个测试。
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", url, statusCode)
		{
			resp, err := http.Get(url)
			if err != nil {
				// t.Fatal 方法，让测试框架知道这个测试失败了。
				// t.Fatal 方法不但报告这个单元测试已经失败，而且会向测试输出写一些消息，
				// 而后立刻停止这个测试函数的执行。如果除了这个函数外还有其他没有执行的测试函数，
				// 会继续执行其他测试函数。这个方法对应的格式化版本名为 t.Fatalf 。
				t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call", checkMark)

			defer resp.Body.Close()

			if resp.StatusCode == statusCode {
				t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
			} else {
				// 如果需要报告测试失败，但是并不想停止当前测试函数的执行，可以使用 t.Error 系列方法
				t.Errorf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
			}
		}
	}
}
