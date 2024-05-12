/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.2.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelSettingsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSettingsResponse{}

// ModelSettingsResponse struct for ModelSettingsResponse
type ModelSettingsResponse struct {
	Description string `json:"description"`
	Id int32 `json:"id"`
	Key string `json:"key"`
	Label string `json:"label"`
	Value interface{} `json:"value"`
}

type _ModelSettingsResponse ModelSettingsResponse

// NewModelSettingsResponse instantiates a new ModelSettingsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSettingsResponse(description string, id int32, key string, label string, value interface{}) *ModelSettingsResponse {
	this := ModelSettingsResponse{}
	this.Description = description
	this.Id = id
	this.Key = key
	this.Label = label
	this.Value = value
	return &this
}

// NewModelSettingsResponseWithDefaults instantiates a new ModelSettingsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSettingsResponseWithDefaults() *ModelSettingsResponse {
	this := ModelSettingsResponse{}
	return &this
}

// GetDescription returns the Description field value
func (o *ModelSettingsResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ModelSettingsResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ModelSettingsResponse) SetDescription(v string) {
	o.Description = v
}

// GetId returns the Id field value
func (o *ModelSettingsResponse) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ModelSettingsResponse) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ModelSettingsResponse) SetId(v int32) {
	o.Id = v
}

// GetKey returns the Key field value
func (o *ModelSettingsResponse) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ModelSettingsResponse) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *ModelSettingsResponse) SetKey(v string) {
	o.Key = v
}

// GetLabel returns the Label field value
func (o *ModelSettingsResponse) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *ModelSettingsResponse) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *ModelSettingsResponse) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ModelSettingsResponse) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSettingsResponse) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *ModelSettingsResponse) SetValue(v interface{}) {
	o.Value = v
}

func (o ModelSettingsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSettingsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["id"] = o.Id
	toSerialize["key"] = o.Key
	toSerialize["label"] = o.Label
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

func (o *ModelSettingsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"id",
		"key",
		"label",
		"value",
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

	varModelSettingsResponse := _ModelSettingsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelSettingsResponse)

	if err != nil {
		return err
	}

	*o = ModelSettingsResponse(varModelSettingsResponse)

	return err
}

type NullableModelSettingsResponse struct {
	value *ModelSettingsResponse
	isSet bool
}

func (v NullableModelSettingsResponse) Get() *ModelSettingsResponse {
	return v.value
}

func (v *NullableModelSettingsResponse) Set(val *ModelSettingsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSettingsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSettingsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSettingsResponse(val *ModelSettingsResponse) *NullableModelSettingsResponse {
	return &NullableModelSettingsResponse{value: val, isSet: true}
}

func (v NullableModelSettingsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSettingsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


