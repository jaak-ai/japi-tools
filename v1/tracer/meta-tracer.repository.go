package tracer

type MetaRequestTracer struct {
	Request  interface{} `bson:"request" json:"request"`
	Response interface{} `bson:"response" json:"response"`
}

func (model *MetaRequestTracer) SetRequest(value interface{}) {
	model.Request = value
}

func (model *MetaRequestTracer) SetResponse(value interface{}) {
	model.Response = value
}
