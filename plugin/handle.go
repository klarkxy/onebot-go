package plugin

import "github.com/klarkxy/onebot-go/interfaces"

type HandleFunc func(*interfaces.Contexter)

func (h HandleFunc) Do(ctx *interfaces.Contexter) {
	h(ctx)
}

func MakeFunc(handle func(*interfaces.Contexter)) HandleFunc {
	return HandleFunc(handle)
}
