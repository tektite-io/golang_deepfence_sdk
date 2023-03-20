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

// checks if the GraphProviderThreatGraph type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphProviderThreatGraph{}

// GraphProviderThreatGraph struct for GraphProviderThreatGraph
type GraphProviderThreatGraph struct {
	CloudComplianceCount int32 `json:"cloud_compliance_count"`
	ComplianceCount int32 `json:"compliance_count"`
	Resources []GraphThreatNodeInfo `json:"resources"`
	SecretsCount int32 `json:"secrets_count"`
	VulnerabilityCount int32 `json:"vulnerability_count"`
}

// NewGraphProviderThreatGraph instantiates a new GraphProviderThreatGraph object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphProviderThreatGraph(cloudComplianceCount int32, complianceCount int32, resources []GraphThreatNodeInfo, secretsCount int32, vulnerabilityCount int32) *GraphProviderThreatGraph {
	this := GraphProviderThreatGraph{}
	this.CloudComplianceCount = cloudComplianceCount
	this.ComplianceCount = complianceCount
	this.Resources = resources
	this.SecretsCount = secretsCount
	this.VulnerabilityCount = vulnerabilityCount
	return &this
}

// NewGraphProviderThreatGraphWithDefaults instantiates a new GraphProviderThreatGraph object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphProviderThreatGraphWithDefaults() *GraphProviderThreatGraph {
	this := GraphProviderThreatGraph{}
	return &this
}

// GetCloudComplianceCount returns the CloudComplianceCount field value
func (o *GraphProviderThreatGraph) GetCloudComplianceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CloudComplianceCount
}

// GetCloudComplianceCountOk returns a tuple with the CloudComplianceCount field value
// and a boolean to check if the value has been set.
func (o *GraphProviderThreatGraph) GetCloudComplianceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudComplianceCount, true
}

// SetCloudComplianceCount sets field value
func (o *GraphProviderThreatGraph) SetCloudComplianceCount(v int32) {
	o.CloudComplianceCount = v
}

// GetComplianceCount returns the ComplianceCount field value
func (o *GraphProviderThreatGraph) GetComplianceCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ComplianceCount
}

// GetComplianceCountOk returns a tuple with the ComplianceCount field value
// and a boolean to check if the value has been set.
func (o *GraphProviderThreatGraph) GetComplianceCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ComplianceCount, true
}

// SetComplianceCount sets field value
func (o *GraphProviderThreatGraph) SetComplianceCount(v int32) {
	o.ComplianceCount = v
}

// GetResources returns the Resources field value
// If the value is explicit nil, the zero value for []GraphThreatNodeInfo will be returned
func (o *GraphProviderThreatGraph) GetResources() []GraphThreatNodeInfo {
	if o == nil {
		var ret []GraphThreatNodeInfo
		return ret
	}

	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphProviderThreatGraph) GetResourcesOk() ([]GraphThreatNodeInfo, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// SetResources sets field value
func (o *GraphProviderThreatGraph) SetResources(v []GraphThreatNodeInfo) {
	o.Resources = v
}

// GetSecretsCount returns the SecretsCount field value
func (o *GraphProviderThreatGraph) GetSecretsCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SecretsCount
}

// GetSecretsCountOk returns a tuple with the SecretsCount field value
// and a boolean to check if the value has been set.
func (o *GraphProviderThreatGraph) GetSecretsCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecretsCount, true
}

// SetSecretsCount sets field value
func (o *GraphProviderThreatGraph) SetSecretsCount(v int32) {
	o.SecretsCount = v
}

// GetVulnerabilityCount returns the VulnerabilityCount field value
func (o *GraphProviderThreatGraph) GetVulnerabilityCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VulnerabilityCount
}

// GetVulnerabilityCountOk returns a tuple with the VulnerabilityCount field value
// and a boolean to check if the value has been set.
func (o *GraphProviderThreatGraph) GetVulnerabilityCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VulnerabilityCount, true
}

// SetVulnerabilityCount sets field value
func (o *GraphProviderThreatGraph) SetVulnerabilityCount(v int32) {
	o.VulnerabilityCount = v
}

func (o GraphProviderThreatGraph) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphProviderThreatGraph) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cloud_compliance_count"] = o.CloudComplianceCount
	toSerialize["compliance_count"] = o.ComplianceCount
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	toSerialize["secrets_count"] = o.SecretsCount
	toSerialize["vulnerability_count"] = o.VulnerabilityCount
	return toSerialize, nil
}

type NullableGraphProviderThreatGraph struct {
	value *GraphProviderThreatGraph
	isSet bool
}

func (v NullableGraphProviderThreatGraph) Get() *GraphProviderThreatGraph {
	return v.value
}

func (v *NullableGraphProviderThreatGraph) Set(val *GraphProviderThreatGraph) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphProviderThreatGraph) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphProviderThreatGraph) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphProviderThreatGraph(val *GraphProviderThreatGraph) *NullableGraphProviderThreatGraph {
	return &NullableGraphProviderThreatGraph{value: val, isSet: true}
}

func (v NullableGraphProviderThreatGraph) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphProviderThreatGraph) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


