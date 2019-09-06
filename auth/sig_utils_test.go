package auth

import (
	"crypto/ecdsa"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

// Should generate a good public key from private
func TestAccount(t *testing.T) {
	privateKey, _ := crypto.HexToECDSA("78066589963545A353FD4768E12EAEE034292FA746DEDF23611D1AB8B24F2A79") // Test Priv
	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
}

func TestSig(t *testing.T) {
	privateKey, _ := crypto.HexToECDSA("78066589963545A353FD4768E12EAEE034292FA746DEDF23611D1AB8B24F2A79") // Test Priv
	slc := []byte("Sign into as 0x2db59226316905ae6e00fa795a5990d44af30f13 to Frontier backend services. By signing that you accept our terms and conditions and you understand that Frontier and all the related entities have no responsibilities on your private key, your eth, or your tokens.")
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(slc), slc)
	hash := crypto.Keccak256Hash([]byte(msg))
	signature, _ := crypto.Sign(hash.Bytes(), privateKey)
	if(signature != )
}

func TestVerifySig(t *testing.T) {
	account := "0xAEd65e60B25cF40ff7F8A564079aAF3fbF437d9d"
	msg := []byte("Sign into as 0x2db59226316905ae6e00fa795a5990d44af30f13 to Frontier backend services. By signing that you accept our terms and conditions and you understand that Frontier and all the related entities have no responsibilities on your private key, your eth, or your tokens.")
	sig := "0x1ede05802d2cc74b2c65f3a0e61c408af2c4b0801f5383e0c68f30472fac896e1222663a623666eaee32b32ecefa73bb3b6b1583135cb2cc104ef3954c59d4461c"

