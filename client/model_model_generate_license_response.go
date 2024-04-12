/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.2.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelGenerateLicenseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelGenerateLicenseResponse{}

// ModelGenerateLicenseResponse struct for ModelGenerateLicenseResponse
type ModelGenerateLicenseResponse struct {
	GenerateLicenseLink *string `json:"generate_license_link,omitempty"`
	Message string `json:"message"`
	Success bool `json:"success"`
}

type _ModelGenerateLicenseResponse ModelGenerateLicenseResponse

// NewModelGenerateLicenseResponse instantiates a new ModelGenerateLicenseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelGenerateLicenseResponse(message string, success bool) *ModelGenerateLicenseResponse {
	this := ModelGenerateLicenseResponse{}
	this.Message = message
	this.Success = success
	return &this
}

// NewModelGenerateLicenseResponseWithDefaults instantiates a new ModelGenerateLicenseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelGenerateLicenseResponseWithDefaults() *ModelGenerateLicenseResponse {
	this := ModelGenerateLicenseResponse{}
	return &this
}

// GetGenerateLicenseLink returns the GenerateLicenseLink field value if set, zero value otherwise.
func (o *ModelGenerateLicenseResponse) GetGenerateLicenseLink() string {
	if o == nil || IsNil(o.GenerateLicenseLink) {
		var ret string
		return ret
	}
	return *o.GenerateLicenseLink
}

// GetGenerateLicenseLinkOk returns a tuple with the GenerateLicenseLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGenerateLicenseResponse) GetGenerateLicenseLinkOk() (*string, bool) {
	if o == nil || IsNil(o.GenerateLicenseLink) {
		return nil, false
	}
	return o.GenerateLicenseLink, true
}

// HasGenerateLicenseLink returns a boolean if a field has been set.
func (o *ModelGenerateLicenseResponse) HasGenerateLicenseLink() bool {
	if o != nil && !IsNil(o.GenerateLicenseLink) {
		return true
	}

	return false
}

// SetGenerateLicenseLink gets a reference to the given string and assigns it to the GenerateLicenseLink field.
func (o *ModelGenerateLicenseResponse) SetGenerateLicenseLink(v string) {
	o.GenerateLicenseLink = &v
}

// GetMessage returns the Message field value
func (o *ModelGenerateLicenseResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ModelGenerateLicenseResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ModelGenerateLicenseResponse) SetMessage(v string) {
	o.Message = v
}

// GetSuccess returns the Success field value
func (o *ModelGenerateLicenseResponse) GetSuccess() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Success
}

// GetSuccessOk returns a tuple with the Success field value
// and a boolean to check if the value has been set.
func (o *ModelGenerateLicenseResponse) GetSuccessOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Success, true
}

// SetSuccess sets field value
func (o *ModelGenerateLicenseResponse) SetSuccess(v bool) {
	o.Success = v
}

func (o ModelGenerateLicenseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelGenerateLicenseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GenerateLicenseLink) {
		toSerialize["generate_license_link"] = o.GenerateLicenseLink
	}
	toSerialize["message"] = o.Message
	toSerialize["success"] = o.Success
	return toSerialize, nil
}

func (o *ModelGenerateLicenseResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"message",
		"success",
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

	varModelGenerateLicenseResponse := _ModelGenerateLicenseResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelGenerateLicenseResponse)

	if err != nil {
		return err
	}

	*o = ModelGenerateLicenseResponse(varModelGenerateLicenseResponse)

	return err
}

type NullableModelGenerateLicenseResponse struct {
	value *ModelGenerateLicenseResponse
	isSet bool
}

func (v NullableModelGenerateLicenseResponse) Get() *ModelGenerateLicenseResponse {
	return v.value
}

func (v *NullableModelGenerateLicenseResponse) Set(val *ModelGenerateLicenseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelGenerateLicenseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelGenerateLicenseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelGenerateLicenseResponse(val *ModelGenerateLicenseResponse) *NullableModelGenerateLicenseResponse {
	return &NullableModelGenerateLicenseResponse{value: val, isSet: true}
}

func (v NullableModelGenerateLicenseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelGenerateLicenseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


