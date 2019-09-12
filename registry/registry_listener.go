package listener

import (
	"context"
	"fmt"
	"log"
	"math/big"

	// for demo

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Watcher struct {
	hex      string
	client   EthClient
	wsclient EthClient
}

func NewWatcher() Watcher {
	client, err := ethclient.Dial(HomeRPCUrl)
	if err != nil {
		log.Fatal(err)
	}
	wsclient, err := ethclient.Dial(HomeSubscriptionURL)
	return Watcher{os.Getenv("REGISTRY_ADDRESS"), client, wsclient}
}

func (w Watcher) PrevLogs() {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(StartBlock),
		Addresses: []common.Address{
			common.HexToAddress(os.Getenv("REGISTRY_ADDRESS")),
		},
	}
	prvlogs, err := w.client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range prvlogs {
		resolve(vLog)
	}
}

func (w Watcher) Subscribe() {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(StartBlock),
		Addresses: []common.Address{
			common.HexToAddress(os.Getenv("REGISTRY_ADDRESS")),
		},
	}
	slogs := make(chan types.Log)
	_, err := w.wsclient.SubscribeFilterLogs(context.Background(), query, slogs)
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println("New Event received", <-slogs)
	}
}
