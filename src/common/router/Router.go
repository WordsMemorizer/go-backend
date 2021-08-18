package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

// HandlerFunc is function signature for API controller's methods
// first value from result is what ypu want to return from controller methods
// and the next value is error
//
// What the code will do with response is defined in Mapper
type HandlerFunc func(w http.ResponseWriter, r *http.Request) (interface{}, error)

// HttpHandlerFunc is just an alias on http.HandlerFunc
type HttpHandlerFunc http.HandlerFunc

// Wrap wraps handler by middleware
//
// You can use it like that
//
// ```go
//
// func handler(w http.ResponseWriter, r *http.Request) {
//    fmt.Println(handler)
// }
//
// func InitController(r *Router) {
//    HandlerFunc(handler).Wrap(func(next HandlerFunc) HandlerFunc {
//       return func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
//      	fmt.Println(1)
//       	return next(w, r)
//       }
//    }).Wrap(func(next HandlerFunc) HandlerFunc {
//         return func(w http.ResponseWriter, r *http.Request) (interface{}, error) {
//        	fmt.Println(2)
//         	return next(w, r)
//         }
//    })
// }
// ```
//
// This code will print
// ```
// 2
// 1
// handler
//
// So processing is going from the bottom to the top
func (h HandlerFunc) Wrap(middleware func(next HandlerFunc) HandlerFunc) HandlerFunc {
	return middleware(h)
}

// Map maps internal type HandlerFunc to http.HandlerFunc
// the idea is to map result type from business method (interface{}, error) to http.ResponseWriter
//
// You can use whatever mapper you need
// For example ypu can map response in JSON or in FormUrl or in XML
// For example ypu can map response in JSON or in FormUrl or in XML
func (h HandlerFunc) Map(mapper func(next HandlerFunc) HttpHandlerFunc) HttpHandlerFunc {
	return mapper(h)
}

// Proxy works just like Wrap but for other type of function
func (h HttpHandlerFunc) Proxy(middleware func(next HttpHandlerFunc) HttpHandlerFunc) HttpHandlerFunc {
	return middleware(h)
}

// RegisterIn will add method to router
func (h HttpHandlerFunc) RegisterIn(router Router, method, pattern string) HttpHandlerFunc {
	router.Method(method, pattern, http.HandlerFunc(h))
	return h
}

// Router is simple interface for a http request router
type Router interface {
	http.Handler
	// Method could be used for whatever http method you need
	Method(method, pattern string, h http.HandlerFunc)

	// CreateMethod does nothing. Just returns function. Use it to avoid direct casting and build pipeline
	CreateMethod(method HandlerFunc) HandlerFunc
}

// ChiRouterWrapper implementation of Router based on `chi` package
type ChiRouterWrapper struct {
	chi.Router
}
