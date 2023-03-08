package keeper_test

import (
	"testing"

	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/x/evidence/keeper"
	evidencetestutil "cosmossdk.io/x/evidence/testutil"
	"cosmossdk.io/x/evidence/types"
	"gotest.tools/v3/assert"

	"github.com/golang/mock/gomock"

	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	// The default power validators are initialized to have within tests
	initAmt   = sdk.TokensFromConsensusPower(200, sdk.DefaultPowerReduction)
	initCoins = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, initAmt))
)

type TestSuite struct {
	ctx sdk.Context

	evidenceKeeper keeper.Keeper
	bankKeeper     *evidencetestutil.MockBankKeeper
}

func (s *KeeperTestSuite) TestSetup() {
	key := storetypes.NewKVStoreKey(types.StoreKey)
	tkey := storetypes.NewTransientStoreKey("evidence_transient_store")
	testCtx := testutil.DefaultContextWithDB(&testing.T{}, key, tkey)
	s.ctx = testCtx.Ctx

	ctrl := gomock.NewController(&testing.T{})
	bankKeeper := evidencetestutil.NewMockBankKeeper(ctrl)
	s.bankKeeper = bankKeeper
}

func (s *KeeperTestSuite) TestHandleDoubleSign() {
	// t.Parallel()

	// ctx := s.ctx.WithIsCheckTx(false).WithBlockHeight(1)
	s.populateValidators(s.T())
}

func (s *KeeperTestSuite) populateValidators(t assert.TestingT) {
	// add accounts and set total supply
	totalSupplyAmt := initAmt.MulRaw(int64(len(valAddresses)))
	totalSupply := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, totalSupplyAmt))
	s.bankKeeper.EXPECT().MintCoins(gomock.Any(), "mint", totalSupply).Return(nil)
	assert.NilError(t, s.bankKeeper.MintCoins(s.ctx, "mint", totalSupply))

	for _, addr := range valAddresses {
		s.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(s.ctx, "mint", (sdk.AccAddress)(addr), initCoins).Return(nil)
		assert.NilError(t, s.bankKeeper.SendCoinsFromModuleToAccount(s.ctx, "mint", (sdk.AccAddress)(addr), initCoins))
	}
}
