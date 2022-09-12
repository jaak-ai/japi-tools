package bracer

import (
	"github.com/jaak-ai/jaak-japi/v1/tracer"
)

func New() *Bracer {
	b := &Bracer{
		Tracer: tracer.New(),
	}
	return b
}
