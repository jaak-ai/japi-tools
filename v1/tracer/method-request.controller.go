package tracer

import (
	"strings"
)

func MapperMethodRequest(method string) MethodRequestTracer {
	switch strings.ToLower(method) {
	case "get":
		return GetStatusTracer
	case "post":
		return PostStatusTracer
	case "delete":
		return DeleteStatusTracer
	case "put":
		return PutStatusTracer
	case "patch":
		return PatchStatusTracer
	}

	return ""
}
