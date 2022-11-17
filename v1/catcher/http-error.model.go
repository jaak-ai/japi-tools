package catcher

import (
	"fmt"
)

type Error struct {
	EventID string `json:"eventId"`
	Status  int    `json:"status"`
	Code    int32  `json:"code"`
	Message string `json:"message"`
}

func (model *Error) Error() string {
	return fmt.Sprintf("%s", model.Message)
}
