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

// checks if the ModelGraphResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelGraphResult{}

// ModelGraphResult struct for ModelGraphResult
type ModelGraphResult struct {
	Edges map[string]DetailedConnectionSummary `json:"edges"`
	Nodes map[string]DetailedNodeSummary `json:"nodes"`
}

type _ModelGraphResult ModelGraphResult

// NewModelGraphResult instantiates a new ModelGraphResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelGraphResult(edges map[string]DetailedConnectionSummary, nodes map[string]DetailedNodeSummary) *ModelGraphResult {
	this := ModelGraphResult{}
	this.Edges = edges
	this.Nodes = nodes
	return &this
}

// NewModelGraphResultWithDefaults instantiates a new ModelGraphResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelGraphResultWithDefaults() *ModelGraphResult {
	this := ModelGraphResult{}
	return &this
}

// GetEdges returns the Edges field value
func (o *ModelGraphResult) GetEdges() map[string]DetailedConnectionSummary {
	if o == nil {
		var ret map[string]DetailedConnectionSummary
		return ret
	}

	return o.Edges
}

// GetEdgesOk returns a tuple with the Edges field value
// and a boolean to check if the value has been set.
func (o *ModelGraphResult) GetEdgesOk() (*map[string]DetailedConnectionSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Edges, true
}

// SetEdges sets field value
func (o *ModelGraphResult) SetEdges(v map[string]DetailedConnectionSummary) {
	o.Edges = v
}

// GetNodes returns the Nodes field value
func (o *ModelGraphResult) GetNodes() map[string]DetailedNodeSummary {
	if o == nil {
		var ret map[string]DetailedNodeSummary
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *ModelGraphResult) GetNodesOk() (*map[string]DetailedNodeSummary, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nodes, true
}

// SetNodes sets field value
func (o *ModelGraphResult) SetNodes(v map[string]DetailedNodeSummary) {
	o.Nodes = v
}

func (o ModelGraphResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelGraphResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["edges"] = o.Edges
	toSerialize["nodes"] = o.Nodes
	return toSerialize, nil
}

func (o *ModelGraphResult) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"edges",
		"nodes",
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

	varModelGraphResult := _ModelGraphResult{}

	err = json.Unmarshal(bytes, &varModelGraphResult)

	if err != nil {
		return err
	}

	*o = ModelGraphResult(varModelGraphResult)

	return err
}

type NullableModelGraphResult struct {
	value *ModelGraphResult
	isSet bool
}

func (v NullableModelGraphResult) Get() *ModelGraphResult {
	return v.value
}

func (v *NullableModelGraphResult) Set(val *ModelGraphResult) {
	v.value = val
	v.isSet = true
}

func (v NullableModelGraphResult) IsSet() bool {
	return v.isSet
}

func (v *NullableModelGraphResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelGraphResult(val *ModelGraphResult) *NullableModelGraphResult {
	return &NullableModelGraphResult{value: val, isSet: true}
}

func (v NullableModelGraphResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelGraphResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


