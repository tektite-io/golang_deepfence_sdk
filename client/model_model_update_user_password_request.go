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

// checks if the ModelUpdateUserPasswordRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelUpdateUserPasswordRequest{}

// ModelUpdateUserPasswordRequest struct for ModelUpdateUserPasswordRequest
type ModelUpdateUserPasswordRequest struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

type _ModelUpdateUserPasswordRequest ModelUpdateUserPasswordRequest

// NewModelUpdateUserPasswordRequest instantiates a new ModelUpdateUserPasswordRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelUpdateUserPasswordRequest(newPassword string, oldPassword string) *ModelUpdateUserPasswordRequest {
	this := ModelUpdateUserPasswordRequest{}
	this.NewPassword = newPassword
	this.OldPassword = oldPassword
	return &this
}

// NewModelUpdateUserPasswordRequestWithDefaults instantiates a new ModelUpdateUserPasswordRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelUpdateUserPasswordRequestWithDefaults() *ModelUpdateUserPasswordRequest {
	this := ModelUpdateUserPasswordRequest{}
	return &this
}

// GetNewPassword returns the NewPassword field value
func (o *ModelUpdateUserPasswordRequest) GetNewPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value
// and a boolean to check if the value has been set.
func (o *ModelUpdateUserPasswordRequest) GetNewPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewPassword, true
}

// SetNewPassword sets field value
func (o *ModelUpdateUserPasswordRequest) SetNewPassword(v string) {
	o.NewPassword = v
}

// GetOldPassword returns the OldPassword field value
func (o *ModelUpdateUserPasswordRequest) GetOldPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OldPassword
}

// GetOldPasswordOk returns a tuple with the OldPassword field value
// and a boolean to check if the value has been set.
func (o *ModelUpdateUserPasswordRequest) GetOldPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OldPassword, true
}

// SetOldPassword sets field value
func (o *ModelUpdateUserPasswordRequest) SetOldPassword(v string) {
	o.OldPassword = v
}

func (o ModelUpdateUserPasswordRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelUpdateUserPasswordRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["new_password"] = o.NewPassword
	toSerialize["old_password"] = o.OldPassword
	return toSerialize, nil
}

func (o *ModelUpdateUserPasswordRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"new_password",
		"old_password",
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

	varModelUpdateUserPasswordRequest := _ModelUpdateUserPasswordRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelUpdateUserPasswordRequest)

	if err != nil {
		return err
	}

	*o = ModelUpdateUserPasswordRequest(varModelUpdateUserPasswordRequest)

	return err
}

type NullableModelUpdateUserPasswordRequest struct {
	value *ModelUpdateUserPasswordRequest
	isSet bool
}

func (v NullableModelUpdateUserPasswordRequest) Get() *ModelUpdateUserPasswordRequest {
	return v.value
}

func (v *NullableModelUpdateUserPasswordRequest) Set(val *ModelUpdateUserPasswordRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelUpdateUserPasswordRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelUpdateUserPasswordRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelUpdateUserPasswordRequest(val *ModelUpdateUserPasswordRequest) *NullableModelUpdateUserPasswordRequest {
	return &NullableModelUpdateUserPasswordRequest{value: val, isSet: true}
}

func (v NullableModelUpdateUserPasswordRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelUpdateUserPasswordRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


