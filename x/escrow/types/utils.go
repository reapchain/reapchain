package types

import (
	"github.com/reapchain/reapchain-core/crypto/tmhash"
	"regexp"
)

func (escrowSupply EscrowSupply) GetID() []byte {
	id := escrowSupply.Denom
	return tmhash.Sum([]byte(id))
}

func (denom RegisteredDenom) GetID() []byte {
	id := denom.Denom
	return tmhash.Sum([]byte(id))
}

func removePrefix(denom string) string {
	re := regexp.MustCompile(DenomPrefix)
	return re.ReplaceAllString(denom, "")
}
