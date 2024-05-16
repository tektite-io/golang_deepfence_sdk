/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.2.1
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ModelCloudNodeControlResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCloudNodeControlResp{}

// ModelCloudNodeControlResp struct for ModelCloudNodeControlResp
type ModelCloudNodeControlResp struct {
	Controls []ModelCloudNodeComplianceControl `json:"controls,omitempty"`
}

// NewModelCloudNodeControlResp instantiates a new ModelCloudNodeControlResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCloudNodeControlResp() *ModelCloudNodeControlResp {
	this := ModelCloudNodeControlResp{}
	return &this
}

// NewModelCloudNodeControlRespWithDefaults instantiates a new ModelCloudNodeControlResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCloudNodeControlRespWithDefaults() *ModelCloudNodeControlResp {
	this := ModelCloudNodeControlResp{}
	return &this
}

// GetControls returns the Controls field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCloudNodeControlResp) GetControls() []ModelCloudNodeComplianceControl {
	if o == nil {
		var ret []ModelCloudNodeComplianceControl
		return ret
	}
	return o.Controls
}

// GetControlsOk returns a tuple with the Controls field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCloudNodeControlResp) GetControlsOk() ([]ModelCloudNodeComplianceControl, bool) {
	if o == nil || IsNil(o.Controls) {
		return nil, false
	}
	return o.Controls, true
}

// HasControls returns a boolean if a field has been set.
func (o *ModelCloudNodeControlResp) HasControls() bool {
	if o != nil && !IsNil(o.Controls) {
		return true
	}

	return false
}

// SetControls gets a reference to the given []ModelCloudNodeComplianceControl and assigns it to the Controls field.
func (o *ModelCloudNodeControlResp) SetControls(v []ModelCloudNodeComplianceControl) {
	o.Controls = v
}

func (o ModelCloudNodeControlResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCloudNodeControlResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Controls != nil {
		toSerialize["controls"] = o.Controls
	}
	return toSerialize, nil
}

type NullableModelCloudNodeControlResp struct {
	value *ModelCloudNodeControlResp
	isSet bool
}

func (v NullableModelCloudNodeControlResp) Get() *ModelCloudNodeControlResp {
	return v.value
}

func (v *NullableModelCloudNodeControlResp) Set(val *ModelCloudNodeControlResp) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCloudNodeControlResp) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCloudNodeControlResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCloudNodeControlResp(val *ModelCloudNodeControlResp) *NullableModelCloudNodeControlResp {
	return &NullableModelCloudNodeControlResp{value: val, isSet: true}
}

func (v NullableModelCloudNodeControlResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCloudNodeControlResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


