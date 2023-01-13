package auction

import (
	"net/http"
)

type RestGet interface {
	Index(r *http.Request) interface{}
}

type RestPost interface {
	Create(r *http.Request) interface{}
}

type RestPut interface {
	Update(r *http.Request) interface{}
}
