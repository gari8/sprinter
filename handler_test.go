package sprinter

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	url = "/test"
)

func TestHandle(t *testing.T) {
	reqBody := bytes.NewBufferString("request body")
	req := httptest.NewRequest(http.MethodPost, url, reqBody)

	got := httptest.NewRecorder()

	fn := func(ctx context.Context, r *http.Request) Response {
		target := struct {
			Param string `json:"param"`
		}{
			Param: "sprinter",
		}
		GetInputByJson(r.Body, target)
		return Response{
			Code:   StatusOK,
			Object: target,
		}
	}

	handlerFunc := Handle(fn)
	handlerFunc.ServeHTTP(got, req)
	expected, err := json.Marshal(Response{
		Code: StatusOK,
		Object: struct {
			Param string `json:"param"`
		}{
			Param: "sprinter",
		},
	})
	assert.NoError(t, err)

	body, err := ioutil.ReadAll(got.Body)
	assert.NoError(t, err)

	assert.Equal(t, StatusOK, got.Code)
	assert.Equal(t, string(expected), string(body))
}

func TestGetInputByJson(t *testing.T) {
	expected := struct {
		Param string `json:"param"`
	}{
		Param: "sprinter",
	}

	reqBody, err := json.Marshal(expected)
	assert.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(reqBody))

	target := struct {
		Param string `json:"param"`
	}{}
	GetInputByJson(req.Body, &target)

	assert.Equal(t, expected, target)
}
