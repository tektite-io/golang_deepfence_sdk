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
)

// checks if the ModelImageStub type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelImageStub{}

// ModelImageStub struct for ModelImageStub
type ModelImageStub struct {
	Id *string `json:"id,omitempty"`
	Images *int32 `json:"images,omitempty"`
	Name *string `json:"name,omitempty"`
	Tags []string `json:"tags,omitempty"`
}

// NewModelImageStub instantiates a new ModelImageStub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelImageStub() *ModelImageStub {
	this := ModelImageStub{}
	return &this
}

// NewModelImageStubWithDefaults instantiates a new ModelImageStub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelImageStubWithDefaults() *ModelImageStub {
	this := ModelImageStub{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelImageStub) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelImageStub) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelImageStub) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ModelImageStub) SetId(v string) {
	o.Id = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *ModelImageStub) GetImages() int32 {
	if o == nil || IsNil(o.Images) {
		var ret int32
		return ret
	}
	return *o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelImageStub) GetImagesOk() (*int32, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *ModelImageStub) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given int32 and assigns it to the Images field.
func (o *ModelImageStub) SetImages(v int32) {
	o.Images = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelImageStub) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelImageStub) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelImageStub) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelImageStub) SetName(v string) {
	o.Name = &v
}

// GetTags returns the Tags field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelImageStub) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelImageStub) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ModelImageStub) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ModelImageStub) SetTags(v []string) {
	o.Tags = v
}

func (o ModelImageStub) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelImageStub) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableModelImageStub struct {
	value *ModelImageStub
	isSet bool
}

func (v NullableModelImageStub) Get() *ModelImageStub {
	return v.value
}

func (v *NullableModelImageStub) Set(val *ModelImageStub) {
	v.value = val
	v.isSet = true
}

func (v NullableModelImageStub) IsSet() bool {
	return v.isSet
}

func (v *NullableModelImageStub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelImageStub(val *ModelImageStub) *NullableModelImageStub {
	return &NullableModelImageStub{value: val, isSet: true}
}

func (v NullableModelImageStub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelImageStub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


