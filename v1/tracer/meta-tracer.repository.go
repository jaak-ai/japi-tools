package tracer

type MetaRequestTracer struct {
	Body interface{} `bson:"body" json:"body"`
}

func (model *MetaRequestTracer) SetBody(value interface{}) {
	model.Body = value
}
