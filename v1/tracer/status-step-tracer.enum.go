package tracer

type StatusStepTracer string

const (
	SuccessStatusStepTracer     StatusStepTracer = "success"
	FailedStatusStepTracer      StatusStepTracer = "failed"
	UnprocessedStatusStepTracer StatusStepTracer = "unprocessed"
)
