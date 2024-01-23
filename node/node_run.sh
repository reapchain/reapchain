#!/bin/bash

if [ $# -lt 2 ] ; then
	echo $'\nPlease indicate if you want the node to run in foreground or background and the directory of the node data'
	echo $'\n./node_run.sh [FOREGROUND OR BACKGROUND] [DIRECTORY]\n'
	echo $'\n./node_run.sh f data \n'
	echo $'\n./node_run.sh b reapchain/nodes/node-data\n'
	exit
fi



REAPCHAIND="reapchaind"

RUN_TYPE=$1
DATA_DIR=$2

CURRENT_DATA_DIR="$DATA_DIR/data"
CURRENT_LOG_DIR="$DATA_DIR/log"

if [ $RUN_TYPE = "f" ]; then
  echo "$REAPCHAIND start --home $CURRENT_DATA_DIR"
  $REAPCHAIND start --home $CURRENT_DATA_DIR

else
  echo "$REAPCHAIND start --home $CURRENT_DATA_DIR &> $CURRENT_LOG_DIR/log.log &"
  $REAPCHAIND start --home $CURRENT_DATA_DIR &> $CURRENT_LOG_DIR/log.log &
fi