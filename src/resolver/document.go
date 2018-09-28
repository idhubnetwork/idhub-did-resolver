package resolver

import "github.com/ethereum/go-ethereum/common"

type DIDLog struct {
	DidPublicKeyLogs      []LogPublicKeyChanged
	DidAuthenticationLogs []LogAuthenticationChanged
	DidAttributeLogs      []LogAttributeChanged
}

func (r *resolver) GetDIDLogs(address string) *DIDLog {
	identity := common.HexToAddress(address)
	publicKeyChanged := r.GetPublicKeyChanged(address)
	authenticationChanged := r.GetAuthenticationChanged(address)
	attributeChanged := r.GetAttributeChanged(address)
	var DidPublicKeyLogs = make([]LogPublicKeyChanged, 0)
	for publicKeyChanged.Sign() != 0 {
		logPublicKeyChangeds := r.EventPublicKeyChanged(publicKeyChanged)
		for _, logV := range logPublicKeyChangeds {
			if identity == logV.Identity {
				DidPublicKeyLogs = append(DidPublicKeyLogs, logV)
				publicKeyChanged = logV.PreviousChange
			} else {
				continue
			}
		}
	}
	var DidAuthenticationLogs = make([]LogAuthenticationChanged, 0)
	for authenticationChanged.Sign() != 0 {
		logAuthenticationChangeds := r.EventAuthenticationChanged(authenticationChanged)
		for _, logV := range logAuthenticationChangeds {
			if identity == logV.Identity {
				DidAuthenticationLogs = append(DidAuthenticationLogs, logV)
				authenticationChanged = logV.PreviousChange
			} else {
				continue
			}
		}
	}
	var DidAttributeLogs = make([]LogAttributeChanged, 0)
	for attributeChanged.Sign() != 0 {
		logAttributeChangeds := r.EventAttributeChanged(attributeChanged)
		for _, logV := range logAttributeChangeds {
			if identity == logV.Identity {
				DidAttributeLogs = append(DidAttributeLogs, logV)
				attributeChanged = logV.PreviousChange
			} else {
				continue
			}
		}
	}
	return &DIDLog{DidPublicKeyLogs, DidAuthenticationLogs, DidAttributeLogs}
}
