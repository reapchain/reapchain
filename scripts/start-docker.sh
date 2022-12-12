#!/bin/bash

KEY="mykey"
CHAINID="mercury_2022-3"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t reapchain-datadir.XXXXX)

echo "create and add new keys"
./reapchaind keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Reapchain with moniker=$MONIKER and chain-id=$CHAINID"
./reapchaind init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./reapchaind add-genesis-account \
"$(./reapchaind keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000areap,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./reapchaind gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./reapchaind collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./reapchaind validate-genesis --home $DATA_DIR

echo "starting reapchain node $i in background ..."
./reapchaind start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started reapchain node"
tail -f /dev/null