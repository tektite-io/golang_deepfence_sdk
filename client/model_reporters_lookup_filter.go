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

// checks if the ReportersLookupFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportersLookupFilter{}

// ReportersLookupFilter struct for ReportersLookupFilter
type ReportersLookupFilter struct {
	InFieldFilter []string `json:"in_field_filter"`
	NodeIds []string `json:"node_ids"`
}

// NewReportersLookupFilter instantiates a new ReportersLookupFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportersLookupFilter(inFieldFilter []string, nodeIds []string) *ReportersLookupFilter {
	this := ReportersLookupFilter{}
	this.InFieldFilter = inFieldFilter
	this.NodeIds = nodeIds
	return &this
}

// NewReportersLookupFilterWithDefaults instantiates a new ReportersLookupFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportersLookupFilterWithDefaults() *ReportersLookupFilter {
	this := ReportersLookupFilter{}
	return &this
}

// GetInFieldFilter returns the InFieldFilter field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ReportersLookupFilter) GetInFieldFilter() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.InFieldFilter
}

// GetInFieldFilterOk returns a tuple with the InFieldFilter field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportersLookupFilter) GetInFieldFilterOk() ([]string, bool) {
	if o == nil || isNil(o.InFieldFilter) {
		return nil, false
	}
	return o.InFieldFilter, true
}

// SetInFieldFilter sets field value
func (o *ReportersLookupFilter) SetInFieldFilter(v []string) {
	o.InFieldFilter = v
}

// GetNodeIds returns the NodeIds field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ReportersLookupFilter) GetNodeIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NodeIds
}

// GetNodeIdsOk returns a tuple with the NodeIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportersLookupFilter) GetNodeIdsOk() ([]string, bool) {
	if o == nil || isNil(o.NodeIds) {
		return nil, false
	}
	return o.NodeIds, true
}

// SetNodeIds sets field value
func (o *ReportersLookupFilter) SetNodeIds(v []string) {
	o.NodeIds = v
}

func (o ReportersLookupFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportersLookupFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.InFieldFilter != nil {
		toSerialize["in_field_filter"] = o.InFieldFilter
	}
	if o.NodeIds != nil {
		toSerialize["node_ids"] = o.NodeIds
	}
	return toSerialize, nil
}

type NullableReportersLookupFilter struct {
	value *ReportersLookupFilter
	isSet bool
}

func (v NullableReportersLookupFilter) Get() *ReportersLookupFilter {
	return v.value
}

func (v *NullableReportersLookupFilter) Set(val *ReportersLookupFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableReportersLookupFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableReportersLookupFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportersLookupFilter(val *ReportersLookupFilter) *NullableReportersLookupFilter {
	return &NullableReportersLookupFilter{value: val, isSet: true}
}

func (v NullableReportersLookupFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportersLookupFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


