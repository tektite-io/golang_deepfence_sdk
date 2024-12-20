/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.2
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelDeleteIntegrationReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelDeleteIntegrationReq{}

// ModelDeleteIntegrationReq struct for ModelDeleteIntegrationReq
type ModelDeleteIntegrationReq struct {
	IntegrationIds []int32 `json:"integration_ids"`
}

type _ModelDeleteIntegrationReq ModelDeleteIntegrationReq

// NewModelDeleteIntegrationReq instantiates a new ModelDeleteIntegrationReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelDeleteIntegrationReq(integrationIds []int32) *ModelDeleteIntegrationReq {
	this := ModelDeleteIntegrationReq{}
	this.IntegrationIds = integrationIds
	return &this
}

// NewModelDeleteIntegrationReqWithDefaults instantiates a new ModelDeleteIntegrationReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelDeleteIntegrationReqWithDefaults() *ModelDeleteIntegrationReq {
	this := ModelDeleteIntegrationReq{}
	return &this
}

// GetIntegrationIds returns the IntegrationIds field value
// If the value is explicit nil, the zero value for []int32 will be returned
func (o *ModelDeleteIntegrationReq) GetIntegrationIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.IntegrationIds
}

// GetIntegrationIdsOk returns a tuple with the IntegrationIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelDeleteIntegrationReq) GetIntegrationIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.IntegrationIds) {
		return nil, false
	}
	return o.IntegrationIds, true
}

// SetIntegrationIds sets field value
func (o *ModelDeleteIntegrationReq) SetIntegrationIds(v []int32) {
	o.IntegrationIds = v
}

func (o ModelDeleteIntegrationReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelDeleteIntegrationReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IntegrationIds != nil {
		toSerialize["integration_ids"] = o.IntegrationIds
	}
	return toSerialize, nil
}

func (o *ModelDeleteIntegrationReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"integration_ids",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varModelDeleteIntegrationReq := _ModelDeleteIntegrationReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelDeleteIntegrationReq)

	if err != nil {
		return err
	}

	*o = ModelDeleteIntegrationReq(varModelDeleteIntegrationReq)

	return err
}

type NullableModelDeleteIntegrationReq struct {
	value *ModelDeleteIntegrationReq
	isSet bool
}

func (v NullableModelDeleteIntegrationReq) Get() *ModelDeleteIntegrationReq {
	return v.value
}

func (v *NullableModelDeleteIntegrationReq) Set(val *ModelDeleteIntegrationReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelDeleteIntegrationReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelDeleteIntegrationReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelDeleteIntegrationReq(val *ModelDeleteIntegrationReq) *NullableModelDeleteIntegrationReq {
	return &NullableModelDeleteIntegrationReq{value: val, isSet: true}
}

func (v NullableModelDeleteIntegrationReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelDeleteIntegrationReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


