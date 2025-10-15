package usecase_test

import (
	"errors"
	"github.com/isurucuma/hello/internal/domain"
	"github.com/isurucuma/hello/internal/usecase"
	"log/slog"
	"os"
	"testing"
)

func TestHelloImpl_GetHello(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelError,
	}))

	helloUsecase := usecase.NewHelloImpl(logger)

	tests := []struct {
		name          string
		input         domain.HelloReqeustData
		expectedResp  domain.HelloResponseData
		expectedError error
	}{
		{
			name: "Valid name starting with a-m",
			input: domain.HelloReqeustData{
				Name: "alice",
			},
			expectedResp: domain.HelloResponseData{
				Message: "Hello alice",
			},
			expectedError: nil,
		},
		{
			name: "Valid name with uppercase starting with A-M",
			input: domain.HelloReqeustData{
				Name: "Alice",
			},
			expectedResp: domain.HelloResponseData{
				Message: "Hello Alice",
			},
			expectedError: nil,
		},
		{
			name: "Valid name with spaces starting with a-m",
			input: domain.HelloReqeustData{
				Name: "  john  ",
			},
			expectedResp: domain.HelloResponseData{
				Message: "Hello john",
			},
			expectedError: nil,
		},
		{
			name: "Name starting with n-z",
			input: domain.HelloReqeustData{
				Name: "zack",
			},
			expectedResp:  domain.HelloResponseData{},
			expectedError: domain.InvalidInputError,
		},
		{
			name: "Empty name",
			input: domain.HelloReqeustData{
				Name: "",
			},
			expectedResp:  domain.HelloResponseData{},
			expectedError: domain.InvalidInputError,
		},
		{
			name: "Name with spaces only",
			input: domain.HelloReqeustData{
				Name: "   ",
			},
			expectedResp:  domain.HelloResponseData{},
			expectedError: domain.InvalidInputError,
		},
		{
			name: "Name starting with number",
			input: domain.HelloReqeustData{
				Name: "123John",
			},
			expectedResp:  domain.HelloResponseData{},
			expectedError: domain.InvalidInputError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp, err := helloUsecase.GetHello(tc.input)

			if !errors.Is(err, tc.expectedError) {
				t.Errorf("expected error %v, got %v", tc.expectedError, err)
			}

			if err == nil && resp.Message != tc.expectedResp.Message {
				t.Errorf("expected message %q, got %q", tc.expectedResp.Message, resp.Message)
			}
		})
	}
}
