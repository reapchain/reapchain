<!--
order: 2
-->

# Tokens

Learn about the the different types of tokens available in reapchain. {synopsis}

## Introduction

reapchain is a Cosmos-based chain with full Ethereum Virtual Machine (EVM) support. Because of this [architecture](./../technical_concepts/architecture.md), tokens and assets in the network may come from different independent sources.

## The reapchain Token

The denomination used for staking, governance and gas consumption on the EVM is the reapchain. The reapchain provides the utility of: securing the Proof-of-Stake chain, token used for governance proposals, distribution of fees to validator and users, and as a mean of gas for running smart contracts on the EVM.

reapchain uses [Atto](https://en.wikipedia.org/wiki/Atto-) reapchain as the base denomination to maintain parity with Ethereum.

1 reapchain = 10<sup>18</sup> areap

This matches Ethereum denomination of:

1 ETH = 10<sup>18</sup> wei

## Cosmos Coins

Accounts can own Cosmos coins in their balance, which are used for operations with other Cosmos and transactions. Examples of these are using the coins for staking, IBC transfers, governance deposits and EVM.

## EVM Tokens

reapchain is compatible with ERC20 tokens and other non-fungible token standards (EIP721, EIP1155)
that are natively supported by the EVM.
