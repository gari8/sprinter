package sprinter

import (
	"context"
	"encoding/json"
	"net/http"
)

func Handle (fn func(ctx context.Context) Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := fn(r.Context())
		b, _ := json.Marshal(response)
		w.Write(b)
	}
}
