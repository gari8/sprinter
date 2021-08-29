package sprinter

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

func Handle(fn func(ctx context.Context, r *http.Request) Response) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := fn(r.Context(), r)
		w.WriteHeader(response.Code)
		w.Header().Add("Content-Type", response.ContentType)
		if response.Code < 400 {
			body, _ := json.Marshal(response.Object)
			_, _ = w.Write(body)
		} else {
			body, _ := json.Marshal(response.Err)
			_, _ = w.Write(body)
		}
	}
}

func GetInputByJson(body io.ReadCloser, target interface{}) {
	decoder := json.NewDecoder(body)
	_ = decoder.Decode(&target)
}
