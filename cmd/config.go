package main

type config struct {
	EthereumAddress   string `split_words:"true" default:"https://rpc.sepolia.org/" desc:"Адрес сети Ethereum"`
	GRPCServerAddress string `split_words:"true" default:"0.0.0.0:90" desc:"Адрес, на котором будет запущен GRPC сервер (host:port)"`
}
