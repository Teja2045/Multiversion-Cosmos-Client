package app

import (
	"log"

	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
)

type RpcClient struct {
	client *rpchttp.HTTP
}

func InitializeRPCClient(uri string) (*RpcClient, error) {
	log.Println("rpc client url", uri)
	cl, err := rpchttp.New(uri, "/websocket")
	if err != nil {
		log.Println("Error while connecting web socket", err)
		return nil, err
	}

	return &RpcClient{client: cl}, nil
}

func (c *RpcClient) Start() error {
	return c.client.Start()
}
