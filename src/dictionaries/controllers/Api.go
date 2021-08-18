package controllers

import (
	"encoding/json"
	"github.com/WordsMemorizer/go-backend/common/errors"
	"github.com/WordsMemorizer/go-backend/common/router"
	"github.com/WordsMemorizer/go-backend/common/router/middlewares"
	"github.com/WordsMemorizer/go-backend/dictionaries/entries"
	"net/http"
)

type DictionaryService interface {
	GetDictionary(id string) (entries.Dictionary, error)
	DeleteDictionary(id string) error
	CreateDictionary(dictionary entries.Dictionary) (entries.Dictionary, error)
}

type Api struct {
	DictionaryService
}

func (a *Api) Init(r router.Router) {
	router.HandlerFunc(a.GetSpecificDictionaryHandler).
		Map(middlewares.JsonMapperMiddleware).
		Proxy(middlewares.StdoutLoggerMiddleware).
		RegisterIn(r, http.MethodGet, "/dictionaries/{id}")

	router.HandlerFunc(a.DeleteSpecificDictionaryHandler).
		Map(middlewares.JsonMapperMiddleware).
		Proxy(middlewares.StdoutLoggerMiddleware).
		RegisterIn(r, http.MethodDelete, "/dictionaries/{id}")

	router.HandlerFunc(a.CreateNewDictionaryHandler).
		Map(middlewares.JsonMapperMiddleware).
		Proxy(middlewares.StdoutLoggerMiddleware).
		RegisterIn(r, http.MethodPost, "/dictionaries")
}

func (a *Api) GetSpecificDictionaryHandler(_ http.ResponseWriter, r *http.Request) (interface{}, error) {
	panic("Implement me")
}

func (a *Api) DeleteSpecificDictionaryHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	panic("Implement me")
}

func (a *Api) CreateNewDictionaryHandler(w http.ResponseWriter, r *http.Request) (interface{}, error) {

	var dict = entries.Dictionary{}

	if err := json.NewDecoder(r.Body).Decode(&dict); err != nil {
		return nil, &errors.HttpError{
			HttpStatusCode:   http.StatusBadRequest,
			DeveloperMessage: "Dictionary json parsing fall with error" + err.Error(),
		}
	}

	return a.DictionaryService.CreateDictionary(dict)
}
