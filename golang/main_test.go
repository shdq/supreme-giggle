package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleFib(t *testing.T) {
	tests := []struct {
		name       string
		queryParam string
		wantCode   int
		wantBody   string
	}{
		{
			name:       "Valid Fibonacci Request",
			queryParam: "5",
			wantCode:   http.StatusOK,
			wantBody:   "The Fibonacci number at position 5 is 5\n",
		},
		{
			name:       "Negative Query Parameter",
			queryParam: "-5",
			wantCode:   http.StatusBadRequest,
			wantBody:   "Query parameter 'n' has to be a positive integer\n",
		},
		{
			name:       "Non-integer Query Parameter",
			queryParam: "abc",
			wantCode:   http.StatusBadRequest,
			wantBody:   "Query parameter 'n' has to be a positive integer\n",
		},
		{
			name:       "Missing Query Parameter",
			queryParam: "",
			wantCode:   http.StatusBadRequest,
			wantBody:   "Query parameter 'n' is required\n",
		},
		{
			name:       "Zero Query Parameter",
			queryParam: "0",
			wantCode:   http.StatusBadRequest,
			wantBody:   "Query parameter 'n' has to be a positive integer\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/fib?n="+tt.queryParam, nil)
			w := httptest.NewRecorder()

			handleFib(w, req)

			res := w.Result()
			body := w.Body.String()

			if res.StatusCode != tt.wantCode {
				t.Errorf("Got status: %d, want status: %d", res.StatusCode, tt.wantCode)
			}

			if body != tt.wantBody {
				t.Errorf("Got body %q, want body %q", body, tt.wantBody)
			}
		})
	}
}
