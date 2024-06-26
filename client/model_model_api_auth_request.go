/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.3.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelAPIAuthRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelAPIAuthRequest{}

// ModelAPIAuthRequest struct for ModelAPIAuthRequest
type ModelAPIAuthRequest struct {
	ApiToken string `json:"api_token"`
}

type _ModelAPIAuthRequest ModelAPIAuthRequest

// NewModelAPIAuthRequest instantiates a new ModelAPIAuthRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelAPIAuthRequest(apiToken string) *ModelAPIAuthRequest {
	this := ModelAPIAuthRequest{}
	this.ApiToken = apiToken
	return &this
}

// NewModelAPIAuthRequestWithDefaults instantiates a new ModelAPIAuthRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelAPIAuthRequestWithDefaults() *ModelAPIAuthRequest {
	this := ModelAPIAuthRequest{}
	return &this
}

// GetApiToken returns the ApiToken field value
func (o *ModelAPIAuthRequest) GetApiToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiToken
}

// GetApiTokenOk returns a tuple with the ApiToken field value
// and a boolean to check if the value has been set.
func (o *ModelAPIAuthRequest) GetApiTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiToken, true
}

// SetApiToken sets field value
func (o *ModelAPIAuthRequest) SetApiToken(v string) {
	o.ApiToken = v
}

func (o ModelAPIAuthRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelAPIAuthRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["api_token"] = o.ApiToken
	return toSerialize, nil
}

func (o *ModelAPIAuthRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"api_token",
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

	varModelAPIAuthRequest := _ModelAPIAuthRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelAPIAuthRequest)

	if err != nil {
		return err
	}

	*o = ModelAPIAuthRequest(varModelAPIAuthRequest)

	return err
}

type NullableModelAPIAuthRequest struct {
	value *ModelAPIAuthRequest
	isSet bool
}

func (v NullableModelAPIAuthRequest) Get() *ModelAPIAuthRequest {
	return v.value
}

func (v *NullableModelAPIAuthRequest) Set(val *ModelAPIAuthRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelAPIAuthRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelAPIAuthRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelAPIAuthRequest(val *ModelAPIAuthRequest) *NullableModelAPIAuthRequest {
	return &NullableModelAPIAuthRequest{value: val, isSet: true}
}

func (v NullableModelAPIAuthRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelAPIAuthRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


