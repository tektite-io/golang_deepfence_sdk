/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.2.1
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ReportersFieldsFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportersFieldsFilters{}

// ReportersFieldsFilters struct for ReportersFieldsFilters
type ReportersFieldsFilters struct {
	CompareFilter []ReportersCompareFilter `json:"compare_filter"`
	ContainsFilter ReportersContainsFilter `json:"contains_filter"`
	ContainsInArrayFilter *ReportersContainsFilter `json:"contains_in_array_filter,omitempty"`
	MatchFilter ReportersMatchFilter `json:"match_filter"`
	MatchInArrayFilter *ReportersMatchFilter `json:"match_in_array_filter,omitempty"`
	NotContainsFilter *ReportersContainsFilter `json:"not_contains_filter,omitempty"`
	OrderFilter ReportersOrderFilter `json:"order_filter"`
}

type _ReportersFieldsFilters ReportersFieldsFilters

// NewReportersFieldsFilters instantiates a new ReportersFieldsFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportersFieldsFilters(compareFilter []ReportersCompareFilter, containsFilter ReportersContainsFilter, matchFilter ReportersMatchFilter, orderFilter ReportersOrderFilter) *ReportersFieldsFilters {
	this := ReportersFieldsFilters{}
	this.CompareFilter = compareFilter
	this.ContainsFilter = containsFilter
	this.MatchFilter = matchFilter
	this.OrderFilter = orderFilter
	return &this
}

// NewReportersFieldsFiltersWithDefaults instantiates a new ReportersFieldsFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportersFieldsFiltersWithDefaults() *ReportersFieldsFilters {
	this := ReportersFieldsFilters{}
	return &this
}

// GetCompareFilter returns the CompareFilter field value
// If the value is explicit nil, the zero value for []ReportersCompareFilter will be returned
func (o *ReportersFieldsFilters) GetCompareFilter() []ReportersCompareFilter {
	if o == nil {
		var ret []ReportersCompareFilter
		return ret
	}

	return o.CompareFilter
}

// GetCompareFilterOk returns a tuple with the CompareFilter field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportersFieldsFilters) GetCompareFilterOk() ([]ReportersCompareFilter, bool) {
	if o == nil || IsNil(o.CompareFilter) {
		return nil, false
	}
	return o.CompareFilter, true
}

// SetCompareFilter sets field value
func (o *ReportersFieldsFilters) SetCompareFilter(v []ReportersCompareFilter) {
	o.CompareFilter = v
}

// GetContainsFilter returns the ContainsFilter field value
func (o *ReportersFieldsFilters) GetContainsFilter() ReportersContainsFilter {
	if o == nil {
		var ret ReportersContainsFilter
		return ret
	}

	return o.ContainsFilter
}

// GetContainsFilterOk returns a tuple with the ContainsFilter field value
// and a boolean to check if the value has been set.
func (o *ReportersFieldsFilters) GetContainsFilterOk() (*ReportersContainsFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContainsFilter, true
}

// SetContainsFilter sets field value
func (o *ReportersFieldsFilters) SetContainsFilter(v ReportersContainsFilter) {
	o.ContainsFilter = v
}

// GetContainsInArrayFilter returns the ContainsInArrayFilter field value if set, zero value otherwise.
func (o *ReportersFieldsFilters) GetContainsInArrayFilter() ReportersContainsFilter {
	if o == nil || IsNil(o.ContainsInArrayFilter) {
		var ret ReportersContainsFilter
		return ret
	}
	return *o.ContainsInArrayFilter
}

// GetContainsInArrayFilterOk returns a tuple with the ContainsInArrayFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportersFieldsFilters) GetContainsInArrayFilterOk() (*ReportersContainsFilter, bool) {
	if o == nil || IsNil(o.ContainsInArrayFilter) {
		return nil, false
	}
	return o.ContainsInArrayFilter, true
}

// HasContainsInArrayFilter returns a boolean if a field has been set.
func (o *ReportersFieldsFilters) HasContainsInArrayFilter() bool {
	if o != nil && !IsNil(o.ContainsInArrayFilter) {
		return true
	}

	return false
}

// SetContainsInArrayFilter gets a reference to the given ReportersContainsFilter and assigns it to the ContainsInArrayFilter field.
func (o *ReportersFieldsFilters) SetContainsInArrayFilter(v ReportersContainsFilter) {
	o.ContainsInArrayFilter = &v
}

// GetMatchFilter returns the MatchFilter field value
func (o *ReportersFieldsFilters) GetMatchFilter() ReportersMatchFilter {
	if o == nil {
		var ret ReportersMatchFilter
		return ret
	}

	return o.MatchFilter
}

// GetMatchFilterOk returns a tuple with the MatchFilter field value
// and a boolean to check if the value has been set.
func (o *ReportersFieldsFilters) GetMatchFilterOk() (*ReportersMatchFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MatchFilter, true
}

