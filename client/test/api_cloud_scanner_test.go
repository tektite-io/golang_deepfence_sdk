/*
Deepfence ThreatMapper

Testing CloudScannerAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/deepfence/golang_deepfence_sdk/client"
)

func Test_client_CloudScannerAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CloudScannerAPIService CountResultsCloudComplianceScan", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CloudScannerAPI.CountResultsCloudComplianceScan(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudScannerAPIService IngestCloudComplianceScanStatus", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CloudScannerAPI.IngestCloudComplianceScanStatus(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudScannerAPIService IngestCloudCompliances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.CloudScannerAPI.IngestCloudCompliances(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudScannerAPIService ListCloudComplianceScan", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CloudScannerAPI.ListCloudComplianceScan(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudScannerAPIService ResultsCloudComplianceScan", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CloudScannerAPI.ResultsCloudComplianceScan(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CloudScannerAPIService StatusCloudComplianceScan", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CloudScannerAPI.StatusCloudComplianceScan(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
