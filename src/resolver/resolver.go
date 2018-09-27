package resolver

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	did "./contracts"
)

type resolver struct {
	client   *ethclient.Client
	contract *did.Did
}

var urls map[string]string
var err error

func init() {
	urls = make(map[string]string, 5)
	urls["infuraMainnet"] = "https://mainnet.infura.io"
	urls["infuraRopsten"] = "https://ropsten.infura.io"
	urls["infuraRinkeby"] = "https://rinkeby.infura.io"
}

func NewResolver(net string, address string) *resolver {
	r := new(resolver)
	r.client, err = ethclient.Dial(urls[net])
	if err != nil {
	}
	contractAddr := common.HexToAddress(address)
	r.contract, err = did.NewDid(contractAddr, r.client)
	if err != nil {
	}
	return r
}
