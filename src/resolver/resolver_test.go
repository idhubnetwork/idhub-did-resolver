package resolver

import (
	"testing"
)
import "fmt"

var contractAddress string = "0x23fb5aE3228800E4FB4bDC8e717abD50fb31c4eA"
var identity string = "0x49dBa8f906c745B0a82f4D21E02BAFD7Df1a0be4"

func TestNewResolver(t *testing.T) {
	_, err := NewResolver("infuraRopsten", contractAddress)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	t.Log("func NewResolver success")
}

/*
func TestGetAuthenticationChanged(t *testing.T) {
	r, _ := NewResolver("infuraRopsten", contractAddress)
	block, err := r.GetAuthenticationChanged(identity)
	fmt.Println(block)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

func TestIdentityOwner(t *testing.T) {
	r, _ := NewResolver("infuraRopsten", contractAddress)
	addr, err := r.IdentityOwner(identity)
	fmt.Println(addr)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}
*/

func TestValidPublicKey(t *testing.T) {
	r, _ := NewResolver("infuraRopsten", contractAddress)
	ok, err := r.ValidPublicKey(identity, "veriKey", "0xa8d31fac615e26049614a3423eeb554bcada91024cf21059e806888c2cf4a756")
	fmt.Println(ok)
	if err != nil {
		t.Error(err)
		t.Fail()
	}
}

/*
func TestEventPublicKeyChanged(t *testing.T) {
	r, _ := NewResolver("infuraRopsten", contractAddress)
	logs, err := r.EventPublicKeyChanged(big.NewInt(4198003))
	fmt.Println(logs)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Log("success")
	}
}

func TestEventAuthenticationChanged(t *testing.T) {
	r, _ := NewResolver("infuraRopsten", contractAddress)
	logs, err := r.EventAuthenticationChanged(big.NewInt(4198033))
	fmt.Println(logs)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Log("success")
	}
}

func TestEventAttributeChanged(t *testing.T) {
	r, _ := NewResolver("infuraRopsten", contractAddress)
	logs, err := r.EventAttributeChanged(big.NewInt(4198057))
	fmt.Println(logs)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Log("success")
	}
}

func TestGetDIDLogs(t *testing.T) {
	r, _ := NewResolver("infuraRopsten", contractAddress)
	logs, err := r.GetDIDLogs(identity)
	fmt.Println(logs)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Log("success")
	}
}
*/

/*
func TestGetDIDDocument(t *testing.T) {
	r, _ := NewResolver("infuraRopsten", contractAddress)
	doc, err := r.getDIDDocument(identity)
	fmt.Println(doc)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Log("success")
	}
}
*/
func TestGetDocument(t *testing.T) {
	r, _ := NewResolver("infuraRopsten", contractAddress)
	doc, err := r.GetDocument(identity)
	fmt.Println(doc)
	if err != nil {
		t.Error(err)
		t.Fail()
	} else {
		t.Log("success")
	}
}
