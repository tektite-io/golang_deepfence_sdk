/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.2.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelScanResultsActionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelScanResultsActionRequest{}

// ModelScanResultsActionRequest struct for ModelScanResultsActionRequest
type ModelScanResultsActionRequest struct {
	IntegrationIds []int32 `json:"integration_ids,omitempty"`
	NotifyIndividual *bool `json:"notify_individual,omitempty"`
	ResultIds []string `json:"result_ids"`
	ScanId string `json:"scan_id"`
	ScanType string `json:"scan_type"`
}

type _ModelScanResultsActionRequest ModelScanResultsActionRequest

// NewModelScanResultsActionRequest instantiates a new ModelScanResultsActionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelScanResultsActionRequest(resultIds []string, scanId string, scanType string) *ModelScanResultsActionRequest {
	this := ModelScanResultsActionRequest{}
	this.ResultIds = resultIds
	this.ScanId = scanId
	this.ScanType = scanType
	return &this
}

// NewModelScanResultsActionRequestWithDefaults instantiates a new ModelScanResultsActionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelScanResultsActionRequestWithDefaults() *ModelScanResultsActionRequest {
	this := ModelScanResultsActionRequest{}
	return &this
}

// GetIntegrationIds returns the IntegrationIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelScanResultsActionRequest) GetIntegrationIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}
	return o.IntegrationIds
}

// GetIntegrationIdsOk returns a tuple with the IntegrationIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelScanResultsActionRequest) GetIntegrationIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.IntegrationIds) {
		return nil, false
	}
	return o.IntegrationIds, true
}

// HasIntegrationIds returns a boolean if a field has been set.
func (o *ModelScanResultsActionRequest) HasIntegrationIds() bool {
	if o != nil && !IsNil(o.IntegrationIds) {
		return true
	}

	return false
}

// SetIntegrationIds gets a reference to the given []int32 and assigns it to the IntegrationIds field.
func (o *ModelScanResultsActionRequest) SetIntegrationIds(v []int32) {
	o.IntegrationIds = v
}

// GetNotifyIndividual returns the NotifyIndividual field value if set, zero value otherwise.
func (o *ModelScanResultsActionRequest) GetNotifyIndividual() bool {
	if o == nil || IsNil(o.NotifyIndividual) {
		var ret bool
		return ret
	}
	return *o.NotifyIndividual
}

// GetNotifyIndividualOk returns a tuple with the NotifyIndividual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelScanResultsActionRequest) GetNotifyIndividualOk() (*bool, bool) {
	if o == nil || IsNil(o.NotifyIndividual) {
		return nil, false
	}
	return o.NotifyIndividual, true
}

// HasNotifyIndividual returns a boolean if a field has been set.
func (o *ModelScanResultsActionRequest) HasNotifyIndividual() bool {
	if o != nil && !IsNil(o.NotifyIndividual) {
		return true
	}

	return false
}

// SetNotifyIndividual gets a reference to the given bool and assigns it to the NotifyIndividual field.
func (o *ModelScanResultsActionRequest) SetNotifyIndividual(v bool) {
	o.NotifyIndividual = &v
}

// GetResultIds returns the ResultIds field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelScanResultsActionRequest) GetResultIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResultIds
}

// GetResultIdsOk returns a tuple with the ResultIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelScanResultsActionRequest) GetResultIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ResultIds) {
		return nil, false
	}
	return o.ResultIds, true
}

// SetResultIds sets field value
func (o *ModelScanResultsActionRequest) SetResultIds(v []string) {
	o.ResultIds = v
}

// GetScanId returns the ScanId field value
func (o *ModelScanResultsActionRequest) GetScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsActionRequest) GetScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanId, true
}

// SetScanId sets field value
func (o *ModelScanResultsActionRequest) SetScanId(v string) {
	o.ScanId = v
}

// GetScanType returns the ScanType field value
func (o *ModelScanResultsActionRequest) GetScanType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScanType
}

// GetScanTypeOk returns a tuple with the ScanType field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsActionRequest) GetScanTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanType, true
}

// SetScanType sets field value
func (o *ModelScanResultsActionRequest) SetScanType(v string) {
	o.ScanType = v
}

func (o ModelScanResultsActionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelScanResultsActionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.IntegrationIds != nil {
		toSerialize["integration_ids"] = o.IntegrationIds
	}
	if !IsNil(o.NotifyIndividual) {
		toSerialize["notify_individual"] = o.NotifyIndividual
	}
	if o.ResultIds != nil {
		toSerialize["result_ids"] = o.ResultIds
	}
	toSerialize["scan_id"] = o.ScanId
	toSerialize["scan_type"] = o.ScanType
	return toSerialize, nil
}

func (o *ModelScanResultsActionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"result_ids",
		"scan_id",
		"scan_type",
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

	varModelScanResultsActionRequest := _ModelScanResultsActionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelScanResultsActionRequest)

	if err != nil {
		return err
	}

	*o = ModelScanResultsActionRequest(varModelScanResultsActionRequest)

	return err
}

type NullableModelScanResultsActionRequest struct {
	value *ModelScanResultsActionRequest
	isSet bool
}

func (v NullableModelScanResultsActionRequest) Get() *ModelScanResultsActionRequest {
	return v.value
}

func (v *NullableModelScanResultsActionRequest) Set(val *ModelScanResultsActionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelScanResultsActionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelScanResultsActionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelScanResultsActionRequest(val *ModelScanResultsActionRequest) *NullableModelScanResultsActionRequest {
	return &NullableModelScanResultsActionRequest{value: val, isSet: true}
}

func (v NullableModelScanResultsActionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelScanResultsActionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


