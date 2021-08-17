package sprinter

import (
	"context"
	"encoding/json"
	"io"
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

func GetInputByJson(body io.ReadCloser, target *interface{}) {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&target)
	if err != nil {
		panic(err)
	}
}
