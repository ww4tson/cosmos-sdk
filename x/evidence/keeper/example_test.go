package keeper_test

import (
	"testing"

	storetypes "cosmossdk.io/store/types"
	"cosmossdk.io/x/evidence"
	"cosmossdk.io/x/evidence/keeper"
	evidencetestutil "cosmossdk.io/x/evidence/testutil"
	"cosmossdk.io/x/evidence/types"
	"gotest.tools/v3/assert"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/golang/mock/gomock"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtestutil "github.com/cosmos/cosmos-sdk/x/staking/testutil"
)

var (
	// The default power validators are initialized to have within tests
	initAmt   = sdk.TokensFromConsensusPower(200, sdk.DefaultPowerReduction)
	initCoins = sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, initAmt))
)

type ReplaceModules struct {
	EvidenceKeeper keeper.Keeper
	// Add any other needed fixtures here
	StakingKeeper  *evidencetestutil.MockStakingKeeper
	SlashingKeeper *evidencetestutil.MockSlashingKeeper
}

type fixture struct {
	ctx sdk.Context
	app *runtime.App

	evidenceKeeper    keeper.Keeper
	bankKeeper        *evidencetestutil.MockBankKeeper
	accountKeeper     *evidencetestutil.MockAccountKeeper
	slashingKeeper    *evidencetestutil.MockSlashingKeeper
	stakingKeeper     *evidencetestutil.MockStakingKeeper
	stakeKeeper       *stakingkeeper.Keeper
	interfaceRegistry codectypes.InterfaceRegistry
}

// SetupApp returns a fixture required for testing
// Here, depModules are minimal dependency modules list required for testing current module and
// replaceModules contains current module + if any other modules which needs custom changes for testing.

// func SetupApp(depModules []string, replaceModules ReplaceModules, config depinject.Config) (*fixture, error) {
func SetupApp(depModules []string, replaceModules ReplaceModules) (*fixture, error) {
	f := &fixture{}

	app, err := simtestutil.Setup(evidencetestutil.AppConfig,
		&f.evidenceKeeper,
		&f.interfaceRegistry,
		&f.stakeKeeper,
	)
	if err != nil {
		return nil, err
	}
	f.stakingKeeper = replaceModules.StakingKeeper
	f.slashingKeeper = replaceModules.SlashingKeeper

	f.ctx = app.BaseApp.NewContext(false, cmtproto.Header{Height: 1})
	f.app = app

	// err := depinject.Provide(replaceModules)
	// for _, depModule := range depModules {
	// }

	return f, nil
}

