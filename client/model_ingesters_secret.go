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
	"time"
)

// checks if the IngestersSecret type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestersSecret{}

// IngestersSecret struct for IngestersSecret
type IngestersSecret struct {
	Timestamp *time.Time `json:"@timestamp,omitempty"`
	ImageLayerId *string `json:"ImageLayerId,omitempty"`
	Match *IngestersSecretMatch `json:"Match,omitempty"`
	Rule *IngestersSecretRule `json:"Rule,omitempty"`
	Severity *IngestersSecretSeverity `json:"Severity,omitempty"`
	ContainerName *string `json:"container_name,omitempty"`
	HostName *string `json:"host_name,omitempty"`
	KubernetesClusterName *string `json:"kubernetes_cluster_name,omitempty"`
	Masked *string `json:"masked,omitempty"`
	NodeId *string `json:"node_id,omitempty"`
	NodeName *string `json:"node_name,omitempty"`
	NodeType *string `json:"node_type,omitempty"`
	ScanId *string `json:"scan_id,omitempty"`
}

// NewIngestersSecret instantiates a new IngestersSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestersSecret() *IngestersSecret {
	this := IngestersSecret{}
	return &this
}

// NewIngestersSecretWithDefaults instantiates a new IngestersSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestersSecretWithDefaults() *IngestersSecret {
	this := IngestersSecret{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *IngestersSecret) GetTimestamp() time.Time {
	if o == nil || isNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecret) GetTimestampOk() (*time.Time, bool) {
	if o == nil || isNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *IngestersSecret) HasTimestamp() bool {
	if o != nil && !isNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *IngestersSecret) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetImageLayerId returns the ImageLayerId field value if set, zero value otherwise.
func (o *IngestersSecret) GetImageLayerId() string {
	if o == nil || isNil(o.ImageLayerId) {
		var ret string
		return ret
	}
	return *o.ImageLayerId
}

// GetImageLayerIdOk returns a tuple with the ImageLayerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecret) GetImageLayerIdOk() (*string, bool) {
	if o == nil || isNil(o.ImageLayerId) {
		return nil, false
	}
	return o.ImageLayerId, true
}

// HasImageLayerId returns a boolean if a field has been set.
func (o *IngestersSecret) HasImageLayerId() bool {
	if o != nil && !isNil(o.ImageLayerId) {
		return true
	}

	return false
}

