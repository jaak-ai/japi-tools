package tracer

type RequestTracer struct {
	Ip         string              `bson:"ip" json:"ip"`
	Method     MethodRequestTracer `bson:"method" json:"method"`
	StatusCode int16               `bson:"status_code" json:"statusCode"`
	Path       string              `bson:"status_code" json:"path"`
	Meta       *MetaRequestTracer  `bson:"meta" json:"meta"`
}

func (model *RequestTracer) SetIp(value string) {
	model.Ip = value
}

func (model *RequestTracer) SetMethod(value MethodRequestTracer) {
	model.Method = value
}

func (model *RequestTracer) SetStatusCode(value int16) {
	model.StatusCode = value
}

func (model *RequestTracer) SetPath(value string) {
	model.Path = value
}

func (model *RequestTracer) SetMeta(value *MetaRequestTracer) {
	model.Meta = value
}
