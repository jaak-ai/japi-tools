package tracer

type MetaStepTracer struct {
	ProcessTime float32     `bson:"process_time" json:"processTime"`
	Extra       interface{} `bson:"extra" json:"extra"`
}

func (model *MetaStepTracer) SetProcessTime(value float32) {
	model.ProcessTime = value
}

func (model *MetaStepTracer) SetExtra(value interface{}) {
	model.Extra = value
}
