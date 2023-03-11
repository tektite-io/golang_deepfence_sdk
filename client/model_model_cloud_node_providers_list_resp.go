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

// checks if the ModelCloudNodeProvidersListResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCloudNodeProvidersListResp{}

// ModelCloudNodeProvidersListResp struct for ModelCloudNodeProvidersListResp
type ModelCloudNodeProvidersListResp struct {
	Providers []ModelPostureProvider `json:"providers"`
}

// NewModelCloudNodeProvidersListResp instantiates a new ModelCloudNodeProvidersListResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCloudNodeProvidersListResp(providers []ModelPostureProvider) *ModelCloudNodeProvidersListResp {
	this := ModelCloudNodeProvidersListResp{}
	this.Providers = providers
	return &this
}

// NewModelCloudNodeProvidersListRespWithDefaults instantiates a new ModelCloudNodeProvidersListResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCloudNodeProvidersListRespWithDefaults() *ModelCloudNodeProvidersListResp {
	this := ModelCloudNodeProvidersListResp{}
	return &this
}

// GetProviders returns the Providers field value
// If the value is explicit nil, the zero value for []ModelPostureProvider will be returned
func (o *ModelCloudNodeProvidersListResp) GetProviders() []ModelPostureProvider {
	if o == nil {
		var ret []ModelPostureProvider
		return ret
	}

	return o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCloudNodeProvidersListResp) GetProvidersOk() ([]ModelPostureProvider, bool) {
	if o == nil || IsNil(o.Providers) {
		return nil, false
	}
	return o.Providers, true
}

// SetProviders sets field value
func (o *ModelCloudNodeProvidersListResp) SetProviders(v []ModelPostureProvider) {
	o.Providers = v
}

func (o ModelCloudNodeProvidersListResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCloudNodeProvidersListResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Providers != nil {
		toSerialize["providers"] = o.Providers
	}
	return toSerialize, nil
}

type NullableModelCloudNodeProvidersListResp struct {
	value *ModelCloudNodeProvidersListResp
	isSet bool
}

func (v NullableModelCloudNodeProvidersListResp) Get() *ModelCloudNodeProvidersListResp {
	return v.value
}

func (v *NullableModelCloudNodeProvidersListResp) Set(val *ModelCloudNodeProvidersListResp) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCloudNodeProvidersListResp) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCloudNodeProvidersListResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCloudNodeProvidersListResp(val *ModelCloudNodeProvidersListResp) *NullableModelCloudNodeProvidersListResp {
	return &NullableModelCloudNodeProvidersListResp{value: val, isSet: true}
}

func (v NullableModelCloudNodeProvidersListResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCloudNodeProvidersListResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


