/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.2.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelScanInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelScanInfo{}

// ModelScanInfo struct for ModelScanInfo
type ModelScanInfo struct {
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

type _ModelScanInfo ModelScanInfo

// NewModelScanInfo instantiates a new ModelScanInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelScanInfo(createdAt int64, nodeId string, nodeName string, nodeType string, scanId string, severityCounts map[string]int32, status string, statusMessage string, updatedAt int64) *ModelScanInfo {
	this := ModelScanInfo{}
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

// NewModelScanInfoWithDefaults instantiates a new ModelScanInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelScanInfoWithDefaults() *ModelScanInfo {
	this := ModelScanInfo{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *ModelScanInfo) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelScanInfo) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ModelScanInfo) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetNodeId returns the NodeId field value
func (o *ModelScanInfo) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelScanInfo) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelScanInfo) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeName returns the NodeName field value
func (o *ModelScanInfo) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *ModelScanInfo) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *ModelScanInfo) SetNodeName(v string) {
	o.NodeName = v
}

// GetNodeType returns the NodeType field value
func (o *ModelScanInfo) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *ModelScanInfo) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *ModelScanInfo) SetNodeType(v string) {
	o.NodeType = v
}

// GetScanId returns the ScanId field value
func (o *ModelScanInfo) GetScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value
// and a boolean to check if the value has been set.
func (o *ModelScanInfo) GetScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanId, true
}

// SetScanId sets field value
func (o *ModelScanInfo) SetScanId(v string) {
	o.ScanId = v
}

// GetSeverityCounts returns the SeverityCounts field value
// If the value is explicit nil, the zero value for map[string]int32 will be returned
func (o *ModelScanInfo) GetSeverityCounts() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.SeverityCounts
}

// GetSeverityCountsOk returns a tuple with the SeverityCounts field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelScanInfo) GetSeverityCountsOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.SeverityCounts) {
		return nil, false
	}
	return &o.SeverityCounts, true
}

// SetSeverityCounts sets field value
func (o *ModelScanInfo) SetSeverityCounts(v map[string]int32) {
	o.SeverityCounts = v
}

// GetStatus returns the Status field value
func (o *ModelScanInfo) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ModelScanInfo) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ModelScanInfo) SetStatus(v string) {
	o.Status = v
}

// GetStatusMessage returns the StatusMessage field value
func (o *ModelScanInfo) GetStatusMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value
// and a boolean to check if the value has been set.
func (o *ModelScanInfo) GetStatusMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusMessage, true
}

// SetStatusMessage sets field value
func (o *ModelScanInfo) SetStatusMessage(v string) {
	o.StatusMessage = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ModelScanInfo) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelScanInfo) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ModelScanInfo) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

func (o ModelScanInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelScanInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

func (o *ModelScanInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varModelScanInfo := _ModelScanInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelScanInfo)

	if err != nil {
		return err
	}

	*o = ModelScanInfo(varModelScanInfo)

	return err
}

type NullableModelScanInfo struct {
	value *ModelScanInfo
	isSet bool
}

func (v NullableModelScanInfo) Get() *ModelScanInfo {
	return v.value
}

func (v *NullableModelScanInfo) Set(val *ModelScanInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelScanInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelScanInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelScanInfo(val *ModelScanInfo) *NullableModelScanInfo {
	return &NullableModelScanInfo{value: val, isSet: true}
}

func (v NullableModelScanInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelScanInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


