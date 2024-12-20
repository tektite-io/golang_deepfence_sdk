/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.2
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the DetailedNodeSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DetailedNodeSummary{}

// DetailedNodeSummary struct for DetailedNodeSummary
type DetailedNodeSummary struct {
	Adjacency []string `json:"adjacency,omitempty"`
	Id *string `json:"id,omitempty"`
	Ids []string `json:"ids,omitempty"`
	ImmediateParentId *string `json:"immediate_parent_id,omitempty"`
	Label *string `json:"label,omitempty"`
	Metadata *ReportMetadata `json:"metadata,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewDetailedNodeSummary instantiates a new DetailedNodeSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDetailedNodeSummary() *DetailedNodeSummary {
	this := DetailedNodeSummary{}
	return &this
}

// NewDetailedNodeSummaryWithDefaults instantiates a new DetailedNodeSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDetailedNodeSummaryWithDefaults() *DetailedNodeSummary {
	this := DetailedNodeSummary{}
	return &this
}

// GetAdjacency returns the Adjacency field value if set, zero value otherwise.
func (o *DetailedNodeSummary) GetAdjacency() []string {
	if o == nil || IsNil(o.Adjacency) {
		var ret []string
		return ret
	}
	return o.Adjacency
}

// GetAdjacencyOk returns a tuple with the Adjacency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedNodeSummary) GetAdjacencyOk() ([]string, bool) {
	if o == nil || IsNil(o.Adjacency) {
		return nil, false
	}
	return o.Adjacency, true
}

// HasAdjacency returns a boolean if a field has been set.
func (o *DetailedNodeSummary) HasAdjacency() bool {
	if o != nil && !IsNil(o.Adjacency) {
		return true
	}

	return false
}

// SetAdjacency gets a reference to the given []string and assigns it to the Adjacency field.
func (o *DetailedNodeSummary) SetAdjacency(v []string) {
	o.Adjacency = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DetailedNodeSummary) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedNodeSummary) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DetailedNodeSummary) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DetailedNodeSummary) SetId(v string) {
	o.Id = &v
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *DetailedNodeSummary) GetIds() []string {
	if o == nil || IsNil(o.Ids) {
		var ret []string
		return ret
	}
	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedNodeSummary) GetIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ids) {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *DetailedNodeSummary) HasIds() bool {
	if o != nil && !IsNil(o.Ids) {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *DetailedNodeSummary) SetIds(v []string) {
	o.Ids = v
}

// GetImmediateParentId returns the ImmediateParentId field value if set, zero value otherwise.
func (o *DetailedNodeSummary) GetImmediateParentId() string {
	if o == nil || IsNil(o.ImmediateParentId) {
		var ret string
		return ret
	}
	return *o.ImmediateParentId
}

// GetImmediateParentIdOk returns a tuple with the ImmediateParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedNodeSummary) GetImmediateParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImmediateParentId) {
		return nil, false
	}
	return o.ImmediateParentId, true
}

// HasImmediateParentId returns a boolean if a field has been set.
func (o *DetailedNodeSummary) HasImmediateParentId() bool {
	if o != nil && !IsNil(o.ImmediateParentId) {
		return true
	}

	return false
}

// SetImmediateParentId gets a reference to the given string and assigns it to the ImmediateParentId field.
func (o *DetailedNodeSummary) SetImmediateParentId(v string) {
	o.ImmediateParentId = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *DetailedNodeSummary) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedNodeSummary) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *DetailedNodeSummary) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *DetailedNodeSummary) SetLabel(v string) {
	o.Label = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DetailedNodeSummary) GetMetadata() ReportMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret ReportMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedNodeSummary) GetMetadataOk() (*ReportMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DetailedNodeSummary) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ReportMetadata and assigns it to the Metadata field.
func (o *DetailedNodeSummary) SetMetadata(v ReportMetadata) {
	o.Metadata = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DetailedNodeSummary) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DetailedNodeSummary) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DetailedNodeSummary) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DetailedNodeSummary) SetType(v string) {
	o.Type = &v
}

func (o DetailedNodeSummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DetailedNodeSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Adjacency) {
		toSerialize["adjacency"] = o.Adjacency
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ids) {
		toSerialize["ids"] = o.Ids
	}
	if !IsNil(o.ImmediateParentId) {
		toSerialize["immediate_parent_id"] = o.ImmediateParentId
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableDetailedNodeSummary struct {
	value *DetailedNodeSummary
	isSet bool
}

func (v NullableDetailedNodeSummary) Get() *DetailedNodeSummary {
	return v.value
}

func (v *NullableDetailedNodeSummary) Set(val *DetailedNodeSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableDetailedNodeSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableDetailedNodeSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetailedNodeSummary(val *DetailedNodeSummary) *NullableDetailedNodeSummary {
	return &NullableDetailedNodeSummary{value: val, isSet: true}
}

func (v NullableDetailedNodeSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetailedNodeSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


