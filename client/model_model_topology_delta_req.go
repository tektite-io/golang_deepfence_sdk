/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelTopologyDeltaReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelTopologyDeltaReq{}

// ModelTopologyDeltaReq struct for ModelTopologyDeltaReq
type ModelTopologyDeltaReq struct {
	Addition bool `json:"addition"`
	AdditionTimestamp int64 `json:"addition_timestamp"`
	Deletion bool `json:"deletion"`
	DeletionTimestamp int64 `json:"deletion_timestamp"`
	EntityTypes []string `json:"entity_types"`
}

type _ModelTopologyDeltaReq ModelTopologyDeltaReq

// NewModelTopologyDeltaReq instantiates a new ModelTopologyDeltaReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelTopologyDeltaReq(addition bool, additionTimestamp int64, deletion bool, deletionTimestamp int64, entityTypes []string) *ModelTopologyDeltaReq {
	this := ModelTopologyDeltaReq{}
	this.Addition = addition
	this.AdditionTimestamp = additionTimestamp
	this.Deletion = deletion
	this.DeletionTimestamp = deletionTimestamp
	this.EntityTypes = entityTypes
	return &this
}

// NewModelTopologyDeltaReqWithDefaults instantiates a new ModelTopologyDeltaReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelTopologyDeltaReqWithDefaults() *ModelTopologyDeltaReq {
	this := ModelTopologyDeltaReq{}
	return &this
}

// GetAddition returns the Addition field value
func (o *ModelTopologyDeltaReq) GetAddition() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Addition
}

// GetAdditionOk returns a tuple with the Addition field value
// and a boolean to check if the value has been set.
func (o *ModelTopologyDeltaReq) GetAdditionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Addition, true
}

// SetAddition sets field value
func (o *ModelTopologyDeltaReq) SetAddition(v bool) {
	o.Addition = v
}

// GetAdditionTimestamp returns the AdditionTimestamp field value
func (o *ModelTopologyDeltaReq) GetAdditionTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AdditionTimestamp
}

// GetAdditionTimestampOk returns a tuple with the AdditionTimestamp field value
// and a boolean to check if the value has been set.
func (o *ModelTopologyDeltaReq) GetAdditionTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdditionTimestamp, true
}

// SetAdditionTimestamp sets field value
func (o *ModelTopologyDeltaReq) SetAdditionTimestamp(v int64) {
	o.AdditionTimestamp = v
}

// GetDeletion returns the Deletion field value
func (o *ModelTopologyDeltaReq) GetDeletion() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deletion
}

// GetDeletionOk returns a tuple with the Deletion field value
// and a boolean to check if the value has been set.
func (o *ModelTopologyDeltaReq) GetDeletionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deletion, true
}

// SetDeletion sets field value
func (o *ModelTopologyDeltaReq) SetDeletion(v bool) {
	o.Deletion = v
}

// GetDeletionTimestamp returns the DeletionTimestamp field value
func (o *ModelTopologyDeltaReq) GetDeletionTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DeletionTimestamp
}

// GetDeletionTimestampOk returns a tuple with the DeletionTimestamp field value
// and a boolean to check if the value has been set.
func (o *ModelTopologyDeltaReq) GetDeletionTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletionTimestamp, true
}

// SetDeletionTimestamp sets field value
func (o *ModelTopologyDeltaReq) SetDeletionTimestamp(v int64) {
	o.DeletionTimestamp = v
}

// GetEntityTypes returns the EntityTypes field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelTopologyDeltaReq) GetEntityTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EntityTypes
}

// GetEntityTypesOk returns a tuple with the EntityTypes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelTopologyDeltaReq) GetEntityTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.EntityTypes) {
		return nil, false
	}
	return o.EntityTypes, true
}

// SetEntityTypes sets field value
func (o *ModelTopologyDeltaReq) SetEntityTypes(v []string) {
	o.EntityTypes = v
}

func (o ModelTopologyDeltaReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelTopologyDeltaReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addition"] = o.Addition
	toSerialize["addition_timestamp"] = o.AdditionTimestamp
	toSerialize["deletion"] = o.Deletion
	toSerialize["deletion_timestamp"] = o.DeletionTimestamp
	if o.EntityTypes != nil {
		toSerialize["entity_types"] = o.EntityTypes
	}
	return toSerialize, nil
}

func (o *ModelTopologyDeltaReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"addition",
		"addition_timestamp",
		"deletion",
		"deletion_timestamp",
		"entity_types",
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

	varModelTopologyDeltaReq := _ModelTopologyDeltaReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelTopologyDeltaReq)

	if err != nil {
		return err
	}

	*o = ModelTopologyDeltaReq(varModelTopologyDeltaReq)

	return err
}

type NullableModelTopologyDeltaReq struct {
	value *ModelTopologyDeltaReq
	isSet bool
}

func (v NullableModelTopologyDeltaReq) Get() *ModelTopologyDeltaReq {
	return v.value
}

func (v *NullableModelTopologyDeltaReq) Set(val *ModelTopologyDeltaReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelTopologyDeltaReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelTopologyDeltaReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelTopologyDeltaReq(val *ModelTopologyDeltaReq) *NullableModelTopologyDeltaReq {
	return &NullableModelTopologyDeltaReq{value: val, isSet: true}
}

func (v NullableModelTopologyDeltaReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelTopologyDeltaReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


