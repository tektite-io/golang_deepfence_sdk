/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.0.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package deepfence_server_client

import (
	"encoding/json"
)

// checks if the ModelLoginRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelLoginRequest{}

// ModelLoginRequest struct for ModelLoginRequest
type ModelLoginRequest struct {
	Email *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

// NewModelLoginRequest instantiates a new ModelLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelLoginRequest() *ModelLoginRequest {
	this := ModelLoginRequest{}
	return &this
}

// NewModelLoginRequestWithDefaults instantiates a new ModelLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelLoginRequestWithDefaults() *ModelLoginRequest {
	this := ModelLoginRequest{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ModelLoginRequest) GetEmail() string {
	if o == nil || isNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLoginRequest) GetEmailOk() (*string, bool) {
	if o == nil || isNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ModelLoginRequest) HasEmail() bool {
	if o != nil && !isNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ModelLoginRequest) SetEmail(v string) {
	o.Email = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *ModelLoginRequest) GetPassword() string {
	if o == nil || isNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLoginRequest) GetPasswordOk() (*string, bool) {
	if o == nil || isNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *ModelLoginRequest) HasPassword() bool {
	if o != nil && !isNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *ModelLoginRequest) SetPassword(v string) {
	o.Password = &v
}

func (o ModelLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelLoginRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !isNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	return toSerialize, nil
}

type NullableModelLoginRequest struct {
	value *ModelLoginRequest
	isSet bool
}

func (v NullableModelLoginRequest) Get() *ModelLoginRequest {
	return v.value
}

func (v *NullableModelLoginRequest) Set(val *ModelLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelLoginRequest(val *ModelLoginRequest) *NullableModelLoginRequest {
	return &NullableModelLoginRequest{value: val, isSet: true}
}

func (v NullableModelLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


