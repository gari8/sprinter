package middleware

type middleware struct{}

type Middleware interface{}

func NewMiddleware() Middleware {
	return &middleware{}
}
