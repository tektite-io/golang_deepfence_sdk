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

// checks if the ModelLoginResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelLoginResponse{}

// ModelLoginResponse struct for ModelLoginResponse
type ModelLoginResponse struct {
	AccessToken string `json:"access_token"`
	EmailDomain string `json:"email_domain"`
	LicenseKey string `json:"license_key"`
	LicenseRegistered bool `json:"license_registered"`
	OnboardingRequired bool `json:"onboarding_required"`
	PasswordInvalidated bool `json:"password_invalidated"`
	RefreshToken string `json:"refresh_token"`
}

type _ModelLoginResponse ModelLoginResponse

// NewModelLoginResponse instantiates a new ModelLoginResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelLoginResponse(accessToken string, emailDomain string, licenseKey string, licenseRegistered bool, onboardingRequired bool, passwordInvalidated bool, refreshToken string) *ModelLoginResponse {
	this := ModelLoginResponse{}
	this.AccessToken = accessToken
	this.EmailDomain = emailDomain
	this.LicenseKey = licenseKey
	this.LicenseRegistered = licenseRegistered
	this.OnboardingRequired = onboardingRequired
	this.PasswordInvalidated = passwordInvalidated
	this.RefreshToken = refreshToken
	return &this
}

// NewModelLoginResponseWithDefaults instantiates a new ModelLoginResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelLoginResponseWithDefaults() *ModelLoginResponse {
	this := ModelLoginResponse{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *ModelLoginResponse) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *ModelLoginResponse) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *ModelLoginResponse) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetEmailDomain returns the EmailDomain field value
func (o *ModelLoginResponse) GetEmailDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailDomain
}

// GetEmailDomainOk returns a tuple with the EmailDomain field value
// and a boolean to check if the value has been set.
func (o *ModelLoginResponse) GetEmailDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailDomain, true
}

// SetEmailDomain sets field value
func (o *ModelLoginResponse) SetEmailDomain(v string) {
	o.EmailDomain = v
}

// GetLicenseKey returns the LicenseKey field value
func (o *ModelLoginResponse) GetLicenseKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LicenseKey
}

// GetLicenseKeyOk returns a tuple with the LicenseKey field value
// and a boolean to check if the value has been set.
func (o *ModelLoginResponse) GetLicenseKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseKey, true
}

// SetLicenseKey sets field value
func (o *ModelLoginResponse) SetLicenseKey(v string) {
	o.LicenseKey = v
}

// GetLicenseRegistered returns the LicenseRegistered field value
func (o *ModelLoginResponse) GetLicenseRegistered() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LicenseRegistered
}

// GetLicenseRegisteredOk returns a tuple with the LicenseRegistered field value
// and a boolean to check if the value has been set.
func (o *ModelLoginResponse) GetLicenseRegisteredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LicenseRegistered, true
}

// SetLicenseRegistered sets field value
func (o *ModelLoginResponse) SetLicenseRegistered(v bool) {
	o.LicenseRegistered = v
}

// GetOnboardingRequired returns the OnboardingRequired field value
func (o *ModelLoginResponse) GetOnboardingRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OnboardingRequired
}

// GetOnboardingRequiredOk returns a tuple with the OnboardingRequired field value
// and a boolean to check if the value has been set.
func (o *ModelLoginResponse) GetOnboardingRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OnboardingRequired, true
}

// SetOnboardingRequired sets field value
func (o *ModelLoginResponse) SetOnboardingRequired(v bool) {
	o.OnboardingRequired = v
}

// GetPasswordInvalidated returns the PasswordInvalidated field value
func (o *ModelLoginResponse) GetPasswordInvalidated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PasswordInvalidated
}

// GetPasswordInvalidatedOk returns a tuple with the PasswordInvalidated field value
// and a boolean to check if the value has been set.
func (o *ModelLoginResponse) GetPasswordInvalidatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordInvalidated, true
}

// SetPasswordInvalidated sets field value
func (o *ModelLoginResponse) SetPasswordInvalidated(v bool) {
	o.PasswordInvalidated = v
}

// GetRefreshToken returns the RefreshToken field value
func (o *ModelLoginResponse) GetRefreshToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RefreshToken
}

// GetRefreshTokenOk returns a tuple with the RefreshToken field value
// and a boolean to check if the value has been set.
func (o *ModelLoginResponse) GetRefreshTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RefreshToken, true
}

// SetRefreshToken sets field value
func (o *ModelLoginResponse) SetRefreshToken(v string) {
	o.RefreshToken = v
}

func (o ModelLoginResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelLoginResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["access_token"] = o.AccessToken
	toSerialize["email_domain"] = o.EmailDomain
	toSerialize["license_key"] = o.LicenseKey
	toSerialize["license_registered"] = o.LicenseRegistered
	toSerialize["onboarding_required"] = o.OnboardingRequired
	toSerialize["password_invalidated"] = o.PasswordInvalidated
	toSerialize["refresh_token"] = o.RefreshToken
	return toSerialize, nil
}

func (o *ModelLoginResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"access_token",
		"email_domain",
		"license_key",
		"license_registered",
		"onboarding_required",
		"password_invalidated",
		"refresh_token",
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

	varModelLoginResponse := _ModelLoginResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelLoginResponse)

	if err != nil {
		return err
	}

	*o = ModelLoginResponse(varModelLoginResponse)

	return err
}

type NullableModelLoginResponse struct {
	value *ModelLoginResponse
	isSet bool
}

func (v NullableModelLoginResponse) Get() *ModelLoginResponse {
	return v.value
}

func (v *NullableModelLoginResponse) Set(val *ModelLoginResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelLoginResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelLoginResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelLoginResponse(val *ModelLoginResponse) *NullableModelLoginResponse {
	return &NullableModelLoginResponse{value: val, isSet: true}
}

func (v NullableModelLoginResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelLoginResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


