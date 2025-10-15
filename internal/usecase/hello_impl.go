package usecase

import (
	"github.com/isurucuma/hello/internal/domain"
	"log/slog"
	"strings"
)

type HelloImpl struct {
	logger *slog.Logger
}

func (h HelloImpl) GetHello(data domain.HelloReqeustData) (domain.HelloResponseData, error) {
	nameSpaceRemoved := strings.TrimSpace(data.Name)
	nameLower := strings.ToLower(nameSpaceRemoved)
	if nameLower == "" || !(nameLower[0] >= 'a' && nameLower[0] <= 'z') {
		h.logger.Info("Invalid name provided", slog.String("name", data.Name))
		return domain.HelloResponseData{}, domain.InvalidInputError
	}
	if nameLower[0] >= 'a' && nameLower[0] <= 'm' {
		return domain.HelloResponseData{
			Message: "Hello " + nameSpaceRemoved,
		}, nil
	}
	return domain.HelloResponseData{}, domain.InvalidInputError
}

func NewHelloImpl(logger *slog.Logger) Hello {
	return HelloImpl{
		logger: logger,
	}
}