// SetImageLayerId gets a reference to the given string and assigns it to the ImageLayerId field.
func (o *IngestersSecret) SetImageLayerId(v string) {
	o.ImageLayerId = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *IngestersSecret) GetMatch() IngestersSecretMatch {
	if o == nil || isNil(o.Match) {
		var ret IngestersSecretMatch
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecret) GetMatchOk() (*IngestersSecretMatch, bool) {
	if o == nil || isNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *IngestersSecret) HasMatch() bool {
	if o != nil && !isNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given IngestersSecretMatch and assigns it to the Match field.
func (o *IngestersSecret) SetMatch(v IngestersSecretMatch) {
	o.Match = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *IngestersSecret) GetRule() IngestersSecretRule {
	if o == nil || isNil(o.Rule) {
		var ret IngestersSecretRule
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecret) GetRuleOk() (*IngestersSecretRule, bool) {
	if o == nil || isNil(o.Rule) {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *IngestersSecret) HasRule() bool {
	if o != nil && !isNil(o.Rule) {
		return true
	}

	return false
}

// SetRule gets a reference to the given IngestersSecretRule and assigns it to the Rule field.
func (o *IngestersSecret) SetRule(v IngestersSecretRule) {
	o.Rule = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *IngestersSecret) GetSeverity() IngestersSecretSeverity {
	if o == nil || isNil(o.Severity) {
		var ret IngestersSecretSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecret) GetSeverityOk() (*IngestersSecretSeverity, bool) {
	if o == nil || isNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *IngestersSecret) HasSeverity() bool {
	if o != nil && !isNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given IngestersSecretSeverity and assigns it to the Severity field.
func (o *IngestersSecret) SetSeverity(v IngestersSecretSeverity) {
	o.Severity = &v
}

// GetContainerName returns the ContainerName field value if set, zero value otherwise.
func (o *IngestersSecret) GetContainerName() string {
	if o == nil || isNil(o.ContainerName) {
		var ret string
		return ret
	}
	return *o.ContainerName
}

// GetContainerNameOk returns a tuple with the ContainerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecret) GetContainerNameOk() (*string, bool) {
	if o == nil || isNil(o.ContainerName) {
		return nil, false
	}
	return o.ContainerName, true
}

// HasContainerName returns a boolean if a field has been set.
func (o *IngestersSecret) HasContainerName() bool {
	if o != nil && !isNil(o.ContainerName) {
		return true
	}

	return false
}

// SetContainerName gets a reference to the given string and assigns it to the ContainerName field.
func (o *IngestersSecret) SetContainerName(v string) {
	o.ContainerName = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *IngestersSecret) GetHostName() string {
	if o == nil || isNil(o.HostName) {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecret) GetHostNameOk() (*string, bool) {
	if o == nil || isNil(o.HostName) {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *IngestersSecret) HasHostName() bool {
	if o != nil && !isNil(o.HostName) {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *IngestersSecret) SetHostName(v string) {
	o.HostName = &v
}

// GetKubernetesClusterName returns the KubernetesClusterName field value if set, zero value otherwise.
func (o *IngestersSecret) GetKubernetesClusterName() string {
	if o == nil || isNil(o.KubernetesClusterName) {
		var ret string
		return ret
	}
	return *o.KubernetesClusterName
}

// GetKubernetesClusterNameOk returns a tuple with the KubernetesClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecret) GetKubernetesClusterNameOk() (*string, bool) {
	if o == nil || isNil(o.KubernetesClusterName) {
		return nil, false
	}
	return o.KubernetesClusterName, true
}

// HasKubernetesClusterName returns a boolean if a field has been set.
func (o *IngestersSecret) HasKubernetesClusterName() bool {
	if o != nil && !isNil(o.KubernetesClusterName) {
		return true
	}

	return false
}

// SetKubernetesClusterName gets a reference to the given string and assigns it to the KubernetesClusterName field.
func (o *IngestersSecret) SetKubernetesClusterName(v string) {
	o.KubernetesClusterName = &v
}

// GetMasked returns the Masked field value if set, zero value otherwise.
func (o *IngestersSecret) GetMasked() string {
	if o == nil || isNil(o.Masked) {
		var ret string
		return ret
	}
	return *o.Masked
}

// GetMaskedOk returns a tuple with the Masked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecret) GetMaskedOk() (*string, bool) {
	if o == nil || isNil(o.Masked) {
		return nil, false
	}
	return o.Masked, true
}

// HasMasked returns a boolean if a field has been set.
func (o *IngestersSecret) HasMasked() bool {
	if o != nil && !isNil(o.Masked) {
		return true
	}

	return false
}

// SetMasked gets a reference to the given string and assigns it to the Masked field.
func (o *IngestersSecret) SetMasked(v string) {
	o.Masked = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *IngestersSecret) GetNodeId() string {
	if o == nil || isNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecret) GetNodeIdOk() (*string, bool) {
	if o == nil || isNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *IngestersSecret) HasNodeId() bool {
	if o != nil && !isNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *IngestersSecret) SetNodeId(v string) {
	o.NodeId = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *IngestersSecret) GetNodeName() string {
	if o == nil || isNil(o.NodeName) {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecret) GetNodeNameOk() (*string, bool) {
	if o == nil || isNil(o.NodeName) {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *IngestersSecret) HasNodeName() bool {
	if o != nil && !isNil(o.NodeName) {
		return true
	}

	return false
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *IngestersSecret) SetNodeName(v string) {
	o.NodeName = &v
}

// GetNodeType returns the NodeType field value if set, zero value otherwise.
func (o *IngestersSecret) GetNodeType() string {
	if o == nil || isNil(o.NodeType) {
		var ret string
		return ret
	}
	return *o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecret) GetNodeTypeOk() (*string, bool) {
	if o == nil || isNil(o.NodeType) {
		return nil, false
	}
	return o.NodeType, true
}

// HasNodeType returns a boolean if a field has been set.
func (o *IngestersSecret) HasNodeType() bool {
	if o != nil && !isNil(o.NodeType) {
		return true
	}

	return false
}

// SetNodeType gets a reference to the given string and assigns it to the NodeType field.
func (o *IngestersSecret) SetNodeType(v string) {
	o.NodeType = &v
}

// GetScanId returns the ScanId field value if set, zero value otherwise.
func (o *IngestersSecret) GetScanId() string {
	if o == nil || isNil(o.ScanId) {
		var ret string
		return ret
	}
	return *o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecret) GetScanIdOk() (*string, bool) {
	if o == nil || isNil(o.ScanId) {
		return nil, false
	}
	return o.ScanId, true
}

// HasScanId returns a boolean if a field has been set.
func (o *IngestersSecret) HasScanId() bool {
	if o != nil && !isNil(o.ScanId) {
		return true
	}

	return false
}

// SetScanId gets a reference to the given string and assigns it to the ScanId field.
func (o *IngestersSecret) SetScanId(v string) {
	o.ScanId = &v
}

func (o IngestersSecret) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestersSecret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Timestamp) {
		toSerialize["@timestamp"] = o.Timestamp
	}
	if !isNil(o.ImageLayerId) {
		toSerialize["ImageLayerId"] = o.ImageLayerId
	}
	if !isNil(o.Match) {
		toSerialize["Match"] = o.Match
	}
	if !isNil(o.Rule) {
		toSerialize["Rule"] = o.Rule
	}
	if !isNil(o.Severity) {
		toSerialize["Severity"] = o.Severity
	}
	if !isNil(o.ContainerName) {
		toSerialize["container_name"] = o.ContainerName
	}
	if !isNil(o.HostName) {
		toSerialize["host_name"] = o.HostName
	}
	if !isNil(o.KubernetesClusterName) {
		toSerialize["kubernetes_cluster_name"] = o.KubernetesClusterName
	}
	if !isNil(o.Masked) {
		toSerialize["masked"] = o.Masked
	}
	if !isNil(o.NodeId) {
		toSerialize["node_id"] = o.NodeId
	}
	if !isNil(o.NodeName) {
		toSerialize["node_name"] = o.NodeName
	}
	if !isNil(o.NodeType) {
		toSerialize["node_type"] = o.NodeType
	}
	if !isNil(o.ScanId) {
		toSerialize["scan_id"] = o.ScanId
	}
	return toSerialize, nil
}

type NullableIngestersSecret struct {
	value *IngestersSecret
	isSet bool
}

func (v NullableIngestersSecret) Get() *IngestersSecret {
	return v.value
}

func (v *NullableIngestersSecret) Set(val *IngestersSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestersSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestersSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestersSecret(val *IngestersSecret) *NullableIngestersSecret {
	return &NullableIngestersSecret{value: val, isSet: true}
}

func (v NullableIngestersSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestersSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


