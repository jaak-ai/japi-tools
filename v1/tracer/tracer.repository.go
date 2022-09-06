package tracer

import (
	"encoding/json"
)

type Tracer struct {
	Request     *RequestTracer `bson:"request" json:"request"`
	Action      string         `bson:"action" json:"action"`
	Status      StatusTracer   `bson:"status" json:"status"`
	Message     string         `bson:"message" json:"message"`
	ProcessTime float32        `bson:"process_time" json:"processTime"`
	Flow        []*StepTracer  `bson:"flow" json:"flow"`
}

func (model *Tracer) SetRequest(value *RequestTracer) {
	model.Request = value
}

func (model *Tracer) SetStatus(value StatusTracer) {
	model.Status = value
}

func (model *Tracer) SetMessage(value string) {
	model.Message = value
}

func (model *Tracer) SetProcessTime(value float32) {
	model.ProcessTime = value
}

func (model *Tracer) AddStep(value *StepTracer) {
	model.Flow = append(model.Flow, value)
}

func (model *Tracer) ToString(indentList ...int) (string, error) {
	//if model.CreatedAt.IsZero() {
	//	model.CreatedAt = time.Now()
	//}
	//
	//if model.UpdatedAt.IsZero() {
	//	model.UpdatedAt = time.Now()
	//}

	indent := 0
	if len(indentList) > 0 {
		indent = indentList[0]
	}

	if indent <= 0 {
		str, err := json.Marshal(model)
		if err != nil {
			return "", err
		}

		return string(str), nil
	}

	indentString := ""
	for i := 0; i < indent; i++ {
		indentString = indentString + " "
	}

	str, err := json.MarshalIndent(model, "", indentString)
	if err != nil {
		return "", err
	}
	return string(str), nil
}
