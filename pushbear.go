package pushbear

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// URLPushbearService api url
const (
	URLPushbearService string = "https://pushbear.ftqq.com/sub"
)

// Client ...
type Client interface {
	Send(m Message) (*Result, error)
}

// Result ...
type Result struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Created string `json:"created"`
}

// Message ...
type Message struct {
	// Title 标题，必填。不超过80个字
	Title string
	// Desp 长文本内容，选填。用户通过点击短信里的链接，打开浏览器阅读。支持Markdown语法，不超过64K
	Desp string
}

// Pushbear client
type Pushbear struct {
	SendKey    string
	httpClient *http.Client
}

var _ Client = &Pushbear{}

// New create new pushbear service client
func New(key string) Client {
	return Pushbear{SendKey: key, httpClient: &http.Client{}}
}

// Send sends message
func (p Pushbear) Send(m Message) (*Result, error) {
	res := Result{}
	req, err := http.NewRequest("GET", URLPushbearService, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("sendkey", p.SendKey)
	q.Add("text", m.Title)
	q.Add("desp", m.Desp)
	req.URL.RawQuery = q.Encode()

	resp, err := p.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer func() {
		if err = resp.Body.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}
