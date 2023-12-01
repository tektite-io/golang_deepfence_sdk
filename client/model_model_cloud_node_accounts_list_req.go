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
	"fmt"
)

// checks if the ModelCloudNodeAccountsListReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCloudNodeAccountsListReq{}

// ModelCloudNodeAccountsListReq struct for ModelCloudNodeAccountsListReq
type ModelCloudNodeAccountsListReq struct {
	CloudProvider *string `json:"cloud_provider,omitempty"`
	Window ModelFetchWindow `json:"window"`
}

type _ModelCloudNodeAccountsListReq ModelCloudNodeAccountsListReq

// NewModelCloudNodeAccountsListReq instantiates a new ModelCloudNodeAccountsListReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCloudNodeAccountsListReq(window ModelFetchWindow) *ModelCloudNodeAccountsListReq {
	this := ModelCloudNodeAccountsListReq{}
	this.Window = window
	return &this
}

// NewModelCloudNodeAccountsListReqWithDefaults instantiates a new ModelCloudNodeAccountsListReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCloudNodeAccountsListReqWithDefaults() *ModelCloudNodeAccountsListReq {
	this := ModelCloudNodeAccountsListReq{}
	return &this
}

// GetCloudProvider returns the CloudProvider field value if set, zero value otherwise.
func (o *ModelCloudNodeAccountsListReq) GetCloudProvider() string {
	if o == nil || IsNil(o.CloudProvider) {
		var ret string
		return ret
	}
	return *o.CloudProvider
}

// GetCloudProviderOk returns a tuple with the CloudProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeAccountsListReq) GetCloudProviderOk() (*string, bool) {
	if o == nil || IsNil(o.CloudProvider) {
		return nil, false
	}
	return o.CloudProvider, true
}

// HasCloudProvider returns a boolean if a field has been set.
func (o *ModelCloudNodeAccountsListReq) HasCloudProvider() bool {
	if o != nil && !IsNil(o.CloudProvider) {
		return true
	}

	return false
}

// SetCloudProvider gets a reference to the given string and assigns it to the CloudProvider field.
func (o *ModelCloudNodeAccountsListReq) SetCloudProvider(v string) {
	o.CloudProvider = &v
}

// GetWindow returns the Window field value
func (o *ModelCloudNodeAccountsListReq) GetWindow() ModelFetchWindow {
	if o == nil {
		var ret ModelFetchWindow
		return ret
	}

	return o.Window
}

// GetWindowOk returns a tuple with the Window field value
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeAccountsListReq) GetWindowOk() (*ModelFetchWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Window, true
}

// SetWindow sets field value
func (o *ModelCloudNodeAccountsListReq) SetWindow(v ModelFetchWindow) {
	o.Window = v
}

func (o ModelCloudNodeAccountsListReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCloudNodeAccountsListReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CloudProvider) {
		toSerialize["cloud_provider"] = o.CloudProvider
	}
	toSerialize["window"] = o.Window
	return toSerialize, nil
}

func (o *ModelCloudNodeAccountsListReq) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"window",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varModelCloudNodeAccountsListReq := _ModelCloudNodeAccountsListReq{}

	err = json.Unmarshal(bytes, &varModelCloudNodeAccountsListReq)

	if err != nil {
		return err
	}

	*o = ModelCloudNodeAccountsListReq(varModelCloudNodeAccountsListReq)

	return err
}

type NullableModelCloudNodeAccountsListReq struct {
	value *ModelCloudNodeAccountsListReq
	isSet bool
}

func (v NullableModelCloudNodeAccountsListReq) Get() *ModelCloudNodeAccountsListReq {
	return v.value
}

func (v *NullableModelCloudNodeAccountsListReq) Set(val *ModelCloudNodeAccountsListReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCloudNodeAccountsListReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCloudNodeAccountsListReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCloudNodeAccountsListReq(val *ModelCloudNodeAccountsListReq) *NullableModelCloudNodeAccountsListReq {
	return &NullableModelCloudNodeAccountsListReq{value: val, isSet: true}
}

func (v NullableModelCloudNodeAccountsListReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCloudNodeAccountsListReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


