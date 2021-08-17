package middlewares

import (
	"github.com/WordsMemorizer/go-backend/common/errors"
	"github.com/WordsMemorizer/go-backend/common/router"
	"net/http"
)

// CommonErrorMapperMiddleware converts all errors which are not errors.HttpError to errors.HttpError
var CommonErrorMapperMiddleware = func(next router.HandlerFunc) router.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
		val, err := next(w, r)

		if err == nil {
			return val, err
		}

		if _, ok := err.(*errors.HttpError); !ok {
			return val, &errors.HttpError{
				HttpStatusCode:     http.StatusInternalServerError,
				BusinessStatusCode: errors.ErrorCodeServerBroke,
				DeveloperMessage:   err.Error(),
			}
		}

		return val, err
	}
}
