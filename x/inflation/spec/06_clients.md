<!--
order: 8
-->

# Clients

A user can query the `x/incentives` module using the CLI, JSON-RPC, gRPC or
REST.

## CLI

Find below a list of `reapchaind` commands added with the `x/inflation` module. You
can obtain the full list by using the `reapchaind -h` command.

### Queries

The `query` commands allow users to query `inflation` state.

**`period`**

Allows users to query the current inflation period.

```go
reapchaind query inflation period [flags]
```

**`epoch-mint-provision`**

Allows users to query the current inflation epoch provisions value.

```go
reapchaind query inflation epoch-mint-provision [flags]
```

**`skipped-epochs`**

Allows users to query the current number of skipped epochs.

```go
reapchaind query inflation skipped-epochs [flags]
```

**`total-supply`**

Allows users to query the total supply of tokens in circulation.

```go
reapchaind query inflation total-supply [flags]
```

**`inflation-rate`**

Allows users to query the inflation rate of the current period.

```go
reapchaind query inflation inflation-rate [flags]
```

**`params`**

Allows users to query the current inflation parameters.

```go
reapchaind query inflation params [flags]
```

### Proposals

The `tx gov submit-proposal` commands allow users to query create a proposal
using the governance module CLI:

**`param-change`**

Allows users to submit a `ParameterChangeProposal`.

```bash
reapchaind tx gov submit-proposal param-change [proposal-file] [flags]
```

## gRPC

### Queries

| Verb   | Method                                        | Description                                   |
| ------ | --------------------------------------------- | --------------------------------------------- |
| `gRPC` | `reapchain.inflation.v1.Query/Period`             | Gets current inflation period                 |
| `gRPC` | `reapchain.inflation.v1.Query/EpochMintProvision` | Gets current inflation epoch provisions value |
| `gRPC` | `reapchain.inflation.v1.Query/Params`             | Gets current inflation parameters             |
| `gRPC` | `reapchain.inflation.v1.Query/SkippedEpochs`      | Gets current number of skipped epochs         |
| `gRPC` | `reapchain.inflation.v1.Query/TotalSupply`        | Gets current total supply                     |
| `gRPC` | `reapchain.inflation.v1.Query/InflationRate`      | Gets current inflation rate                   |
| `GET`  | `/reapchain/inflation/v1/period`                  | Gets current inflation period                 |
| `GET`  | `/reapchain/inflation/v1/epoch_mint_provision`    | Gets current inflation epoch provisions value |
| `GET`  | `/reapchain/inflation/v1/skipped_epochs`          | Gets current number of skipped epochs         |
| `GET`  | `/reapchain/inflation/v1/total_supply`          | Gets current total supply                     |
| `GET`  | `/reapchain/inflation/v1/inflation_rate`          | Gets current inflation rate                   |
| `GET`  | `/reapchain/inflation/v1/params`                  | Gets current inflation parameters             |
