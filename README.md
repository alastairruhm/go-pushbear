# go-pushbear

## example

```golang
package main

import (
	"fmt"

	"github.com/alastairruhm/go-pushbear"
)

func main() {
	key := "1666-8a2e7c92921db9a4e563a19e77c0c730"
	p := pushbear.New(key)
	res, err := p.Send(pushbear.Message{Title: "test1", Desp: "test"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v", res)
}

```