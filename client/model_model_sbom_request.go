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
	"fmt"
)

// checks if the ModelSbomRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSbomRequest{}

// ModelSbomRequest struct for ModelSbomRequest
type ModelSbomRequest struct {
	ScanId string `json:"scan_id"`
}

type _ModelSbomRequest ModelSbomRequest

// NewModelSbomRequest instantiates a new ModelSbomRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSbomRequest(scanId string) *ModelSbomRequest {
	this := ModelSbomRequest{}
	this.ScanId = scanId
	return &this
}

// NewModelSbomRequestWithDefaults instantiates a new ModelSbomRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSbomRequestWithDefaults() *ModelSbomRequest {
	this := ModelSbomRequest{}
	return &this
}

// GetScanId returns the ScanId field value
func (o *ModelSbomRequest) GetScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value
// and a boolean to check if the value has been set.
func (o *ModelSbomRequest) GetScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanId, true
}

// SetScanId sets field value
func (o *ModelSbomRequest) SetScanId(v string) {
	o.ScanId = v
}

func (o ModelSbomRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSbomRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["scan_id"] = o.ScanId
	return toSerialize, nil
}

func (o *ModelSbomRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"scan_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varModelSbomRequest := _ModelSbomRequest{}

	err = json.Unmarshal(bytes, &varModelSbomRequest)

	if err != nil {
		return err
	}

	*o = ModelSbomRequest(varModelSbomRequest)

	return err
}

type NullableModelSbomRequest struct {
	value *ModelSbomRequest
	isSet bool
}

func (v NullableModelSbomRequest) Get() *ModelSbomRequest {
	return v.value
}

func (v *NullableModelSbomRequest) Set(val *ModelSbomRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSbomRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSbomRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSbomRequest(val *ModelSbomRequest) *NullableModelSbomRequest {
	return &NullableModelSbomRequest{value: val, isSet: true}
}

func (v NullableModelSbomRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSbomRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


