package tracer

type StepTracer struct {
	Resource ResourceStepTracer `bson:"resource" json:"resource"`
	Status   StatusStepTracer   `bson:"status" json:"status"`
	Message  string             `bson:"message" json:"message"`
	Meta     *MetaStepTracer    `bson:"meta" json:"meta"`
}

func (model *StepTracer) SetResource(value ResourceStepTracer) {
	model.Resource = value
}

func (model *StepTracer) SetStatus(value StatusStepTracer) {
	model.Status = value
}

func (model *StepTracer) SetMessage(value string) {
	model.Message = value
}

func (model *StepTracer) SetMeta(value *MetaStepTracer) {
	model.Meta = value
}
