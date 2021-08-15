package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/WordsMemorizer/go-backend/common/errors"
	"github.com/WordsMemorizer/go-backend/common/router"
	"net/http"
)

// JsonMapperMiddleware contains stateless function which firstly checks error
// if error != nil then casts it to errors.HttpError
// if cast is success then marshall it to JSON and send to Client
// if error isn't errors.HttpError then it will create errors.HttpError from this errors
//
// if error is nil then it will marshall response to json and send it
var JsonMapperMiddleware = func(next router.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		value, err := next(writer, request)

		if err != nil {
			if casted, ok := err.(*errors.HttpError); ok {
				toJson(casted, writer)
				return
			}

			toJson(&errors.HttpError{
				HttpStatusCode:     http.StatusInternalServerError,
				BusinessStatusCode: errors.ErrorCodeServerBroke,
				UserMessage:        "Something went wrong. Please be patient (:",
				DeveloperMessage:   "Unexpected error occurred: " + err.Error(),
			}, writer)
			return
		}

		toJson(value, writer)
	}
}

func toJson(object interface{}, w http.ResponseWriter) {
	raw, err := json.Marshal(object)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		msg := fmt.Sprintf("Can't convert error %v to json with error %s", object, err.Error())
		resultStr := fmt.Sprintf("{\"msg\":\"%s\"", msg)
		writeAndHandleError([]byte(resultStr), w)
		return
	}

	writeAndHandleError(raw, w)
}

func writeAndHandleError(data []byte, w http.ResponseWriter) {
	if _, err := w.Write(data); err != nil {
		fmt.Println("[FATAL] can't write to response writer with error", err.Error())
	}
}
