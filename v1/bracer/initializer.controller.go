package bracer

import (
	"github.com/jaak-ai/jaak-japi/v1/tracer"
)

func New(action string) *Bracer {
	b := &Bracer{
		Tracer: tracer.New(action),
	}
	return b
}
