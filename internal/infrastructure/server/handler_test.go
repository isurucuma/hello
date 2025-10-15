package server

import (
	"encoding/json"
	"github.com/isurucuma/hello/internal/domain"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type MockHello struct {
	GetHelloFn func(data domain.HelloReqeustData) (domain.HelloResponseData, error)
}

func (m MockHello) GetHello(data domain.HelloReqeustData) (domain.HelloResponseData, error) {
	return m.GetHelloFn(data)
}

func TestHandler_GetHello(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelError,
	}))

	tests := []struct {
		name               string
		queryParam         string
		mockResponse       domain.HelloResponseData
		mockError          error
		expectedStatusCode int
		expectedResponse   map[string]interface{}
	}{
		{
			name:               "Success response",
			queryParam:         "name=alice",
			mockResponse:       domain.HelloResponseData{Message: "Hello alice"},
			mockError:          nil,
			expectedStatusCode: http.StatusOK,
			expectedResponse: map[string]interface{}{
				"message": "Hello alice",
			},
		},
		{
			name:               "Invalid input from usecase",
			queryParam:         "name=zack",
			mockResponse:       domain.HelloResponseData{},
			mockError:          domain.InvalidInputError,
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: map[string]interface{}{
				"error": "Invalid Input",
			},
		},
		{
			name:               "No name parameter",
			queryParam:         "",
			mockResponse:       domain.HelloResponseData{},
			mockError:          nil,
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse: map[string]interface{}{
				"error": "Invalid Input",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockHello := MockHello{
				GetHelloFn: func(data domain.HelloReqeustData) (domain.HelloResponseData, error) {
					return tc.mockResponse, tc.mockError
				},
			}

			handler := NewHandler(mockHello, logger)

			url := "/hello-world"
			if tc.queryParam != "" {
				url = "/hello-world?" + tc.queryParam
			}
			req := httptest.NewRequest(http.MethodGet, url, nil)
			rec := httptest.NewRecorder()

			handler.getHello(rec, req)

			if rec.Code != tc.expectedStatusCode {
				t.Errorf("expected status code %d, got %d", tc.expectedStatusCode, rec.Code)
			}

			var response map[string]interface{}
			if err := json.NewDecoder(rec.Body).Decode(&response); err != nil {
				t.Fatalf("error decoding response body: %v", err)
			}

			for k, v := range tc.expectedResponse {
				if response[k] != v {
					t.Errorf("expected %s to be %v, got %v", k, v, response[k])
				}
			}
		})
	}
}
