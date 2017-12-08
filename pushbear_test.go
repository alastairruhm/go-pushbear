package pushbear

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

var (
	mux *http.ServeMux

	client *Client

	server *httptest.Server
)

func TestSend(t *testing.T) {
	Convey("Test Send function", t, func() {
		// some global val
		push := New("9999-99999999999999999999999999999999")
		Convey("message title should not be empty", func() {
			_, err := push.Send(Message{Title: "", Desp: "test"})
			So(err, ShouldEqual, ErrorParamTitleEmpty)
		})
	})
}

func TestMockServer(t *testing.T) {
	Convey("fake server", t, func() {
		mux = http.NewServeMux()
		server = httptest.NewServer(mux)
		url, _ := url.Parse(server.URL)
		key := "9999-99999999999999999999999999999999"
		push := New(key)
		push.BaseURL = url.String()

		// Some difference: func() -> func(c C) ?
		Convey("message title should not be empty", func(c C) {
			res := `
			{
				"code": 0,
				"message": "",
				"data": "1条消息已成功推送到发送队列",
				"created": "2017-12-08 11:02:02"
			}`
			mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
				q := r.URL.Query()
				c.So(q.Get("sendkey"), ShouldEqual, key)
				c.So(q.Get("text"), ShouldEqual, "title")
				c.So(q.Get("desp"), ShouldEqual, "desp")
				fmt.Fprint(w, res)
			})
			result, err := push.Send(Message{Title: "title", Desp: "desp"})
			So(result.Code, ShouldEqual, 0)
			So(err, ShouldBeNil)
		})

		Reset(func() {
			server.Close()
		})
	})
}
