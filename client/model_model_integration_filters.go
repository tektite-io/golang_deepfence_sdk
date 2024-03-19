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
	"bytes"
	"fmt"
)

// checks if the ModelIntegrationFilters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelIntegrationFilters{}

// ModelIntegrationFilters struct for ModelIntegrationFilters
type ModelIntegrationFilters struct {
	ContainerNames []string `json:"container_names,omitempty"`
	FieldsFilters *ReportersFieldsFilters `json:"fields_filters,omitempty"`
	NodeIds []ModelNodeIdentifier `json:"node_ids"`
}

type _ModelIntegrationFilters ModelIntegrationFilters

// NewModelIntegrationFilters instantiates a new ModelIntegrationFilters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelIntegrationFilters(nodeIds []ModelNodeIdentifier) *ModelIntegrationFilters {
	this := ModelIntegrationFilters{}
	this.NodeIds = nodeIds
	return &this
}

// NewModelIntegrationFiltersWithDefaults instantiates a new ModelIntegrationFilters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelIntegrationFiltersWithDefaults() *ModelIntegrationFilters {
	this := ModelIntegrationFilters{}
	return &this
}

// GetContainerNames returns the ContainerNames field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelIntegrationFilters) GetContainerNames() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ContainerNames
}

// GetContainerNamesOk returns a tuple with the ContainerNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelIntegrationFilters) GetContainerNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.ContainerNames) {
		return nil, false
	}
	return o.ContainerNames, true
}

// HasContainerNames returns a boolean if a field has been set.
func (o *ModelIntegrationFilters) HasContainerNames() bool {
	if o != nil && !IsNil(o.ContainerNames) {
		return true
	}

	return false
}

// SetContainerNames gets a reference to the given []string and assigns it to the ContainerNames field.
func (o *ModelIntegrationFilters) SetContainerNames(v []string) {
	o.ContainerNames = v
}

// GetFieldsFilters returns the FieldsFilters field value if set, zero value otherwise.
func (o *ModelIntegrationFilters) GetFieldsFilters() ReportersFieldsFilters {
	if o == nil || IsNil(o.FieldsFilters) {
		var ret ReportersFieldsFilters
		return ret
	}
	return *o.FieldsFilters
}

// GetFieldsFiltersOk returns a tuple with the FieldsFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelIntegrationFilters) GetFieldsFiltersOk() (*ReportersFieldsFilters, bool) {
	if o == nil || IsNil(o.FieldsFilters) {
		return nil, false
	}
	return o.FieldsFilters, true
}

// HasFieldsFilters returns a boolean if a field has been set.
func (o *ModelIntegrationFilters) HasFieldsFilters() bool {
	if o != nil && !IsNil(o.FieldsFilters) {
		return true
	}

	return false
}

// SetFieldsFilters gets a reference to the given ReportersFieldsFilters and assigns it to the FieldsFilters field.
func (o *ModelIntegrationFilters) SetFieldsFilters(v ReportersFieldsFilters) {
	o.FieldsFilters = &v
}

// GetNodeIds returns the NodeIds field value
// If the value is explicit nil, the zero value for []ModelNodeIdentifier will be returned
func (o *ModelIntegrationFilters) GetNodeIds() []ModelNodeIdentifier {
	if o == nil {
		var ret []ModelNodeIdentifier
		return ret
	}

	return o.NodeIds
}

// GetNodeIdsOk returns a tuple with the NodeIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelIntegrationFilters) GetNodeIdsOk() ([]ModelNodeIdentifier, bool) {
	if o == nil || IsNil(o.NodeIds) {
		return nil, false
	}
	return o.NodeIds, true
}

// SetNodeIds sets field value
func (o *ModelIntegrationFilters) SetNodeIds(v []ModelNodeIdentifier) {
	o.NodeIds = v
}

func (o ModelIntegrationFilters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelIntegrationFilters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.ContainerNames != nil {
		toSerialize["container_names"] = o.ContainerNames
	}
	if !IsNil(o.FieldsFilters) {
		toSerialize["fields_filters"] = o.FieldsFilters
	}
	if o.NodeIds != nil {
		toSerialize["node_ids"] = o.NodeIds
	}
	return toSerialize, nil
}

func (o *ModelIntegrationFilters) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"node_ids",
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

	varModelIntegrationFilters := _ModelIntegrationFilters{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelIntegrationFilters)

	if err != nil {
		return err
	}

	*o = ModelIntegrationFilters(varModelIntegrationFilters)

	return err
}

type NullableModelIntegrationFilters struct {
	value *ModelIntegrationFilters
	isSet bool
}

func (v NullableModelIntegrationFilters) Get() *ModelIntegrationFilters {
	return v.value
}

func (v *NullableModelIntegrationFilters) Set(val *ModelIntegrationFilters) {
	v.value = val
	v.isSet = true
}

func (v NullableModelIntegrationFilters) IsSet() bool {
	return v.isSet
}

func (v *NullableModelIntegrationFilters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelIntegrationFilters(val *ModelIntegrationFilters) *NullableModelIntegrationFilters {
	return &NullableModelIntegrationFilters{value: val, isSet: true}
}

func (v NullableModelIntegrationFilters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelIntegrationFilters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


