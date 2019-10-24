// 这个示例程序展示如何内部模仿 HTTP GET 调用
package thetest

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// feed 模仿了我们期望接收的 XML 文档
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
	<title>Going Go Programming</title>
	<description>Golang : https://github.com/goinggo</description>
	<link>http://www.goinggo.net/</link>
	<item>
		<pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
		<title>Object Oriented Programming Mechanics</title>
		<description>Go is an object oriented language.</description>
		<link>http://www.goinggo.net/2015/03/object-oriented</link>
	</item>
</channel>
</rss>`

// mockServer 返回用来处理请求的服务器的指针
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("content-type", "application/xml")
		fmt.Fprintln(w, feed)
	}

	// type HandlerFunc func(ResponseWriter, *Request)
	// HandlerFunc 类型是一个适配器，允许常规函数作为 HTTP 的处理函数使用。
	// 如果函数 f 具有合适的签名，HandlerFunc(f)就是一个处理 HTTP 请求的 Handler 对象，
	// 内部通过调用 f 处理请求
	return httptest.NewServer(http.HandlerFunc(f))
}

// TestDownloadMocking 确认 http 包的 Get 函数可以下载内容，并且内容可以被正确地反序列化并关闭
func TestDownloadMocking(t *testing.T) {
	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close()

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"", server.URL, statusCode)
		{
			// 由 httptest.Server 值提供了请求的 URL
			// 包 http 与包 httptest 和模仿服务器结合在一起，知道如何通过 URL 路由到我们自己的处理函数
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.", ballotX, err)
			}
			t.Log("\t\tShould be able to make the Get call.", checkMark)

			defer resp.Body.Close()

			if resp.StatusCode != statusCode {
				t.Fatalf("\t\tShould receive a \"%d\" status. %v %v", statusCode, ballotX, resp.StatusCode)
			}
			t.Logf("\t\tShould receive a \"%d\" status. %v", statusCode, checkMark)
		}
	}
}
