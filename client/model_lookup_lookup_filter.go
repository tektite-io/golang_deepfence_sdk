/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.2.1
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the LookupLookupFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LookupLookupFilter{}

// LookupLookupFilter struct for LookupLookupFilter
type LookupLookupFilter struct {
	InFieldFilter []string `json:"in_field_filter"`
	NodeIds []string `json:"node_ids"`
	Window ModelFetchWindow `json:"window"`
}

type _LookupLookupFilter LookupLookupFilter

// NewLookupLookupFilter instantiates a new LookupLookupFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLookupLookupFilter(inFieldFilter []string, nodeIds []string, window ModelFetchWindow) *LookupLookupFilter {
	this := LookupLookupFilter{}
	this.InFieldFilter = inFieldFilter
	this.NodeIds = nodeIds
	this.Window = window
	return &this
}

// NewLookupLookupFilterWithDefaults instantiates a new LookupLookupFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLookupLookupFilterWithDefaults() *LookupLookupFilter {
	this := LookupLookupFilter{}
	return &this
}

// GetInFieldFilter returns the InFieldFilter field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *LookupLookupFilter) GetInFieldFilter() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.InFieldFilter
}

// GetInFieldFilterOk returns a tuple with the InFieldFilter field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LookupLookupFilter) GetInFieldFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.InFieldFilter) {
		return nil, false
	}
	return o.InFieldFilter, true
}

// SetInFieldFilter sets field value
func (o *LookupLookupFilter) SetInFieldFilter(v []string) {
	o.InFieldFilter = v
}

// GetNodeIds returns the NodeIds field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *LookupLookupFilter) GetNodeIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NodeIds
}

// GetNodeIdsOk returns a tuple with the NodeIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *LookupLookupFilter) GetNodeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.NodeIds) {
		return nil, false
	}
	return o.NodeIds, true
}

// SetNodeIds sets field value
func (o *LookupLookupFilter) SetNodeIds(v []string) {
	o.NodeIds = v
}

// GetWindow returns the Window field value
func (o *LookupLookupFilter) GetWindow() ModelFetchWindow {
	if o == nil {
		var ret ModelFetchWindow
		return ret
	}

	return o.Window
}

// GetWindowOk returns a tuple with the Window field value
// and a boolean to check if the value has been set.
func (o *LookupLookupFilter) GetWindowOk() (*ModelFetchWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Window, true
}

// SetWindow sets field value
func (o *LookupLookupFilter) SetWindow(v ModelFetchWindow) {
	o.Window = v
}

func (o LookupLookupFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LookupLookupFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.InFieldFilter != nil {
		toSerialize["in_field_filter"] = o.InFieldFilter
	}
	if o.NodeIds != nil {
		toSerialize["node_ids"] = o.NodeIds
	}
	toSerialize["window"] = o.Window
	return toSerialize, nil
}

func (o *LookupLookupFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"in_field_filter",
		"node_ids",
		"window",
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

	varLookupLookupFilter := _LookupLookupFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLookupLookupFilter)

	if err != nil {
		return err
	}

	*o = LookupLookupFilter(varLookupLookupFilter)

	return err
}

type NullableLookupLookupFilter struct {
	value *LookupLookupFilter
	isSet bool
}

func (v NullableLookupLookupFilter) Get() *LookupLookupFilter {
	return v.value
}

func (v *NullableLookupLookupFilter) Set(val *LookupLookupFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableLookupLookupFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableLookupLookupFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLookupLookupFilter(val *LookupLookupFilter) *NullableLookupLookupFilter {
	return &NullableLookupLookupFilter{value: val, isSet: true}
}

func (v NullableLookupLookupFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLookupLookupFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


