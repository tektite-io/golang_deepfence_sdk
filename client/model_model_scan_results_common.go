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
	"bytes"
	"fmt"
)

// checks if the ModelScanResultsCommon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelScanResultsCommon{}

// ModelScanResultsCommon struct for ModelScanResultsCommon
type ModelScanResultsCommon struct {
	CloudAccountId string `json:"cloud_account_id"`
	CreatedAt int64 `json:"created_at"`
	DockerContainerName string `json:"docker_container_name"`
	DockerImageName string `json:"docker_image_name"`
	HostName string `json:"host_name"`
	KubernetesClusterName string `json:"kubernetes_cluster_name"`
	NodeId string `json:"node_id"`
	NodeName string `json:"node_name"`
	NodeType string `json:"node_type"`
	ScanId string `json:"scan_id"`
	UpdatedAt int64 `json:"updated_at"`
}

type _ModelScanResultsCommon ModelScanResultsCommon

// NewModelScanResultsCommon instantiates a new ModelScanResultsCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelScanResultsCommon(cloudAccountId string, createdAt int64, dockerContainerName string, dockerImageName string, hostName string, kubernetesClusterName string, nodeId string, nodeName string, nodeType string, scanId string, updatedAt int64) *ModelScanResultsCommon {
	this := ModelScanResultsCommon{}
	this.CloudAccountId = cloudAccountId
	this.CreatedAt = createdAt
	this.DockerContainerName = dockerContainerName
	this.DockerImageName = dockerImageName
	this.HostName = hostName
	this.KubernetesClusterName = kubernetesClusterName
	this.NodeId = nodeId
	this.NodeName = nodeName
	this.NodeType = nodeType
	this.ScanId = scanId
	this.UpdatedAt = updatedAt
	return &this
}

// NewModelScanResultsCommonWithDefaults instantiates a new ModelScanResultsCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelScanResultsCommonWithDefaults() *ModelScanResultsCommon {
	this := ModelScanResultsCommon{}
	return &this
}

// GetCloudAccountId returns the CloudAccountId field value
func (o *ModelScanResultsCommon) GetCloudAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloudAccountId
}

// GetCloudAccountIdOk returns a tuple with the CloudAccountId field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsCommon) GetCloudAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudAccountId, true
}

// SetCloudAccountId sets field value
func (o *ModelScanResultsCommon) SetCloudAccountId(v string) {
	o.CloudAccountId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ModelScanResultsCommon) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsCommon) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ModelScanResultsCommon) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetDockerContainerName returns the DockerContainerName field value
func (o *ModelScanResultsCommon) GetDockerContainerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerContainerName
}

// GetDockerContainerNameOk returns a tuple with the DockerContainerName field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsCommon) GetDockerContainerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerContainerName, true
}

// SetDockerContainerName sets field value
func (o *ModelScanResultsCommon) SetDockerContainerName(v string) {
	o.DockerContainerName = v
}

// GetDockerImageName returns the DockerImageName field value
func (o *ModelScanResultsCommon) GetDockerImageName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DockerImageName
}

// GetDockerImageNameOk returns a tuple with the DockerImageName field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsCommon) GetDockerImageNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DockerImageName, true
}

// SetDockerImageName sets field value
func (o *ModelScanResultsCommon) SetDockerImageName(v string) {
	o.DockerImageName = v
}

// GetHostName returns the HostName field value
func (o *ModelScanResultsCommon) GetHostName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsCommon) GetHostNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostName, true
}

// SetHostName sets field value
func (o *ModelScanResultsCommon) SetHostName(v string) {
	o.HostName = v
}

// GetKubernetesClusterName returns the KubernetesClusterName field value
func (o *ModelScanResultsCommon) GetKubernetesClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KubernetesClusterName
}

// GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsCommon) GetKubernetesClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KubernetesClusterName, true
}

// SetKubernetesClusterName sets field value
func (o *ModelScanResultsCommon) SetKubernetesClusterName(v string) {
	o.KubernetesClusterName = v
}

// GetNodeId returns the NodeId field value
func (o *ModelScanResultsCommon) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsCommon) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelScanResultsCommon) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeName returns the NodeName field value
func (o *ModelScanResultsCommon) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsCommon) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *ModelScanResultsCommon) SetNodeName(v string) {
	o.NodeName = v
}

// GetNodeType returns the NodeType field value
func (o *ModelScanResultsCommon) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsCommon) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *ModelScanResultsCommon) SetNodeType(v string) {
	o.NodeType = v
}

// GetScanId returns the ScanId field value
func (o *ModelScanResultsCommon) GetScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsCommon) GetScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanId, true
}

// SetScanId sets field value
func (o *ModelScanResultsCommon) SetScanId(v string) {
	o.ScanId = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ModelScanResultsCommon) GetUpdatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsCommon) GetUpdatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ModelScanResultsCommon) SetUpdatedAt(v int64) {
	o.UpdatedAt = v
}

func (o ModelScanResultsCommon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelScanResultsCommon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloud_account_id"] = o.CloudAccountId
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["docker_container_name"] = o.DockerContainerName
	toSerialize["docker_image_name"] = o.DockerImageName
	toSerialize["host_name"] = o.HostName
	toSerialize["kubernetes_cluster_name"] = o.KubernetesClusterName
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_name"] = o.NodeName
	toSerialize["node_type"] = o.NodeType
	toSerialize["scan_id"] = o.ScanId
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *ModelScanResultsCommon) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloud_account_id",
		"created_at",
		"docker_container_name",
		"docker_image_name",
		"host_name",
		"kubernetes_cluster_name",
		"node_id",
		"node_name",
		"node_type",
		"scan_id",
		"updated_at",
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

	varModelScanResultsCommon := _ModelScanResultsCommon{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelScanResultsCommon)

	if err != nil {
		return err
	}

	*o = ModelScanResultsCommon(varModelScanResultsCommon)

	return err
}

type NullableModelScanResultsCommon struct {
	value *ModelScanResultsCommon
	isSet bool
}

func (v NullableModelScanResultsCommon) Get() *ModelScanResultsCommon {
	return v.value
}

func (v *NullableModelScanResultsCommon) Set(val *ModelScanResultsCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableModelScanResultsCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableModelScanResultsCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelScanResultsCommon(val *ModelScanResultsCommon) *NullableModelScanResultsCommon {
	return &NullableModelScanResultsCommon{value: val, isSet: true}
}

func (v NullableModelScanResultsCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelScanResultsCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


