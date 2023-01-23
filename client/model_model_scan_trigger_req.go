/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.0.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package deepfence_server_client

import (
	"encoding/json"
)

// checks if the ModelScanTriggerReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelScanTriggerReq{}

// ModelScanTriggerReq struct for ModelScanTriggerReq
type ModelScanTriggerReq struct {
	NodeId string `json:"node_id"`
	NodeType string `json:"node_type"`
}

// NewModelScanTriggerReq instantiates a new ModelScanTriggerReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelScanTriggerReq(nodeId string, nodeType string) *ModelScanTriggerReq {
	this := ModelScanTriggerReq{}
	this.NodeId = nodeId
	this.NodeType = nodeType
	return &this
}

// NewModelScanTriggerReqWithDefaults instantiates a new ModelScanTriggerReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelScanTriggerReqWithDefaults() *ModelScanTriggerReq {
	this := ModelScanTriggerReq{}
	return &this
}

// GetNodeId returns the NodeId field value
func (o *ModelScanTriggerReq) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelScanTriggerReq) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelScanTriggerReq) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeType returns the NodeType field value
func (o *ModelScanTriggerReq) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *ModelScanTriggerReq) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *ModelScanTriggerReq) SetNodeType(v string) {
	o.NodeType = v
}

func (o ModelScanTriggerReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelScanTriggerReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_type"] = o.NodeType
	return toSerialize, nil
}

type NullableModelScanTriggerReq struct {
	value *ModelScanTriggerReq
	isSet bool
}

func (v NullableModelScanTriggerReq) Get() *ModelScanTriggerReq {
	return v.value
}

func (v *NullableModelScanTriggerReq) Set(val *ModelScanTriggerReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelScanTriggerReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelScanTriggerReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelScanTriggerReq(val *ModelScanTriggerReq) *NullableModelScanTriggerReq {
	return &NullableModelScanTriggerReq{value: val, isSet: true}
}

func (v NullableModelScanTriggerReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelScanTriggerReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


