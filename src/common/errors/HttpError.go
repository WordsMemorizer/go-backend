package errors

import "fmt"

type HttpError struct {
	internal       error
	HttpStatusCode int

	// CommonError part from common/errors.yaml

	BusinessStatusCode string `json:"code"`
	UserMessage        string `json:"user_message"`
	DeveloperMessage   string `json:"developer_message"`
}

func (h *HttpError) Error() string {
	return fmt.Sprintf(
		"Internal %s\nHttpStatusCode: %v\nBusinessStatusCode: %v\nUserMessage: %v\nDeveloperMessage: %v",
		h.internal.Error(), h.HttpStatusCode, h.BusinessStatusCode, h.UserMessage, h.DeveloperMessage)
}
