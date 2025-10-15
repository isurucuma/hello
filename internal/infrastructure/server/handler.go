package server

import (
	"encoding/json"
	"errors"
	"github.com/isurucuma/hello/internal/domain"
	"github.com/isurucuma/hello/internal/usecase"
	"log/slog"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type HelloResponse struct {
	Message string `json:"message"`
}

type Handler struct {
	hello  usecase.Hello
	logger *slog.Logger
}

func NewHandler(hello usecase.Hello, logger *slog.Logger) Handler {
	return Handler{
		hello:  hello,
		logger: logger,
	}
}

func (h Handler) getHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		sendInvalidInputResponse(w)
		return
	}

	reqData := domain.HelloReqeustData{
		Name: name,
	}

	respData, err := h.hello.GetHello(reqData)
	if err != nil && errors.Is(err, domain.InvalidInputError) {
		sendInvalidInputResponse(w)
		return
	}

	responseData := HelloResponse{
		Message: respData.Message,
	}

	writeResponse(w, responseData, http.StatusOK)
}

func writeResponse(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func sendInvalidInputResponse(w http.ResponseWriter) {
	errResponse := ErrorResponse{
		Error: "Invalid Input",
	}
	writeResponse(w, errResponse, http.StatusBadRequest)
}
