package tracer

type StatusTracer string

const (
	SuccessStatusTracer    StatusTracer = "success"
	FailedStatusTracer     StatusTracer = "failed"
	ProcessingStatusTracer StatusTracer = "processing"
)
