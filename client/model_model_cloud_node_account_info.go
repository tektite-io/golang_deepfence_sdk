/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.0.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ModelCloudNodeAccountInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCloudNodeAccountInfo{}

// ModelCloudNodeAccountInfo struct for ModelCloudNodeAccountInfo
type ModelCloudNodeAccountInfo struct {
	Active *bool `json:"active,omitempty"`
	CloudProvider *string `json:"cloud_provider,omitempty"`
	CompliancePercentage *float32 `json:"compliance_percentage,omitempty"`
	LastScanId *string `json:"last_scan_id,omitempty"`
	LastScanStatus *string `json:"last_scan_status,omitempty"`
	NodeId *string `json:"node_id,omitempty"`
	NodeName *string `json:"node_name,omitempty"`
}

// NewModelCloudNodeAccountInfo instantiates a new ModelCloudNodeAccountInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCloudNodeAccountInfo() *ModelCloudNodeAccountInfo {
	this := ModelCloudNodeAccountInfo{}
	return &this
}

// NewModelCloudNodeAccountInfoWithDefaults instantiates a new ModelCloudNodeAccountInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCloudNodeAccountInfoWithDefaults() *ModelCloudNodeAccountInfo {
	this := ModelCloudNodeAccountInfo{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *ModelCloudNodeAccountInfo) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeAccountInfo) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountInfo) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *ModelCloudNodeAccountInfo) SetActive(v bool) {
	o.Active = &v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *ModelCloudNodeAccountInfo) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeAccountInfo) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountInfo) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *ModelCloudNodeAccountInfo) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetCompliancePercentage returns the CompliancePercentage field value if set, zero value otherwise.
func (o *ModelCloudNodeAccountInfo) GetCompliancePercentage() float32 {
	if o == nil || IsNil(o.CompliancePercentage) {
		var ret float32
		return ret
	}
	return *o.CompliancePercentage
}

// GetCompliancePercentageOk returns a tuple with the CompliancePercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeAccountInfo) GetCompliancePercentageOk() (*float32, bool) {
	if o == nil || IsNil(o.CompliancePercentage) {
		return nil, false
	}
	return o.CompliancePercentage, true
}

// HasCompliancePercentage returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountInfo) HasCompliancePercentage() bool {
	if o != nil && !IsNil(o.CompliancePercentage) {
		return true
	}

	return false
}

// SetCompliancePercentage gets a reference to the given float32 and assigns it to the CompliancePercentage field.
func (o *ModelCloudNodeAccountInfo) SetCompliancePercentage(v float32) {
	o.CompliancePercentage = &v
}

// GetLastScanId returns the LastScanId field value if set, zero value otherwise.
func (o *ModelCloudNodeAccountInfo) GetLastScanId() string {
	if o == nil || IsNil(o.LastScanId) {
		var ret string
		return ret
	}
	return *o.LastScanId
}

// GetLastScanIdOk returns a tuple with the LastScanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeAccountInfo) GetLastScanIdOk() (*string, bool) {
	if o == nil || IsNil(o.LastScanId) {
		return nil, false
	}
	return o.LastScanId, true
}

// HasLastScanId returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountInfo) HasLastScanId() bool {
	if o != nil && !IsNil(o.LastScanId) {
		return true
	}

	return false
}

// SetLastScanId gets a reference to the given string and assigns it to the LastScanId field.
func (o *ModelCloudNodeAccountInfo) SetLastScanId(v string) {
	o.LastScanId = &v
}

// GetLastScanStatus returns the LastScanStatus field value if set, zero value otherwise.
func (o *ModelCloudNodeAccountInfo) GetLastScanStatus() string {
	if o == nil || IsNil(o.LastScanStatus) {
		var ret string
		return ret
	}
	return *o.LastScanStatus
}

// GetLastScanStatusOk returns a tuple with the LastScanStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeAccountInfo) GetLastScanStatusOk() (*string, bool) {
	if o == nil || IsNil(o.LastScanStatus) {
		return nil, false
	}
	return o.LastScanStatus, true
}

// HasLastScanStatus returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountInfo) HasLastScanStatus() bool {
	if o != nil && !IsNil(o.LastScanStatus) {
		return true
	}

	return false
}

// SetLastScanStatus gets a reference to the given string and assigns it to the LastScanStatus field.
func (o *ModelCloudNodeAccountInfo) SetLastScanStatus(v string) {
	o.LastScanStatus = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *ModelCloudNodeAccountInfo) GetNodeId() string {
	if o == nil || IsNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeAccountInfo) GetNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountInfo) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *ModelCloudNodeAccountInfo) SetNodeId(v string) {
	o.NodeId = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *ModelCloudNodeAccountInfo) GetNodeName() string {
	if o == nil || IsNil(o.NodeName) {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeAccountInfo) GetNodeNameOk() (*string, bool) {
	if o == nil || IsNil(o.NodeName) {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountInfo) HasNodeName() bool {
	if o != nil && !IsNil(o.NodeName) {
		return true
	}

	return false
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *ModelCloudNodeAccountInfo) SetNodeName(v string) {
	o.NodeName = &v
}

func (o ModelCloudNodeAccountInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCloudNodeAccountInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if !IsNil(o.CompliancePercentage) {
		toSerialize["compliance_percentage"] = o.CompliancePercentage
	}
	if !IsNil(o.LastScanId) {
		toSerialize["last_scan_id"] = o.LastScanId
	}
	if !IsNil(o.LastScanStatus) {
		toSerialize["last_scan_status"] = o.LastScanStatus
	}
	if !IsNil(o.NodeId) {
		toSerialize["node_id"] = o.NodeId
	}
	if !IsNil(o.NodeName) {
		toSerialize["node_name"] = o.NodeName
	}
	return toSerialize, nil
}

type NullableModelCloudNodeAccountInfo struct {
	value *ModelCloudNodeAccountInfo
	isSet bool
}

func (v NullableModelCloudNodeAccountInfo) Get() *ModelCloudNodeAccountInfo {
	return v.value
}

func (v *NullableModelCloudNodeAccountInfo) Set(val *ModelCloudNodeAccountInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCloudNodeAccountInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCloudNodeAccountInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCloudNodeAccountInfo(val *ModelCloudNodeAccountInfo) *NullableModelCloudNodeAccountInfo {
	return &NullableModelCloudNodeAccountInfo{value: val, isSet: true}
}

func (v NullableModelCloudNodeAccountInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCloudNodeAccountInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


