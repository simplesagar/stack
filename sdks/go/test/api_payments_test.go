/*
Formance Stack API

Testing PaymentsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package formance

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    client "./openapi"
)

func Test_formance_PaymentsApiService(t *testing.T) {

    configuration := client.NewConfiguration()
    apiClient := client.NewAPIClient(configuration)

    t.Run("Test PaymentsApiService ConnectorsStripeTransfer", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.PaymentsApi.ConnectorsStripeTransfer(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PaymentsApiService GetAllConnectors", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.PaymentsApi.GetAllConnectors(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PaymentsApiService GetAllConnectorsConfigs", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.PaymentsApi.GetAllConnectorsConfigs(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PaymentsApiService GetConnectorTask", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var connector Connectors
        var taskId string

        resp, httpRes, err := apiClient.PaymentsApi.GetConnectorTask(context.Background(), connector, taskId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PaymentsApiService GetPayment", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var paymentId string

        resp, httpRes, err := apiClient.PaymentsApi.GetPayment(context.Background(), paymentId).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PaymentsApiService InstallConnector", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var connector Connectors

        resp, httpRes, err := apiClient.PaymentsApi.InstallConnector(context.Background(), connector).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PaymentsApiService ListConnectorTasks", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var connector Connectors

        resp, httpRes, err := apiClient.PaymentsApi.ListConnectorTasks(context.Background(), connector).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PaymentsApiService ListPayments", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.PaymentsApi.ListPayments(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PaymentsApiService ReadConnectorConfig", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var connector Connectors

        resp, httpRes, err := apiClient.PaymentsApi.ReadConnectorConfig(context.Background(), connector).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PaymentsApiService ResetConnector", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var connector Connectors

        resp, httpRes, err := apiClient.PaymentsApi.ResetConnector(context.Background(), connector).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test PaymentsApiService UninstallConnector", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var connector Connectors

        resp, httpRes, err := apiClient.PaymentsApi.UninstallConnector(context.Background(), connector).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
