/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.3.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GraphCloudProviderFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphCloudProviderFilter{}

// GraphCloudProviderFilter struct for GraphCloudProviderFilter
type GraphCloudProviderFilter struct {
	AccountIds []string `json:"account_ids"`
}

type _GraphCloudProviderFilter GraphCloudProviderFilter

// NewGraphCloudProviderFilter instantiates a new GraphCloudProviderFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphCloudProviderFilter(accountIds []string) *GraphCloudProviderFilter {
	this := GraphCloudProviderFilter{}
	this.AccountIds = accountIds
	return &this
}

// NewGraphCloudProviderFilterWithDefaults instantiates a new GraphCloudProviderFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphCloudProviderFilterWithDefaults() *GraphCloudProviderFilter {
	this := GraphCloudProviderFilter{}
	return &this
}

// GetAccountIds returns the AccountIds field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *GraphCloudProviderFilter) GetAccountIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *GraphCloudProviderFilter) GetAccountIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AccountIds) {
		return nil, false
	}
	return o.AccountIds, true
}

// SetAccountIds sets field value
func (o *GraphCloudProviderFilter) SetAccountIds(v []string) {
	o.AccountIds = v
}

func (o GraphCloudProviderFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphCloudProviderFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountIds != nil {
		toSerialize["account_ids"] = o.AccountIds
	}
	return toSerialize, nil
}

func (o *GraphCloudProviderFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_ids",
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

	varGraphCloudProviderFilter := _GraphCloudProviderFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGraphCloudProviderFilter)

	if err != nil {
		return err
	}

	*o = GraphCloudProviderFilter(varGraphCloudProviderFilter)

	return err
}

type NullableGraphCloudProviderFilter struct {
	value *GraphCloudProviderFilter
	isSet bool
}

func (v NullableGraphCloudProviderFilter) Get() *GraphCloudProviderFilter {
	return v.value
}

func (v *NullableGraphCloudProviderFilter) Set(val *GraphCloudProviderFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphCloudProviderFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphCloudProviderFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphCloudProviderFilter(val *GraphCloudProviderFilter) *NullableGraphCloudProviderFilter {
	return &NullableGraphCloudProviderFilter{value: val, isSet: true}
}

func (v NullableGraphCloudProviderFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphCloudProviderFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


