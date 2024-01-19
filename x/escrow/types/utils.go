package types

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/reapchain-core/crypto/tmhash"
	"regexp"
)

func GetID(escrowPool sdk.Coin) []byte {
	id := escrowPool.Denom
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
