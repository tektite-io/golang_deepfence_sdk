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

// checks if the ModelUserRegisterRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelUserRegisterRequest{}

// ModelUserRegisterRequest struct for ModelUserRegisterRequest
type ModelUserRegisterRequest struct {
	Company string `json:"company"`
	ConsoleUrl string `json:"console_url"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	IsTemporaryPassword *bool `json:"is_temporary_password,omitempty"`
	LastName string `json:"last_name"`
	Password string `json:"password"`
}

type _ModelUserRegisterRequest ModelUserRegisterRequest

// NewModelUserRegisterRequest instantiates a new ModelUserRegisterRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelUserRegisterRequest(company string, consoleUrl string, email string, firstName string, lastName string, password string) *ModelUserRegisterRequest {
	this := ModelUserRegisterRequest{}
	this.Company = company
	this.ConsoleUrl = consoleUrl
	this.Email = email
	this.FirstName = firstName
	this.LastName = lastName
	this.Password = password
	return &this
}

// NewModelUserRegisterRequestWithDefaults instantiates a new ModelUserRegisterRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelUserRegisterRequestWithDefaults() *ModelUserRegisterRequest {
	this := ModelUserRegisterRequest{}
	return &this
}

// GetCompany returns the Company field value
func (o *ModelUserRegisterRequest) GetCompany() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Company
}

// GetCompanyOk returns a tuple with the Company field value
// and a boolean to check if the value has been set.
func (o *ModelUserRegisterRequest) GetCompanyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Company, true
}

// SetCompany sets field value
func (o *ModelUserRegisterRequest) SetCompany(v string) {
	o.Company = v
}

// GetConsoleUrl returns the ConsoleUrl field value
func (o *ModelUserRegisterRequest) GetConsoleUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConsoleUrl
}

// GetConsoleUrlOk returns a tuple with the ConsoleUrl field value
// and a boolean to check if the value has been set.
func (o *ModelUserRegisterRequest) GetConsoleUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConsoleUrl, true
}

// SetConsoleUrl sets field value
func (o *ModelUserRegisterRequest) SetConsoleUrl(v string) {
	o.ConsoleUrl = v
}

// GetEmail returns the Email field value
func (o *ModelUserRegisterRequest) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *ModelUserRegisterRequest) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *ModelUserRegisterRequest) SetEmail(v string) {
	o.Email = v
}

// GetFirstName returns the FirstName field value
func (o *ModelUserRegisterRequest) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *ModelUserRegisterRequest) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *ModelUserRegisterRequest) SetFirstName(v string) {
	o.FirstName = v
}

// GetIsTemporaryPassword returns the IsTemporaryPassword field value if set, zero value otherwise.
func (o *ModelUserRegisterRequest) GetIsTemporaryPassword() bool {
	if o == nil || IsNil(o.IsTemporaryPassword) {
		var ret bool
		return ret
	}
	return *o.IsTemporaryPassword
}

// GetIsTemporaryPasswordOk returns a tuple with the IsTemporaryPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelUserRegisterRequest) GetIsTemporaryPasswordOk() (*bool, bool) {
	if o == nil || IsNil(o.IsTemporaryPassword) {
		return nil, false
	}
	return o.IsTemporaryPassword, true
}

// HasIsTemporaryPassword returns a boolean if a field has been set.
func (o *ModelUserRegisterRequest) HasIsTemporaryPassword() bool {
	if o != nil && !IsNil(o.IsTemporaryPassword) {
		return true
	}

	return false
}

// SetIsTemporaryPassword gets a reference to the given bool and assigns it to the IsTemporaryPassword field.
func (o *ModelUserRegisterRequest) SetIsTemporaryPassword(v bool) {
	o.IsTemporaryPassword = &v
}

// GetLastName returns the LastName field value
func (o *ModelUserRegisterRequest) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *ModelUserRegisterRequest) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *ModelUserRegisterRequest) SetLastName(v string) {
	o.LastName = v
}

// GetPassword returns the Password field value
func (o *ModelUserRegisterRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ModelUserRegisterRequest) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ModelUserRegisterRequest) SetPassword(v string) {
	o.Password = v
}

func (o ModelUserRegisterRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelUserRegisterRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["company"] = o.Company
	toSerialize["console_url"] = o.ConsoleUrl
	toSerialize["email"] = o.Email
	toSerialize["first_name"] = o.FirstName
	if !IsNil(o.IsTemporaryPassword) {
		toSerialize["is_temporary_password"] = o.IsTemporaryPassword
	}
	toSerialize["last_name"] = o.LastName
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *ModelUserRegisterRequest) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"company",
		"console_url",
		"email",
		"first_name",
		"last_name",
		"password",
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

	varModelUserRegisterRequest := _ModelUserRegisterRequest{}

	err = json.Unmarshal(bytes, &varModelUserRegisterRequest)

	if err != nil {
		return err
	}

	*o = ModelUserRegisterRequest(varModelUserRegisterRequest)

	return err
}

type NullableModelUserRegisterRequest struct {
	value *ModelUserRegisterRequest
	isSet bool
}

func (v NullableModelUserRegisterRequest) Get() *ModelUserRegisterRequest {
	return v.value
}

func (v *NullableModelUserRegisterRequest) Set(val *ModelUserRegisterRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelUserRegisterRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelUserRegisterRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelUserRegisterRequest(val *ModelUserRegisterRequest) *NullableModelUserRegisterRequest {
	return &NullableModelUserRegisterRequest{value: val, isSet: true}
}

func (v NullableModelUserRegisterRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelUserRegisterRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


