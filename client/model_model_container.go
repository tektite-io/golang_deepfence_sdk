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

// checks if the ModelContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelContainer{}

// ModelContainer struct for ModelContainer
type ModelContainer struct {
	DockerLabels map[string]interface{} `json:"docker_labels"`
	HostName string `json:"host_name"`
	Image ModelContainerImage `json:"image"`
	Metadata map[string]interface{} `json:"metadata"`
	Metrics ModelComputeMetrics `json:"metrics"`
	Name string `json:"name"`
	NodeId string `json:"node_id"`
	Processes []ModelProcess `json:"processes"`
}

// NewModelContainer instantiates a new ModelContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelContainer(dockerLabels map[string]interface{}, hostName string, image ModelContainerImage, metadata map[string]interface{}, metrics ModelComputeMetrics, name string, nodeId string, processes []ModelProcess) *ModelContainer {
	this := ModelContainer{}
	this.DockerLabels = dockerLabels
	this.HostName = hostName
	this.Image = image
	this.Metadata = metadata
	this.Metrics = metrics
	this.Name = name
	this.NodeId = nodeId
	this.Processes = processes
	return &this
}

// NewModelContainerWithDefaults instantiates a new ModelContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelContainerWithDefaults() *ModelContainer {
	this := ModelContainer{}
	return &this
}

// GetDockerLabels returns the DockerLabels field value
func (o *ModelContainer) GetDockerLabels() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.DockerLabels
}

// GetDockerLabelsOk returns a tuple with the DockerLabels field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetDockerLabelsOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.DockerLabels, true
}

// SetDockerLabels sets field value
func (o *ModelContainer) SetDockerLabels(v map[string]interface{}) {
	o.DockerLabels = v
}

// GetHostName returns the HostName field value
func (o *ModelContainer) GetHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostName, true
}

// SetHostName sets field value
func (o *ModelContainer) SetHostName(v string) {
	o.HostName = v
}

// GetImage returns the Image field value
func (o *ModelContainer) GetImage() ModelContainerImage {
	if o == nil {
		var ret ModelContainerImage
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetImageOk() (*ModelContainerImage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *ModelContainer) SetImage(v ModelContainerImage) {
	o.Image = v
}

// GetMetadata returns the Metadata field value
func (o *ModelContainer) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *ModelContainer) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMetrics returns the Metrics field value
func (o *ModelContainer) GetMetrics() ModelComputeMetrics {
	if o == nil {
		var ret ModelComputeMetrics
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetMetricsOk() (*ModelComputeMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value
func (o *ModelContainer) SetMetrics(v ModelComputeMetrics) {
	o.Metrics = v
}

// GetName returns the Name field value
func (o *ModelContainer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModelContainer) SetName(v string) {
	o.Name = v
}

// GetNodeId returns the NodeId field value
func (o *ModelContainer) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelContainer) SetNodeId(v string) {
	o.NodeId = v
}

// GetProcesses returns the Processes field value
// If the value is explicit nil, the zero value for []ModelProcess will be returned
func (o *ModelContainer) GetProcesses() []ModelProcess {
	if o == nil {
		var ret []ModelProcess
		return ret
	}

	return o.Processes
}

// GetProcessesOk returns a tuple with the Processes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelContainer) GetProcessesOk() ([]ModelProcess, bool) {
	if o == nil || IsNil(o.Processes) {
		return nil, false
	}
	return o.Processes, true
}

// SetProcesses sets field value
func (o *ModelContainer) SetProcesses(v []ModelProcess) {
	o.Processes = v
}

func (o ModelContainer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["docker_labels"] = o.DockerLabels
	toSerialize["host_name"] = o.HostName
	toSerialize["image"] = o.Image
	toSerialize["metadata"] = o.Metadata
	toSerialize["metrics"] = o.Metrics
	toSerialize["name"] = o.Name
	toSerialize["node_id"] = o.NodeId
	if o.Processes != nil {
		toSerialize["processes"] = o.Processes
	}
	return toSerialize, nil
}

type NullableModelContainer struct {
	value *ModelContainer
	isSet bool
}

func (v NullableModelContainer) Get() *ModelContainer {
	return v.value
}

func (v *NullableModelContainer) Set(val *ModelContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableModelContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableModelContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelContainer(val *ModelContainer) *NullableModelContainer {
	return &NullableModelContainer{value: val, isSet: true}
}

func (v NullableModelContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


