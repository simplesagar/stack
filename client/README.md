# Go API client for client

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v0.2.1
- Package version: latest
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import client "github.com/formancehq/webhooks/client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), client.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), client.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), client.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), client.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://.o.formance.cloud/webhooks*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ConfigsApi* | [**ActivateOneConfig**](docs/ConfigsApi.md#activateoneconfig) | **Put** /configs/{id}/activate | Activate one config
*ConfigsApi* | [**ChangeOneConfigSecret**](docs/ConfigsApi.md#changeoneconfigsecret) | **Put** /configs/{id}/secret/change | Change the signing secret of a config
*ConfigsApi* | [**DeactivateOneConfig**](docs/ConfigsApi.md#deactivateoneconfig) | **Put** /configs/{id}/deactivate | Deactivate one config
*ConfigsApi* | [**DeleteOneConfig**](docs/ConfigsApi.md#deleteoneconfig) | **Delete** /configs/{id} | Delete one config
*ConfigsApi* | [**GetManyConfigs**](docs/ConfigsApi.md#getmanyconfigs) | **Get** /configs | Get many configs
*ConfigsApi* | [**InsertOneConfig**](docs/ConfigsApi.md#insertoneconfig) | **Post** /configs | Insert a new config 
*HealthApi* | [**HealthCheck**](docs/HealthApi.md#healthcheck) | **Get** /_healthcheck | Health check of the server


## Documentation For Models

 - [ChangeOneConfigSecretRequest](docs/ChangeOneConfigSecretRequest.md)
 - [Config](docs/Config.md)
 - [ConfigUser](docs/ConfigUser.md)
 - [Cursor](docs/Cursor.md)
 - [GetManyConfigs200Response](docs/GetManyConfigs200Response.md)
 - [GetManyConfigs200ResponseCursor](docs/GetManyConfigs200ResponseCursor.md)
 - [GetManyConfigs200ResponseCursorAllOf](docs/GetManyConfigs200ResponseCursorAllOf.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


