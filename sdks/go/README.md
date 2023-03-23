<div align="center">
    <picture>
        <source srcset="https://user-images.githubusercontent.com/6267663/221572723-e77f55a3-5d19-4a13-94f8-e7b0b340d71e.svg" media="(prefers-color-scheme: dark)">
        <img src="https://user-images.githubusercontent.com/6267663/221572726-6982541c-d1cf-4d9f-9bbf-cd774a2713e6.svg">
    </picture>
   <h1>Formance Go SDK</h1>
   <p><strong>Open Source Ledger for money-moving platforms</strong></p>
   <p>Build and track custom fit money flows on a scalable financial infrastructure.</p>
   <a href="https://docs.formance.com"><img src="https://img.shields.io/static/v1?label=Docs&message=Docs&color=000&style=for-the-badge" /></a>
   <a href="https://join.slack.com/t/formance-community/shared_invite/zt-1of48xmgy-Jc6RH8gzcWf5D0qD2HBPQA"><img src="https://img.shields.io/static/v1?label=Slack&message=Join&color=7289da&style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/formancehq/formance-sdk-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/formancehq/formance-sdk-go"
    "github.com/formancehq/formance-sdk-go/pkg/models/shared"
    "github.com/formancehq/formance-sdk-go/pkg/models/operations"
)

func main() {
    s := formance.New(
        formance.WithSecurity(shared.Security{
            Authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.GetServerInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.ServerInfo != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## SDK Available Operations

### Formance SDK

* `GetServerInfo` - Get server info
* `PaymentsgetServerInfo` - Get server info
* `SearchgetServerInfo` - Get server info

### Auth

* `AddScopeToClient` - Add scope to client
* `AddTransientScope` - Add a transient scope to a scope
* `CreateClient` - Create client
* `CreateScope` - Create scope
* `CreateSecret` - Add a secret to a client
* `DeleteClient` - Delete client
* `DeleteScope` - Delete scope
* `DeleteScopeFromClient` - Delete scope from client
* `DeleteSecret` - Delete a secret from a client
* `DeleteTransientScope` - Delete a transient scope from a scope
* `ListClients` - List clients
* `ListScopes` - List scopes
* `ListUsers` - List users
* `ReadClient` - Read client
* `ReadScope` - Read scope
* `ReadUser` - Read user
* `UpdateClient` - Update client
* `UpdateScope` - Update scope

### Ledger

* `CreateTransactions` - Create a new batch of transactions to a ledger
* `AddMetadataOnTransaction` - Set the metadata of a transaction by its ID
* `AddMetadataToAccount` - Add metadata to an account
* `CountAccounts` - Count the accounts from a ledger
* `CountTransactions` - Count the transactions from a ledger
* `CreateTransaction` - Create a new transaction to a ledger
* `GetAccount` - Get account by its address
* `GetBalances` - Get the balances from a ledger's account
* `GetBalancesAggregated` - Get the aggregated balances from selected accounts
* `GetInfo` - Show server information
* `GetLedgerInfo` - Get information about a ledger
* `GetMapping` - Get the mapping of a ledger
* `GetTransaction` - Get transaction from a ledger by its ID
* `ListAccounts` - List accounts from a ledger
* `ListLogs` - List the logs from a ledger
* `ListTransactions` - List transactions from a ledger
* `ReadStats` - Get statistics from a ledger
* `RevertTransaction` - Revert a ledger transaction by its ID
* `RunScript` - Execute a Numscript
* `UpdateMapping` - Update the mapping of a ledger

### Orchestration

* `CancelEvent` - Cancel a running workflow
* `CreateWorkflow` - Create workflow
* `GetInstance` - Get a workflow instance by id
* `GetInstanceHistory` - Get a workflow instance history by id
* `GetInstanceStageHistory` - Get a workflow instance stage history
* `GetWorkflow` - Get a flow by id
* `ListInstances` - List instances of a workflow
* `ListWorkflows` - List registered workflows
* `OrchestrationgetServerInfo` - Get server info
* `RunWorkflow` - Run workflow
* `SendEvent` - Send an event to a running workflow

### Payments

* `ConnectorsStripeTransfer` - Transfer funds between Stripe accounts
* `ConnectorsTransfer` - Transfer funds between Connector accounts
* `GetConnectorTask` - Read a specific task of the connector
* `GetPayment` - Get a payment
* `InstallConnector` - Install a connector
* `ListAllConnectors` - List all installed connectors
* `ListConfigsAvailableConnectors` - List the configs of each available connector
* `ListConnectorTasks` - List tasks from a connector
* `ListConnectorsTransfers` - List transfers and their statuses
* `ListPayments` - List payments
* `PaymentslistAccounts` - List accounts
* `ReadConnectorConfig` - Read the config of a connector
* `ResetConnector` - Reset a connector
* `UninstallConnector` - Uninstall a connector
* `UpdateMetadata` - Update metadata

### Search

* `Search` - Search

### Wallets

* `ConfirmHold` - Confirm a hold
* `CreateBalance` - Create a balance
* `CreateWallet` - Create a new wallet
* `CreditWallet` - Credit a wallet
* `DebitWallet` - Debit a wallet
* `GetBalance` - Get detailed balance
* `GetHold` - Get a hold
* `GetHolds` - Get all holds for a wallet
* `GetTransactions`
* `GetWallet` - Get a wallet
* `ListBalances` - List balances of a wallet
* `ListWallets` - List all wallets
* `UpdateWallet` - Update a wallet
* `VoidHold` - Cancel a hold
* `WalletsgetServerInfo` - Get server info

### Webhooks

* `ActivateConfig` - Activate one config
* `ChangeConfigSecret` - Change the signing secret of a config
* `DeactivateConfig` - Deactivate one config
* `DeleteConfig` - Delete one config
* `GetManyConfigs` - Get many configs
* `InsertConfig` - Insert a new config
* `TestConfig` - Test one config
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
