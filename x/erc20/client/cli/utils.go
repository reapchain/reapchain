package cli

import (
	"os"
	"path/filepath"

	"github.com/reapchain/cosmos-sdk/codec"
	banktypes "github.com/reapchain/cosmos-sdk/x/bank/types"
)

// ParseRegisterCoinProposal reads and parses a ParseRegisterCoinProposal from a file.
func ParseMetadata(cdc codec.JSONCodec, metadataFile string) (banktypes.Metadata, error) {
	metadata := banktypes.Metadata{}

	contents, err := os.ReadFile(filepath.Clean(metadataFile))
	if err != nil {
		return metadata, err
	}

	if err = cdc.UnmarshalJSON(contents, &metadata); err != nil {
		return metadata, err
	}

	return metadata, nil
}