// SetMatchFilter sets field value
func (o *ReportersFieldsFilters) SetMatchFilter(v ReportersMatchFilter) {
	o.MatchFilter = v
}

// GetMatchInArrayFilter returns the MatchInArrayFilter field value if set, zero value otherwise.
func (o *ReportersFieldsFilters) GetMatchInArrayFilter() ReportersMatchFilter {
	if o == nil || IsNil(o.MatchInArrayFilter) {
		var ret ReportersMatchFilter
		return ret
	}
	return *o.MatchInArrayFilter
}

// GetMatchInArrayFilterOk returns a tuple with the MatchInArrayFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportersFieldsFilters) GetMatchInArrayFilterOk() (*ReportersMatchFilter, bool) {
	if o == nil || IsNil(o.MatchInArrayFilter) {
		return nil, false
	}
	return o.MatchInArrayFilter, true
}

// HasMatchInArrayFilter returns a boolean if a field has been set.
func (o *ReportersFieldsFilters) HasMatchInArrayFilter() bool {
	if o != nil && !IsNil(o.MatchInArrayFilter) {
		return true
	}

	return false
}

// SetMatchInArrayFilter gets a reference to the given ReportersMatchFilter and assigns it to the MatchInArrayFilter field.
func (o *ReportersFieldsFilters) SetMatchInArrayFilter(v ReportersMatchFilter) {
	o.MatchInArrayFilter = &v
}

// GetNotContainsFilter returns the NotContainsFilter field value if set, zero value otherwise.
func (o *ReportersFieldsFilters) GetNotContainsFilter() ReportersContainsFilter {
	if o == nil || IsNil(o.NotContainsFilter) {
		var ret ReportersContainsFilter
		return ret
	}
	return *o.NotContainsFilter
}

// GetNotContainsFilterOk returns a tuple with the NotContainsFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportersFieldsFilters) GetNotContainsFilterOk() (*ReportersContainsFilter, bool) {
	if o == nil || IsNil(o.NotContainsFilter) {
		return nil, false
	}
	return o.NotContainsFilter, true
}

// HasNotContainsFilter returns a boolean if a field has been set.
func (o *ReportersFieldsFilters) HasNotContainsFilter() bool {
	if o != nil && !IsNil(o.NotContainsFilter) {
		return true
	}

	return false
}

// SetNotContainsFilter gets a reference to the given ReportersContainsFilter and assigns it to the NotContainsFilter field.
func (o *ReportersFieldsFilters) SetNotContainsFilter(v ReportersContainsFilter) {
	o.NotContainsFilter = &v
}

// GetOrderFilter returns the OrderFilter field value
func (o *ReportersFieldsFilters) GetOrderFilter() ReportersOrderFilter {
	if o == nil {
		var ret ReportersOrderFilter
		return ret
	}

	return o.OrderFilter
}

// GetOrderFilterOk returns a tuple with the OrderFilter field value
// and a boolean to check if the value has been set.
func (o *ReportersFieldsFilters) GetOrderFilterOk() (*ReportersOrderFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderFilter, true
}

// SetOrderFilter sets field value
func (o *ReportersFieldsFilters) SetOrderFilter(v ReportersOrderFilter) {
	o.OrderFilter = v
}

func (o ReportersFieldsFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportersFieldsFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CompareFilter != nil {
		toSerialize["compare_filter"] = o.CompareFilter
	}
	toSerialize["contains_filter"] = o.ContainsFilter
	if !IsNil(o.ContainsInArrayFilter) {
		toSerialize["contains_in_array_filter"] = o.ContainsInArrayFilter
	}
	toSerialize["match_filter"] = o.MatchFilter
	if !IsNil(o.MatchInArrayFilter) {
		toSerialize["match_in_array_filter"] = o.MatchInArrayFilter
	}
	if !IsNil(o.NotContainsFilter) {
		toSerialize["not_contains_filter"] = o.NotContainsFilter
	}
	toSerialize["order_filter"] = o.OrderFilter
	return toSerialize, nil
}

func (o *ReportersFieldsFilters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"compare_filter",
		"contains_filter",
		"match_filter",
		"order_filter",
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

	varReportersFieldsFilters := _ReportersFieldsFilters{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReportersFieldsFilters)

	if err != nil {
		return err
	}

	*o = ReportersFieldsFilters(varReportersFieldsFilters)

	return err
}

type NullableReportersFieldsFilters struct {
	value *ReportersFieldsFilters
	isSet bool
}

func (v NullableReportersFieldsFilters) Get() *ReportersFieldsFilters {
	return v.value
}

func (v *NullableReportersFieldsFilters) Set(val *ReportersFieldsFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableReportersFieldsFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableReportersFieldsFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportersFieldsFilters(val *ReportersFieldsFilters) *NullableReportersFieldsFilters {
	return &NullableReportersFieldsFilters{value: val, isSet: true}
}

func (v NullableReportersFieldsFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportersFieldsFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


