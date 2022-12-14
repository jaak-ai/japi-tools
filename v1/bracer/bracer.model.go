package bracer

import (
	"github.com/jaak-ai/jaak-japi/v1/tracer"
	"github.com/kamva/mgm/v3"
)

type Bracer struct {
	mgm.DefaultModel `bson:",inline"`
	*tracer.Tracer   `bson:",inline"`
}

func (model *Bracer) CollectionName() string {
	return "tracer"
}

func (model *Bracer) Save() error {
	return mgm.Coll(model).Create(model)
}
