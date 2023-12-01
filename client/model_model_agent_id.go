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

// checks if the ModelAgentID type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelAgentID{}

// ModelAgentID struct for ModelAgentID
type ModelAgentID struct {
	AvailableWorkload int32 `json:"available_workload"`
	NodeId string `json:"node_id"`
}

type _ModelAgentID ModelAgentID

// NewModelAgentID instantiates a new ModelAgentID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelAgentID(availableWorkload int32, nodeId string) *ModelAgentID {
	this := ModelAgentID{}
	this.AvailableWorkload = availableWorkload
	this.NodeId = nodeId
	return &this
}

// NewModelAgentIDWithDefaults instantiates a new ModelAgentID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelAgentIDWithDefaults() *ModelAgentID {
	this := ModelAgentID{}
	return &this
}

// GetAvailableWorkload returns the AvailableWorkload field value
func (o *ModelAgentID) GetAvailableWorkload() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AvailableWorkload
}

// GetAvailableWorkloadOk returns a tuple with the AvailableWorkload field value
// and a boolean to check if the value has been set.
func (o *ModelAgentID) GetAvailableWorkloadOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableWorkload, true
}

// SetAvailableWorkload sets field value
func (o *ModelAgentID) SetAvailableWorkload(v int32) {
	o.AvailableWorkload = v
}

// GetNodeId returns the NodeId field value
func (o *ModelAgentID) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelAgentID) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelAgentID) SetNodeId(v string) {
	o.NodeId = v
}

func (o ModelAgentID) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelAgentID) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["available_workload"] = o.AvailableWorkload
	toSerialize["node_id"] = o.NodeId
	return toSerialize, nil
}

func (o *ModelAgentID) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"available_workload",
		"node_id",
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

	varModelAgentID := _ModelAgentID{}

	err = json.Unmarshal(bytes, &varModelAgentID)

	if err != nil {
		return err
	}

	*o = ModelAgentID(varModelAgentID)

	return err
}

type NullableModelAgentID struct {
	value *ModelAgentID
	isSet bool
}

func (v NullableModelAgentID) Get() *ModelAgentID {
	return v.value
}

func (v *NullableModelAgentID) Set(val *ModelAgentID) {
	v.value = val
	v.isSet = true
}

func (v NullableModelAgentID) IsSet() bool {
	return v.isSet
}

func (v *NullableModelAgentID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelAgentID(val *ModelAgentID) *NullableModelAgentID {
	return &NullableModelAgentID{value: val, isSet: true}
}

func (v NullableModelAgentID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelAgentID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


