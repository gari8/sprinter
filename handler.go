package sprinter

import (
	"context"
	"encoding/json"
	"net/http"
)

func Handle (fn func(ctx context.Context, r *http.Request) Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := fn(r.Context(), r)
		b, _ := json.Marshal(response)
		w.WriteHeader(response.Code)
		_, _ = w.Write(b)
	}
}
