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
	"fmt"
)

// checks if the SearchSearchNodeReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchSearchNodeReq{}

// SearchSearchNodeReq struct for SearchSearchNodeReq
type SearchSearchNodeReq struct {
	ExtendedNodeFilter *SearchSearchFilter `json:"extended_node_filter,omitempty"`
	NodeFilter SearchSearchFilter `json:"node_filter"`
	RelatedNodeFilter *SearchChainedSearchFilter `json:"related_node_filter,omitempty"`
	Window ModelFetchWindow `json:"window"`
}

type _SearchSearchNodeReq SearchSearchNodeReq

// NewSearchSearchNodeReq instantiates a new SearchSearchNodeReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchSearchNodeReq(nodeFilter SearchSearchFilter, window ModelFetchWindow) *SearchSearchNodeReq {
	this := SearchSearchNodeReq{}
	this.NodeFilter = nodeFilter
	this.Window = window
	return &this
}

// NewSearchSearchNodeReqWithDefaults instantiates a new SearchSearchNodeReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchSearchNodeReqWithDefaults() *SearchSearchNodeReq {
	this := SearchSearchNodeReq{}
	return &this
}

// GetExtendedNodeFilter returns the ExtendedNodeFilter field value if set, zero value otherwise.
func (o *SearchSearchNodeReq) GetExtendedNodeFilter() SearchSearchFilter {
	if o == nil || IsNil(o.ExtendedNodeFilter) {
		var ret SearchSearchFilter
		return ret
	}
	return *o.ExtendedNodeFilter
}

// GetExtendedNodeFilterOk returns a tuple with the ExtendedNodeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSearchNodeReq) GetExtendedNodeFilterOk() (*SearchSearchFilter, bool) {
	if o == nil || IsNil(o.ExtendedNodeFilter) {
		return nil, false
	}
	return o.ExtendedNodeFilter, true
}

// HasExtendedNodeFilter returns a boolean if a field has been set.
func (o *SearchSearchNodeReq) HasExtendedNodeFilter() bool {
	if o != nil && !IsNil(o.ExtendedNodeFilter) {
		return true
	}

	return false
}

// SetExtendedNodeFilter gets a reference to the given SearchSearchFilter and assigns it to the ExtendedNodeFilter field.
func (o *SearchSearchNodeReq) SetExtendedNodeFilter(v SearchSearchFilter) {
	o.ExtendedNodeFilter = &v
}

// GetNodeFilter returns the NodeFilter field value
func (o *SearchSearchNodeReq) GetNodeFilter() SearchSearchFilter {
	if o == nil {
		var ret SearchSearchFilter
		return ret
	}

	return o.NodeFilter
}

// GetNodeFilterOk returns a tuple with the NodeFilter field value
// and a boolean to check if the value has been set.
func (o *SearchSearchNodeReq) GetNodeFilterOk() (*SearchSearchFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeFilter, true
}

// SetNodeFilter sets field value
func (o *SearchSearchNodeReq) SetNodeFilter(v SearchSearchFilter) {
	o.NodeFilter = v
}

// GetRelatedNodeFilter returns the RelatedNodeFilter field value if set, zero value otherwise.
func (o *SearchSearchNodeReq) GetRelatedNodeFilter() SearchChainedSearchFilter {
	if o == nil || IsNil(o.RelatedNodeFilter) {
		var ret SearchChainedSearchFilter
		return ret
	}
	return *o.RelatedNodeFilter
}

// GetRelatedNodeFilterOk returns a tuple with the RelatedNodeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSearchNodeReq) GetRelatedNodeFilterOk() (*SearchChainedSearchFilter, bool) {
	if o == nil || IsNil(o.RelatedNodeFilter) {
		return nil, false
	}
	return o.RelatedNodeFilter, true
}

// HasRelatedNodeFilter returns a boolean if a field has been set.
func (o *SearchSearchNodeReq) HasRelatedNodeFilter() bool {
	if o != nil && !IsNil(o.RelatedNodeFilter) {
		return true
	}

	return false
}

// SetRelatedNodeFilter gets a reference to the given SearchChainedSearchFilter and assigns it to the RelatedNodeFilter field.
func (o *SearchSearchNodeReq) SetRelatedNodeFilter(v SearchChainedSearchFilter) {
	o.RelatedNodeFilter = &v
}

// GetWindow returns the Window field value
func (o *SearchSearchNodeReq) GetWindow() ModelFetchWindow {
	if o == nil {
		var ret ModelFetchWindow
		return ret
	}

	return o.Window
}

// GetWindowOk returns a tuple with the Window field value
// and a boolean to check if the value has been set.
func (o *SearchSearchNodeReq) GetWindowOk() (*ModelFetchWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Window, true
}

// SetWindow sets field value
func (o *SearchSearchNodeReq) SetWindow(v ModelFetchWindow) {
	o.Window = v
}

func (o SearchSearchNodeReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchSearchNodeReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ExtendedNodeFilter) {
		toSerialize["extended_node_filter"] = o.ExtendedNodeFilter
	}
	toSerialize["node_filter"] = o.NodeFilter
	if !IsNil(o.RelatedNodeFilter) {
		toSerialize["related_node_filter"] = o.RelatedNodeFilter
	}
	toSerialize["window"] = o.Window
	return toSerialize, nil
}

func (o *SearchSearchNodeReq) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"node_filter",
		"window",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSearchSearchNodeReq := _SearchSearchNodeReq{}

	err = json.Unmarshal(bytes, &varSearchSearchNodeReq)

	if err != nil {
		return err
	}

	*o = SearchSearchNodeReq(varSearchSearchNodeReq)

	return err
}

type NullableSearchSearchNodeReq struct {
	value *SearchSearchNodeReq
	isSet bool
}

func (v NullableSearchSearchNodeReq) Get() *SearchSearchNodeReq {
	return v.value
}

func (v *NullableSearchSearchNodeReq) Set(val *SearchSearchNodeReq) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSearchNodeReq) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSearchNodeReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSearchNodeReq(val *SearchSearchNodeReq) *NullableSearchSearchNodeReq {
	return &NullableSearchSearchNodeReq{value: val, isSet: true}
}

func (v NullableSearchSearchNodeReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSearchNodeReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


