/*
这个示例程序展示如何测试内部服务端点的执行效果

1. 这次包的名字也使用 _test 结尾。如果包使用这种方式命名，测试代码只能访问包里公开的标识符。
即便测试代码文件和被测试的代码放在同一个文件夹 中，也只能访问公开的标识符。
*/
package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"notes.goinaction/chapter09/endpoint/handlers"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// 如果没有在单元测试运行之前初始化路由，那么测试就会遇到 http.StatusNotFound 错误而失败。
func init() {
	handlers.Routes()
}

// TestSendJSON 测试/sendjson 内部服务端点
func TestSendJSON(t *testing.T) {
	t.Log("Given the need to test the SendJSON endpoint.")
	{
		req, err := http.NewRequest("GET", "/sendjson", nil)
		if err != nil {
			t.Fatal("\tShould be able to create a request.", ballotX, err)
		}
		t.Log("\tShould be able to create a request.", checkMark)

		// 调用 httptest.NewRecoder 函数来创建一个 http.ResponseRecorder 值。
		// 有了 http.Request 和 http.ResponseRecoder 这两个值，就直接调用服务默认的
		// 多路选择器（mux）的 ServeHttp 方法。调用这个方法模仿了外部客户端对 /sendjson
		// 服务端点的请求。一旦 ServeHTTP 方法调用完成，http.ResponseRecorder 值
		// 就包含了 SendJSON 处理函数的响应。
		rw := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(rw, req)

		if rw.Code != 200 {
			t.Fatal("\tShould receive \"200\"", ballotX, rw.Code)
		}
		t.Log("\tShould receive \"200\"", checkMark)

		u := struct {
			Name  string
			Email string
		}{}

		if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
			t.Fatal("\tShould decode the response.", ballotX)
		}
		t.Log("\tShould decode the response.", checkMark)

		if u.Name == "李四" {
			t.Log("\tShould have a Name.", checkMark)
		} else {
			t.Error("\tShould have a Name.", ballotX, u.Name)
		}

		if u.Email == "lisi@example.com" {
			t.Log("\tShould have an Email.", checkMark)
		} else {
			t.Error("\tShould have an Email.", ballotX, u.Email)
		}
	}
}
