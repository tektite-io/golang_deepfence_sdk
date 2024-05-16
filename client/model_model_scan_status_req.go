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

// checks if the ModelScanStatusReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelScanStatusReq{}

// ModelScanStatusReq struct for ModelScanStatusReq
type ModelScanStatusReq struct {
	BulkScanId string `json:"bulk_scan_id"`
	ScanIds []string `json:"scan_ids"`
}

type _ModelScanStatusReq ModelScanStatusReq

// NewModelScanStatusReq instantiates a new ModelScanStatusReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelScanStatusReq(bulkScanId string, scanIds []string) *ModelScanStatusReq {
	this := ModelScanStatusReq{}
	this.BulkScanId = bulkScanId
	this.ScanIds = scanIds
	return &this
}

// NewModelScanStatusReqWithDefaults instantiates a new ModelScanStatusReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelScanStatusReqWithDefaults() *ModelScanStatusReq {
	this := ModelScanStatusReq{}
	return &this
}

// GetBulkScanId returns the BulkScanId field value
func (o *ModelScanStatusReq) GetBulkScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BulkScanId
}

// GetBulkScanIdOk returns a tuple with the BulkScanId field value
// and a boolean to check if the value has been set.
func (o *ModelScanStatusReq) GetBulkScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BulkScanId, true
}

// SetBulkScanId sets field value
func (o *ModelScanStatusReq) SetBulkScanId(v string) {
	o.BulkScanId = v
}

// GetScanIds returns the ScanIds field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelScanStatusReq) GetScanIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ScanIds
}

// GetScanIdsOk returns a tuple with the ScanIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelScanStatusReq) GetScanIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ScanIds) {
		return nil, false
	}
	return o.ScanIds, true
}

// SetScanIds sets field value
func (o *ModelScanStatusReq) SetScanIds(v []string) {
	o.ScanIds = v
}

func (o ModelScanStatusReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelScanStatusReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bulk_scan_id"] = o.BulkScanId
	if o.ScanIds != nil {
		toSerialize["scan_ids"] = o.ScanIds
	}
	return toSerialize, nil
}

func (o *ModelScanStatusReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bulk_scan_id",
		"scan_ids",
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

	varModelScanStatusReq := _ModelScanStatusReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelScanStatusReq)

	if err != nil {
		return err
	}

	*o = ModelScanStatusReq(varModelScanStatusReq)

	return err
}

type NullableModelScanStatusReq struct {
	value *ModelScanStatusReq
	isSet bool
}

func (v NullableModelScanStatusReq) Get() *ModelScanStatusReq {
	return v.value
}

func (v *NullableModelScanStatusReq) Set(val *ModelScanStatusReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelScanStatusReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelScanStatusReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelScanStatusReq(val *ModelScanStatusReq) *NullableModelScanStatusReq {
	return &NullableModelScanStatusReq{value: val, isSet: true}
}

func (v NullableModelScanStatusReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelScanStatusReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


