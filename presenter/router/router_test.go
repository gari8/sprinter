package router_test

import (
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"reflect"
	"sprinter/presenter/router"
	"testing"
)

func TestNewRouter(t *testing.T) {
	route := chi.NewRouter()
	wantNewRouterForm := &router.Server{
		Route: route,
	}

	r := router.NewRouter()

	v := reflect.ValueOf(r)
	w := reflect.ValueOf(wantNewRouterForm)

	assert.Equal(t, v.Type(), w.Type())
}