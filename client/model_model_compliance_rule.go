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

// checks if the ModelComplianceRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelComplianceRule{}

// ModelComplianceRule struct for ModelComplianceRule
type ModelComplianceRule struct {
	Description string `json:"description"`
	Masked bool `json:"masked"`
	TestCategory string `json:"test_category"`
	TestDesc string `json:"test_desc"`
	TestNumber string `json:"test_number"`
	TestRationale string `json:"test_rationale"`
	TestSeverity string `json:"test_severity"`
	UpdatedAt int32 `json:"updated_at"`
}

type _ModelComplianceRule ModelComplianceRule

// NewModelComplianceRule instantiates a new ModelComplianceRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelComplianceRule(description string, masked bool, testCategory string, testDesc string, testNumber string, testRationale string, testSeverity string, updatedAt int32) *ModelComplianceRule {
	this := ModelComplianceRule{}
	this.Description = description
	this.Masked = masked
	this.TestCategory = testCategory
	this.TestDesc = testDesc
	this.TestNumber = testNumber
	this.TestRationale = testRationale
	this.TestSeverity = testSeverity
	this.UpdatedAt = updatedAt
	return &this
}

// NewModelComplianceRuleWithDefaults instantiates a new ModelComplianceRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelComplianceRuleWithDefaults() *ModelComplianceRule {
	this := ModelComplianceRule{}
	return &this
}

// GetDescription returns the Description field value
func (o *ModelComplianceRule) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ModelComplianceRule) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ModelComplianceRule) SetDescription(v string) {
	o.Description = v
}

// GetMasked returns the Masked field value
func (o *ModelComplianceRule) GetMasked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Masked
}

// GetMaskedOk returns a tuple with the Masked field value
// and a boolean to check if the value has been set.
func (o *ModelComplianceRule) GetMaskedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Masked, true
}

// SetMasked sets field value
func (o *ModelComplianceRule) SetMasked(v bool) {
	o.Masked = v
}

// GetTestCategory returns the TestCategory field value
func (o *ModelComplianceRule) GetTestCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestCategory
}

// GetTestCategoryOk returns a tuple with the TestCategory field value
// and a boolean to check if the value has been set.
func (o *ModelComplianceRule) GetTestCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestCategory, true
}

// SetTestCategory sets field value
func (o *ModelComplianceRule) SetTestCategory(v string) {
	o.TestCategory = v
}

// GetTestDesc returns the TestDesc field value
func (o *ModelComplianceRule) GetTestDesc() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestDesc
}

// GetTestDescOk returns a tuple with the TestDesc field value
// and a boolean to check if the value has been set.
func (o *ModelComplianceRule) GetTestDescOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestDesc, true
}

// SetTestDesc sets field value
func (o *ModelComplianceRule) SetTestDesc(v string) {
	o.TestDesc = v
}

// GetTestNumber returns the TestNumber field value
func (o *ModelComplianceRule) GetTestNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestNumber
}

// GetTestNumberOk returns a tuple with the TestNumber field value
// and a boolean to check if the value has been set.
func (o *ModelComplianceRule) GetTestNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestNumber, true
}

// SetTestNumber sets field value
func (o *ModelComplianceRule) SetTestNumber(v string) {
	o.TestNumber = v
}

// GetTestRationale returns the TestRationale field value
func (o *ModelComplianceRule) GetTestRationale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestRationale
}

// GetTestRationaleOk returns a tuple with the TestRationale field value
// and a boolean to check if the value has been set.
func (o *ModelComplianceRule) GetTestRationaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestRationale, true
}

// SetTestRationale sets field value
func (o *ModelComplianceRule) SetTestRationale(v string) {
	o.TestRationale = v
}

// GetTestSeverity returns the TestSeverity field value
func (o *ModelComplianceRule) GetTestSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TestSeverity
}

// GetTestSeverityOk returns a tuple with the TestSeverity field value
// and a boolean to check if the value has been set.
func (o *ModelComplianceRule) GetTestSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestSeverity, true
}

// SetTestSeverity sets field value
func (o *ModelComplianceRule) SetTestSeverity(v string) {
	o.TestSeverity = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ModelComplianceRule) GetUpdatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ModelComplianceRule) GetUpdatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ModelComplianceRule) SetUpdatedAt(v int32) {
	o.UpdatedAt = v
}

func (o ModelComplianceRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelComplianceRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["masked"] = o.Masked
	toSerialize["test_category"] = o.TestCategory
	toSerialize["test_desc"] = o.TestDesc
	toSerialize["test_number"] = o.TestNumber
	toSerialize["test_rationale"] = o.TestRationale
	toSerialize["test_severity"] = o.TestSeverity
	toSerialize["updated_at"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *ModelComplianceRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"description",
		"masked",
		"test_category",
		"test_desc",
		"test_number",
		"test_rationale",
		"test_severity",
		"updated_at",
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

	varModelComplianceRule := _ModelComplianceRule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelComplianceRule)

	if err != nil {
		return err
	}

	*o = ModelComplianceRule(varModelComplianceRule)

	return err
}

type NullableModelComplianceRule struct {
	value *ModelComplianceRule
	isSet bool
}

func (v NullableModelComplianceRule) Get() *ModelComplianceRule {
	return v.value
}

func (v *NullableModelComplianceRule) Set(val *ModelComplianceRule) {
	v.value = val
	v.isSet = true
}

func (v NullableModelComplianceRule) IsSet() bool {
	return v.isSet
}

func (v *NullableModelComplianceRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelComplianceRule(val *ModelComplianceRule) *NullableModelComplianceRule {
	return &NullableModelComplianceRule{value: val, isSet: true}
}

func (v NullableModelComplianceRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelComplianceRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


