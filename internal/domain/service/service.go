package service

import (
	"ethereum-test/internal/domain"
)

type Service struct {
	ethereum domain.EthereumClient
}

func New(ethClient domain.EthereumClient) *Service {
	return &Service{
		ethereum: ethClient,
	}
}
