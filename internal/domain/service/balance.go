package service

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// Balance returns wei balance of the provided account.
func (s *Service) Balance(ctx context.Context, account common.Address) (*big.Int, error) {
	return s.ethereum.BalanceAt(ctx, account, nil)
}
