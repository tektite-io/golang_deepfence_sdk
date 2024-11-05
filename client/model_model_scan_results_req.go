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

// checks if the ModelScanResultsReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelScanResultsReq{}

// ModelScanResultsReq struct for ModelScanResultsReq
type ModelScanResultsReq struct {
	FieldsFilter ReportersFieldsFilters `json:"fields_filter"`
	ScanId string `json:"scan_id"`
	Window ModelFetchWindow `json:"window"`
}

type _ModelScanResultsReq ModelScanResultsReq

// NewModelScanResultsReq instantiates a new ModelScanResultsReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelScanResultsReq(fieldsFilter ReportersFieldsFilters, scanId string, window ModelFetchWindow) *ModelScanResultsReq {
	this := ModelScanResultsReq{}
	this.FieldsFilter = fieldsFilter
	this.ScanId = scanId
	this.Window = window
	return &this
}

// NewModelScanResultsReqWithDefaults instantiates a new ModelScanResultsReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelScanResultsReqWithDefaults() *ModelScanResultsReq {
	this := ModelScanResultsReq{}
	return &this
}

// GetFieldsFilter returns the FieldsFilter field value
func (o *ModelScanResultsReq) GetFieldsFilter() ReportersFieldsFilters {
	if o == nil {
		var ret ReportersFieldsFilters
		return ret
	}

	return o.FieldsFilter
}

// GetFieldsFilterOk returns a tuple with the FieldsFilter field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsReq) GetFieldsFilterOk() (*ReportersFieldsFilters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldsFilter, true
}

// SetFieldsFilter sets field value
func (o *ModelScanResultsReq) SetFieldsFilter(v ReportersFieldsFilters) {
	o.FieldsFilter = v
}

// GetScanId returns the ScanId field value
func (o *ModelScanResultsReq) GetScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsReq) GetScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanId, true
}

// SetScanId sets field value
func (o *ModelScanResultsReq) SetScanId(v string) {
	o.ScanId = v
}

// GetWindow returns the Window field value
func (o *ModelScanResultsReq) GetWindow() ModelFetchWindow {
	if o == nil {
		var ret ModelFetchWindow
		return ret
	}

	return o.Window
}

// GetWindowOk returns a tuple with the Window field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsReq) GetWindowOk() (*ModelFetchWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Window, true
}

// SetWindow sets field value
func (o *ModelScanResultsReq) SetWindow(v ModelFetchWindow) {
	o.Window = v
}

func (o ModelScanResultsReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelScanResultsReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fields_filter"] = o.FieldsFilter
	toSerialize["scan_id"] = o.ScanId
	toSerialize["window"] = o.Window
	return toSerialize, nil
}

func (o *ModelScanResultsReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fields_filter",
		"scan_id",
		"window",
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

	varModelScanResultsReq := _ModelScanResultsReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelScanResultsReq)

	if err != nil {
		return err
	}

	*o = ModelScanResultsReq(varModelScanResultsReq)

	return err
}

type NullableModelScanResultsReq struct {
	value *ModelScanResultsReq
	isSet bool
}

func (v NullableModelScanResultsReq) Get() *ModelScanResultsReq {
	return v.value
}

func (v *NullableModelScanResultsReq) Set(val *ModelScanResultsReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelScanResultsReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelScanResultsReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelScanResultsReq(val *ModelScanResultsReq) *NullableModelScanResultsReq {
	return &NullableModelScanResultsReq{value: val, isSet: true}
}

func (v NullableModelScanResultsReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelScanResultsReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


