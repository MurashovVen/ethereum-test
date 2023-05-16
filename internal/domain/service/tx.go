package service

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
)

// SendTransaction takes JSON encoded bytes of types.Transaction and sends tx to the Ethereum network.
func (s *Service) SendTransaction(ctx context.Context, txBytes []byte) error {
	tx := new(types.Transaction)
	if err := tx.UnmarshalJSON(txBytes); err != nil {
		return err
	}

	return s.ethereum.SendTransaction(ctx, tx)
}
