package v2

import (
	"fmt"

	"github.com/reapchain/reapchain-core/crypto/tmhash"

	"github.com/reapchain/reapchain/v8/x/gravity/types"
)

// Hash implements WithdrawBatch.Hash
func MsgBatchSendToEthClaimHash(msg types.MsgBatchSendToEthClaim) ([]byte, error) {
	path := fmt.Sprintf("%s/%d/%d/%s", msg.TokenContract, msg.BatchNonce, msg.EventNonce, msg.TokenContract)
	return tmhash.Sum([]byte(path)), nil
}
