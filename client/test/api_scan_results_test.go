/*
Deepfence ThreatMapper

Testing ScanResultsAPIService

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

func Test_client_ScanResultsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ScanResultsAPIService BulkDeleteScans", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ScanResultsAPI.BulkDeleteScans(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScanResultsAPIService DeleteScanResult", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ScanResultsAPI.DeleteScanResult(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScanResultsAPIService DeleteScanResultsForScanID", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scanId string
		var scanType string

		httpRes, err := apiClient.ScanResultsAPI.DeleteScanResultsForScanID(context.Background(), scanId, scanType).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScanResultsAPIService DownloadScanResults", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var scanId string
		var scanType string

		resp, httpRes, err := apiClient.ScanResultsAPI.DownloadScanResults(context.Background(), scanId, scanType).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScanResultsAPIService GetAllNodesInScanResults", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ScanResultsAPI.GetAllNodesInScanResults(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScanResultsAPIService MaskScanResult", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ScanResultsAPI.MaskScanResult(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScanResultsAPIService NotifyScanResult", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ScanResultsAPI.NotifyScanResult(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ScanResultsAPIService UnmaskScanResult", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.ScanResultsAPI.UnmaskScanResult(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
