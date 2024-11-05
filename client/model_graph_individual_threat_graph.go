/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the GraphIndividualThreatGraph type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphIndividualThreatGraph{}

// GraphIndividualThreatGraph struct for GraphIndividualThreatGraph
type GraphIndividualThreatGraph struct {
	AttackPath [][]string `json:"attack_path,omitempty"`
	CveAttackVector *string `json:"cve_attack_vector,omitempty"`
	CveId []string `json:"cve_id,omitempty"`
	Ports []interface{} `json:"ports,omitempty"`
}

// NewGraphIndividualThreatGraph instantiates a new GraphIndividualThreatGraph object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphIndividualThreatGraph() *GraphIndividualThreatGraph {
	this := GraphIndividualThreatGraph{}
	return &this
}

// NewGraphIndividualThreatGraphWithDefaults instantiates a new GraphIndividualThreatGraph object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphIndividualThreatGraphWithDefaults() *GraphIndividualThreatGraph {
	this := GraphIndividualThreatGraph{}
	return &this
}

// GetAttackPath returns the AttackPath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphIndividualThreatGraph) GetAttackPath() [][]string {
	if o == nil {
		var ret [][]string
		return ret
	}
	return o.AttackPath
}

// GetAttackPathOk returns a tuple with the AttackPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphIndividualThreatGraph) GetAttackPathOk() ([][]string, bool) {
	if o == nil || IsNil(o.AttackPath) {
		return nil, false
	}
	return o.AttackPath, true
}

// HasAttackPath returns a boolean if a field has been set.
func (o *GraphIndividualThreatGraph) HasAttackPath() bool {
	if o != nil && !IsNil(o.AttackPath) {
		return true
	}

	return false
}

// SetAttackPath gets a reference to the given [][]string and assigns it to the AttackPath field.
func (o *GraphIndividualThreatGraph) SetAttackPath(v [][]string) {
	o.AttackPath = v
}

// GetCveAttackVector returns the CveAttackVector field value if set, zero value otherwise.
func (o *GraphIndividualThreatGraph) GetCveAttackVector() string {
	if o == nil || IsNil(o.CveAttackVector) {
		var ret string
		return ret
	}
	return *o.CveAttackVector
}

// GetCveAttackVectorOk returns a tuple with the CveAttackVector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphIndividualThreatGraph) GetCveAttackVectorOk() (*string, bool) {
	if o == nil || IsNil(o.CveAttackVector) {
		return nil, false
	}
	return o.CveAttackVector, true
}

// HasCveAttackVector returns a boolean if a field has been set.
func (o *GraphIndividualThreatGraph) HasCveAttackVector() bool {
	if o != nil && !IsNil(o.CveAttackVector) {
		return true
	}

	return false
}

// SetCveAttackVector gets a reference to the given string and assigns it to the CveAttackVector field.
func (o *GraphIndividualThreatGraph) SetCveAttackVector(v string) {
	o.CveAttackVector = &v
}

// GetCveId returns the CveId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphIndividualThreatGraph) GetCveId() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CveId
}

// GetCveIdOk returns a tuple with the CveId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphIndividualThreatGraph) GetCveIdOk() ([]string, bool) {
	if o == nil || IsNil(o.CveId) {
		return nil, false
	}
	return o.CveId, true
}

// HasCveId returns a boolean if a field has been set.
func (o *GraphIndividualThreatGraph) HasCveId() bool {
	if o != nil && !IsNil(o.CveId) {
		return true
	}

	return false
}

// SetCveId gets a reference to the given []string and assigns it to the CveId field.
func (o *GraphIndividualThreatGraph) SetCveId(v []string) {
	o.CveId = v
}

// GetPorts returns the Ports field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *GraphIndividualThreatGraph) GetPorts() []interface{} {
	if o == nil {
		var ret []interface{}
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphIndividualThreatGraph) GetPortsOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *GraphIndividualThreatGraph) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []interface{} and assigns it to the Ports field.
func (o *GraphIndividualThreatGraph) SetPorts(v []interface{}) {
	o.Ports = v
}

func (o GraphIndividualThreatGraph) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphIndividualThreatGraph) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AttackPath != nil {
		toSerialize["attack_path"] = o.AttackPath
	}
	if !IsNil(o.CveAttackVector) {
		toSerialize["cve_attack_vector"] = o.CveAttackVector
	}
	if o.CveId != nil {
		toSerialize["cve_id"] = o.CveId
	}
	if o.Ports != nil {
		toSerialize["ports"] = o.Ports
	}
	return toSerialize, nil
}

type NullableGraphIndividualThreatGraph struct {
	value *GraphIndividualThreatGraph
	isSet bool
}

func (v NullableGraphIndividualThreatGraph) Get() *GraphIndividualThreatGraph {
	return v.value
}

func (v *NullableGraphIndividualThreatGraph) Set(val *GraphIndividualThreatGraph) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphIndividualThreatGraph) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphIndividualThreatGraph) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphIndividualThreatGraph(val *GraphIndividualThreatGraph) *NullableGraphIndividualThreatGraph {
	return &NullableGraphIndividualThreatGraph{value: val, isSet: true}
}

func (v NullableGraphIndividualThreatGraph) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphIndividualThreatGraph) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


