<!--
parent:
  order: false
-->

<div align="center">
  <h1> Reapchain </h1>
</div>

<!-- TODO: add banner -->
<!-- ![banner](docs/ethermint.jpg) -->

<div align="center">
  <a href="https://github.com/reapchain/reapchain/releases/latest">
    <img alt="Version" src="https://img.shields.io/github/tag/reapchain/reapchain.svg" />
  </a>
  <a href="https://github.com/reapchain/reapchain/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/reapchain/reapchain.svg" />
  </a>
</div>
<div align="center">
</div>

This repository hosts `ReapChain`. This repository is forked from [evmos](https://github.com/evmos/evmos) at 2022-05-03. ReapChain is a mainnet app implementation using [reapchain cosmos-sdk](https://github.com/reapchain/cosmos-sdk), [reapchain-core](https://github.com/reapchain/reapchain-core), [ethermint](https://github.com/reapchain/ethermint) and [ibc-go](https://github.com/reapchain/ibc-go).

**Node**: Requires [Go 1.18+](https://golang.org/dl/)

# Quick Start

**Build**
```
make build
make install 
```

**Check version**
```
reapchaind version
```

**Change folder**
```
cd node
```
**Configure**
```
sh node_initialize.sh
```

**Run**
```
# sh node_run.sh [FOREGROUND or BACKGROUND] [DATA DIR]
sh node_run.sh b reapchain-node-data
```


# Check Node
**Check process**
```
ps -ef | grep reapchaind
```
**Check log**
```
tail -f reapchain-node-data/log/log.log
```
**visit with your browser**
* Node: http://[local-IP]:27000/

## Ref Initial Setup
- https://reapchain.gitbook.io/mainnet/user-guides/initial-setup

## Docker
**Warnings**: Initial development is in progress, but there has not yet been a stable.

# How to contribute
check out [CONTRIBUTING.md](CONTRIBUTING.md) for our guidelines & policies for how we develop ReapChain. Thank you to all those who have contributed!

# Guide && Resources
- [ReapchainGuide](https://reapchain.gitbook.io)
- [Mainnet Dashboard](https://dashboard.reapchain.org)
- [Testnet Dashboard](https://test-dashboard.reapchain.org)
- [Faucet](https://test-faucet.reapchain.org)
