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

// checks if the ModelScanListResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelScanListResp{}

// ModelScanListResp struct for ModelScanListResp
type ModelScanListResp struct {
	ScansInfo []ModelScanInfo `json:"scans_info"`
}

// NewModelScanListResp instantiates a new ModelScanListResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelScanListResp(scansInfo []ModelScanInfo) *ModelScanListResp {
	this := ModelScanListResp{}
	this.ScansInfo = scansInfo
	return &this
}

// NewModelScanListRespWithDefaults instantiates a new ModelScanListResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelScanListRespWithDefaults() *ModelScanListResp {
	this := ModelScanListResp{}
	return &this
}

// GetScansInfo returns the ScansInfo field value
// If the value is explicit nil, the zero value for []ModelScanInfo will be returned
func (o *ModelScanListResp) GetScansInfo() []ModelScanInfo {
	if o == nil {
		var ret []ModelScanInfo
		return ret
	}

	return o.ScansInfo
}

// GetScansInfoOk returns a tuple with the ScansInfo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelScanListResp) GetScansInfoOk() ([]ModelScanInfo, bool) {
	if o == nil || IsNil(o.ScansInfo) {
		return nil, false
	}
	return o.ScansInfo, true
}

// SetScansInfo sets field value
func (o *ModelScanListResp) SetScansInfo(v []ModelScanInfo) {
	o.ScansInfo = v
}

func (o ModelScanListResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelScanListResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ScansInfo != nil {
		toSerialize["scans_info"] = o.ScansInfo
	}
	return toSerialize, nil
}

type NullableModelScanListResp struct {
	value *ModelScanListResp
	isSet bool
}

func (v NullableModelScanListResp) Get() *ModelScanListResp {
	return v.value
}

func (v *NullableModelScanListResp) Set(val *ModelScanListResp) {
	v.value = val
	v.isSet = true
}

func (v NullableModelScanListResp) IsSet() bool {
	return v.isSet
}

func (v *NullableModelScanListResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelScanListResp(val *ModelScanListResp) *NullableModelScanListResp {
	return &NullableModelScanListResp{value: val, isSet: true}
}

func (v NullableModelScanListResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelScanListResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


