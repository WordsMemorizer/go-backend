package middlewares

import (
	"fmt"
	"github.com/WordsMemorizer/go-backend/common/router"
	"log"
	"net/http"
)

type stdoutLoggerWriterProxy struct {
	http.ResponseWriter
	code int
	data []byte
}

func (s *stdoutLoggerWriterProxy) WriteHeader(code int) {
	s.code = code
	s.ResponseWriter.WriteHeader(code)
}

func (s *stdoutLoggerWriterProxy) Write(data []byte) (int, error) {
	if s.isError() {
		s.data = data
	}
	return s.ResponseWriter.Write(data)
}

func (s *stdoutLoggerWriterProxy) isError() bool { return s.code >= 400 }

// StdoutLoggerMiddleware is http.HandleFunc level proxy
// it will print information about request and error (if response code is error code)
var StdoutLoggerMiddleware = func(next router.HttpHandlerFunc) router.HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		proxy := &stdoutLoggerWriterProxy{w, 0, nil}

		next(proxy, r)

		req := fmt.Sprintf("%s %s -> %v; Headers: %v", r.Method, r.URL, proxy.code, r.Header)

		if proxy.isError() {

			errStr := ""

			if len(proxy.data) != 0 {
				errStr = string(proxy.data)
			}

			response := fmt.Sprintf("Headers: %v; Error: %s", w.Header(), errStr)

			log.Println(req, response)
		}
		log.Println(req)
	}
}
