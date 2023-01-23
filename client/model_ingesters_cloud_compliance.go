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

// checks if the IngestersCloudCompliance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestersCloudCompliance{}

// IngestersCloudCompliance struct for IngestersCloudCompliance
type IngestersCloudCompliance struct {
	Timestamp *string `json:"@timestamp,omitempty"`
	AccountId *string `json:"account_id,omitempty"`
	CloudProvider *string `json:"cloud_provider,omitempty"`
	ComplianceCheckType *string `json:"compliance_check_type,omitempty"`
	ControlId *string `json:"control_id,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Description *string `json:"description,omitempty"`
	DocId *string `json:"doc_id,omitempty"`
	Group *string `json:"group,omitempty"`
	Masked *string `json:"masked,omitempty"`
	NodeId *string `json:"node_id,omitempty"`
	NodeName *string `json:"node_name,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Region *string `json:"region,omitempty"`
	Resource *string `json:"resource,omitempty"`
	ScanId *string `json:"scan_id,omitempty"`
	Service *string `json:"service,omitempty"`
	Severity *string `json:"severity,omitempty"`
	Status *string `json:"status,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewIngestersCloudCompliance instantiates a new IngestersCloudCompliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestersCloudCompliance() *IngestersCloudCompliance {
	this := IngestersCloudCompliance{}
	return &this
}

// NewIngestersCloudComplianceWithDefaults instantiates a new IngestersCloudCompliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestersCloudComplianceWithDefaults() *IngestersCloudCompliance {
	this := IngestersCloudCompliance{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetTimestamp() string {
	if o == nil || isNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetTimestampOk() (*string, bool) {
	if o == nil || isNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *IngestersCloudCompliance) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetAccountId() string {
	if o == nil || isNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetAccountIdOk() (*string, bool) {
	if o == nil || isNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasAccountId() bool {
	if o != nil && !isNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *IngestersCloudCompliance) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetCloudProvider() string {
	if o == nil || isNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetCloudProviderOk() (*string, bool) {
	if o == nil || isNil(o.CloudProvider) {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasCloudProvider() bool {
	if o != nil && !isNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *IngestersCloudCompliance) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetComplianceCheckType returns the ComplianceCheckType field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetComplianceCheckType() string {
	if o == nil || isNil(o.ComplianceCheckType) {
		var ret string
		return ret
	}
	return *o.ComplianceCheckType
}

// GetComplianceCheckTypeOk returns a tuple with the ComplianceCheckType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetComplianceCheckTypeOk() (*string, bool) {
	if o == nil || isNil(o.ComplianceCheckType) {
		return nil, false
	}
	return o.ComplianceCheckType, true
}

// HasComplianceCheckType returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasComplianceCheckType() bool {
	if o != nil && !isNil(o.ComplianceCheckType) {
		return true
	}

	return false
}

// SetComplianceCheckType gets a reference to the given string and assigns it to the ComplianceCheckType field.
func (o *IngestersCloudCompliance) SetComplianceCheckType(v string) {
	o.ComplianceCheckType = &v
}

// GetControlId returns the ControlId field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetControlId() string {
	if o == nil || isNil(o.ControlId) {
		var ret string
		return ret
	}
	return *o.ControlId
}

// GetControlIdOk returns a tuple with the ControlId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetControlIdOk() (*string, bool) {
	if o == nil || isNil(o.ControlId) {
		return nil, false
	}
	return o.ControlId, true
}

// HasControlId returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasControlId() bool {
	if o != nil && !isNil(o.ControlId) {
		return true
	}

	return false
}

// SetControlId gets a reference to the given string and assigns it to the ControlId field.
func (o *IngestersCloudCompliance) SetControlId(v string) {
	o.ControlId = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetCount() int32 {
	if o == nil || isNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetCountOk() (*int32, bool) {
	if o == nil || isNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasCount() bool {
	if o != nil && !isNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *IngestersCloudCompliance) SetCount(v int32) {
	o.Count = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IngestersCloudCompliance) SetDescription(v string) {
	o.Description = &v
}

// GetDocId returns the DocId field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetDocId() string {
	if o == nil || isNil(o.DocId) {
		var ret string
		return ret
	}
	return *o.DocId
}

// GetDocIdOk returns a tuple with the DocId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetDocIdOk() (*string, bool) {
	if o == nil || isNil(o.DocId) {
		return nil, false
	}
	return o.DocId, true
}

// HasDocId returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasDocId() bool {
	if o != nil && !isNil(o.DocId) {
		return true
	}

	return false
}

// SetDocId gets a reference to the given string and assigns it to the DocId field.
func (o *IngestersCloudCompliance) SetDocId(v string) {
	o.DocId = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetGroup() string {
	if o == nil || isNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetGroupOk() (*string, bool) {
	if o == nil || isNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasGroup() bool {
	if o != nil && !isNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *IngestersCloudCompliance) SetGroup(v string) {
	o.Group = &v
}

// GetMasked returns the Masked field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetMasked() string {
	if o == nil || isNil(o.Masked) {
		var ret string
		return ret
	}
	return *o.Masked
}

// GetMaskedOk returns a tuple with the Masked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetMaskedOk() (*string, bool) {
	if o == nil || isNil(o.Masked) {
		return nil, false
	}
	return o.Masked, true
}

// HasMasked returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasMasked() bool {
	if o != nil && !isNil(o.Masked) {
		return true
	}

	return false
}

// SetMasked gets a reference to the given string and assigns it to the Masked field.
func (o *IngestersCloudCompliance) SetMasked(v string) {
	o.Masked = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetNodeId() string {
	if o == nil || isNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetNodeIdOk() (*string, bool) {
	if o == nil || isNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasNodeId() bool {
	if o != nil && !isNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *IngestersCloudCompliance) SetNodeId(v string) {
	o.NodeId = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetNodeName() string {
	if o == nil || isNil(o.NodeName) {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetNodeNameOk() (*string, bool) {
	if o == nil || isNil(o.NodeName) {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasNodeName() bool {
	if o != nil && !isNil(o.NodeName) {
		return true
	}

	return false
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *IngestersCloudCompliance) SetNodeName(v string) {
	o.NodeName = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *IngestersCloudCompliance) SetReason(v string) {
	o.Reason = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetRegion() string {
	if o == nil || isNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetRegionOk() (*string, bool) {
	if o == nil || isNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasRegion() bool {
	if o != nil && !isNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *IngestersCloudCompliance) SetRegion(v string) {
	o.Region = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetResource() string {
	if o == nil || isNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetResourceOk() (*string, bool) {
	if o == nil || isNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasResource() bool {
	if o != nil && !isNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *IngestersCloudCompliance) SetResource(v string) {
	o.Resource = &v
}

// GetScanId returns the ScanId field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetScanId() string {
	if o == nil || isNil(o.ScanId) {
		var ret string
		return ret
	}
	return *o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetScanIdOk() (*string, bool) {
	if o == nil || isNil(o.ScanId) {
		return nil, false
	}
	return o.ScanId, true
}

// HasScanId returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasScanId() bool {
	if o != nil && !isNil(o.ScanId) {
		return true
	}

	return false
}

// SetScanId gets a reference to the given string and assigns it to the ScanId field.
func (o *IngestersCloudCompliance) SetScanId(v string) {
	o.ScanId = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetService() string {
	if o == nil || isNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetServiceOk() (*string, bool) {
	if o == nil || isNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasService() bool {
	if o != nil && !isNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *IngestersCloudCompliance) SetService(v string) {
	o.Service = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetSeverity() string {
	if o == nil || isNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetSeverityOk() (*string, bool) {
	if o == nil || isNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasSeverity() bool {
	if o != nil && !isNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *IngestersCloudCompliance) SetSeverity(v string) {
	o.Severity = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IngestersCloudCompliance) SetStatus(v string) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *IngestersCloudCompliance) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IngestersCloudCompliance) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudCompliance) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IngestersCloudCompliance) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IngestersCloudCompliance) SetType(v string) {
	o.Type = &v
}

func (o IngestersCloudCompliance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestersCloudCompliance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Timestamp) {
		toSerialize["@timestamp"] = o.Timestamp
	}
	if !isNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !isNil(o.CloudProvider) {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	if !isNil(o.ComplianceCheckType) {
		toSerialize["compliance_check_type"] = o.ComplianceCheckType
	}
	if !isNil(o.ControlId) {
		toSerialize["control_id"] = o.ControlId
	}
	if !isNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.DocId) {
		toSerialize["doc_id"] = o.DocId
	}
	if !isNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !isNil(o.Masked) {
		toSerialize["masked"] = o.Masked
	}
	if !isNil(o.NodeId) {
		toSerialize["node_id"] = o.NodeId
	}
	if !isNil(o.NodeName) {
		toSerialize["node_name"] = o.NodeName
	}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !isNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !isNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !isNil(o.ScanId) {
		toSerialize["scan_id"] = o.ScanId
	}
	if !isNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !isNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableIngestersCloudCompliance struct {
	value *IngestersCloudCompliance
	isSet bool
}

func (v NullableIngestersCloudCompliance) Get() *IngestersCloudCompliance {
	return v.value
}

func (v *NullableIngestersCloudCompliance) Set(val *IngestersCloudCompliance) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestersCloudCompliance) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestersCloudCompliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestersCloudCompliance(val *IngestersCloudCompliance) *NullableIngestersCloudCompliance {
	return &NullableIngestersCloudCompliance{value: val, isSet: true}
}

func (v NullableIngestersCloudCompliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestersCloudCompliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


