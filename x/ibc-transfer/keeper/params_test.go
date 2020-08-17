package keeper_test

import "github.com/cosmos/cosmos-sdk/x/ibc-transfer/types"

func (suite *KeeperTestSuite) TestParams() {
	expParams := types.DefaultParams()

	params := suite.chainA.App.TransferKeeper.GetParams(suite.chainA.GetContext())
	suite.Require().Equal(expParams, params)

	expParams.TransfersEnabled = false
	suite.chainA.App.TransferKeeper.SetParams(suite.chainA.GetContext(), expParams)
	params = suite.chainA.App.TransferKeeper.GetParams(suite.chainA.GetContext())
	suite.Require().Equal(expParams, params)
}
