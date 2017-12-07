# go-pushbear

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
	res, err := p.SendMessage(pushbear.Message{Title: "test", Desp: "test"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", res)
}

```