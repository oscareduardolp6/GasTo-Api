package share

import (
	"net/http"
)

type MethodRelation = map[string]http.HandlerFunc

func DependingOnMethod(handlers MethodRelation) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler, found := handlers[r.Method]
		if !found {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}
