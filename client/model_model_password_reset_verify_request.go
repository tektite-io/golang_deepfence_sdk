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

// checks if the ModelPasswordResetVerifyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelPasswordResetVerifyRequest{}

// ModelPasswordResetVerifyRequest struct for ModelPasswordResetVerifyRequest
type ModelPasswordResetVerifyRequest struct {
	Code string `json:"code"`
	Password string `json:"password"`
}

// NewModelPasswordResetVerifyRequest instantiates a new ModelPasswordResetVerifyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelPasswordResetVerifyRequest(code string, password string) *ModelPasswordResetVerifyRequest {
	this := ModelPasswordResetVerifyRequest{}
	this.Code = code
	this.Password = password
	return &this
}

// NewModelPasswordResetVerifyRequestWithDefaults instantiates a new ModelPasswordResetVerifyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelPasswordResetVerifyRequestWithDefaults() *ModelPasswordResetVerifyRequest {
	this := ModelPasswordResetVerifyRequest{}
	return &this
}

// GetCode returns the Code field value
func (o *ModelPasswordResetVerifyRequest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ModelPasswordResetVerifyRequest) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ModelPasswordResetVerifyRequest) SetCode(v string) {
	o.Code = v
}

// GetPassword returns the Password field value
func (o *ModelPasswordResetVerifyRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ModelPasswordResetVerifyRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ModelPasswordResetVerifyRequest) SetPassword(v string) {
	o.Password = v
}

func (o ModelPasswordResetVerifyRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelPasswordResetVerifyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

type NullableModelPasswordResetVerifyRequest struct {
	value *ModelPasswordResetVerifyRequest
	isSet bool
}

func (v NullableModelPasswordResetVerifyRequest) Get() *ModelPasswordResetVerifyRequest {
	return v.value
}

func (v *NullableModelPasswordResetVerifyRequest) Set(val *ModelPasswordResetVerifyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelPasswordResetVerifyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelPasswordResetVerifyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelPasswordResetVerifyRequest(val *ModelPasswordResetVerifyRequest) *NullableModelPasswordResetVerifyRequest {
	return &NullableModelPasswordResetVerifyRequest{value: val, isSet: true}
}

func (v NullableModelPasswordResetVerifyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelPasswordResetVerifyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


