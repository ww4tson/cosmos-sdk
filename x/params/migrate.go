package params

import (
	"fmt"

	storetypes "cosmossdk.io/core/store"
	"github.com/cosmos/gogoproto/proto"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params/types"
)

type SubspaceI interface {
	GetParamSet(ctx sdk.Context, ps types.ParamSet)
}

// MigrateModuleFromParams is an helper function to migrate a module from the x/params module.
// The params module has been deprecated since v0.46.0 and the migration of all core SDK modules
// away from x/params has happened in v0.47.0. This function is meant to be used by chains that have
// not yet migrated SDK modules away from x/params and skipped v0.47.x.
func MigrateModuleFromParams(
	ctx sdk.Context,
	storeService storetypes.KVStoreService,
	legacySubspace SubspaceI,
	cdc codec.BinaryCodec,
	migrateParams any,
	targetStoreParamKey []byte,
) error {
	store := storeService.OpenKVStore(ctx)

	currParams, ok := migrateParams.(types.ParamSet)
	if !ok {
		return fmt.Errorf("the params to be migrated must be a ParamSet, got %T", migrateParams)
	}

	legacySubspace.GetParamSet(ctx, currParams)

	// hasValidate is an interface that all parameter types should implement.
	// It is used to validate the parameters.
	type hasValidate interface {
		Validate() error
	}

	// Validate the params if the current params type implements the hasValidate interface.
	if params, ok := currParams.(hasValidate); ok {
		if err := params.Validate(); err != nil {
			return err
		}
	}

	newParams, ok := currParams.(proto.Message)
	if !ok {
		return fmt.Errorf("the params to be migrated must be a proto.Message, got %T", currParams)
	}

	bz := cdc.MustMarshal(newParams)
	store.Set(targetStoreParamKey, bz)

	return nil
}
