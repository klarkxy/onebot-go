package plugin

import "github.com/klarkxy/onebot-go/interfaces"

type Plugin struct {
	handles []interfaces.Handler
}

func (p *Plugin) Do(ctx *interfaces.Contexter) {
	for _, handle := range p.handles {
		handle.Do(ctx)
	}
}
func (p *Plugin) AddHandler(handles ...interfaces.Handler) {
	p.handles = append(p.handles, handles...)
}
