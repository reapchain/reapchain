package types

import (
	"testing"

	codectypes "github.com/reapchain/cosmos-sdk/codec/types"
	sdk "github.com/reapchain/cosmos-sdk/types"

	"github.com/stretchr/testify/suite"
)

type CodecTestSuite struct {
	suite.Suite
}

func TestCodecSuite(t *testing.T) {
	suite.Run(t, new(CodecTestSuite))
}

func (suite *CodecTestSuite) TestRegisterInterfaces() {
	registry := codectypes.NewInterfaceRegistry()
	registry.RegisterInterface(sdk.MsgInterfaceProtoName, (*sdk.Msg)(nil))
	RegisterInterfaces(registry)

	impls := registry.ListImplementations(sdk.MsgInterfaceProtoName)
	suite.Require().Equal(3, len(impls))
	suite.Require().ElementsMatch([]string{
		"/reapchain.feesplit.v1.MsgRegisterFeeSplit",
		"/reapchain.feesplit.v1.MsgCancelFeeSplit",
		"/reapchain.feesplit.v1.MsgUpdateFeeSplit",
	}, impls)
}
