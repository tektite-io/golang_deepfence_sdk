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
	"bytes"
	"fmt"
)

// checks if the ControlsAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControlsAction{}

// ControlsAction struct for ControlsAction
type ControlsAction struct {
	Id int32 `json:"id"`
	RequestPayload string `json:"request_payload"`
}

type _ControlsAction ControlsAction

// NewControlsAction instantiates a new ControlsAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControlsAction(id int32, requestPayload string) *ControlsAction {
	this := ControlsAction{}
	this.Id = id
	this.RequestPayload = requestPayload
	return &this
}

// NewControlsActionWithDefaults instantiates a new ControlsAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControlsActionWithDefaults() *ControlsAction {
	this := ControlsAction{}
	return &this
}

// GetId returns the Id field value
func (o *ControlsAction) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ControlsAction) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ControlsAction) SetId(v int32) {
	o.Id = v
}

// GetRequestPayload returns the RequestPayload field value
func (o *ControlsAction) GetRequestPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestPayload
}

// GetRequestPayloadOk returns a tuple with the RequestPayload field value
// and a boolean to check if the value has been set.
func (o *ControlsAction) GetRequestPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestPayload, true
}

// SetRequestPayload sets field value
func (o *ControlsAction) SetRequestPayload(v string) {
	o.RequestPayload = v
}

func (o ControlsAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControlsAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["request_payload"] = o.RequestPayload
	return toSerialize, nil
}

func (o *ControlsAction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"request_payload",
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

	varControlsAction := _ControlsAction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varControlsAction)

	if err != nil {
		return err
	}

	*o = ControlsAction(varControlsAction)

	return err
}

type NullableControlsAction struct {
	value *ControlsAction
	isSet bool
}

func (v NullableControlsAction) Get() *ControlsAction {
	return v.value
}

func (v *NullableControlsAction) Set(val *ControlsAction) {
	v.value = val
	v.isSet = true
}

func (v NullableControlsAction) IsSet() bool {
	return v.isSet
}

func (v *NullableControlsAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControlsAction(val *ControlsAction) *NullableControlsAction {
	return &NullableControlsAction{value: val, isSet: true}
}

func (v NullableControlsAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControlsAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


