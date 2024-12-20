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

// checks if the ModelRegisterInvitedUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelRegisterInvitedUserRequest{}

// ModelRegisterInvitedUserRequest struct for ModelRegisterInvitedUserRequest
type ModelRegisterInvitedUserRequest struct {
	Code string `json:"code"`
	FirstName string `json:"first_name"`
	IsTemporaryPassword *bool `json:"is_temporary_password,omitempty"`
	LastName string `json:"last_name"`
	Namespace string `json:"namespace"`
	Password string `json:"password"`
}

type _ModelRegisterInvitedUserRequest ModelRegisterInvitedUserRequest

// NewModelRegisterInvitedUserRequest instantiates a new ModelRegisterInvitedUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelRegisterInvitedUserRequest(code string, firstName string, lastName string, namespace string, password string) *ModelRegisterInvitedUserRequest {
	this := ModelRegisterInvitedUserRequest{}
	this.Code = code
	this.FirstName = firstName
	this.LastName = lastName
	this.Namespace = namespace
	this.Password = password
	return &this
}

// NewModelRegisterInvitedUserRequestWithDefaults instantiates a new ModelRegisterInvitedUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelRegisterInvitedUserRequestWithDefaults() *ModelRegisterInvitedUserRequest {
	this := ModelRegisterInvitedUserRequest{}
	return &this
}

// GetCode returns the Code field value
func (o *ModelRegisterInvitedUserRequest) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ModelRegisterInvitedUserRequest) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ModelRegisterInvitedUserRequest) SetCode(v string) {
	o.Code = v
}

// GetFirstName returns the FirstName field value
func (o *ModelRegisterInvitedUserRequest) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *ModelRegisterInvitedUserRequest) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *ModelRegisterInvitedUserRequest) SetFirstName(v string) {
	o.FirstName = v
}

// GetIsTemporaryPassword returns the IsTemporaryPassword field value if set, zero value otherwise.
func (o *ModelRegisterInvitedUserRequest) GetIsTemporaryPassword() bool {
	if o == nil || IsNil(o.IsTemporaryPassword) {
		var ret bool
		return ret
	}
	return *o.IsTemporaryPassword
}

// GetIsTemporaryPasswordOk returns a tuple with the IsTemporaryPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelRegisterInvitedUserRequest) GetIsTemporaryPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTemporaryPassword) {
		return nil, false
	}
	return o.IsTemporaryPassword, true
}

// HasIsTemporaryPassword returns a boolean if a field has been set.
func (o *ModelRegisterInvitedUserRequest) HasIsTemporaryPassword() bool {
	if o != nil && !IsNil(o.IsTemporaryPassword) {
		return true
	}

	return false
}

// SetIsTemporaryPassword gets a reference to the given bool and assigns it to the IsTemporaryPassword field.
func (o *ModelRegisterInvitedUserRequest) SetIsTemporaryPassword(v bool) {
	o.IsTemporaryPassword = &v
}

// GetLastName returns the LastName field value
func (o *ModelRegisterInvitedUserRequest) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *ModelRegisterInvitedUserRequest) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *ModelRegisterInvitedUserRequest) SetLastName(v string) {
	o.LastName = v
}

// GetNamespace returns the Namespace field value
func (o *ModelRegisterInvitedUserRequest) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *ModelRegisterInvitedUserRequest) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *ModelRegisterInvitedUserRequest) SetNamespace(v string) {
	o.Namespace = v
}

// GetPassword returns the Password field value
func (o *ModelRegisterInvitedUserRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ModelRegisterInvitedUserRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ModelRegisterInvitedUserRequest) SetPassword(v string) {
	o.Password = v
}

func (o ModelRegisterInvitedUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelRegisterInvitedUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["first_name"] = o.FirstName
	if !IsNil(o.IsTemporaryPassword) {
		toSerialize["is_temporary_password"] = o.IsTemporaryPassword
	}
	toSerialize["last_name"] = o.LastName
	toSerialize["namespace"] = o.Namespace
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *ModelRegisterInvitedUserRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"first_name",
		"last_name",
		"namespace",
		"password",
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

	varModelRegisterInvitedUserRequest := _ModelRegisterInvitedUserRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelRegisterInvitedUserRequest)

	if err != nil {
		return err
	}

	*o = ModelRegisterInvitedUserRequest(varModelRegisterInvitedUserRequest)

	return err
}

type NullableModelRegisterInvitedUserRequest struct {
	value *ModelRegisterInvitedUserRequest
	isSet bool
}

func (v NullableModelRegisterInvitedUserRequest) Get() *ModelRegisterInvitedUserRequest {
	return v.value
}

func (v *NullableModelRegisterInvitedUserRequest) Set(val *ModelRegisterInvitedUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRegisterInvitedUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRegisterInvitedUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRegisterInvitedUserRequest(val *ModelRegisterInvitedUserRequest) *NullableModelRegisterInvitedUserRequest {
	return &NullableModelRegisterInvitedUserRequest{value: val, isSet: true}
}

func (v NullableModelRegisterInvitedUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRegisterInvitedUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


