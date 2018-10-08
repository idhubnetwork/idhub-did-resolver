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

type LogAuthenticationChanged struct {
	Identity           common.Address
	AuthenticationType [32]byte
	Authentication     [32]byte
	ValidTO            *big.Int
	PreviousChange     *big.Int
}

type LogAttributeChanged struct {
	Identity       common.Address
	Name           [32]byte
	Value          []byte
	ValidTO        *big.Int
	PreviousChange *big.Int
}

func (r *resolver) EventPublicKeyChanged(blockNumber *big.Int) []LogPublicKeyChanged {
	query := ethereum.FilterQuery{
		FromBlock: blockNumber,
		ToBlock:   blockNumber,
		Addresses: []common.Address{
			r.address,
		},
	}
	logs, err := r.client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	var logPublicKeyChangeds = make([]LogPublicKeyChanged, 0)
	for _, vLog := range logs {
		var logPublicKeyChanged LogPublicKeyChanged
		err := r.abi.Unpack(&logPublicKeyChanged, "DIDPublicKeyChanged", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		logPublicKeyChanged.Identity = common.HexToAddress(vLog.Topics[1].Hex())
		logPublicKeyChangeds = append(logPublicKeyChangeds, logPublicKeyChanged)
	}
	return logPublicKeyChangeds
}

func (r *resolver) EventAuthenticationChanged(blockNumber *big.Int) []LogAuthenticationChanged {
	query := ethereum.FilterQuery{
		FromBlock: blockNumber,
		ToBlock:   blockNumber,
		Addresses: []common.Address{
			r.address,
		},
	}
	logs, err := r.client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	var logAuthenticationChangeds = make([]LogAuthenticationChanged, 0)
	for _, vLog := range logs {
		var logAuthenticationChanged LogAuthenticationChanged
		err := r.abi.Unpack(&logAuthenticationChanged, "DIDAuthenticationChanged", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		logAuthenticationChanged.Identity = common.HexToAddress(vLog.Topics[1].Hex())
		logAuthenticationChangeds = append(logAuthenticationChangeds, logAuthenticationChanged)
	}
	return logAuthenticationChangeds
}

func (r *resolver) EventAttributeChanged(blockNumber *big.Int) []LogAttributeChanged {
	query := ethereum.FilterQuery{
		FromBlock: blockNumber,
		ToBlock:   blockNumber,
		Addresses: []common.Address{
			r.address,
		},
	}
	logs, err := r.client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	var logAttributeChangeds = make([]LogAttributeChanged, 0)
	for _, vLog := range logs {
		var logAttributeChanged LogAttributeChanged
		err := r.abi.Unpack(&logAttributeChanged, "DIDAttributeChanged", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		logAttributeChanged.Identity = common.HexToAddress(vLog.Topics[1].Hex())
		logAttributeChangeds = append(logAttributeChangeds, logAttributeChanged)
	}
	return logAttributeChangeds
}
