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

// checks if the ModelGenerateLicenseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelGenerateLicenseRequest{}

// ModelGenerateLicenseRequest struct for ModelGenerateLicenseRequest
type ModelGenerateLicenseRequest struct {
	Company string `json:"company"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	ResendEmail bool `json:"resend_email"`
}

type _ModelGenerateLicenseRequest ModelGenerateLicenseRequest

// NewModelGenerateLicenseRequest instantiates a new ModelGenerateLicenseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelGenerateLicenseRequest(company string, email string, firstName string, lastName string, resendEmail bool) *ModelGenerateLicenseRequest {
	this := ModelGenerateLicenseRequest{}
	this.Company = company
	this.Email = email
	this.FirstName = firstName
	this.LastName = lastName
	this.ResendEmail = resendEmail
	return &this
}

// NewModelGenerateLicenseRequestWithDefaults instantiates a new ModelGenerateLicenseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelGenerateLicenseRequestWithDefaults() *ModelGenerateLicenseRequest {
	this := ModelGenerateLicenseRequest{}
	return &this
}

// GetCompany returns the Company field value
func (o *ModelGenerateLicenseRequest) GetCompany() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value
// and a boolean to check if the value has been set.
func (o *ModelGenerateLicenseRequest) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Company, true
}

// SetCompany sets field value
func (o *ModelGenerateLicenseRequest) SetCompany(v string) {
	o.Company = v
}

// GetEmail returns the Email field value
func (o *ModelGenerateLicenseRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ModelGenerateLicenseRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ModelGenerateLicenseRequest) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value
func (o *ModelGenerateLicenseRequest) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *ModelGenerateLicenseRequest) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *ModelGenerateLicenseRequest) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *ModelGenerateLicenseRequest) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *ModelGenerateLicenseRequest) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *ModelGenerateLicenseRequest) SetLastName(v string) {
	o.LastName = v
}

// GetResendEmail returns the ResendEmail field value
func (o *ModelGenerateLicenseRequest) GetResendEmail() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ResendEmail
}

// GetResendEmailOk returns a tuple with the ResendEmail field value
// and a boolean to check if the value has been set.
func (o *ModelGenerateLicenseRequest) GetResendEmailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResendEmail, true
}

// SetResendEmail sets field value
func (o *ModelGenerateLicenseRequest) SetResendEmail(v bool) {
	o.ResendEmail = v
}

func (o ModelGenerateLicenseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelGenerateLicenseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["company"] = o.Company
	toSerialize["email"] = o.Email
	toSerialize["first_name"] = o.FirstName
	toSerialize["last_name"] = o.LastName
	toSerialize["resend_email"] = o.ResendEmail
	return toSerialize, nil
}

func (o *ModelGenerateLicenseRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"company",
		"email",
		"first_name",
		"last_name",
		"resend_email",
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

	varModelGenerateLicenseRequest := _ModelGenerateLicenseRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelGenerateLicenseRequest)

	if err != nil {
		return err
	}

	*o = ModelGenerateLicenseRequest(varModelGenerateLicenseRequest)

	return err
}

type NullableModelGenerateLicenseRequest struct {
	value *ModelGenerateLicenseRequest
	isSet bool
}

func (v NullableModelGenerateLicenseRequest) Get() *ModelGenerateLicenseRequest {
	return v.value
}

func (v *NullableModelGenerateLicenseRequest) Set(val *ModelGenerateLicenseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelGenerateLicenseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelGenerateLicenseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelGenerateLicenseRequest(val *ModelGenerateLicenseRequest) *NullableModelGenerateLicenseRequest {
	return &NullableModelGenerateLicenseRequest{value: val, isSet: true}
}

func (v NullableModelGenerateLicenseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelGenerateLicenseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


