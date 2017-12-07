package pushbear

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
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
	Created time.Time
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

// SendMessage ...
func (p *Pushbear) SendMessage(m Message) (*Result, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", URLPushbearService, nil)
	if err != nil {
		return nil, ErrorParamTitleEmpty
	}

	q := req.URL.Query()
	q.Add("sendkey", p.SendKey)
	q.Add("text", m.Title)
	q.Add("desp", m.Desp)

	resp, err := client.Do(req)

	if err != nil {
		return nil, ErrorParamTitleEmpty
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, ErrorParamTitleEmpty
	}

	res := Result{}

	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, ErrorParamTitleEmpty
	}

	return &res, nil
}

var _ Client = &Pushbear{}
