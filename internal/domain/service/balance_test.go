package service

import (
	"context"
	"errors"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"ethereum-test/internal/domain/mocks"
)

var errSimple = errors.New("simple error")

func TestService_Balance(t *testing.T) {
	t.Parallel()

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		var (
			ctrl          = gomock.NewController(t)
			ethClientMock = mocks.NewMockEthereumClient(ctrl)

			account = common.HexToAddress("0x0000006916a87b82333f4245046623b23794C65C")

			balanceExp = big.NewInt(100)
		)

		ethClientMock.EXPECT().BalanceAt(gomock.Any(), account, nil).Return(balanceExp, nil)

		// testing

		svc := New(ethClientMock)
		balanceGot, err := svc.Balance(context.Background(), account)
		require.NoError(t, err)
		require.Equal(t, balanceExp, balanceGot)
	})

	t.Run("ERR", func(t *testing.T) {
		t.Parallel()

		var (
			ctrl          = gomock.NewController(t)
			ethClientMock = mocks.NewMockEthereumClient(ctrl)

			account = common.HexToAddress("0x0000006916a87b82333f4245046623b23794C65C")
		)

		ethClientMock.EXPECT().BalanceAt(gomock.Any(), account, nil).Return(nil, errSimple)

		// testing

		svc := New(ethClientMock)
		_, err := svc.Balance(context.Background(), account)
		require.ErrorIs(t, err, errSimple)
	})
}
