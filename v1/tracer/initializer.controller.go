package tracer

func New(action string) *Tracer {
	return &Tracer{
		Request:     nil,
		Action:      action,
		Status:      FailedStatusTracer,
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
