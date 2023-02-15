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
	ComplianceScanStatus *string `json:"compliance_scan_status,omitempty"`
	CompliancesCount *int32 `json:"compliances_count,omitempty"`
	DockerLabels map[string]interface{} `json:"docker_labels"`
	HostName string `json:"host_name"`
	Image ModelContainerImage `json:"image"`
	MalwareScanStatus *string `json:"malware_scan_status,omitempty"`
	MalwaresCount *int32 `json:"malwares_count,omitempty"`
	Metadata map[string]interface{} `json:"metadata"`
	Metrics ModelComputeMetrics `json:"metrics"`
	Name string `json:"name"`
	NodeId string `json:"node_id"`
	Processes []ModelProcess `json:"processes"`
	SecretScanStatus *string `json:"secret_scan_status,omitempty"`
	SecretsCount *int32 `json:"secrets_count,omitempty"`
	VulnerabilitiesCount *int32 `json:"vulnerabilities_count,omitempty"`
	VulnerabilityScanStatus *string `json:"vulnerability_scan_status,omitempty"`
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

// GetComplianceScanStatus returns the ComplianceScanStatus field value if set, zero value otherwise.
func (o *ModelContainer) GetComplianceScanStatus() string {
	if o == nil || IsNil(o.ComplianceScanStatus) {
		var ret string
		return ret
	}
	return *o.ComplianceScanStatus
}

// GetComplianceScanStatusOk returns a tuple with the ComplianceScanStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetComplianceScanStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ComplianceScanStatus) {
		return nil, false
	}
	return o.ComplianceScanStatus, true
}

// HasComplianceScanStatus returns a boolean if a field has been set.
func (o *ModelContainer) HasComplianceScanStatus() bool {
	if o != nil && !IsNil(o.ComplianceScanStatus) {
		return true
	}

	return false
}

// SetComplianceScanStatus gets a reference to the given string and assigns it to the ComplianceScanStatus field.
func (o *ModelContainer) SetComplianceScanStatus(v string) {
	o.ComplianceScanStatus = &v
}

// GetCompliancesCount returns the CompliancesCount field value if set, zero value otherwise.
func (o *ModelContainer) GetCompliancesCount() int32 {
	if o == nil || IsNil(o.CompliancesCount) {
		var ret int32
		return ret
	}
	return *o.CompliancesCount
}

// GetCompliancesCountOk returns a tuple with the CompliancesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetCompliancesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.CompliancesCount) {
		return nil, false
	}
	return o.CompliancesCount, true
}

// HasCompliancesCount returns a boolean if a field has been set.
func (o *ModelContainer) HasCompliancesCount() bool {
	if o != nil && !IsNil(o.CompliancesCount) {
		return true
	}

	return false
}

// SetCompliancesCount gets a reference to the given int32 and assigns it to the CompliancesCount field.
func (o *ModelContainer) SetCompliancesCount(v int32) {
	o.CompliancesCount = &v
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

// GetMalwareScanStatus returns the MalwareScanStatus field value if set, zero value otherwise.
func (o *ModelContainer) GetMalwareScanStatus() string {
	if o == nil || IsNil(o.MalwareScanStatus) {
		var ret string
		return ret
	}
	return *o.MalwareScanStatus
}

// GetMalwareScanStatusOk returns a tuple with the MalwareScanStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetMalwareScanStatusOk() (*string, bool) {
	if o == nil || IsNil(o.MalwareScanStatus) {
		return nil, false
	}
	return o.MalwareScanStatus, true
}

// HasMalwareScanStatus returns a boolean if a field has been set.
func (o *ModelContainer) HasMalwareScanStatus() bool {
	if o != nil && !IsNil(o.MalwareScanStatus) {
		return true
	}

	return false
}

// SetMalwareScanStatus gets a reference to the given string and assigns it to the MalwareScanStatus field.
func (o *ModelContainer) SetMalwareScanStatus(v string) {
	o.MalwareScanStatus = &v
}

// GetMalwaresCount returns the MalwaresCount field value if set, zero value otherwise.
func (o *ModelContainer) GetMalwaresCount() int32 {
	if o == nil || IsNil(o.MalwaresCount) {
		var ret int32
		return ret
	}
	return *o.MalwaresCount
}

// GetMalwaresCountOk returns a tuple with the MalwaresCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetMalwaresCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MalwaresCount) {
		return nil, false
	}
	return o.MalwaresCount, true
}

// HasMalwaresCount returns a boolean if a field has been set.
func (o *ModelContainer) HasMalwaresCount() bool {
	if o != nil && !IsNil(o.MalwaresCount) {
		return true
	}

	return false
}

// SetMalwaresCount gets a reference to the given int32 and assigns it to the MalwaresCount field.
func (o *ModelContainer) SetMalwaresCount(v int32) {
	o.MalwaresCount = &v
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

// GetSecretScanStatus returns the SecretScanStatus field value if set, zero value otherwise.
func (o *ModelContainer) GetSecretScanStatus() string {
	if o == nil || IsNil(o.SecretScanStatus) {
		var ret string
		return ret
	}
	return *o.SecretScanStatus
}

// GetSecretScanStatusOk returns a tuple with the SecretScanStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetSecretScanStatusOk() (*string, bool) {
	if o == nil || IsNil(o.SecretScanStatus) {
		return nil, false
	}
	return o.SecretScanStatus, true
}

