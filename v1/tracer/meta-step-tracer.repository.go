package tracer

type MetaStepTracer struct {
	ProcessTime float32 `bson:"process_time" json:"processTime"`
}

func (model *MetaStepTracer) setProcessTime(value float32) {
	model.ProcessTime = value
}
