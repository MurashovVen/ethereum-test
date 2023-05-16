package grpc

import (
	"context"

	"github.com/ethereum/go-ethereum/common"

	"ethereum-test/internal/domain/service"
	"ethereum-test/pkg/grpc"
)

type ethereumService struct {
	svc *service.Service
	*grpc.UnimplementedEthereumServer
}

func (s *ethereumService) BalanceGet(ctx context.Context, req *grpc.BalanceReq) (*grpc.Balance, error) {
	address := common.HexToAddress(req.Address)
	balance, err := s.svc.Balance(ctx, address)
	if err != nil {
		return nil, err
	}

	return &grpc.Balance{
		Value: balance.String(),
	}, nil
}

func (s *ethereumService) TransactionSend(ctx context.Context, req *grpc.TransactionSendReq) (*grpc.TransactionSendResp, error) {
	return nil, s.svc.SendTransaction(ctx, req.TxBytes)
}
