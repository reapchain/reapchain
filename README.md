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

## Local

**Build**
```
make build
make install 
```

**Configure**
```
sh init_single.sh
```
or
```
sh init_single.sh testnet  # for testnet
```

**Run**
```
reapchaind start                # Run a node
```

**visit with your browser**
* Node: http://localhost:26657/

## Localnet with 4 nodes

**Run**
```
make localnet-start
```

**Stop**
```
make localnet-stop
```
## Docker
**Warnings**: Initial development is in progress, but there has not yet been a stable.

# How to contribute
check out [CONTRIBUTING.md](CONTRIBUTING.md) for our guidelines & policies for how we develop ReapChain. Thank you to all those who have contributed!

# Guide && Resources
- [ReapchainGuide](https://reapchain.gitbook.io/mainnet)
- [Mainnet Dashboard](https://dashboard.reapchain.org)
- [Testnet Dashboard](https://dashboard.reapchain.org)
- [Faucet](https://test-faucet.reapchain.org)
