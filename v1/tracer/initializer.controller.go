package tracer

func New() *Tracer {
	return &Tracer{
		Request:     nil,
		Action:      "",
		Status:      ProcessingStatusTracer,
		Message:     "",
		ProcessTime: 0,
		Flow:        nil,
	}
}

func NewStep(resource ResourceStepTracer) *StepTracer {
	return &StepTracer{
		Resource: resource,
		Status:   UnprocessedStatusStepTracer,
		Message:  "",
		Meta:     nil,
	}
}
