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

// checks if the ModelGetAuditLogsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelGetAuditLogsRequest{}

// ModelGetAuditLogsRequest struct for ModelGetAuditLogsRequest
type ModelGetAuditLogsRequest struct {
	Window ModelFetchWindow `json:"window"`
}

// NewModelGetAuditLogsRequest instantiates a new ModelGetAuditLogsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelGetAuditLogsRequest(window ModelFetchWindow) *ModelGetAuditLogsRequest {
	this := ModelGetAuditLogsRequest{}
	this.Window = window
	return &this
}

// NewModelGetAuditLogsRequestWithDefaults instantiates a new ModelGetAuditLogsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelGetAuditLogsRequestWithDefaults() *ModelGetAuditLogsRequest {
	this := ModelGetAuditLogsRequest{}
	return &this
}

// GetWindow returns the Window field value
func (o *ModelGetAuditLogsRequest) GetWindow() ModelFetchWindow {
	if o == nil {
		var ret ModelFetchWindow
		return ret
	}

	return o.Window
}

// GetWindowOk returns a tuple with the Window field value
// and a boolean to check if the value has been set.
func (o *ModelGetAuditLogsRequest) GetWindowOk() (*ModelFetchWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Window, true
}

// SetWindow sets field value
func (o *ModelGetAuditLogsRequest) SetWindow(v ModelFetchWindow) {
	o.Window = v
}

func (o ModelGetAuditLogsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelGetAuditLogsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["window"] = o.Window
	return toSerialize, nil
}

type NullableModelGetAuditLogsRequest struct {
	value *ModelGetAuditLogsRequest
	isSet bool
}

func (v NullableModelGetAuditLogsRequest) Get() *ModelGetAuditLogsRequest {
	return v.value
}

func (v *NullableModelGetAuditLogsRequest) Set(val *ModelGetAuditLogsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelGetAuditLogsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelGetAuditLogsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelGetAuditLogsRequest(val *ModelGetAuditLogsRequest) *NullableModelGetAuditLogsRequest {
	return &NullableModelGetAuditLogsRequest{value: val, isSet: true}
}

func (v NullableModelGetAuditLogsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelGetAuditLogsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


