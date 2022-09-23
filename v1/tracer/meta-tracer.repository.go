package tracer

type MetaTracer struct {
	Extra interface{} `bson:"extra" json:"extra"`
}

func (model *MetaTracer) SeExtra(value interface{}) {
	model.Extra = value
}
