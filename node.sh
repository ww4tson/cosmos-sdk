#!/usr/bin/env bash
set -e

# This script is for running a local node, mostly for testing and demo purposes.
# It creates a chain with one validator, called `VALIDATOR`.

# Build the SDK's blockchain node binary, called "simapp"
make build

# Set variables
CFG_DIR=~/.simapp
BUILD_CMD=./build/simd
VALIDATOR=validator
KEY1=key1
KEY2=key2
KEY3=key3
KEY4=key4
MULTIKEY=multikey
CHAIN_ID=test-chain
MONIKER=likhita
HOME=~/multisig-test

# Cleanup previous installations, if any.
rm -rf $HOME

# Add VALIDATOR's key to the keyring.
$BUILD_CMD keys add $VALIDATOR --keyring-backend test --home $HOME 
$BUILD_CMD keys add $KEY1 --keyring-backend test --home $HOME 
$BUILD_CMD keys add $KEY2 --keyring-backend test --home $HOME 
$BUILD_CMD keys add $KEY3 --keyring-backend test --home $HOME 

# Create multi type key `multikey`.
$BUILD_CMD keys add $MULTIKEY --multisig "key1,key2" --multisig-threshold 1 --keyring-backend test --home $HOME

VALIDATOR_ADDRESS=$($BUILD_CMD keys show $VALIDATOR -a --keyring-backend test --home $HOME)

# Initialize the genesis file. It is available under $HOME/config/genesis.json.
$BUILD_CMD init $MONIKER --chain-id $CHAIN_ID --home $HOME
$BUILD_CMD genesis add-genesis-account $VALIDATOR_ADDRESS 10000000000stake --home $HOME --keyring-backend test
$BUILD_CMD genesis gentx $VALIDATOR 1000000000stake --keyring-backend test --chain-id $CHAIN_ID --home $HOME
$BUILD_CMD genesis collect-gentxs --home $HOME

# Run the node.
$BUILD_CMD start --home $HOME

# Leave the node running in a terminal. Now execute the script to test multisign-batch command.