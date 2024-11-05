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
)

// checks if the ModelCloudNodeAccountInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCloudNodeAccountInfo{}

// ModelCloudNodeAccountInfo struct for ModelCloudNodeAccountInfo
type ModelCloudNodeAccountInfo struct {
	AccountName *string `json:"account_name,omitempty"`
	Active *bool `json:"active,omitempty"`
	CloudProvider *string `json:"cloud_provider,omitempty"`
	CompliancePercentage *float32 `json:"compliance_percentage,omitempty"`
	HostNodeId *string `json:"host_node_id,omitempty"`
	LastScanId *string `json:"last_scan_id,omitempty"`
	LastScanStatus *string `json:"last_scan_status,omitempty"`
	NodeId *string `json:"node_id,omitempty"`
	NodeName *string `json:"node_name,omitempty"`
	RefreshMessage *string `json:"refresh_message,omitempty"`
	RefreshMetadata *string `json:"refresh_metadata,omitempty"`
	RefreshStatus *string `json:"refresh_status,omitempty"`
	RefreshStatusMap map[string]int32 `json:"refresh_status_map,omitempty"`
	ScanStatusMap map[string]int32 `json:"scan_status_map,omitempty"`
	Version *string `json:"version,omitempty"`
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

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *ModelCloudNodeAccountInfo) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeAccountInfo) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountInfo) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *ModelCloudNodeAccountInfo) SetAccountName(v string) {
	o.AccountName = &v
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

// GetHostNodeId returns the HostNodeId field value if set, zero value otherwise.
func (o *ModelCloudNodeAccountInfo) GetHostNodeId() string {
	if o == nil || IsNil(o.HostNodeId) {
		var ret string
		return ret
	}
	return *o.HostNodeId
}

// GetHostNodeIdOk returns a tuple with the HostNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeAccountInfo) GetHostNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.HostNodeId) {
		return nil, false
	}
	return o.HostNodeId, true
}

// HasHostNodeId returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountInfo) HasHostNodeId() bool {
	if o != nil && !IsNil(o.HostNodeId) {
		return true
	}

	return false
}

// SetHostNodeId gets a reference to the given string and assigns it to the HostNodeId field.
func (o *ModelCloudNodeAccountInfo) SetHostNodeId(v string) {
	o.HostNodeId = &v
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

// GetRefreshMessage returns the RefreshMessage field value if set, zero value otherwise.
func (o *ModelCloudNodeAccountInfo) GetRefreshMessage() string {
	if o == nil || IsNil(o.RefreshMessage) {
		var ret string
		return ret
	}
	return *o.RefreshMessage
}

// GetRefreshMessageOk returns a tuple with the RefreshMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeAccountInfo) GetRefreshMessageOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshMessage) {
		return nil, false
	}
	return o.RefreshMessage, true
}

// HasRefreshMessage returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountInfo) HasRefreshMessage() bool {
	if o != nil && !IsNil(o.RefreshMessage) {
		return true
	}

	return false
}

// SetRefreshMessage gets a reference to the given string and assigns it to the RefreshMessage field.
func (o *ModelCloudNodeAccountInfo) SetRefreshMessage(v string) {
	o.RefreshMessage = &v
}

// GetRefreshMetadata returns the RefreshMetadata field value if set, zero value otherwise.
func (o *ModelCloudNodeAccountInfo) GetRefreshMetadata() string {
	if o == nil || IsNil(o.RefreshMetadata) {
		var ret string
		return ret
	}
	return *o.RefreshMetadata
}

// GetRefreshMetadataOk returns a tuple with the RefreshMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeAccountInfo) GetRefreshMetadataOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshMetadata) {
		return nil, false
	}
	return o.RefreshMetadata, true
}

// HasRefreshMetadata returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountInfo) HasRefreshMetadata() bool {
	if o != nil && !IsNil(o.RefreshMetadata) {
		return true
	}

	return false
}

