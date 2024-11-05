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

// checks if the SearchChainedSearchFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchChainedSearchFilter{}

// SearchChainedSearchFilter struct for SearchChainedSearchFilter
type SearchChainedSearchFilter struct {
	NextFilter *SearchChainedSearchFilter `json:"next_filter,omitempty"`
	NodeFilter SearchSearchFilter `json:"node_filter"`
	RelationShip string `json:"relation_ship"`
}

type _SearchChainedSearchFilter SearchChainedSearchFilter

// NewSearchChainedSearchFilter instantiates a new SearchChainedSearchFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchChainedSearchFilter(nodeFilter SearchSearchFilter, relationShip string) *SearchChainedSearchFilter {
	this := SearchChainedSearchFilter{}
	this.NodeFilter = nodeFilter
	this.RelationShip = relationShip
	return &this
}

// NewSearchChainedSearchFilterWithDefaults instantiates a new SearchChainedSearchFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchChainedSearchFilterWithDefaults() *SearchChainedSearchFilter {
	this := SearchChainedSearchFilter{}
	return &this
}

// GetNextFilter returns the NextFilter field value if set, zero value otherwise.
func (o *SearchChainedSearchFilter) GetNextFilter() SearchChainedSearchFilter {
	if o == nil || IsNil(o.NextFilter) {
		var ret SearchChainedSearchFilter
		return ret
	}
	return *o.NextFilter
}

// GetNextFilterOk returns a tuple with the NextFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchChainedSearchFilter) GetNextFilterOk() (*SearchChainedSearchFilter, bool) {
	if o == nil || IsNil(o.NextFilter) {
		return nil, false
	}
	return o.NextFilter, true
}

// HasNextFilter returns a boolean if a field has been set.
func (o *SearchChainedSearchFilter) HasNextFilter() bool {
	if o != nil && !IsNil(o.NextFilter) {
		return true
	}

	return false
}

// SetNextFilter gets a reference to the given SearchChainedSearchFilter and assigns it to the NextFilter field.
func (o *SearchChainedSearchFilter) SetNextFilter(v SearchChainedSearchFilter) {
	o.NextFilter = &v
}

// GetNodeFilter returns the NodeFilter field value
func (o *SearchChainedSearchFilter) GetNodeFilter() SearchSearchFilter {
	if o == nil {
		var ret SearchSearchFilter
		return ret
	}

	return o.NodeFilter
}

// GetNodeFilterOk returns a tuple with the NodeFilter field value
// and a boolean to check if the value has been set.
func (o *SearchChainedSearchFilter) GetNodeFilterOk() (*SearchSearchFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeFilter, true
}

// SetNodeFilter sets field value
func (o *SearchChainedSearchFilter) SetNodeFilter(v SearchSearchFilter) {
	o.NodeFilter = v
}

// GetRelationShip returns the RelationShip field value
func (o *SearchChainedSearchFilter) GetRelationShip() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RelationShip
}

// GetRelationShipOk returns a tuple with the RelationShip field value
// and a boolean to check if the value has been set.
func (o *SearchChainedSearchFilter) GetRelationShipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelationShip, true
}

// SetRelationShip sets field value
func (o *SearchChainedSearchFilter) SetRelationShip(v string) {
	o.RelationShip = v
}

func (o SearchChainedSearchFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchChainedSearchFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NextFilter) {
		toSerialize["next_filter"] = o.NextFilter
	}
	toSerialize["node_filter"] = o.NodeFilter
	toSerialize["relation_ship"] = o.RelationShip
	return toSerialize, nil
}

func (o *SearchChainedSearchFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"node_filter",
		"relation_ship",
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

	varSearchChainedSearchFilter := _SearchChainedSearchFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchChainedSearchFilter)

	if err != nil {
		return err
	}

	*o = SearchChainedSearchFilter(varSearchChainedSearchFilter)

	return err
}

type NullableSearchChainedSearchFilter struct {
	value *SearchChainedSearchFilter
	isSet bool
}

func (v NullableSearchChainedSearchFilter) Get() *SearchChainedSearchFilter {
	return v.value
}

func (v *NullableSearchChainedSearchFilter) Set(val *SearchChainedSearchFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchChainedSearchFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchChainedSearchFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchChainedSearchFilter(val *SearchChainedSearchFilter) *NullableSearchChainedSearchFilter {
	return &NullableSearchChainedSearchFilter{value: val, isSet: true}
}

func (v NullableSearchChainedSearchFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchChainedSearchFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


