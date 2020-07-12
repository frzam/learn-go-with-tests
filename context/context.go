package context

import (
	"context"
	"net/http"
)

type Store interface {
	Fetch(context context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			return
		}
		w.Write([]byte(data))
	}
}
