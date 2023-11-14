package chad_v2

import "github.com/ethereum/go-ethereum/crypto"

type ecdsaSigner struct {
}

func NewEcdsa() *ecdsaSigner {
	return &ecdsaSigner{}
}

func (es *ecdsaSigner) Sign(data []byte, acc *fixtureAcc) ([]byte, error) {
	return crypto.Sign(data, acc.Private)
}
