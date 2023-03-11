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

// checks if the ModelCloudNodeControlReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCloudNodeControlReq{}

// ModelCloudNodeControlReq struct for ModelCloudNodeControlReq
type ModelCloudNodeControlReq struct {
	CloudProvider string `json:"cloud_provider"`
	ComplianceType string `json:"compliance_type"`
	NodeId *string `json:"node_id,omitempty"`
}

// NewModelCloudNodeControlReq instantiates a new ModelCloudNodeControlReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCloudNodeControlReq(cloudProvider string, complianceType string) *ModelCloudNodeControlReq {
	this := ModelCloudNodeControlReq{}
	this.CloudProvider = cloudProvider
	this.ComplianceType = complianceType
	return &this
}

// NewModelCloudNodeControlReqWithDefaults instantiates a new ModelCloudNodeControlReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCloudNodeControlReqWithDefaults() *ModelCloudNodeControlReq {
	this := ModelCloudNodeControlReq{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value
func (o *ModelCloudNodeControlReq) GetCloudProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeControlReq) GetCloudProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudProvider, true
}

// SetCloudProvider sets field value
func (o *ModelCloudNodeControlReq) SetCloudProvider(v string) {
	o.CloudProvider = v
}

// GetComplianceType returns the ComplianceType field value
func (o *ModelCloudNodeControlReq) GetComplianceType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ComplianceType
}

// GetComplianceTypeOk returns a tuple with the ComplianceType field value
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeControlReq) GetComplianceTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceType, true
}

// SetComplianceType sets field value
func (o *ModelCloudNodeControlReq) SetComplianceType(v string) {
	o.ComplianceType = v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *ModelCloudNodeControlReq) GetNodeId() string {
	if o == nil || IsNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeControlReq) GetNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *ModelCloudNodeControlReq) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *ModelCloudNodeControlReq) SetNodeId(v string) {
	o.NodeId = &v
}

func (o ModelCloudNodeControlReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCloudNodeControlReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloud_provider"] = o.CloudProvider
	toSerialize["compliance_type"] = o.ComplianceType
	if !IsNil(o.NodeId) {
		toSerialize["node_id"] = o.NodeId
	}
	return toSerialize, nil
}

type NullableModelCloudNodeControlReq struct {
	value *ModelCloudNodeControlReq
	isSet bool
}

func (v NullableModelCloudNodeControlReq) Get() *ModelCloudNodeControlReq {
	return v.value
}

func (v *NullableModelCloudNodeControlReq) Set(val *ModelCloudNodeControlReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCloudNodeControlReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCloudNodeControlReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCloudNodeControlReq(val *ModelCloudNodeControlReq) *NullableModelCloudNodeControlReq {
	return &NullableModelCloudNodeControlReq{value: val, isSet: true}
}

func (v NullableModelCloudNodeControlReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCloudNodeControlReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


