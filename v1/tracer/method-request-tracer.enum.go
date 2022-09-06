package tracer

type MethodRequestTracer string

const (
	GetStatusTracer    MethodRequestTracer = "get"
	PostStatusTracer   MethodRequestTracer = "post"
	DeleteStatusTracer MethodRequestTracer = "delete"
	PutStatusTracer    MethodRequestTracer = "put"
	PatchStatusTracer  MethodRequestTracer = "patch"
)
