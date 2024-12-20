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

// checks if the ModelInviteUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelInviteUserRequest{}

// ModelInviteUserRequest struct for ModelInviteUserRequest
type ModelInviteUserRequest struct {
	Action string `json:"action"`
	Email string `json:"email"`
	Role string `json:"role"`
}

type _ModelInviteUserRequest ModelInviteUserRequest

// NewModelInviteUserRequest instantiates a new ModelInviteUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelInviteUserRequest(action string, email string, role string) *ModelInviteUserRequest {
	this := ModelInviteUserRequest{}
	this.Action = action
	this.Email = email
	this.Role = role
	return &this
}

// NewModelInviteUserRequestWithDefaults instantiates a new ModelInviteUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelInviteUserRequestWithDefaults() *ModelInviteUserRequest {
	this := ModelInviteUserRequest{}
	return &this
}

// GetAction returns the Action field value
func (o *ModelInviteUserRequest) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ModelInviteUserRequest) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ModelInviteUserRequest) SetAction(v string) {
	o.Action = v
}

// GetEmail returns the Email field value
func (o *ModelInviteUserRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ModelInviteUserRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ModelInviteUserRequest) SetEmail(v string) {
	o.Email = v
}

// GetRole returns the Role field value
func (o *ModelInviteUserRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ModelInviteUserRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ModelInviteUserRequest) SetRole(v string) {
	o.Role = v
}

func (o ModelInviteUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelInviteUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["email"] = o.Email
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

func (o *ModelInviteUserRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"email",
		"role",
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

	varModelInviteUserRequest := _ModelInviteUserRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelInviteUserRequest)

	if err != nil {
		return err
	}

	*o = ModelInviteUserRequest(varModelInviteUserRequest)

	return err
}

type NullableModelInviteUserRequest struct {
	value *ModelInviteUserRequest
	isSet bool
}

func (v NullableModelInviteUserRequest) Get() *ModelInviteUserRequest {
	return v.value
}

func (v *NullableModelInviteUserRequest) Set(val *ModelInviteUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelInviteUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelInviteUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelInviteUserRequest(val *ModelInviteUserRequest) *NullableModelInviteUserRequest {
	return &NullableModelInviteUserRequest{value: val, isSet: true}
}

func (v NullableModelInviteUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelInviteUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


