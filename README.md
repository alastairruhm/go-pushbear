# go-pushbear

go wrapper for pushbear service: https://pushbear.ftqq.com/admin/#/


## example

```golang
package main

import (
	"fmt"

	"github.com/alastairruhm/go-pushbear"
)

func main() {
	key := "xxxx-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	p := pushbear.New(key)
	res, err := p.Send(pushbear.Message{Title: "test1", Desp: "test"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", res)
}

```