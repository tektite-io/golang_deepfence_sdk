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

// checks if the ModelCloudNodeEnableDisableReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCloudNodeEnableDisableReq{}

// ModelCloudNodeEnableDisableReq struct for ModelCloudNodeEnableDisableReq
type ModelCloudNodeEnableDisableReq struct {
	ControlIds []string `json:"control_ids,omitempty"`
	NodeId *string `json:"node_id,omitempty"`
}

// NewModelCloudNodeEnableDisableReq instantiates a new ModelCloudNodeEnableDisableReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCloudNodeEnableDisableReq() *ModelCloudNodeEnableDisableReq {
	this := ModelCloudNodeEnableDisableReq{}
	return &this
}

// NewModelCloudNodeEnableDisableReqWithDefaults instantiates a new ModelCloudNodeEnableDisableReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCloudNodeEnableDisableReqWithDefaults() *ModelCloudNodeEnableDisableReq {
	this := ModelCloudNodeEnableDisableReq{}
	return &this
}

// GetControlIds returns the ControlIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCloudNodeEnableDisableReq) GetControlIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ControlIds
}

// GetControlIdsOk returns a tuple with the ControlIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCloudNodeEnableDisableReq) GetControlIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ControlIds) {
		return nil, false
	}
	return o.ControlIds, true
}

// HasControlIds returns a boolean if a field has been set.
func (o *ModelCloudNodeEnableDisableReq) HasControlIds() bool {
	if o != nil && !IsNil(o.ControlIds) {
		return true
	}

	return false
}

// SetControlIds gets a reference to the given []string and assigns it to the ControlIds field.
func (o *ModelCloudNodeEnableDisableReq) SetControlIds(v []string) {
	o.ControlIds = v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *ModelCloudNodeEnableDisableReq) GetNodeId() string {
	if o == nil || IsNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeEnableDisableReq) GetNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *ModelCloudNodeEnableDisableReq) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *ModelCloudNodeEnableDisableReq) SetNodeId(v string) {
	o.NodeId = &v
}

func (o ModelCloudNodeEnableDisableReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCloudNodeEnableDisableReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ControlIds != nil {
		toSerialize["control_ids"] = o.ControlIds
	}
	if !IsNil(o.NodeId) {
		toSerialize["node_id"] = o.NodeId
	}
	return toSerialize, nil
}

type NullableModelCloudNodeEnableDisableReq struct {
	value *ModelCloudNodeEnableDisableReq
	isSet bool
}

func (v NullableModelCloudNodeEnableDisableReq) Get() *ModelCloudNodeEnableDisableReq {
	return v.value
}

func (v *NullableModelCloudNodeEnableDisableReq) Set(val *ModelCloudNodeEnableDisableReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCloudNodeEnableDisableReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCloudNodeEnableDisableReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCloudNodeEnableDisableReq(val *ModelCloudNodeEnableDisableReq) *NullableModelCloudNodeEnableDisableReq {
	return &NullableModelCloudNodeEnableDisableReq{value: val, isSet: true}
}

func (v NullableModelCloudNodeEnableDisableReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCloudNodeEnableDisableReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


