package lesson

import (
	"net/http"
	"net/http/httptest"
)

// 中间件 http://blog.jobbole.com/53265/

type ModifierMiddleware struct {
	handler http.Handler
}

func (this *ModifierMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rec := httptest.NewRecorder()
	this.handler.ServeHTTP(rec, r)

	for k, v := range rec.Header() {
		w.Header()[k] = v
	}

	w.Header().Set("Go-web-foundation", "vip")
	w.WriteHeader(418)
	w.Write(rec.Body.Bytes())
	w.Write([]byte("hey, this is middleware!"))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func InitHttpServer() {
	mid := &ModifierMiddleware{http.HandlerFunc(myHandler)}
	http.ListenAndServe(":8080", mid)
}
