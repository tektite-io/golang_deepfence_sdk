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

// checks if the ModelCloudNodeMonitoredAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCloudNodeMonitoredAccount{}

// ModelCloudNodeMonitoredAccount struct for ModelCloudNodeMonitoredAccount
type ModelCloudNodeMonitoredAccount struct {
	AccountId string `json:"account_id"`
	AccountName *string `json:"account_name,omitempty"`
	NodeId string `json:"node_id"`
}

type _ModelCloudNodeMonitoredAccount ModelCloudNodeMonitoredAccount

// NewModelCloudNodeMonitoredAccount instantiates a new ModelCloudNodeMonitoredAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCloudNodeMonitoredAccount(accountId string, nodeId string) *ModelCloudNodeMonitoredAccount {
	this := ModelCloudNodeMonitoredAccount{}
	this.AccountId = accountId
	this.NodeId = nodeId
	return &this
}

// NewModelCloudNodeMonitoredAccountWithDefaults instantiates a new ModelCloudNodeMonitoredAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCloudNodeMonitoredAccountWithDefaults() *ModelCloudNodeMonitoredAccount {
	this := ModelCloudNodeMonitoredAccount{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *ModelCloudNodeMonitoredAccount) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeMonitoredAccount) GetAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *ModelCloudNodeMonitoredAccount) SetAccountId(v string) {
	o.AccountId = v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *ModelCloudNodeMonitoredAccount) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeMonitoredAccount) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// HasAccountName returns a boolean if a field has been set.
func (o *ModelCloudNodeMonitoredAccount) HasAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *ModelCloudNodeMonitoredAccount) SetAccountName(v string) {
	o.AccountName = &v
}

// GetNodeId returns the NodeId field value
func (o *ModelCloudNodeMonitoredAccount) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeMonitoredAccount) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelCloudNodeMonitoredAccount) SetNodeId(v string) {
	o.NodeId = v
}

func (o ModelCloudNodeMonitoredAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCloudNodeMonitoredAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["account_id"] = o.AccountId
	if !IsNil(o.AccountName) {
		toSerialize["account_name"] = o.AccountName
	}
	toSerialize["node_id"] = o.NodeId
	return toSerialize, nil
}

func (o *ModelCloudNodeMonitoredAccount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"account_id",
		"node_id",
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

	varModelCloudNodeMonitoredAccount := _ModelCloudNodeMonitoredAccount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelCloudNodeMonitoredAccount)

	if err != nil {
		return err
	}

	*o = ModelCloudNodeMonitoredAccount(varModelCloudNodeMonitoredAccount)

	return err
}

type NullableModelCloudNodeMonitoredAccount struct {
	value *ModelCloudNodeMonitoredAccount
	isSet bool
}

func (v NullableModelCloudNodeMonitoredAccount) Get() *ModelCloudNodeMonitoredAccount {
	return v.value
}

func (v *NullableModelCloudNodeMonitoredAccount) Set(val *ModelCloudNodeMonitoredAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCloudNodeMonitoredAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCloudNodeMonitoredAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCloudNodeMonitoredAccount(val *ModelCloudNodeMonitoredAccount) *NullableModelCloudNodeMonitoredAccount {
	return &NullableModelCloudNodeMonitoredAccount{value: val, isSet: true}
}

func (v NullableModelCloudNodeMonitoredAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCloudNodeMonitoredAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


