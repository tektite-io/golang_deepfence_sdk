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

// checks if the ModelSecretRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSecretRule{}

// ModelSecretRule struct for ModelSecretRule
type ModelSecretRule struct {
	Id *int32 `json:"id,omitempty"`
	Level string `json:"level"`
	Masked bool `json:"masked"`
	Name *string `json:"name,omitempty"`
	Part *string `json:"part,omitempty"`
	SignatureToMatch *string `json:"signature_to_match,omitempty"`
	UpdatedAt int32 `json:"updated_at"`
}

// NewModelSecretRule instantiates a new ModelSecretRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSecretRule(level string, masked bool, updatedAt int32) *ModelSecretRule {
	this := ModelSecretRule{}
	this.Level = level
	this.Masked = masked
	this.UpdatedAt = updatedAt
	return &this
}

// NewModelSecretRuleWithDefaults instantiates a new ModelSecretRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSecretRuleWithDefaults() *ModelSecretRule {
	this := ModelSecretRule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelSecretRule) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSecretRule) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelSecretRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelSecretRule) SetId(v int32) {
	o.Id = &v
}

// GetLevel returns the Level field value
func (o *ModelSecretRule) GetLevel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *ModelSecretRule) GetLevelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *ModelSecretRule) SetLevel(v string) {
	o.Level = v
}

// GetMasked returns the Masked field value
func (o *ModelSecretRule) GetMasked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Masked
}

// GetMaskedOk returns a tuple with the Masked field value
// and a boolean to check if the value has been set.
func (o *ModelSecretRule) GetMaskedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Masked, true
}

// SetMasked sets field value
func (o *ModelSecretRule) SetMasked(v bool) {
	o.Masked = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ModelSecretRule) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSecretRule) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ModelSecretRule) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ModelSecretRule) SetName(v string) {
	o.Name = &v
}

// GetPart returns the Part field value if set, zero value otherwise.
func (o *ModelSecretRule) GetPart() string {
	if o == nil || IsNil(o.Part) {
		var ret string
		return ret
	}
	return *o.Part
}

// GetPartOk returns a tuple with the Part field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSecretRule) GetPartOk() (*string, bool) {
	if o == nil || IsNil(o.Part) {
		return nil, false
	}
	return o.Part, true
}

// HasPart returns a boolean if a field has been set.
func (o *ModelSecretRule) HasPart() bool {
	if o != nil && !IsNil(o.Part) {
		return true
	}

	return false
}

// SetPart gets a reference to the given string and assigns it to the Part field.
func (o *ModelSecretRule) SetPart(v string) {
	o.Part = &v
}

// GetSignatureToMatch returns the SignatureToMatch field value if set, zero value otherwise.
func (o *ModelSecretRule) GetSignatureToMatch() string {
	if o == nil || IsNil(o.SignatureToMatch) {
		var ret string
		return ret
	}
	return *o.SignatureToMatch
}

// GetSignatureToMatchOk returns a tuple with the SignatureToMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSecretRule) GetSignatureToMatchOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureToMatch) {
		return nil, false
	}
	return o.SignatureToMatch, true
}

// HasSignatureToMatch returns a boolean if a field has been set.
func (o *ModelSecretRule) HasSignatureToMatch() bool {
	if o != nil && !IsNil(o.SignatureToMatch) {
		return true
	}

	return false
}

// SetSignatureToMatch gets a reference to the given string and assigns it to the SignatureToMatch field.
func (o *ModelSecretRule) SetSignatureToMatch(v string) {
	o.SignatureToMatch = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ModelSecretRule) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelSecretRule) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ModelSecretRule) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

func (o ModelSecretRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSecretRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["level"] = o.Level
	toSerialize["masked"] = o.Masked
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Part) {
		toSerialize["part"] = o.Part
	}
	if !IsNil(o.SignatureToMatch) {
		toSerialize["signature_to_match"] = o.SignatureToMatch
	}
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

type NullableModelSecretRule struct {
	value *ModelSecretRule
	isSet bool
}

func (v NullableModelSecretRule) Get() *ModelSecretRule {
	return v.value
}

func (v *NullableModelSecretRule) Set(val *ModelSecretRule) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSecretRule) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSecretRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSecretRule(val *ModelSecretRule) *NullableModelSecretRule {
	return &NullableModelSecretRule{value: val, isSet: true}
}

func (v NullableModelSecretRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSecretRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


