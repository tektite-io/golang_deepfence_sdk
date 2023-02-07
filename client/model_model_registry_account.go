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

// checks if the ModelRegistryAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelRegistryAccount{}

// ModelRegistryAccount struct for ModelRegistryAccount
type ModelRegistryAccount struct {
	ContainerImages []ModelContainerImage `json:"container_images"`
	HostName string `json:"host_name"`
	NodeId string `json:"node_id"`
}

// NewModelRegistryAccount instantiates a new ModelRegistryAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelRegistryAccount(containerImages []ModelContainerImage, hostName string, nodeId string) *ModelRegistryAccount {
	this := ModelRegistryAccount{}
	this.ContainerImages = containerImages
	this.HostName = hostName
	this.NodeId = nodeId
	return &this
}

// NewModelRegistryAccountWithDefaults instantiates a new ModelRegistryAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelRegistryAccountWithDefaults() *ModelRegistryAccount {
	this := ModelRegistryAccount{}
	return &this
}

// GetContainerImages returns the ContainerImages field value
// If the value is explicit nil, the zero value for []ModelContainerImage will be returned
func (o *ModelRegistryAccount) GetContainerImages() []ModelContainerImage {
	if o == nil {
		var ret []ModelContainerImage
		return ret
	}

	return o.ContainerImages
}

// GetContainerImagesOk returns a tuple with the ContainerImages field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelRegistryAccount) GetContainerImagesOk() ([]ModelContainerImage, bool) {
	if o == nil || isNil(o.ContainerImages) {
		return nil, false
	}
	return o.ContainerImages, true
}

// SetContainerImages sets field value
func (o *ModelRegistryAccount) SetContainerImages(v []ModelContainerImage) {
	o.ContainerImages = v
}

// GetHostName returns the HostName field value
func (o *ModelRegistryAccount) GetHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value
// and a boolean to check if the value has been set.
func (o *ModelRegistryAccount) GetHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostName, true
}

// SetHostName sets field value
func (o *ModelRegistryAccount) SetHostName(v string) {
	o.HostName = v
}

// GetNodeId returns the NodeId field value
func (o *ModelRegistryAccount) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelRegistryAccount) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelRegistryAccount) SetNodeId(v string) {
	o.NodeId = v
}

func (o ModelRegistryAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelRegistryAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContainerImages != nil {
		toSerialize["container_images"] = o.ContainerImages
	}
	toSerialize["host_name"] = o.HostName
	toSerialize["node_id"] = o.NodeId
	return toSerialize, nil
}

type NullableModelRegistryAccount struct {
	value *ModelRegistryAccount
	isSet bool
}

func (v NullableModelRegistryAccount) Get() *ModelRegistryAccount {
	return v.value
}

func (v *NullableModelRegistryAccount) Set(val *ModelRegistryAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRegistryAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRegistryAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRegistryAccount(val *ModelRegistryAccount) *NullableModelRegistryAccount {
	return &NullableModelRegistryAccount{value: val, isSet: true}
}

func (v NullableModelRegistryAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRegistryAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

