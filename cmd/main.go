package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/kelseyhightower/envconfig"
	"golang.org/x/sync/errgroup"

	"ethereum-test/internal/domain/service"
	"ethereum-test/internal/grpc"
)

func main() {
	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	cfg := config{}
	err := envconfig.Usage("APP", &cfg)
	panicIfErr("вывод конфигурации в консоль", err)

	err = envconfig.Process("APP", &cfg)
	panicIfErr("чтение конфигурации", err)

	eg, egCtx := errgroup.WithContext(ctx)

	eg.Go(terminationSignalWaiter(egCtx))
	eg.Go(grpcServerRunner(egCtx, &cfg))

	if err = eg.Wait(); err != nil && !errors.Is(err, errTerminationSig) && !errors.Is(err, context.Canceled) {
		panic("аварийное завершение: " + err.Error())
	}

	log.Info("сервис успешно завершил работу")
}

func grpcServerRunner(ctx context.Context, cfg *config) func() error {
	return func() error {
		ethClient, err := ethclient.Dial(cfg.EthereumAddress)
		if err != nil {
			return fmt.Errorf("подключение к сети ethereum: %w", err)
		}
		defer ethClient.Close()

		svc := service.New(ethClient)
		server := grpc.New(svc)

		listener, err := net.Listen("tcp", cfg.GRPCServerAddress)
		if err != nil {
			return fmt.Errorf("запуск слушателя сервера: %w", err)
		}

		serverStopped := make(chan error)
		go func() {
			if err := server.Serve(listener); err != nil {
				serverStopped <- err
			}

			close(serverStopped)
		}()

		select {
		case <-ctx.Done():
			err = ctx.Err()

		case err = <-serverStopped:
		}

		server.GracefulStop()

		return err
	}
}

var errTerminationSig = errors.New("termination signal caught")

func terminationSignalWaiter(ctx context.Context) func() error {
	return func() error {
		trap := make(chan os.Signal, 1)
		signal.Notify(trap, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-trap:
			return errTerminationSig

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func panicIfErr(msg string, err error) {
	if err != nil {
		panic(msg + ": " + err.Error())
	}
}
