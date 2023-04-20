package testutil

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/cometbft/cometbft/proto/tendermint/crypto"
	"github.com/cometbft/cometbft/rpc/client/http"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

type E2ETestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func NewE2ETestSuite(cfg network.Config) *E2ETestSuite {
	return &E2ETestSuite{cfg: cfg}
}

func (s *E2ETestSuite) SetupSuite() {
	s.T().Log("setting up e2e test suite")

	if testing.Short() {
		s.T().Skip("skipping test in unit-tests mode.")
	}

	var err error
	s.network, err = network.New(s.T(), s.T().TempDir(), s.cfg)
	s.Require().NoError(err)

	unbond, err := sdk.ParseCoinNormalized("10stake")
	s.Require().NoError(err)

	val := s.network.Validators[0]
	val2 := s.network.Validators[1]

	// redelegate
	out, err := MsgRedelegateExec(
		val.ClientCtx,
		val.Address,
		val.ValAddress,
		val2.ValAddress,
		unbond,
		fmt.Sprintf("--%s=%d", flags.FlagGas, 300000),
	)
	s.Require().NoError(err)
	var txRes sdk.TxResponse
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txRes))
	s.Require().Equal(uint32(0), txRes.Code)
	s.Require().NoError(s.network.WaitForNextBlock())

	unbondingAmount := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(5))

	// unbonding the amount
	out, err = MsgUnbondExec(val.ClientCtx, val.Address, val.ValAddress, unbondingAmount)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txRes))
	s.Require().Equal(uint32(0), txRes.Code)
	s.Require().NoError(s.network.WaitForNextBlock())

	// unbonding the amount
	out, err = MsgUnbondExec(val.ClientCtx, val.Address, val.ValAddress, unbondingAmount)
	s.Require().NoError(err)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &txRes))
	s.Require().Equal(uint32(0), txRes.Code)
	s.Require().NoError(s.network.WaitForNextBlock())
}

func (s *E2ETestSuite) TearDownSuite() {
	s.T().Log("tearing down e2e test suite")
	s.network.Cleanup()
}

func (s *E2ETestSuite) TestGetCmdQueryHistoricalInfo() {
	val := s.network.Validators[0]

	testCases := []struct {
		name  string
		args  []string
		error bool
	}{
		{
			"wrong height",
			[]string{
				"-1",
				fmt.Sprintf("--%s=json", flags.FlagOutput),
			},
			true,
		},
		{
			"valid request",
			[]string{
				"1",
				fmt.Sprintf("--%s=json", flags.FlagOutput),
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryHistoricalInfo()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.error {
				s.Require().Error(err)
			} else {
				var historicalInfo types.HistoricalInfo

				err = val.ClientCtx.Codec.UnmarshalJSON(out.Bytes(), &historicalInfo)
				s.Require().NoError(err)
				s.Require().NotNil(historicalInfo)
			}
		})
	}
}

func (s *E2ETestSuite) TestGetCmdQueryParams() {
	val := s.network.Validators[0]
	testCases := []struct {
		name           string
		args           []string
		expectedOutput string
	}{
		{
			"with text output",
			[]string{fmt.Sprintf("--%s=text", flags.FlagOutput)},
			`bond_denom: stake
historical_entries: 10000
max_entries: 7
max_validators: 100
min_commission_rate: "0.000000000000000000"
unbonding_time: 1814400s`,
		},
		{
			"with json output",
			[]string{fmt.Sprintf("--%s=json", flags.FlagOutput)},
			`{"unbonding_time":"1814400s","max_validators":100,"max_entries":7,"historical_entries":10000,"bond_denom":"stake","min_commission_rate":"0.000000000000000000"}`,
		},
	}
	for _, tc := range testCases {
		tc := tc
		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryParams()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)
			s.Require().NoError(err)
			s.Require().Equal(tc.expectedOutput, strings.TrimSpace(out.String()))
		})
	}
}