func TestHandleDoubleSign(t *testing.T) {
	t.Parallel()

	encCfg := moduletestutil.MakeTestEncodingConfig(evidence.AppModuleBasic{})
	key := storetypes.NewKVStoreKey(types.StoreKey)

	ctrl := gomock.NewController(t)
	stakingKeeper := evidencetestutil.NewMockStakingKeeper(ctrl)
	slashingKeeper := evidencetestutil.NewMockSlashingKeeper(ctrl)
	accountKeeper := evidencetestutil.NewMockAccountKeeper(ctrl)
	bankKeeper := evidencetestutil.NewMockBankKeeper(ctrl)

	evidenceKeeper := keeper.NewKeeper(
		encCfg.Codec,
		key,
		stakingKeeper,
		slashingKeeper,
	)

	router := types.NewRouter()
	router = router.AddRoute(types.RouteEquivocation, testEquivocationHandler(evidenceKeeper))
	evidenceKeeper.SetRouter(router)

	depModules := []string{"auth", "bank", "slashing", "staking"}
	replaceModules := ReplaceModules{
		EvidenceKeeper: *evidenceKeeper,
		StakingKeeper:  stakingKeeper,
	}
	f, err := SetupApp(depModules, replaceModules)
	assert.NilError(t, err)

	f.bankKeeper = bankKeeper
	f.accountKeeper = accountKeeper

	ctx := f.ctx.WithIsCheckTx(false).WithBlockHeight(1)
	populateValidators(t, f)

	power := int64(100)
	f.stakingKeeper.EXPECT().GetParams(ctx).AnyTimes()
	stakingParams := f.stakingKeeper.GetParams(ctx)
	operatorAddr, val := valAddresses[0], pubkeys[0]
	tstaking := stakingtestutil.NewHelper(t, ctx, f.stakeKeeper)

	selfDelegation := tstaking.CreateValidatorWithValPower(operatorAddr, val, power, true)

	// execute end-blocker and verify validator attributes
	staking.EndBlocker(ctx, f.stakeKeeper)

	assert.DeepEqual(t,
		f.bankKeeper.GetAllBalances(ctx, sdk.AccAddress(operatorAddr)).String(),
		sdk.NewCoins(sdk.NewCoin(stakingParams.BondDenom, initAmt.Sub(selfDelegation))).String(),
	)
	assert.DeepEqual(t, selfDelegation, f.stakingKeeper.Validator(ctx, operatorAddr).GetBondedTokens())

	// // // handle a signature to set signing info
	// // f.slashingKeeper.HandleValidatorSignature(ctx, val.Address(), selfDelegation.Int64(), true)

	// // double sign less than max age
	// oldTokens := f.stakingKeeper.Validator(ctx, operatorAddr).GetTokens()
	// evidence := &types.Equivocation{
	// 	Height:           0,
	// 	Time:             time.Unix(0, 0),
	// 	Power:            power,
	// 	ConsensusAddress: sdk.ConsAddress(val.Address()).String(),
	// }
	// f.evidenceKeeper.HandleEquivocationEvidence(ctx, evidence)

	// // should be jailed and tombstoned
	// assert.Assert(t, f.stakingKeeper.Validator(ctx, operatorAddr).IsJailed())
	// assert.Assert(t, f.slashingKeeper.IsTombstoned(ctx, sdk.ConsAddress(val.Address())))

	// // tokens should be decreased
	// newTokens := f.stakingKeeper.Validator(ctx, operatorAddr).GetTokens()
	// assert.Assert(t, newTokens.LT(oldTokens))

	// // submit duplicate evidence
	// f.evidenceKeeper.HandleEquivocationEvidence(ctx, evidence)

	// // tokens should be the same (capped slash)
	// assert.Assert(t, f.stakingKeeper.Validator(ctx, operatorAddr).GetTokens().Equal(newTokens))

	// // jump to past the unbonding period
	// ctx = ctx.WithBlockTime(time.Unix(1, 0).Add(stakingParams.UnbondingTime))

	// // require we cannot unjail
	// assert.Error(t, f.slashingKeeper.Unjail(ctx, operatorAddr), slashingtypes.ErrValidatorJailed.Error())

	// // require we be able to unbond now
	// ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
	// // del, _ := f.stakingKeeper.GetDelegation(ctx, sdk.AccAddress(operatorAddr), operatorAddr)
	// // validator, _ := f.stakingKeeper.GetValidator(ctx, operatorAddr)
	// // totalBond := validator.TokensFromShares(del.GetShares()).TruncateInt()
	// // tstaking.Ctx = ctx
	// // tstaking.Denom = stakingParams.BondDenom
	// // tstaking.Undelegate(sdk.AccAddress(operatorAddr), operatorAddr, totalBond, true)

	// // query evidence from store
	// evidences := f.evidenceKeeper.GetAllEvidence(ctx)
	// assert.Assert(t, len(evidences) == 1)
}

func populateValidators(t assert.TestingT, f *fixture) {
	// add accounts and set total supply
	totalSupplyAmt := initAmt.MulRaw(int64(len(valAddresses)))
	totalSupply := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, totalSupplyAmt))
	f.bankKeeper.EXPECT().MintCoins(gomock.Any(), "mint", gomock.Any()).AnyTimes()

	assert.NilError(t, f.bankKeeper.MintCoins(f.ctx, "mint", totalSupply))

	for _, addr := range valAddresses {
		f.bankKeeper.EXPECT().SendCoinsFromModuleToAccount(f.ctx, "mint", (sdk.AccAddress)(addr), initCoins)
		assert.NilError(t, f.bankKeeper.SendCoinsFromModuleToAccount(f.ctx, "mint", (sdk.AccAddress)(addr), initCoins))
	}
}

// type Helper struct {
// 	t       *testing.T
// 	msgSrvr stakingtypes.MsgServer
// 	k       *keeper.Keeper

// 	Ctx        sdk.Context
// 	Commission stakingtypes.CommissionRates
// 	// Coin Denomination
// 	Denom string
// }

// // NewHelper creates a new instance of Helper.
// func NewHelper(t *testing.T, ctx sdk.Context, k *keeper.Keeper) *Helper {
// 	return &Helper{t, keeper.NewMsgServerImpl(k), k, ctx, ZeroCommission(), sdk.DefaultBondDenom}
// }
