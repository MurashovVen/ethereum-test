package service

import (
	"context"
	"math/big"
	"math/rand"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"ethereum-test/internal/domain/mocks"
)

func TestService_SendTransaction(t *testing.T) {
	t.Parallel()

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		var (
			ctrl          = gomock.NewController(t)
			ethClientMock = mocks.NewMockEthereumClient(ctrl)

			tx = testTx()
		)

		txBytesExp, err := tx.MarshalJSON()
		require.NoError(t, err)

		ethClientMock.EXPECT().
			SendTransaction(gomock.Any(), gomock.Any()).
			Do(
				func(ctx context.Context, tx *types.Transaction) {
					t.Helper()

					txBytesGot, err := tx.MarshalJSON()
					require.NoError(t, err)

					require.Equal(t, txBytesExp, txBytesGot)
				},
			).Return(nil)

		// testing

		svc := New(ethClientMock)
		err = svc.SendTransaction(context.Background(), txBytesExp)
		require.NoError(t, err)
	})

	t.Run("Unmarshalling_ERR", func(t *testing.T) {
		t.Parallel()

		// testing

		svc := New(nil)
		err := svc.SendTransaction(context.Background(), []byte(`invalid`))
		require.Error(t, err)
	})

	t.Run("SendTx_ERR", func(t *testing.T) {
		t.Parallel()

		var (
			ctrl          = gomock.NewController(t)
			ethClientMock = mocks.NewMockEthereumClient(ctrl)

			tx = testTx()
		)

		txBytesExp, err := tx.MarshalJSON()
		require.NoError(t, err)

		ethClientMock.EXPECT().
			SendTransaction(gomock.Any(), gomock.Any()).
			Return(errSimple)

		// testing

		svc := New(ethClientMock)
		err = svc.SendTransaction(context.Background(), txBytesExp)
		require.ErrorIs(t, err, errSimple)
	})
}

func testTx() *types.Transaction {
	address := common.HexToAddress("0x0000006916a87b82333f4245046623b23794C65C")
	return types.NewTx(
		&types.DynamicFeeTx{
			ChainID:   big.NewInt(rand.Int63()),
			Nonce:     rand.Uint64(),
			GasTipCap: big.NewInt(rand.Int63()),
			GasFeeCap: big.NewInt(rand.Int63()),
			Gas:       rand.Uint64(),
			To:        &address,
			Value:     big.NewInt(rand.Int63()),
		},
	)
}
