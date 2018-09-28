package resolver

import (
	"context"
	"log"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type LogPublicKeyChanged struct {
	Identity       common.Address
	PublicKeyType  [32]byte
	PublicKey      [32]byte
	ValidTO        *big.Int
	PreviousChange *big.Int
}

func (r *resolver) EventPublicKeyChanged(blockNumber *big.Int) *LogPublicKeyChanged {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockNumber),
		ToBlock:   big.NewInt(blockNumber),
		Addresses: []common.Address{
			r.address,
		},
	}
	logs, err := r.client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
}
