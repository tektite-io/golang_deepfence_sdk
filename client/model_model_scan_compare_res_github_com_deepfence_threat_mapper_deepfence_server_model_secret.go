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

// checks if the ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret{}

// ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret struct for ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret
type ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret struct {
	New []ModelSecret `json:"new"`
}

type _ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret

// NewModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret instantiates a new ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret(new []ModelSecret) *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret {
	this := ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret{}
	this.New = new
	return &this
}

// NewModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecretWithDefaults instantiates a new ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecretWithDefaults() *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret {
	this := ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret{}
	return &this
}

// GetNew returns the New field value
// If the value is explicit nil, the zero value for []ModelSecret will be returned
func (o *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret) GetNew() []ModelSecret {
	if o == nil {
		var ret []ModelSecret
		return ret
	}

	return o.New
}

// GetNewOk returns a tuple with the New field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret) GetNewOk() ([]ModelSecret, bool) {
	if o == nil || IsNil(o.New) {
		return nil, false
	}
	return o.New, true
}

// SetNew sets field value
func (o *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret) SetNew(v []ModelSecret) {
	o.New = v
}

func (o ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.New != nil {
		toSerialize["new"] = o.New
	}
	return toSerialize, nil
}

func (o *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"new",
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

	varModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret := _ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret{}

	err = json.Unmarshal(bytes, &varModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret)

	if err != nil {
		return err
	}

	*o = ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret(varModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret)

	return err
}

type NullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret struct {
	value *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret
	isSet bool
}

func (v NullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret) Get() *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret {
	return v.value
}

func (v *NullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret) Set(val *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret(val *ModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret) *NullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret {
	return &NullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret{value: val, isSet: true}
}

func (v NullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelScanCompareResGithubComDeepfenceThreatMapperDeepfenceServerModelSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


