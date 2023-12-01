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
	"fmt"
)

// checks if the ModelComplianceScanStatusResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelComplianceScanStatusResp{}

// ModelComplianceScanStatusResp struct for ModelComplianceScanStatusResp
type ModelComplianceScanStatusResp struct {
	Statuses []ModelComplianceScanInfo `json:"statuses"`
}

type _ModelComplianceScanStatusResp ModelComplianceScanStatusResp

// NewModelComplianceScanStatusResp instantiates a new ModelComplianceScanStatusResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelComplianceScanStatusResp(statuses []ModelComplianceScanInfo) *ModelComplianceScanStatusResp {
	this := ModelComplianceScanStatusResp{}
	this.Statuses = statuses
	return &this
}

// NewModelComplianceScanStatusRespWithDefaults instantiates a new ModelComplianceScanStatusResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelComplianceScanStatusRespWithDefaults() *ModelComplianceScanStatusResp {
	this := ModelComplianceScanStatusResp{}
	return &this
}

// GetStatuses returns the Statuses field value
// If the value is explicit nil, the zero value for []ModelComplianceScanInfo will be returned
func (o *ModelComplianceScanStatusResp) GetStatuses() []ModelComplianceScanInfo {
	if o == nil {
		var ret []ModelComplianceScanInfo
		return ret
	}

	return o.Statuses
}

// GetStatusesOk returns a tuple with the Statuses field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelComplianceScanStatusResp) GetStatusesOk() ([]ModelComplianceScanInfo, bool) {
	if o == nil || IsNil(o.Statuses) {
		return nil, false
	}
	return o.Statuses, true
}

// SetStatuses sets field value
func (o *ModelComplianceScanStatusResp) SetStatuses(v []ModelComplianceScanInfo) {
	o.Statuses = v
}

func (o ModelComplianceScanStatusResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelComplianceScanStatusResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Statuses != nil {
		toSerialize["statuses"] = o.Statuses
	}
	return toSerialize, nil
}

func (o *ModelComplianceScanStatusResp) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"statuses",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varModelComplianceScanStatusResp := _ModelComplianceScanStatusResp{}

	err = json.Unmarshal(bytes, &varModelComplianceScanStatusResp)

	if err != nil {
		return err
	}

	*o = ModelComplianceScanStatusResp(varModelComplianceScanStatusResp)

	return err
}

type NullableModelComplianceScanStatusResp struct {
	value *ModelComplianceScanStatusResp
	isSet bool
}

func (v NullableModelComplianceScanStatusResp) Get() *ModelComplianceScanStatusResp {
	return v.value
}

func (v *NullableModelComplianceScanStatusResp) Set(val *ModelComplianceScanStatusResp) {
	v.value = val
	v.isSet = true
}

func (v NullableModelComplianceScanStatusResp) IsSet() bool {
	return v.isSet
}

func (v *NullableModelComplianceScanStatusResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelComplianceScanStatusResp(val *ModelComplianceScanStatusResp) *NullableModelComplianceScanStatusResp {
	return &NullableModelComplianceScanStatusResp{value: val, isSet: true}
}

func (v NullableModelComplianceScanStatusResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelComplianceScanStatusResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


