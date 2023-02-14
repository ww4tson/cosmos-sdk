package types

import (
	"time"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type (
	// StakingKeeper defines the staking module interface contract needed by the
	// evidence module.
	StakingKeeper interface {
		ValidatorByConsAddr(sdk.Context, sdk.ConsAddress) stakingtypes.ValidatorI
		GetParams(ctx sdk.Context) (params stakingtypes.Params)
		GetDelegation(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) (delegation types.Delegation, found bool)
		GetValidator(ctx sdk.Context, addr sdk.ValAddress) (types.Validator, bool)
		Validator(sdk.Context, sdk.ValAddress) stakingtypes.ValidatorI // get a particular validator by operator address
	}

	// SlashingKeeper defines the slashing module interface contract needed by the
	// evidence module.
	SlashingKeeper interface {
		GetPubkey(sdk.Context, cryptotypes.Address) (cryptotypes.PubKey, error)
		IsTombstoned(sdk.Context, sdk.ConsAddress) bool
		HasValidatorSigningInfo(sdk.Context, sdk.ConsAddress) bool
		Tombstone(sdk.Context, sdk.ConsAddress)
		Slash(sdk.Context, sdk.ConsAddress, sdk.Dec, int64, int64)
		SlashWithInfractionReason(sdk.Context, sdk.ConsAddress, sdk.Dec, int64, int64, stakingtypes.Infraction)
		SlashFractionDoubleSign(sdk.Context) sdk.Dec
		Jail(sdk.Context, sdk.ConsAddress)
		JailUntil(sdk.Context, sdk.ConsAddress, time.Time)
		Unjail(ctx sdk.Context, validatorAddr sdk.ValAddress) error
		HandleValidatorSignature(ctx sdk.Context, addr cryptotypes.Address, power int64, signed bool)
	}

	// AccountKeeper define the account keeper interface contracted needed by the evidence module
	AccountKeeper interface {
		SetAccount(ctx sdk.Context, acc sdk.AccountI)
	}

	// BankKeeper define the account keeper interface contracted needed by the evidence module
	BankKeeper interface {
		MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
		SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
		GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	}
)
