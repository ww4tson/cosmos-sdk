package params_test

import (
	"testing"

	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	moduletestutil "github.com/cosmos/cosmos-sdk/types/module/testutil"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/params/types"
)

type mockSubspace struct {
	ps authtypes.Params
}

func (ms mockSubspace) GetParamSet(ctx sdk.Context, ps types.ParamSet) {
	*ps.(*authtypes.Params) = ms.ps
}

func TestMigrateModuleFromParams(t *testing.T) {
	moduleName := "auth"

	encCfg := moduletestutil.MakeTestEncodingConfig(auth.AppModuleBasic{})

	storeKey := storetypes.NewKVStoreKey(moduleName)
	tKey := storetypes.NewTransientStoreKey("transient_test")
	ctx := testutil.DefaultContext(storeKey, tKey)

	params.MigrateModuleFromParams(
		ctx,
		runtime.NewKVStoreService(storeKey),
		mockSubspace{},
		encCfg.Codec,
		authtypes.Params{},
		authtypes.ParamsKey,
	)
}
