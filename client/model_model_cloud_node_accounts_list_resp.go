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

// checks if the ModelCloudNodeAccountsListResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCloudNodeAccountsListResp{}

// ModelCloudNodeAccountsListResp struct for ModelCloudNodeAccountsListResp
type ModelCloudNodeAccountsListResp struct {
	CloudNodeAccountsInfo []ModelCloudNodeAccountInfo `json:"cloud_node_accounts_info"`
	Total int32 `json:"total"`
}

type _ModelCloudNodeAccountsListResp ModelCloudNodeAccountsListResp

// NewModelCloudNodeAccountsListResp instantiates a new ModelCloudNodeAccountsListResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCloudNodeAccountsListResp(cloudNodeAccountsInfo []ModelCloudNodeAccountInfo, total int32) *ModelCloudNodeAccountsListResp {
	this := ModelCloudNodeAccountsListResp{}
	this.CloudNodeAccountsInfo = cloudNodeAccountsInfo
	this.Total = total
	return &this
}

// NewModelCloudNodeAccountsListRespWithDefaults instantiates a new ModelCloudNodeAccountsListResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCloudNodeAccountsListRespWithDefaults() *ModelCloudNodeAccountsListResp {
	this := ModelCloudNodeAccountsListResp{}
	return &this
}

// GetCloudNodeAccountsInfo returns the CloudNodeAccountsInfo field value
// If the value is explicit nil, the zero value for []ModelCloudNodeAccountInfo will be returned
func (o *ModelCloudNodeAccountsListResp) GetCloudNodeAccountsInfo() []ModelCloudNodeAccountInfo {
	if o == nil {
		var ret []ModelCloudNodeAccountInfo
		return ret
	}

	return o.CloudNodeAccountsInfo
}

// GetCloudNodeAccountsInfoOk returns a tuple with the CloudNodeAccountsInfo field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCloudNodeAccountsListResp) GetCloudNodeAccountsInfoOk() ([]ModelCloudNodeAccountInfo, bool) {
	if o == nil || IsNil(o.CloudNodeAccountsInfo) {
		return nil, false
	}
	return o.CloudNodeAccountsInfo, true
}

// SetCloudNodeAccountsInfo sets field value
func (o *ModelCloudNodeAccountsListResp) SetCloudNodeAccountsInfo(v []ModelCloudNodeAccountInfo) {
	o.CloudNodeAccountsInfo = v
}

// GetTotal returns the Total field value
func (o *ModelCloudNodeAccountsListResp) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeAccountsListResp) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ModelCloudNodeAccountsListResp) SetTotal(v int32) {
	o.Total = v
}

func (o ModelCloudNodeAccountsListResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCloudNodeAccountsListResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CloudNodeAccountsInfo != nil {
		toSerialize["cloud_node_accounts_info"] = o.CloudNodeAccountsInfo
	}
	toSerialize["total"] = o.Total
	return toSerialize, nil
}

func (o *ModelCloudNodeAccountsListResp) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cloud_node_accounts_info",
		"total",
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

	varModelCloudNodeAccountsListResp := _ModelCloudNodeAccountsListResp{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelCloudNodeAccountsListResp)

	if err != nil {
		return err
	}

	*o = ModelCloudNodeAccountsListResp(varModelCloudNodeAccountsListResp)

	return err
}

type NullableModelCloudNodeAccountsListResp struct {
	value *ModelCloudNodeAccountsListResp
	isSet bool
}

func (v NullableModelCloudNodeAccountsListResp) Get() *ModelCloudNodeAccountsListResp {
	return v.value
}

func (v *NullableModelCloudNodeAccountsListResp) Set(val *ModelCloudNodeAccountsListResp) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCloudNodeAccountsListResp) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCloudNodeAccountsListResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCloudNodeAccountsListResp(val *ModelCloudNodeAccountsListResp) *NullableModelCloudNodeAccountsListResp {
	return &NullableModelCloudNodeAccountsListResp{value: val, isSet: true}
}

func (v NullableModelCloudNodeAccountsListResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCloudNodeAccountsListResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