// SetRefreshMetadata gets a reference to the given string and assigns it to the RefreshMetadata field.
func (o *ModelCloudNodeAccountInfo) SetRefreshMetadata(v string) {
	o.RefreshMetadata = &v
}

// GetRefreshStatus returns the RefreshStatus field value if set, zero value otherwise.
func (o *ModelCloudNodeAccountInfo) GetRefreshStatus() string {
	if o == nil || IsNil(o.RefreshStatus) {
		var ret string
		return ret
	}
	return *o.RefreshStatus
}

// GetRefreshStatusOk returns a tuple with the RefreshStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeAccountInfo) GetRefreshStatusOk() (*string, bool) {
	if o == nil || IsNil(o.RefreshStatus) {
		return nil, false
	}
	return o.RefreshStatus, true
}

// HasRefreshStatus returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountInfo) HasRefreshStatus() bool {
	if o != nil && !IsNil(o.RefreshStatus) {
		return true
	}

	return false
}

// SetRefreshStatus gets a reference to the given string and assigns it to the RefreshStatus field.
func (o *ModelCloudNodeAccountInfo) SetRefreshStatus(v string) {
	o.RefreshStatus = &v
}

// GetRefreshStatusMap returns the RefreshStatusMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCloudNodeAccountInfo) GetRefreshStatusMap() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}
	return o.RefreshStatusMap
}

// GetRefreshStatusMapOk returns a tuple with the RefreshStatusMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCloudNodeAccountInfo) GetRefreshStatusMapOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.RefreshStatusMap) {
		return nil, false
	}
	return &o.RefreshStatusMap, true
}

// HasRefreshStatusMap returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountInfo) HasRefreshStatusMap() bool {
	if o != nil && !IsNil(o.RefreshStatusMap) {
		return true
	}

	return false
}

// SetRefreshStatusMap gets a reference to the given map[string]int32 and assigns it to the RefreshStatusMap field.
func (o *ModelCloudNodeAccountInfo) SetRefreshStatusMap(v map[string]int32) {
	o.RefreshStatusMap = v
}

// GetScanStatusMap returns the ScanStatusMap field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCloudNodeAccountInfo) GetScanStatusMap() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}
	return o.ScanStatusMap
}

// GetScanStatusMapOk returns a tuple with the ScanStatusMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCloudNodeAccountInfo) GetScanStatusMapOk() (*map[string]int32, bool) {
	if o == nil || IsNil(o.ScanStatusMap) {
		return nil, false
	}
	return &o.ScanStatusMap, true
}

// HasScanStatusMap returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountInfo) HasScanStatusMap() bool {
	if o != nil && !IsNil(o.ScanStatusMap) {
		return true
	}

	return false
}

// SetScanStatusMap gets a reference to the given map[string]int32 and assigns it to the ScanStatusMap field.
func (o *ModelCloudNodeAccountInfo) SetScanStatusMap(v map[string]int32) {
	o.ScanStatusMap = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ModelCloudNodeAccountInfo) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeAccountInfo) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ModelCloudNodeAccountInfo) SetVersion(v string) {
	o.Version = &v
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
	if !IsNil(o.AccountName) {
		toSerialize["account_name"] = o.AccountName
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if !IsNil(o.CompliancePercentage) {
		toSerialize["compliance_percentage"] = o.CompliancePercentage
	}
	if !IsNil(o.HostNodeId) {
		toSerialize["host_node_id"] = o.HostNodeId
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
	if !IsNil(o.RefreshMessage) {
		toSerialize["refresh_message"] = o.RefreshMessage
	}
	if !IsNil(o.RefreshMetadata) {
		toSerialize["refresh_metadata"] = o.RefreshMetadata
	}
	if !IsNil(o.RefreshStatus) {
		toSerialize["refresh_status"] = o.RefreshStatus
	}
	if o.RefreshStatusMap != nil {
		toSerialize["refresh_status_map"] = o.RefreshStatusMap
	}
	if o.ScanStatusMap != nil {
		toSerialize["scan_status_map"] = o.ScanStatusMap
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
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