// TestBlockResults tests that the validator updates correctly show when
// calling the /block_results RPC endpoint.
// ref: https://github.com/cosmos/cosmos-sdk/issues/7401.
func (s *E2ETestSuite) TestBlockResults() {
	require := s.Require()
	val := s.network.Validators[0]

	// Create new account in the keyring.
	k, _, err := val.ClientCtx.Keyring.NewMnemonic("NewDelegator", keyring.English, sdk.FullFundraiserPath, keyring.DefaultBIP39Passphrase, hd.Secp256k1)
	require.NoError(err)
	pub, err := k.GetPubKey()
	require.NoError(err)
	newAddr := sdk.AccAddress(pub.Address())

	// Send some funds to the new account.
	_, err = clitestutil.MsgSendExec(
		val.ClientCtx,
		val.Address,
		newAddr,
		sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(200))), fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	)
	require.NoError(err)
	require.NoError(s.network.WaitForNextBlock())

	// Use CLI to create a delegation from the new account to validator `val`.
	cmd := cli.NewDelegateCmd()
	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, []string{
		val.ValAddress.String(),
		sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(150)).String(),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, newAddr.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	})
	require.NoError(err)
	require.NoError(s.network.WaitForNextBlock())

	// Create a HTTP rpc client.
	rpcClient, err := http.New(val.RPCAddress, "/websocket")
	require.NoError(err)

	// Loop until we find a block result with the correct validator updates.
	// By experience, it happens around 2 blocks after `delHeight`.
	s.network.RetryForBlocks(func() error {
		latestHeight, err := s.network.LatestHeight()
		require.NoError(err)
		res, err := rpcClient.BlockResults(context.Background(), &latestHeight)
		if err != nil {
			return err
		}

		if len(res.ValidatorUpdates) == 0 {
			return errors.New("validator update not found yet")
		}

		valUpdate := res.ValidatorUpdates[0]
		require.Equal(
			valUpdate.GetPubKey().Sum.(*crypto.PublicKey_Ed25519).Ed25519,
			val.PubKey.Bytes(),
		)

		return nil
	}, 10)
}

// https://github.com/cosmos/cosmos-sdk/issues/10660
func (s *E2ETestSuite) TestEditValidatorMoniker() {
	val := s.network.Validators[0]
	require := s.Require()

	txCmd := cli.NewEditValidatorCmd()
	moniker := "testing"
	_, err := clitestutil.ExecTestCLICmd(val.ClientCtx, txCmd, []string{
		val.ValAddress.String(),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=%s", cli.FlagEditMoniker, moniker),
		fmt.Sprintf("--%s=https://newvalidator.io", cli.FlagWebsite),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	})
	require.NoError(err)
	s.Require().NoError(s.network.WaitForNextBlock())

	queryCmd := cli.GetCmdQueryValidator()
	res, err := clitestutil.ExecTestCLICmd(
		val.ClientCtx, queryCmd,
		[]string{val.ValAddress.String(), fmt.Sprintf("--%s=json", flags.FlagOutput)},
	)
	require.NoError(err)
	var result types.Validator
	require.NoError(val.ClientCtx.Codec.UnmarshalJSON(res.Bytes(), &result))
	require.Equal(result.GetMoniker(), moniker)

	_, err = clitestutil.ExecTestCLICmd(val.ClientCtx, txCmd, []string{
		val.ValAddress.String(),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastSync),
		fmt.Sprintf("--%s=https://newvalidator.io", cli.FlagWebsite),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	})
	require.NoError(err)
	s.Require().NoError(s.network.WaitForNextBlock())

	res, err = clitestutil.ExecTestCLICmd(
		val.ClientCtx, queryCmd,
		[]string{val.ValAddress.String(), fmt.Sprintf("--%s=json", flags.FlagOutput)},
	)
	require.NoError(err)
	require.NoError(val.ClientCtx.Codec.UnmarshalJSON(res.Bytes(), &result))
	require.Equal(result.GetMoniker(), moniker)
}
