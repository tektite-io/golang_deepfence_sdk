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
)

// checks if the ModelUpdateUserIDRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelUpdateUserIDRequest{}

// ModelUpdateUserIDRequest struct for ModelUpdateUserIDRequest
type ModelUpdateUserIDRequest struct {
	FirstName *string `json:"first_name,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	Role *string `json:"role,omitempty"`
}

// NewModelUpdateUserIDRequest instantiates a new ModelUpdateUserIDRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelUpdateUserIDRequest() *ModelUpdateUserIDRequest {
	this := ModelUpdateUserIDRequest{}
	return &this
}

// NewModelUpdateUserIDRequestWithDefaults instantiates a new ModelUpdateUserIDRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelUpdateUserIDRequestWithDefaults() *ModelUpdateUserIDRequest {
	this := ModelUpdateUserIDRequest{}
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *ModelUpdateUserIDRequest) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUpdateUserIDRequest) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *ModelUpdateUserIDRequest) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *ModelUpdateUserIDRequest) SetFirstName(v string) {
	o.FirstName = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ModelUpdateUserIDRequest) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUpdateUserIDRequest) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ModelUpdateUserIDRequest) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ModelUpdateUserIDRequest) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *ModelUpdateUserIDRequest) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUpdateUserIDRequest) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *ModelUpdateUserIDRequest) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *ModelUpdateUserIDRequest) SetLastName(v string) {
	o.LastName = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ModelUpdateUserIDRequest) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUpdateUserIDRequest) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ModelUpdateUserIDRequest) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *ModelUpdateUserIDRequest) SetRole(v string) {
	o.Role = &v
}

func (o ModelUpdateUserIDRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelUpdateUserIDRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstName) {
		toSerialize["first_name"] = o.FirstName
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.LastName) {
		toSerialize["last_name"] = o.LastName
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	return toSerialize, nil
}

type NullableModelUpdateUserIDRequest struct {
	value *ModelUpdateUserIDRequest
	isSet bool
}

func (v NullableModelUpdateUserIDRequest) Get() *ModelUpdateUserIDRequest {
	return v.value
}

func (v *NullableModelUpdateUserIDRequest) Set(val *ModelUpdateUserIDRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelUpdateUserIDRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelUpdateUserIDRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelUpdateUserIDRequest(val *ModelUpdateUserIDRequest) *NullableModelUpdateUserIDRequest {
	return &NullableModelUpdateUserIDRequest{value: val, isSet: true}
}

func (v NullableModelUpdateUserIDRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelUpdateUserIDRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


