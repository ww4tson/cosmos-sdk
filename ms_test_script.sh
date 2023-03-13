#!/usr/bin/env bash
set -e

# This script is for reproducing multisign batch issue.
# ref: https://github.com/cosmos/cosmos-sdk/issues/12728

# Set variables
BUILD_CMD=./build/simd
VALIDATOR=validator
KEY1=key1
KEY2=key2
KEY3=key3
MULTIKEY=multikey
CHAIN_ID=test-chain
HOME=~/multisig-test

MULTIKEY_ADDRESS=$($BUILD_CMD keys show $MULTIKEY -a --keyring-backend test --home $HOME)
KEY1_ADDRESS=$($BUILD_CMD keys show $KEY1 -a --keyring-backend test --home $HOME)
KEY2_ADDRESS=$($BUILD_CMD keys show $KEY2 -a --keyring-backend test --home $HOME)
KEY3_ADDRESS=$($BUILD_CMD keys show $KEY3 -a --keyring-backend test --home $HOME)


echo "Send funds to multikey."
$BUILD_CMD tx bank send $VALIDATOR $MULTIKEY_ADDRESS 10000000stake --keyring-backend test --chain-id $CHAIN_ID --home $HOME --broadcast-mode async -y

echo "Send some funds to key1 and key2 which are a part of multikey for signing purpose."
$BUILD_CMD tx bank send $VALIDATOR $KEY1_ADDRESS 10000000stake --keyring-backend test --chain-id $CHAIN_ID --home $HOME --broadcast-mode async -y
$BUILD_CMD tx bank send $VALIDATOR $KEY2_ADDRESS 10000000stake --keyring-backend test --chain-id $CHAIN_ID --home $HOME --broadcast-mode async -y

echo "Generate two multisig transactions with signer as multikey and save it in tx.json"
($BUILD_CMD tx bank send $MULTIKEY $KEY3_ADDRESS 10000000stake --keyring-backend test --chain-id $CHAIN_ID --home $HOME --generate-only && $BUILD_CMD tx bank send $MULTIKEY $KEY3_ADDRESS 10000000stake --keyring-backend test --chain-id $CHAIN_ID --home $HOME --generate-only) > tx.json
echo "---------------------------tx.json---------------------------"
echo `cat tx.json` 
echo "-------------------------------------------------------------"

echo "Sign tx.json using sign-batch with key1 which is a part of mulitkey and save the signature in separate json file sign1.json"
$BUILD_CMD tx sign-batch tx.json --keyring-backend test --chain-id $CHAIN_ID --home $HOME --from $KEY1 --multisig $MULTIKEY --signature-only > sign1.json
echo "---------------------------sign1.json---------------------------"
echo `cat sign1.json` 
echo "-------------------------------------------------------------"

echo "Sign tx.json using sign-batch with key2 which is a part of mulitkey and save the signature in separate json file sign2.json"
$BUILD_CMD tx sign-batch tx.json --keyring-backend test --chain-id $CHAIN_ID --home $HOME --from $KEY2 --multisig $MULTIKEY --signature-only > sign2.json
echo "---------------------------sign2.json---------------------------"
echo `cat sign2.json` 
echo "-------------------------------------------------------------"

echo "Try executing multisign-batch with tx.json and obtained signatures"
$BUILD_CMD tx multisign-batch tx.json $MULTIKEY sign1.json sign2.json --keyring-backend test --chain-id $CHAIN_ID --home $HOME --broadcast-mode async

