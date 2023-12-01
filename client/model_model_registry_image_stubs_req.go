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

// checks if the ModelRegistryImageStubsReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelRegistryImageStubsReq{}

// ModelRegistryImageStubsReq struct for ModelRegistryImageStubsReq
type ModelRegistryImageStubsReq struct {
	ImageFilter ReportersFieldsFilters `json:"image_filter"`
	RegistryId string `json:"registry_id"`
	Window ModelFetchWindow `json:"window"`
}

type _ModelRegistryImageStubsReq ModelRegistryImageStubsReq

// NewModelRegistryImageStubsReq instantiates a new ModelRegistryImageStubsReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelRegistryImageStubsReq(imageFilter ReportersFieldsFilters, registryId string, window ModelFetchWindow) *ModelRegistryImageStubsReq {
	this := ModelRegistryImageStubsReq{}
	this.ImageFilter = imageFilter
	this.RegistryId = registryId
	this.Window = window
	return &this
}

// NewModelRegistryImageStubsReqWithDefaults instantiates a new ModelRegistryImageStubsReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelRegistryImageStubsReqWithDefaults() *ModelRegistryImageStubsReq {
	this := ModelRegistryImageStubsReq{}
	return &this
}

// GetImageFilter returns the ImageFilter field value
func (o *ModelRegistryImageStubsReq) GetImageFilter() ReportersFieldsFilters {
	if o == nil {
		var ret ReportersFieldsFilters
		return ret
	}

	return o.ImageFilter
}

// GetImageFilterOk returns a tuple with the ImageFilter field value
// and a boolean to check if the value has been set.
func (o *ModelRegistryImageStubsReq) GetImageFilterOk() (*ReportersFieldsFilters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageFilter, true
}

// SetImageFilter sets field value
func (o *ModelRegistryImageStubsReq) SetImageFilter(v ReportersFieldsFilters) {
	o.ImageFilter = v
}

// GetRegistryId returns the RegistryId field value
func (o *ModelRegistryImageStubsReq) GetRegistryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegistryId
}

// GetRegistryIdOk returns a tuple with the RegistryId field value
// and a boolean to check if the value has been set.
func (o *ModelRegistryImageStubsReq) GetRegistryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistryId, true
}

// SetRegistryId sets field value
func (o *ModelRegistryImageStubsReq) SetRegistryId(v string) {
	o.RegistryId = v
}

// GetWindow returns the Window field value
func (o *ModelRegistryImageStubsReq) GetWindow() ModelFetchWindow {
	if o == nil {
		var ret ModelFetchWindow
		return ret
	}

	return o.Window
}

// GetWindowOk returns a tuple with the Window field value
// and a boolean to check if the value has been set.
func (o *ModelRegistryImageStubsReq) GetWindowOk() (*ModelFetchWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Window, true
}

// SetWindow sets field value
func (o *ModelRegistryImageStubsReq) SetWindow(v ModelFetchWindow) {
	o.Window = v
}

func (o ModelRegistryImageStubsReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelRegistryImageStubsReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image_filter"] = o.ImageFilter
	toSerialize["registry_id"] = o.RegistryId
	toSerialize["window"] = o.Window
	return toSerialize, nil
}

func (o *ModelRegistryImageStubsReq) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"image_filter",
		"registry_id",
		"window",
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

	varModelRegistryImageStubsReq := _ModelRegistryImageStubsReq{}

	err = json.Unmarshal(bytes, &varModelRegistryImageStubsReq)

	if err != nil {
		return err
	}

	*o = ModelRegistryImageStubsReq(varModelRegistryImageStubsReq)

	return err
}

type NullableModelRegistryImageStubsReq struct {
	value *ModelRegistryImageStubsReq
	isSet bool
}

func (v NullableModelRegistryImageStubsReq) Get() *ModelRegistryImageStubsReq {
	return v.value
}

func (v *NullableModelRegistryImageStubsReq) Set(val *ModelRegistryImageStubsReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRegistryImageStubsReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRegistryImageStubsReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRegistryImageStubsReq(val *ModelRegistryImageStubsReq) *NullableModelRegistryImageStubsReq {
	return &NullableModelRegistryImageStubsReq{value: val, isSet: true}
}

func (v NullableModelRegistryImageStubsReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRegistryImageStubsReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


