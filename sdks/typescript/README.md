<div align="center">
    <picture>
        <source srcset="https://user-images.githubusercontent.com/6267663/221572723-e77f55a3-5d19-4a13-94f8-e7b0b340d71e.svg" media="(prefers-color-scheme: dark)">
        <img src="https://user-images.githubusercontent.com/6267663/221572726-6982541c-d1cf-4d9f-9bbf-cd774a2713e6.svg">
    </picture>
   <h1>Formance Typescript SDK</h1>
   <p><strong>Open Source Ledger for money-moving platforms</strong></p>
   <p>Build and track custom fit money flows on a scalable financial infrastructure.</p>
   <a href="https://docs.formance.com"><img src="https://img.shields.io/static/v1?label=Docs&message=Docs&color=000&style=for-the-badge" /></a>
   <a href="https://join.slack.com/t/formance-community/shared_invite/zt-1of48xmgy-Jc6RH8gzcWf5D0qD2HBPQA"><img src="https://img.shields.io/static/v1?label=Slack&message=Join&color=7289da&style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## SDK Installation

### NPM

```bash
npm add @formance/formance-sdk
```

### Yarn

```bash
yarn add @formance/formance-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```typescript
import {
  GetServerInfoResponse
} from "@formance/formance-sdk/dist/sdk/models/operations";

import { AxiosError } from "axios";
import { SDK } from "@formance/formance-sdk";
const sdk = new SDK({
  security: {
    authorization: "Bearer YOUR_ACCESS_TOKEN_HERE",
  },
});

sdk.getServerInfo().then((res: GetServerInfoResponse | AxiosError) => {
   // handle response
});
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## SDK Available Operations

### SDK SDK

* `getServerInfo` - Get server info
* `paymentsgetServerInfo` - Get server info
* `searchgetServerInfo` - Get server info

### accounts

* `addMetadataToAccount` - Add metadata to an account
* `countAccounts` - Count the accounts from a ledger
* `getAccount` - Get account by its address
* `listAccounts` - List accounts from a ledger

### auth

* `addScopeToClient` - Add scope to client
* `addTransientScope` - Add a transient scope to a scope
* `createClient` - Create client
* `createScope` - Create scope
* `createSecret` - Add a secret to a client
* `deleteClient` - Delete client
* `deleteScope` - Delete scope
* `deleteScopeFromClient` - Delete scope from client
* `deleteSecret` - Delete a secret from a client
* `deleteTransientScope` - Delete a transient scope from a scope
* `listClients` - List clients
* `listScopes` - List scopes
* `listUsers` - List users
* `readClient` - Read client
* `readScope` - Read scope
* `readUser` - Read user
* `updateClient` - Update client
* `updateScope` - Update scope

### balances

* `getBalances` - Get the balances from a ledger's account
* `getBalancesAggregated` - Get the aggregated balances from selected accounts

### clients

* `addScopeToClient` - Add scope to client
* `createClient` - Create client
* `createSecret` - Add a secret to a client
* `deleteClient` - Delete client
* `deleteScopeFromClient` - Delete scope from client
* `deleteSecret` - Delete a secret from a client
* `listClients` - List clients
* `readClient` - Read client
* `updateClient` - Update client

### ledger

* `createTransactions` - Create a new batch of transactions to a ledger
* `addMetadataOnTransaction` - Set the metadata of a transaction by its ID
* `addMetadataToAccount` - Add metadata to an account
* `countAccounts` - Count the accounts from a ledger
* `countTransactions` - Count the transactions from a ledger
* `createTransaction` - Create a new transaction to a ledger
* `getAccount` - Get account by its address
* `getBalances` - Get the balances from a ledger's account
* `getBalancesAggregated` - Get the aggregated balances from selected accounts
* `getInfo` - Show server information
* `getLedgerInfo` - Get information about a ledger
* `getMapping` - Get the mapping of a ledger
* `getTransaction` - Get transaction from a ledger by its ID
* `listAccounts` - List accounts from a ledger
* `listLogs` - List the logs from a ledger
* `listTransactions` - List transactions from a ledger
* `readStats` - Get statistics from a ledger
* `revertTransaction` - Revert a ledger transaction by its ID
* `runScript` - Execute a Numscript
* `updateMapping` - Update the mapping of a ledger

### logs

* `listLogs` - List the logs from a ledger

### mapping

* `getMapping` - Get the mapping of a ledger
* `updateMapping` - Update the mapping of a ledger

### orchestration

* `cancelEvent` - Cancel a running workflow
* `createWorkflow` - Create workflow
* `getInstance` - Get a workflow instance by id
* `getInstanceHistory` - Get a workflow instance history by id
* `getInstanceStageHistory` - Get a workflow instance stage history
* `getWorkflow` - Get a flow by id
* `listInstances` - List instances of a workflow
* `listWorkflows` - List registered workflows
* `orchestrationgetServerInfo` - Get server info
* `runWorkflow` - Run workflow
* `sendEvent` - Send an event to a running workflow

### payments

* `connectorsStripeTransfer` - Transfer funds between Stripe accounts
* `connectorsTransfer` - Transfer funds between Connector accounts
* `getConnectorTask` - Read a specific task of the connector
* `getPayment` - Get a payment
* `installConnector` - Install a connector
* `listAllConnectors` - List all installed connectors
* `listConfigsAvailableConnectors` - List the configs of each available connector
* `listConnectorTasks` - List tasks from a connector
* `listConnectorsTransfers` - List transfers and their statuses
* `listPayments` - List payments
* `paymentslistAccounts` - List accounts
* `readConnectorConfig` - Read the config of a connector
* `resetConnector` - Reset a connector
* `uninstallConnector` - Uninstall a connector
* `updateMetadata` - Update metadata

### scopes

* `addTransientScope` - Add a transient scope to a scope
* `createScope` - Create scope
* `deleteScope` - Delete scope
* `deleteTransientScope` - Delete a transient scope from a scope
* `listScopes` - List scopes
* `readScope` - Read scope
* `updateScope` - Update scope

### script

* `runScript` - Execute a Numscript

### search

* `search` - Search

### server

* `getInfo` - Show server information

### stats

* `readStats` - Get statistics from a ledger

### transactions

* `createTransactions` - Create a new batch of transactions to a ledger
* `addMetadataOnTransaction` - Set the metadata of a transaction by its ID
* `countTransactions` - Count the transactions from a ledger
* `createTransaction` - Create a new transaction to a ledger
* `getTransaction` - Get transaction from a ledger by its ID
* `listTransactions` - List transactions from a ledger
* `revertTransaction` - Revert a ledger transaction by its ID

### users

* `listUsers` - List users
* `readUser` - Read user

### wallets

* `confirmHold` - Confirm a hold
* `createBalance` - Create a balance
* `createWallet` - Create a new wallet
* `creditWallet` - Credit a wallet
* `debitWallet` - Debit a wallet
* `getBalance` - Get detailed balance
* `getHold` - Get a hold
* `getHolds` - Get all holds for a wallet
* `getTransactions`
* `getWallet` - Get a wallet
* `listBalances` - List balances of a wallet
* `listWallets` - List all wallets
* `updateWallet` - Update a wallet
* `voidHold` - Cancel a hold
* `walletsgetServerInfo` - Get server info

### webhooks

* `activateConfig` - Activate one config
* `changeConfigSecret` - Change the signing secret of a config
* `deactivateConfig` - Deactivate one config
* `deleteConfig` - Delete one config
* `getManyConfigs` - Get many configs
* `insertConfig` - Insert a new config
* `testConfig` - Test one config
<!-- End SDK Available Operations -->

### SDK Generated by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)