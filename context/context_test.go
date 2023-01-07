package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type Spystore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *Spystore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

func (s *Spystore) Cancel() {
	s.cancelled = true
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestServer(t *testing.T) {
	data := "Hello, world!"

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &Spystore{response: data, cancelled: false, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		cancel()
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})

	t.Run("returns data from store", func(t *testing.T) {
		store := &Spystore{data, false, t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if !response.written {
			t.Error("response should have been written")
		}

		// store.assertWasNotCancelled()
	})
}

func (s *Spystore) assertWasCancelled() {
	s.t.Helper()

	if !s.cancelled {
		s.t.Error("store was not told to cancel")
	}
}

func (s *Spystore) assertWasNotCancelled() {
	s.t.Helper()

	if s.cancelled {
		s.t.Error("store was told to cancel")
	}
}
