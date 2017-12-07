package pushbear

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// URLPushbearService api url
const URLPushbearService string = "https://pushbear.ftqq.com/sub"

// Client ...
type Client interface {
	SendMessage(m Message) (*Result, error)
}

// Result ...
type Result struct {
	Code    int
	Message string
	Data    string
	Created string
}

// Message ...
type Message struct {
	// Title 标题，必填。不超过80个字
	Title string
	// Desp 长文本内容，选填。用户通过点击短信里的链接，打开浏览器阅读。支持Markdown语法，不超过64K
	Desp string
}

// Pushbear ...
type Pushbear struct {
	SendKey string
}

// New ...
func New(key string) Client {
	return Pushbear{SendKey: key}
}

// SendMessage ...
func (p Pushbear) SendMessage(m Message) (*Result, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", URLPushbearService, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("sendkey", p.SendKey)
	q.Add("text", m.Title)
	q.Add("desp", m.Desp)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	res := Result{}

	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

var _ Client = &Pushbear{}