// HasSecretScanStatus returns a boolean if a field has been set.
func (o *ModelContainer) HasSecretScanStatus() bool {
	if o != nil && !IsNil(o.SecretScanStatus) {
		return true
	}

	return false
}

// SetSecretScanStatus gets a reference to the given string and assigns it to the SecretScanStatus field.
func (o *ModelContainer) SetSecretScanStatus(v string) {
	o.SecretScanStatus = &v
}

// GetSecretsCount returns the SecretsCount field value if set, zero value otherwise.
func (o *ModelContainer) GetSecretsCount() int32 {
	if o == nil || IsNil(o.SecretsCount) {
		var ret int32
		return ret
	}
	return *o.SecretsCount
}

// GetSecretsCountOk returns a tuple with the SecretsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetSecretsCountOk() (*int32, bool) {
	if o == nil || IsNil(o.SecretsCount) {
		return nil, false
	}
	return o.SecretsCount, true
}

// HasSecretsCount returns a boolean if a field has been set.
func (o *ModelContainer) HasSecretsCount() bool {
	if o != nil && !IsNil(o.SecretsCount) {
		return true
	}

	return false
}

// SetSecretsCount gets a reference to the given int32 and assigns it to the SecretsCount field.
func (o *ModelContainer) SetSecretsCount(v int32) {
	o.SecretsCount = &v
}

// GetVulnerabilitiesCount returns the VulnerabilitiesCount field value if set, zero value otherwise.
func (o *ModelContainer) GetVulnerabilitiesCount() int32 {
	if o == nil || IsNil(o.VulnerabilitiesCount) {
		var ret int32
		return ret
	}
	return *o.VulnerabilitiesCount
}

// GetVulnerabilitiesCountOk returns a tuple with the VulnerabilitiesCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetVulnerabilitiesCountOk() (*int32, bool) {
	if o == nil || IsNil(o.VulnerabilitiesCount) {
		return nil, false
	}
	return o.VulnerabilitiesCount, true
}

// HasVulnerabilitiesCount returns a boolean if a field has been set.
func (o *ModelContainer) HasVulnerabilitiesCount() bool {
	if o != nil && !IsNil(o.VulnerabilitiesCount) {
		return true
	}

	return false
}

// SetVulnerabilitiesCount gets a reference to the given int32 and assigns it to the VulnerabilitiesCount field.
func (o *ModelContainer) SetVulnerabilitiesCount(v int32) {
	o.VulnerabilitiesCount = &v
}

// GetVulnerabilityScanStatus returns the VulnerabilityScanStatus field value if set, zero value otherwise.
func (o *ModelContainer) GetVulnerabilityScanStatus() string {
	if o == nil || IsNil(o.VulnerabilityScanStatus) {
		var ret string
		return ret
	}
	return *o.VulnerabilityScanStatus
}

// GetVulnerabilityScanStatusOk returns a tuple with the VulnerabilityScanStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelContainer) GetVulnerabilityScanStatusOk() (*string, bool) {
	if o == nil || IsNil(o.VulnerabilityScanStatus) {
		return nil, false
	}
	return o.VulnerabilityScanStatus, true
}

// HasVulnerabilityScanStatus returns a boolean if a field has been set.
func (o *ModelContainer) HasVulnerabilityScanStatus() bool {
	if o != nil && !IsNil(o.VulnerabilityScanStatus) {
		return true
	}

	return false
}

// SetVulnerabilityScanStatus gets a reference to the given string and assigns it to the VulnerabilityScanStatus field.
func (o *ModelContainer) SetVulnerabilityScanStatus(v string) {
	o.VulnerabilityScanStatus = &v
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
	if !IsNil(o.ComplianceScanStatus) {
		toSerialize["compliance_scan_status"] = o.ComplianceScanStatus
	}
	if !IsNil(o.CompliancesCount) {
		toSerialize["compliances_count"] = o.CompliancesCount
	}
	toSerialize["docker_labels"] = o.DockerLabels
	toSerialize["host_name"] = o.HostName
	toSerialize["image"] = o.Image
	if !IsNil(o.MalwareScanStatus) {
		toSerialize["malware_scan_status"] = o.MalwareScanStatus
	}
	if !IsNil(o.MalwaresCount) {
		toSerialize["malwares_count"] = o.MalwaresCount
	}
	toSerialize["metadata"] = o.Metadata
	toSerialize["metrics"] = o.Metrics
	toSerialize["name"] = o.Name
	toSerialize["node_id"] = o.NodeId
	if o.Processes != nil {
		toSerialize["processes"] = o.Processes
	}
	if !IsNil(o.SecretScanStatus) {
		toSerialize["secret_scan_status"] = o.SecretScanStatus
	}
	if !IsNil(o.SecretsCount) {
		toSerialize["secrets_count"] = o.SecretsCount
	}
	if !IsNil(o.VulnerabilitiesCount) {
		toSerialize["vulnerabilities_count"] = o.VulnerabilitiesCount
	}
	if !IsNil(o.VulnerabilityScanStatus) {
		toSerialize["vulnerability_scan_status"] = o.VulnerabilityScanStatus
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


