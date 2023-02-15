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

// checks if the ModelKeyValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelKeyValues{}

// ModelKeyValues struct for ModelKeyValues
type ModelKeyValues struct {
	Key string `json:"key"`
	Values []string `json:"values"`
}

// NewModelKeyValues instantiates a new ModelKeyValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelKeyValues(key string, values []string) *ModelKeyValues {
	this := ModelKeyValues{}
	this.Key = key
	this.Values = values
	return &this
}

// NewModelKeyValuesWithDefaults instantiates a new ModelKeyValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelKeyValuesWithDefaults() *ModelKeyValues {
	this := ModelKeyValues{}
	return &this
}

// GetKey returns the Key field value
func (o *ModelKeyValues) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ModelKeyValues) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ModelKeyValues) SetKey(v string) {
	o.Key = v
}

// GetValues returns the Values field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelKeyValues) GetValues() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelKeyValues) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// SetValues sets field value
func (o *ModelKeyValues) SetValues(v []string) {
	o.Values = v
}

func (o ModelKeyValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelKeyValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableModelKeyValues struct {
	value *ModelKeyValues
	isSet bool
}

func (v NullableModelKeyValues) Get() *ModelKeyValues {
	return v.value
}

func (v *NullableModelKeyValues) Set(val *ModelKeyValues) {
	v.value = val
	v.isSet = true
}

func (v NullableModelKeyValues) IsSet() bool {
	return v.isSet
}

func (v *NullableModelKeyValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelKeyValues(val *ModelKeyValues) *NullableModelKeyValues {
	return &NullableModelKeyValues{value: val, isSet: true}
}

func (v NullableModelKeyValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelKeyValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

