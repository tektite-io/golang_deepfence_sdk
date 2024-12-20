/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.5.2
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelStopScanRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelStopScanRequest{}

// ModelStopScanRequest struct for ModelStopScanRequest
type ModelStopScanRequest struct {
	ScanIds []string `json:"scan_ids"`
	ScanType string `json:"scan_type"`
}

type _ModelStopScanRequest ModelStopScanRequest

// NewModelStopScanRequest instantiates a new ModelStopScanRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelStopScanRequest(scanIds []string, scanType string) *ModelStopScanRequest {
	this := ModelStopScanRequest{}
	this.ScanIds = scanIds
	this.ScanType = scanType
	return &this
}

// NewModelStopScanRequestWithDefaults instantiates a new ModelStopScanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelStopScanRequestWithDefaults() *ModelStopScanRequest {
	this := ModelStopScanRequest{}
	return &this
}

// GetScanIds returns the ScanIds field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelStopScanRequest) GetScanIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ScanIds
}

// GetScanIdsOk returns a tuple with the ScanIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelStopScanRequest) GetScanIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ScanIds) {
		return nil, false
	}
	return o.ScanIds, true
}

// SetScanIds sets field value
func (o *ModelStopScanRequest) SetScanIds(v []string) {
	o.ScanIds = v
}

// GetScanType returns the ScanType field value
func (o *ModelStopScanRequest) GetScanType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScanType
}

// GetScanTypeOk returns a tuple with the ScanType field value
// and a boolean to check if the value has been set.
func (o *ModelStopScanRequest) GetScanTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanType, true
}

// SetScanType sets field value
func (o *ModelStopScanRequest) SetScanType(v string) {
	o.ScanType = v
}

func (o ModelStopScanRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelStopScanRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ScanIds != nil {
		toSerialize["scan_ids"] = o.ScanIds
	}
	toSerialize["scan_type"] = o.ScanType
	return toSerialize, nil
}

func (o *ModelStopScanRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"scan_ids",
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

	varModelStopScanRequest := _ModelStopScanRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelStopScanRequest)

	if err != nil {
		return err
	}

	*o = ModelStopScanRequest(varModelStopScanRequest)

	return err
}

type NullableModelStopScanRequest struct {
	value *ModelStopScanRequest
	isSet bool
}

func (v NullableModelStopScanRequest) Get() *ModelStopScanRequest {
	return v.value
}

func (v *NullableModelStopScanRequest) Set(val *ModelStopScanRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelStopScanRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelStopScanRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelStopScanRequest(val *ModelStopScanRequest) *NullableModelStopScanRequest {
	return &NullableModelStopScanRequest{value: val, isSet: true}
}

func (v NullableModelStopScanRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelStopScanRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


