/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelComplianceScanInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelComplianceScanInfo{}

// ModelComplianceScanInfo struct for ModelComplianceScanInfo
type ModelComplianceScanInfo struct {
	BenchmarkTypes []string `json:"benchmark_types"`
	CloudProvider string `json:"cloud_provider"`
	CreatedAt int64 `json:"created_at"`
	NodeId string `json:"node_id"`
	NodeName string `json:"node_name"`
	NodeType string `json:"node_type"`
	ScanId string `json:"scan_id"`
	SeverityCounts map[string]int32 `json:"severity_counts"`
	Status string `json:"status"`
	StatusMessage string `json:"status_message"`
	UpdatedAt int64 `json:"updated_at"`
}

type _ModelComplianceScanInfo ModelComplianceScanInfo

// NewModelComplianceScanInfo instantiates a new ModelComplianceScanInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelComplianceScanInfo(benchmarkTypes []string, cloudProvider string, createdAt int64, nodeId string, nodeName string, nodeType string, scanId string, severityCounts map[string]int32, status string, statusMessage string, updatedAt int64) *ModelComplianceScanInfo {
	this := ModelComplianceScanInfo{}
	this.BenchmarkTypes = benchmarkTypes
	this.CloudProvider = cloudProvider
	this.CreatedAt = createdAt
	this.NodeId = nodeId
	this.NodeName = nodeName
	this.NodeType = nodeType
	this.ScanId = scanId
	this.SeverityCounts = severityCounts
	this.Status = status
	this.StatusMessage = statusMessage
	this.UpdatedAt = updatedAt
	return &this
}

// NewModelComplianceScanInfoWithDefaults instantiates a new ModelComplianceScanInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelComplianceScanInfoWithDefaults() *ModelComplianceScanInfo {
	this := ModelComplianceScanInfo{}
	return &this
}

// GetBenchmarkTypes returns the BenchmarkTypes field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelComplianceScanInfo) GetBenchmarkTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BenchmarkTypes
}

// GetBenchmarkTypesOk returns a tuple with the BenchmarkTypes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelComplianceScanInfo) GetBenchmarkTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.BenchmarkTypes) {
		return nil, false
	}
	return o.BenchmarkTypes, true
}

// SetBenchmarkTypes sets field value
func (o *ModelComplianceScanInfo) SetBenchmarkTypes(v []string) {
	o.BenchmarkTypes = v
}

// GetCloudProvider returns the CloudProvider field value
func (o *ModelComplianceScanInfo) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *ModelComplianceScanInfo) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *ModelComplianceScanInfo) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ModelComplianceScanInfo) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelComplianceScanInfo) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ModelComplianceScanInfo) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetNodeId returns the NodeId field value
func (o *ModelComplianceScanInfo) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelComplianceScanInfo) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelComplianceScanInfo) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeName returns the NodeName field value
func (o *ModelComplianceScanInfo) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *ModelComplianceScanInfo) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *ModelComplianceScanInfo) SetNodeName(v string) {
	o.NodeName = v
}

// GetNodeType returns the NodeType field value
func (o *ModelComplianceScanInfo) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *ModelComplianceScanInfo) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *ModelComplianceScanInfo) SetNodeType(v string) {
	o.NodeType = v
}

// GetScanId returns the ScanId field value
func (o *ModelComplianceScanInfo) GetScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value
// and a boolean to check if the value has been set.
func (o *ModelComplianceScanInfo) GetScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanId, true
}

// SetScanId sets field value
func (o *ModelComplianceScanInfo) SetScanId(v string) {
	o.ScanId = v
}

// GetSeverityCounts returns the SeverityCounts field value
// If the value is explicit nil, the zero value for map[string]int32 will be returned
func (o *ModelComplianceScanInfo) GetSeverityCounts() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.SeverityCounts
}

// GetSeverityCountsOk returns a tuple with the SeverityCounts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelComplianceScanInfo) GetSeverityCountsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.SeverityCounts) {
		return nil, false
	}
	return &o.SeverityCounts, true
}

// SetSeverityCounts sets field value
func (o *ModelComplianceScanInfo) SetSeverityCounts(v map[string]int32) {
	o.SeverityCounts = v
}

// GetStatus returns the Status field value
func (o *ModelComplianceScanInfo) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ModelComplianceScanInfo) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ModelComplianceScanInfo) SetStatus(v string) {
	o.Status = v
}

// GetStatusMessage returns the StatusMessage field value
func (o *ModelComplianceScanInfo) GetStatusMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value
// and a boolean to check if the value has been set.
func (o *ModelComplianceScanInfo) GetStatusMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusMessage, true
}

// SetStatusMessage sets field value
func (o *ModelComplianceScanInfo) SetStatusMessage(v string) {
	o.StatusMessage = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ModelComplianceScanInfo) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelComplianceScanInfo) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ModelComplianceScanInfo) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

func (o ModelComplianceScanInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelComplianceScanInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BenchmarkTypes != nil {
		toSerialize["benchmark_types"] = o.BenchmarkTypes
	}
	toSerialize["cloud_provider"] = o.CloudProvider
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_name"] = o.NodeName
	toSerialize["node_type"] = o.NodeType
	toSerialize["scan_id"] = o.ScanId
	if o.SeverityCounts != nil {
		toSerialize["severity_counts"] = o.SeverityCounts
	}
	toSerialize["status"] = o.Status
	toSerialize["status_message"] = o.StatusMessage
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *ModelComplianceScanInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"benchmark_types",
		"cloud_provider",
		"created_at",
		"node_id",
		"node_name",
		"node_type",
		"scan_id",
		"severity_counts",
		"status",
		"status_message",
		"updated_at",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varModelComplianceScanInfo := _ModelComplianceScanInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelComplianceScanInfo)

	if err != nil {
		return err
	}

	*o = ModelComplianceScanInfo(varModelComplianceScanInfo)

	return err
}

type NullableModelComplianceScanInfo struct {
	value *ModelComplianceScanInfo
	isSet bool
}

func (v NullableModelComplianceScanInfo) Get() *ModelComplianceScanInfo {
	return v.value
}

func (v *NullableModelComplianceScanInfo) Set(val *ModelComplianceScanInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelComplianceScanInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelComplianceScanInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelComplianceScanInfo(val *ModelComplianceScanInfo) *NullableModelComplianceScanInfo {
	return &NullableModelComplianceScanInfo{value: val, isSet: true}
}

func (v NullableModelComplianceScanInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelComplianceScanInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


