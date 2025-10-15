package usecase

import "github.com/isurucuma/hello/internal/domain"

type Hello interface {
	GetHello(data domain.HelloReqeustData) (domain.HelloResponseData, error)
}
