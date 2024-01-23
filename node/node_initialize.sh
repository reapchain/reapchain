#!/bin/bash
OSTYPE="`uname | tr '[:upper:]' '[:lower:]'`"
if [[ "$OSTYPE" = "darwin"* ]]; then
    HOST=$(ifconfig en0 | grep 192 | cut -f 2 -d ' ')
elif [[ "$OSTYPE" = "linux"* ]]; then
    HOST=$(hostname -I | cut -f 1 -d ' ')
fi

PERSISTENT_PEER="e680439ea3ce99f020cc3250084b2cac61ac0042@state-sync-p2p.reapchain.org:27100"
REAPCHAIND="reapchaind"

CHAIN_ID="reapchain_221230-1"

echo -n "Please input Data Directory (default: reapchain-node-data): "
read -r DATA_DIR

echo -n "Please input your RPC Port (default: 27000): "
read -r RPC_PORT

echo -n "Please input your P2P Port (default: 27100): "
read -r P2P_PORT

echo -n "Please input APP RPC Port (default: 9090): "
read -r APP_RPC_PORT

echo -n "Please input Moniker (default: reapchaind): "
read -r MONIKER

  
if [ $DATA_DIR == ""]; then
    DATA_DIR=reapchain-node-data
fi

if [ $RPC_PORT == ""]; then
    RPC_PORT=27000
fi

if [ $REAPCHAIN_PORT == ""]; then
    REAPCHAIN_PORT=27100
fi

if [ $APP_RPC_PORT == ""]; then
    APP_RPC_PORT=9090
fi

if [ $MONIKER == ""]; then
    MONIKER="reapchaind"
fi


CURRENT_DATA_DIR="$DATA_DIR/data"
CURRENT_LOG_DIR="$DATA_DIR/log"

rm -rf $CURRENT_DATA_DIR
rm -rf $CURRENT_LOG_DIR
mkdir -p $CURRENT_LOG_DIR

$REAPCHAIND init $MONIKER --home $CURRENT_DATA_DIR --chain-id $CHAIN_ID

echo "RPC_PORT: $RPC_PORT"
echo "REAPCHAIN_PORT: $REAPCHAIN_PORT"

echo "cp ./genesis.json $CURRENT_DATA_DIR/config/genesis.json"
cp ./genesis.json $CURRENT_DATA_DIR/config/genesis.json

# EnableUnsafeCORS defines if CORS should be enabled (unsafe - use it at your own risk).
sed -i.temp 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' $CURRENT_DATA_DIR/config/app.toml

sed -i.temp 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["\*"\]/g' $CURRENT_DATA_DIR/config/config.toml
sed -i.temp 's/cors_allowed_headers = \["Origin", "Accept", "Content-Type", "X-Requested-With", "X-Server-Time", \]/cors_allowed_headers = \["Origin", "Accept", "Content-Type", "X-Requested-With", "X-Server-Time", "Access-Control-Request-Method", "Access-Control-Request-Headers", "Authorization", \]/g' $CURRENT_DATA_DIR/config/config.toml

# set REAPCHAIN adress
sed -i.temp "s/laddr = \"tcp:\/\/0.0.0.0:26656\"/laddr = \"tcp:\/\/${HOST}:${REAPCHAIN_PORT}\"/g" $CURRENT_DATA_DIR/config/config.toml 

# set RCP address
sed -i.temp "s/laddr = \"tcp:\/\/127.0.0.1:26657\"/laddr = \"tcp:\/\/${HOST}:${RPC_PORT}\"/g" $CURRENT_DATA_DIR/config/config.toml

# Moniker
sed -i.temp "s/moniker = \"$MONIKER\"/moniker = \"node$INDEX\"/g" $CURRENT_DATA_DIR/config/config.toml

sed -i.temp 's/allow_duplicate_ip = false/allow_duplicate_ip = true/g' $CURRENT_DATA_DIR/config/config.toml


# [p2p: max_packet_msg_payload_size]
sed -i.temp 's/max_packet_msg_payload_size = 1024/max_packet_msg_payload_size = 10240/g' $CURRENT_DATA_DIR/config/config.toml

# Enable API Server 
sed -i.temp 's/enable = false/enable = true/g' $CURRENT_DATA_DIR/config/app.toml

# [Rosetta disable]
sed -i.temp '138s/enable = true/enable = false/' $CURRENT_DATA_DIR/config/app.toml



# Snapshot Server Initialization
# Faster Syncing
SNAP_RPC="https://state-sync-rpc.reapchain.org:443"
SNAP_RPC_SECOND="https://rpc-network.reapchain.org:443"


LATEST_HEIGHT=$(curl -s $SNAP_RPC/block | jq -r .result.block.header.height); \
BLOCK_HEIGHT=$((LATEST_HEIGHT - 1200)); \
TRUST_HASH=$(curl -s "$SNAP_RPC/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)

sed -i.temp -E "s|^(enable[[:space:]]+=[[:space:]]+).*$|\1true| ; \
s|^(rpc_servers[[:space:]]+=[[:space:]]+).*$|\1\"$SNAP_RPC,$SNAP_RPC_SECOND\"| ; \
s|^(trust_height[[:space:]]+=[[:space:]]+).*$|\1$BLOCK_HEIGHT| ; \
s|^(trust_hash[[:space:]]+=[[:space:]]+).*$|\1\"$TRUST_HASH\"| ; \
s|^(seeds[[:space:]]+=[[:space:]]+).*$|\1\"\"|" $CURRENT_DATA_DIR/config/config.toml


# Persistent Peer
sed -i.temp "s/persistent_peers = .*/persistent_peers = \"$PERSISTENT_PEER\"/g" $CURRENT_DATA_DIR/config/config.toml



rm -rf $CURRENT_DATA_DIR/config/config.toml.temp
rm -rf $CURRENT_DATA_DIR/config/app.toml.temp