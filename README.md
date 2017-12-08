[![Build Status](https://travis-ci.org/alastairruhm/go-pushbear.svg?branch=master)](https://travis-ci.org/alastairruhm/go-pushbear)
[![Coverage Status](https://coveralls.io/repos/github/alastairruhm/go-pushbear/badge.svg?branch=master)](https://coveralls.io/github/alastairruhm/go-pushbear?branch=master)
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