package bracer

import (
	"github.com/jaak-ai/jaak-japi/v1/tracer"
	"github.com/kamva/mgm/v3"
	"sync"
)

type Bracer struct {
	mu               sync.Mutex
	mgm.DefaultModel `bson:",inline"`
	*tracer.Tracer   `bson:",inline"`
}

func (model *Bracer) CollectionName() string {
	return "tracer"
}

func (model *Bracer) Save() error {
	model.mu.Lock()
	defer model.mu.Unlock()

	if model.ID.IsZero() {
		return mgm.Coll(model).Create(model)
	} else {
		return mgm.Coll(model).Update(model)
	}
}
