package main

import (
	"context"
	"fmt"

	"github.com/klarkxy/onebot-go/context"
	"github.com/klarkxy/onebot-go/plugin"
)

func main() {
	p := plugin.Plugin{}
	p.OnMsg(func(c *context.Context) {
		fmt.Println("1")
	})
}
