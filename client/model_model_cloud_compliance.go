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

// checks if the ModelCloudCompliance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCloudCompliance{}

// ModelCloudCompliance struct for ModelCloudCompliance
type ModelCloudCompliance struct {
	AccountId string `json:"account_id"`
	CloudProvider string `json:"cloud_provider"`
	ComplianceCheckType string `json:"compliance_check_type"`
	ControlId string `json:"control_id"`
	Count int32 `json:"count"`
	Description string `json:"description"`
	Group string `json:"group"`
	Masked bool `json:"masked"`
	NodeId string `json:"node_id"`
	NodeName string `json:"node_name"`
	Reason string `json:"reason"`
	Region string `json:"region"`
	Resource string `json:"resource"`
	Service string `json:"service"`
	Severity string `json:"severity"`
	Status string `json:"status"`
	Title string `json:"title"`
	Type string `json:"type"`
	UpdatedAt int32 `json:"updated_at"`
}

// NewModelCloudCompliance instantiates a new ModelCloudCompliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCloudCompliance(accountId string, cloudProvider string, complianceCheckType string, controlId string, count int32, description string, group string, masked bool, nodeId string, nodeName string, reason string, region string, resource string, service string, severity string, status string, title string, type_ string, updatedAt int32) *ModelCloudCompliance {
	this := ModelCloudCompliance{}
	this.AccountId = accountId
	this.CloudProvider = cloudProvider
	this.ComplianceCheckType = complianceCheckType
	this.ControlId = controlId
	this.Count = count
	this.Description = description
	this.Group = group
	this.Masked = masked
	this.NodeId = nodeId
	this.NodeName = nodeName
	this.Reason = reason
	this.Region = region
	this.Resource = resource
	this.Service = service
	this.Severity = severity
	this.Status = status
	this.Title = title
	this.Type = type_
	this.UpdatedAt = updatedAt
	return &this
}

// NewModelCloudComplianceWithDefaults instantiates a new ModelCloudCompliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCloudComplianceWithDefaults() *ModelCloudCompliance {
	this := ModelCloudCompliance{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *ModelCloudCompliance) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *ModelCloudCompliance) SetAccountId(v string) {
	o.AccountId = v
}

// GetCloudProvider returns the CloudProvider field value
func (o *ModelCloudCompliance) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *ModelCloudCompliance) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetComplianceCheckType returns the ComplianceCheckType field value
func (o *ModelCloudCompliance) GetComplianceCheckType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComplianceCheckType
}

// GetComplianceCheckTypeOk returns a tuple with the ComplianceCheckType field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetComplianceCheckTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceCheckType, true
}

// SetComplianceCheckType sets field value
func (o *ModelCloudCompliance) SetComplianceCheckType(v string) {
	o.ComplianceCheckType = v
}

// GetControlId returns the ControlId field value
func (o *ModelCloudCompliance) GetControlId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ControlId
}

// GetControlIdOk returns a tuple with the ControlId field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetControlIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ControlId, true
}

// SetControlId sets field value
func (o *ModelCloudCompliance) SetControlId(v string) {
	o.ControlId = v
}

// GetCount returns the Count field value
func (o *ModelCloudCompliance) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *ModelCloudCompliance) SetCount(v int32) {
	o.Count = v
}

// GetDescription returns the Description field value
func (o *ModelCloudCompliance) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ModelCloudCompliance) SetDescription(v string) {
	o.Description = v
}

// GetGroup returns the Group field value
func (o *ModelCloudCompliance) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *ModelCloudCompliance) SetGroup(v string) {
	o.Group = v
}

// GetMasked returns the Masked field value
func (o *ModelCloudCompliance) GetMasked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Masked
}

// GetMaskedOk returns a tuple with the Masked field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetMaskedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Masked, true
}

// SetMasked sets field value
func (o *ModelCloudCompliance) SetMasked(v bool) {
	o.Masked = v
}

// GetNodeId returns the NodeId field value
func (o *ModelCloudCompliance) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelCloudCompliance) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeName returns the NodeName field value
func (o *ModelCloudCompliance) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *ModelCloudCompliance) SetNodeName(v string) {
	o.NodeName = v
}

// GetReason returns the Reason field value
func (o *ModelCloudCompliance) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *ModelCloudCompliance) SetReason(v string) {
	o.Reason = v
}

// GetRegion returns the Region field value
func (o *ModelCloudCompliance) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *ModelCloudCompliance) SetRegion(v string) {
	o.Region = v
}

// GetResource returns the Resource field value
func (o *ModelCloudCompliance) GetResource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetResourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Resource, true
}

// SetResource sets field value
func (o *ModelCloudCompliance) SetResource(v string) {
	o.Resource = v
}

// GetService returns the Service field value
func (o *ModelCloudCompliance) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetServiceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *ModelCloudCompliance) SetService(v string) {
	o.Service = v
}

// GetSeverity returns the Severity field value
func (o *ModelCloudCompliance) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *ModelCloudCompliance) SetSeverity(v string) {
	o.Severity = v
}

// GetStatus returns the Status field value
func (o *ModelCloudCompliance) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ModelCloudCompliance) SetStatus(v string) {
	o.Status = v
}

// GetTitle returns the Title field value
func (o *ModelCloudCompliance) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ModelCloudCompliance) SetTitle(v string) {
	o.Title = v
}

// GetType returns the Type field value
func (o *ModelCloudCompliance) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ModelCloudCompliance) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ModelCloudCompliance) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelCloudCompliance) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ModelCloudCompliance) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

func (o ModelCloudCompliance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCloudCompliance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	toSerialize["cloud_provider"] = o.CloudProvider
	toSerialize["compliance_check_type"] = o.ComplianceCheckType
	toSerialize["control_id"] = o.ControlId
	toSerialize["count"] = o.Count
	toSerialize["description"] = o.Description
	toSerialize["group"] = o.Group
	toSerialize["masked"] = o.Masked
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_name"] = o.NodeName
	toSerialize["reason"] = o.Reason
	toSerialize["region"] = o.Region
	toSerialize["resource"] = o.Resource
	toSerialize["service"] = o.Service
	toSerialize["severity"] = o.Severity
	toSerialize["status"] = o.Status
	toSerialize["title"] = o.Title
	toSerialize["type"] = o.Type
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableModelCloudCompliance struct {
	value *ModelCloudCompliance
	isSet bool
}

func (v NullableModelCloudCompliance) Get() *ModelCloudCompliance {
	return v.value
}

func (v *NullableModelCloudCompliance) Set(val *ModelCloudCompliance) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCloudCompliance) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCloudCompliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCloudCompliance(val *ModelCloudCompliance) *NullableModelCloudCompliance {
	return &NullableModelCloudCompliance{value: val, isSet: true}
}

func (v NullableModelCloudCompliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCloudCompliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


