package resolver

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func (r *resolver) GetPublicKeyChanged(address string) *big.Int {
	instance := r.contract
	identity := common.HexToAddress(address)
	publicKeyChanged, err := instance.PublicKeyChanged(nil, identity)
	if err != nil {

	}
	return publicKeyChanged
}

func (r *resolver) GetAuthenticationChanged(address string) *big.Int {
	instance := r.contract
	identity := common.HexToAddress(address)
	authenticationChanged, err := instance.AuthenticationChanged(nil, identity)
	if err != nil {

	}
	return authenticationChanged
}

func (r *resolver) GetAttributeChanged(address string) *big.Int {
	instance := r.contract
	identity := common.HexToAddress(address)
	attributeChanged, err := instance.AttributeChanged(nil, identity)
	if err != nil {

	}
	return attributeChanged
}

func (r *resolver) IdentityOwner(address string) string {
	instance := r.contract
	identity := common.HexToAddress(address)
	owner, err := instance.IdentityOwner(nil, identity)
	if err != nil {
	}
	return owner.Hex()
}

func (r *resolver) ValidPublicKey(address, keyType, key string) bool {
	instance := r.contract
	identity := common.HexToAddress(address)
	publickKeyType := [32]byte{}
	copy(publickKeyType[:], keyType)
	publickKey := [32]byte{}
	copy(publickKey[:], key)
	ok, err := instance.ValidPublicKey(nil, identity, publickKeyType,
		publickKey)
	if err != nil {
	}
	return ok
}

func (r *resolver) ValidAuthentication(address, keyType, key string) bool {
	instance := r.contract
	identity := common.HexToAddress(address)
	authenticationType := [32]byte{}
	copy(authenticationType[:], keyType)
	authentication := [32]byte{}
	copy(authentication[:], key)
	ok, err := instance.ValidAuthentication(nil, identity, authenticationType,
		authentication)
	if err != nil {
	}
	return ok
}
