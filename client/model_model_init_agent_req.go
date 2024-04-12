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

// checks if the ModelInitAgentReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelInitAgentReq{}

// ModelInitAgentReq struct for ModelInitAgentReq
type ModelInitAgentReq struct {
	AvailableWorkload int32 `json:"available_workload"`
	NodeId string `json:"node_id"`
	Version string `json:"version"`
}

type _ModelInitAgentReq ModelInitAgentReq

// NewModelInitAgentReq instantiates a new ModelInitAgentReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelInitAgentReq(availableWorkload int32, nodeId string, version string) *ModelInitAgentReq {
	this := ModelInitAgentReq{}
	this.AvailableWorkload = availableWorkload
	this.NodeId = nodeId
	this.Version = version
	return &this
}

// NewModelInitAgentReqWithDefaults instantiates a new ModelInitAgentReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelInitAgentReqWithDefaults() *ModelInitAgentReq {
	this := ModelInitAgentReq{}
	return &this
}

// GetAvailableWorkload returns the AvailableWorkload field value
func (o *ModelInitAgentReq) GetAvailableWorkload() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AvailableWorkload
}

// GetAvailableWorkloadOk returns a tuple with the AvailableWorkload field value
// and a boolean to check if the value has been set.
func (o *ModelInitAgentReq) GetAvailableWorkloadOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableWorkload, true
}

// SetAvailableWorkload sets field value
func (o *ModelInitAgentReq) SetAvailableWorkload(v int32) {
	o.AvailableWorkload = v
}

// GetNodeId returns the NodeId field value
func (o *ModelInitAgentReq) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelInitAgentReq) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelInitAgentReq) SetNodeId(v string) {
	o.NodeId = v
}

// GetVersion returns the Version field value
func (o *ModelInitAgentReq) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ModelInitAgentReq) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ModelInitAgentReq) SetVersion(v string) {
	o.Version = v
}

func (o ModelInitAgentReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelInitAgentReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["available_workload"] = o.AvailableWorkload
	toSerialize["node_id"] = o.NodeId
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

func (o *ModelInitAgentReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"available_workload",
		"node_id",
		"version",
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

	varModelInitAgentReq := _ModelInitAgentReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelInitAgentReq)

	if err != nil {
		return err
	}

	*o = ModelInitAgentReq(varModelInitAgentReq)

	return err
}

type NullableModelInitAgentReq struct {
	value *ModelInitAgentReq
	isSet bool
}

func (v NullableModelInitAgentReq) Get() *ModelInitAgentReq {
	return v.value
}

func (v *NullableModelInitAgentReq) Set(val *ModelInitAgentReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelInitAgentReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelInitAgentReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelInitAgentReq(val *ModelInitAgentReq) *NullableModelInitAgentReq {
	return &NullableModelInitAgentReq{value: val, isSet: true}
}

func (v NullableModelInitAgentReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelInitAgentReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


