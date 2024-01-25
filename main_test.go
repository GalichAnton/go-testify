package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainHandler(t *testing.T) {
	t.Parallel()

	type args struct {
		url    string
		method string
	}

	tests := []struct {
		name           string
		args           args
		expected       string
		expectedStatus int
		count          int
	}{
		{
			name: "when request ok",
			args: args{
				url:    "/cafe?count=2&city=moscow",
				method: "GET",
			},
			expected:       "Мир кофе,Сладкоежка",
			expectedStatus: http.StatusOK,
		},
		{
			name: "when city not allowed",
			args: args{
				url:    "/cafe?count=2&city=NewYork",
				method: "GET",
			},
			expected:       "wrong city value",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "when count more than total",
			args: args{
				url:    "/cafe?count=10&city=moscow",
				method: "GET",
			},
			expected:       "Мир кофе,Сладкоежка,Кофе и завтраки,Сытый студент",
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.args.method, tt.args.url, nil)
			responseRecorder := httptest.NewRecorder()
			handler := http.HandlerFunc(mainHandle)
			handler.ServeHTTP(responseRecorder, req)
			status := responseRecorder.Code
			response := responseRecorder.Body.String()

			assert.Equal(t, tt.expectedStatus, status)
			assert.Equal(t, tt.expected, response)
		})
	}
}
