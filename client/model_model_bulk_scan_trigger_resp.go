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
)

// checks if the ModelBulkScanTriggerResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelBulkScanTriggerResp{}

// ModelBulkScanTriggerResp struct for ModelBulkScanTriggerResp
type ModelBulkScanTriggerResp struct {
	BulkScanId string `json:"bulk_scan_id"`
}

// NewModelBulkScanTriggerResp instantiates a new ModelBulkScanTriggerResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelBulkScanTriggerResp(bulkScanId string) *ModelBulkScanTriggerResp {
	this := ModelBulkScanTriggerResp{}
	this.BulkScanId = bulkScanId
	return &this
}

// NewModelBulkScanTriggerRespWithDefaults instantiates a new ModelBulkScanTriggerResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelBulkScanTriggerRespWithDefaults() *ModelBulkScanTriggerResp {
	this := ModelBulkScanTriggerResp{}
	return &this
}

// GetBulkScanId returns the BulkScanId field value
func (o *ModelBulkScanTriggerResp) GetBulkScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BulkScanId
}

// GetBulkScanIdOk returns a tuple with the BulkScanId field value
// and a boolean to check if the value has been set.
func (o *ModelBulkScanTriggerResp) GetBulkScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BulkScanId, true
}

// SetBulkScanId sets field value
func (o *ModelBulkScanTriggerResp) SetBulkScanId(v string) {
	o.BulkScanId = v
}

func (o ModelBulkScanTriggerResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelBulkScanTriggerResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bulk_scan_id"] = o.BulkScanId
	return toSerialize, nil
}

type NullableModelBulkScanTriggerResp struct {
	value *ModelBulkScanTriggerResp
	isSet bool
}

func (v NullableModelBulkScanTriggerResp) Get() *ModelBulkScanTriggerResp {
	return v.value
}

func (v *NullableModelBulkScanTriggerResp) Set(val *ModelBulkScanTriggerResp) {
	v.value = val
	v.isSet = true
}

func (v NullableModelBulkScanTriggerResp) IsSet() bool {
	return v.isSet
}

func (v *NullableModelBulkScanTriggerResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelBulkScanTriggerResp(val *ModelBulkScanTriggerResp) *NullableModelBulkScanTriggerResp {
	return &NullableModelBulkScanTriggerResp{value: val, isSet: true}
}

func (v NullableModelBulkScanTriggerResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelBulkScanTriggerResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


