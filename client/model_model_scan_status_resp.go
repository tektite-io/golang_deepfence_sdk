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

// checks if the ModelScanStatusResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelScanStatusResp{}

// ModelScanStatusResp struct for ModelScanStatusResp
type ModelScanStatusResp struct {
	Statuses map[string]ModelScanInfo `json:"statuses"`
}

// NewModelScanStatusResp instantiates a new ModelScanStatusResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelScanStatusResp(statuses map[string]ModelScanInfo) *ModelScanStatusResp {
	this := ModelScanStatusResp{}
	this.Statuses = statuses
	return &this
}

// NewModelScanStatusRespWithDefaults instantiates a new ModelScanStatusResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelScanStatusRespWithDefaults() *ModelScanStatusResp {
	this := ModelScanStatusResp{}
	return &this
}

// GetStatuses returns the Statuses field value
// If the value is explicit nil, the zero value for map[string]ModelScanInfo will be returned
func (o *ModelScanStatusResp) GetStatuses() map[string]ModelScanInfo {
	if o == nil {
		var ret map[string]ModelScanInfo
		return ret
	}

	return o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelScanStatusResp) GetStatusesOk() (*map[string]ModelScanInfo, bool) {
	if o == nil || isNil(o.Statuses) {
		return nil, false
	}
	return &o.Statuses, true
}

// SetStatuses sets field value
func (o *ModelScanStatusResp) SetStatuses(v map[string]ModelScanInfo) {
	o.Statuses = v
}

func (o ModelScanStatusResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelScanStatusResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Statuses != nil {
		toSerialize["statuses"] = o.Statuses
	}
	return toSerialize, nil
}

type NullableModelScanStatusResp struct {
	value *ModelScanStatusResp
	isSet bool
}

func (v NullableModelScanStatusResp) Get() *ModelScanStatusResp {
	return v.value
}

func (v *NullableModelScanStatusResp) Set(val *ModelScanStatusResp) {
	v.value = val
	v.isSet = true
}

func (v NullableModelScanStatusResp) IsSet() bool {
	return v.isSet
}

func (v *NullableModelScanStatusResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelScanStatusResp(val *ModelScanStatusResp) *NullableModelScanStatusResp {
	return &NullableModelScanStatusResp{value: val, isSet: true}
}

func (v NullableModelScanStatusResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelScanStatusResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


