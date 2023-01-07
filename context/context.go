package main

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	Cancel()
}

type Substore struct {
	response string
}

func (s Substore) Fetch() (string, error) {
	return s.response, nil
}

func (s Substore) Cancel() {

}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())
		if err != nil {
			return
		}

		fmt.Fprint(w, data)
	}
}

// func ClassicServer(store Store) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		store.Cancel()
// 		fmt.Fprint(w, store.Fetch(r.Context()))
// 	}
// }
