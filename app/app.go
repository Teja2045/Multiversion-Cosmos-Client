package app

import (
	"log"
	"os"
	"time"
)

var (
	RPC_URL  string
	gRPC_URL string
)

type App struct {
	rpcClient  *RpcClient
	grpcClient QueryClient
	codec      EncodingConfig
}

func (app *App) GetGrpcClient() QueryClient {
	return app.grpcClient
}

func (app *App) GetRpcClient() *RpcClient {
	return app.rpcClient
}

func (app *App) init() {
	app.codec = MakeEncodingConfig()
	RPC_URL = os.Getenv("RPC")
	gRPC_URL = os.Getenv("GRPC")
	app.tryConnectChain()
}

func (app *App) tryConnectChain() {
	err := app.connectChain()
	if err != nil {
		time.Sleep(5 * time.Second)
		app.tryConnectChain()
	}
}

func (app *App) connectChain() error {
	var err error
	app.rpcClient, err = InitializeRPCClient(RPC_URL)
	if err != nil {
		log.Println("Error while staring rpc client", err)
		return err
	} else {
		log.Println("RPC client initialized...")
	}

	log.Println("starting rpc client.........")
	if err = app.rpcClient.Start(); err != nil {
		log.Println("Error while staring rpc client", err)
		return err
	} else {
		log.Println("started rpc client.........")
	}

	app.grpcClient, err = NewClient(gRPC_URL, app.codec)
	if err != nil {
		log.Println("Error while staring rpc client", err)
		return err
	} else {
		log.Println("successfully connected GRPC client")
	}
	return nil
}

func StartApp() *App {
	app := new(App)

	/** initialize the app with tendermint and grpc client */
	app.init()
	return app
}
