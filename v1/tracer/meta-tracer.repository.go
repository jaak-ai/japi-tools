package tracer

import "github.com/kamva/mgm/v3"

type MetaRequestTracer struct {
	mgm.DefaultModel
	Body interface{} `bson:"body" json:"body"`
}

func (model *MetaRequestTracer) SetBody(value interface{}) {
	model.Body = value
}
