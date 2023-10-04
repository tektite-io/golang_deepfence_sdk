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

// checks if the CompletionCompletionNodeFieldReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompletionCompletionNodeFieldReq{}

// CompletionCompletionNodeFieldReq struct for CompletionCompletionNodeFieldReq
type CompletionCompletionNodeFieldReq struct {
	Completion string `json:"completion"`
	FieldName string `json:"field_name"`
	Window ModelFetchWindow `json:"window"`
}

// NewCompletionCompletionNodeFieldReq instantiates a new CompletionCompletionNodeFieldReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompletionCompletionNodeFieldReq(completion string, fieldName string, window ModelFetchWindow) *CompletionCompletionNodeFieldReq {
	this := CompletionCompletionNodeFieldReq{}
	this.Completion = completion
	this.FieldName = fieldName
	this.Window = window
	return &this
}

// NewCompletionCompletionNodeFieldReqWithDefaults instantiates a new CompletionCompletionNodeFieldReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompletionCompletionNodeFieldReqWithDefaults() *CompletionCompletionNodeFieldReq {
	this := CompletionCompletionNodeFieldReq{}
	return &this
}

// GetCompletion returns the Completion field value
func (o *CompletionCompletionNodeFieldReq) GetCompletion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Completion
}

// GetCompletionOk returns a tuple with the Completion field value
// and a boolean to check if the value has been set.
func (o *CompletionCompletionNodeFieldReq) GetCompletionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Completion, true
}

// SetCompletion sets field value
func (o *CompletionCompletionNodeFieldReq) SetCompletion(v string) {
	o.Completion = v
}

// GetFieldName returns the FieldName field value
func (o *CompletionCompletionNodeFieldReq) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *CompletionCompletionNodeFieldReq) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *CompletionCompletionNodeFieldReq) SetFieldName(v string) {
	o.FieldName = v
}

// GetWindow returns the Window field value
func (o *CompletionCompletionNodeFieldReq) GetWindow() ModelFetchWindow {
	if o == nil {
		var ret ModelFetchWindow
		return ret
	}

	return o.Window
}

// GetWindowOk returns a tuple with the Window field value
// and a boolean to check if the value has been set.
func (o *CompletionCompletionNodeFieldReq) GetWindowOk() (*ModelFetchWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Window, true
}

// SetWindow sets field value
func (o *CompletionCompletionNodeFieldReq) SetWindow(v ModelFetchWindow) {
	o.Window = v
}

func (o CompletionCompletionNodeFieldReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompletionCompletionNodeFieldReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["completion"] = o.Completion
	toSerialize["field_name"] = o.FieldName
	toSerialize["window"] = o.Window
	return toSerialize, nil
}

type NullableCompletionCompletionNodeFieldReq struct {
	value *CompletionCompletionNodeFieldReq
	isSet bool
}

func (v NullableCompletionCompletionNodeFieldReq) Get() *CompletionCompletionNodeFieldReq {
	return v.value
}

func (v *NullableCompletionCompletionNodeFieldReq) Set(val *CompletionCompletionNodeFieldReq) {
	v.value = val
	v.isSet = true
}

func (v NullableCompletionCompletionNodeFieldReq) IsSet() bool {
	return v.isSet
}

func (v *NullableCompletionCompletionNodeFieldReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompletionCompletionNodeFieldReq(val *CompletionCompletionNodeFieldReq) *NullableCompletionCompletionNodeFieldReq {
	return &NullableCompletionCompletionNodeFieldReq{value: val, isSet: true}
}

func (v NullableCompletionCompletionNodeFieldReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompletionCompletionNodeFieldReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


