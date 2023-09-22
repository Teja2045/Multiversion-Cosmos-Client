//go:build server
// +build server

package main

import (
	"log"
	"multiVersionCodec/app"
)

func main() {
	app.LoadEnv()
	appInstance := app.StartApp()
	log.Println(appInstance)
	resp, err := appInstance.GetGrpcClient().GetAccountBalances("pasg1emkh9v2kk03j4ccs0pnzk78e7ejq6wlz8mnn9u")
	if err != nil {
		log.Println("some error", err)
	} else {
		log.Println("res: ", resp)
	}

}
