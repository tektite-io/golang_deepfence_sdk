/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.3.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ModelLicense type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelLicense{}

// ModelLicense struct for ModelLicense
type ModelLicense struct {
	CurrentHosts *int32 `json:"current_hosts,omitempty"`
	DeepfenceSupportEmail *string `json:"deepfence_support_email,omitempty"`
	Description *string `json:"description,omitempty"`
	EndDate *string `json:"end_date,omitempty"`
	IsActive *bool `json:"is_active,omitempty"`
	Key *string `json:"key,omitempty"`
	LicenseEmail *string `json:"license_email,omitempty"`
	LicenseEmailDomain *string `json:"license_email_domain,omitempty"`
	LicenseType *string `json:"license_type,omitempty"`
	Message *string `json:"message,omitempty"`
	NoOfCloudAccounts *int32 `json:"no_of_cloud_accounts,omitempty"`
	NoOfHosts *int32 `json:"no_of_hosts,omitempty"`
	NoOfImagesInRegistry *int32 `json:"no_of_images_in_registry,omitempty"`
	NoOfRegistries *int32 `json:"no_of_registries,omitempty"`
	NotificationThresholdPercentage *int32 `json:"notification_threshold_percentage,omitempty"`
	NotificationThresholdUpdatedAt *int32 `json:"notification_threshold_updated_at,omitempty"`
	RegistryCredentials *ModelRegistryCredentials `json:"registry_credentials,omitempty"`
	StartDate *string `json:"start_date,omitempty"`
}

// NewModelLicense instantiates a new ModelLicense object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelLicense() *ModelLicense {
	this := ModelLicense{}
	return &this
}

// NewModelLicenseWithDefaults instantiates a new ModelLicense object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelLicenseWithDefaults() *ModelLicense {
	this := ModelLicense{}
	return &this
}

// GetCurrentHosts returns the CurrentHosts field value if set, zero value otherwise.
func (o *ModelLicense) GetCurrentHosts() int32 {
	if o == nil || IsNil(o.CurrentHosts) {
		var ret int32
		return ret
	}
	return *o.CurrentHosts
}

// GetCurrentHostsOk returns a tuple with the CurrentHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLicense) GetCurrentHostsOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentHosts) {
		return nil, false
	}
	return o.CurrentHosts, true
}

// HasCurrentHosts returns a boolean if a field has been set.
func (o *ModelLicense) HasCurrentHosts() bool {
	if o != nil && !IsNil(o.CurrentHosts) {
		return true
	}

	return false
}

// SetCurrentHosts gets a reference to the given int32 and assigns it to the CurrentHosts field.
func (o *ModelLicense) SetCurrentHosts(v int32) {
	o.CurrentHosts = &v
}

// GetDeepfenceSupportEmail returns the DeepfenceSupportEmail field value if set, zero value otherwise.
func (o *ModelLicense) GetDeepfenceSupportEmail() string {
	if o == nil || IsNil(o.DeepfenceSupportEmail) {
		var ret string
		return ret
	}
	return *o.DeepfenceSupportEmail
}

// GetDeepfenceSupportEmailOk returns a tuple with the DeepfenceSupportEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLicense) GetDeepfenceSupportEmailOk() (*string, bool) {
	if o == nil || IsNil(o.DeepfenceSupportEmail) {
		return nil, false
	}
	return o.DeepfenceSupportEmail, true
}

// HasDeepfenceSupportEmail returns a boolean if a field has been set.
func (o *ModelLicense) HasDeepfenceSupportEmail() bool {
	if o != nil && !IsNil(o.DeepfenceSupportEmail) {
		return true
	}

	return false
}

// SetDeepfenceSupportEmail gets a reference to the given string and assigns it to the DeepfenceSupportEmail field.
func (o *ModelLicense) SetDeepfenceSupportEmail(v string) {
	o.DeepfenceSupportEmail = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelLicense) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLicense) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelLicense) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelLicense) SetDescription(v string) {
	o.Description = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ModelLicense) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLicense) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ModelLicense) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *ModelLicense) SetEndDate(v string) {
	o.EndDate = &v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *ModelLicense) GetIsActive() bool {
	if o == nil || IsNil(o.IsActive) {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLicense) GetIsActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.IsActive) {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *ModelLicense) HasIsActive() bool {
	if o != nil && !IsNil(o.IsActive) {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *ModelLicense) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ModelLicense) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLicense) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ModelLicense) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ModelLicense) SetKey(v string) {
	o.Key = &v
}

// GetLicenseEmail returns the LicenseEmail field value if set, zero value otherwise.
func (o *ModelLicense) GetLicenseEmail() string {
	if o == nil || IsNil(o.LicenseEmail) {
		var ret string
		return ret
	}
	return *o.LicenseEmail
}

// GetLicenseEmailOk returns a tuple with the LicenseEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLicense) GetLicenseEmailOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseEmail) {
		return nil, false
	}
	return o.LicenseEmail, true
}

// HasLicenseEmail returns a boolean if a field has been set.
func (o *ModelLicense) HasLicenseEmail() bool {
	if o != nil && !IsNil(o.LicenseEmail) {
		return true
	}

	return false
}

// SetLicenseEmail gets a reference to the given string and assigns it to the LicenseEmail field.
func (o *ModelLicense) SetLicenseEmail(v string) {
	o.LicenseEmail = &v
}

// GetLicenseEmailDomain returns the LicenseEmailDomain field value if set, zero value otherwise.
func (o *ModelLicense) GetLicenseEmailDomain() string {
	if o == nil || IsNil(o.LicenseEmailDomain) {
		var ret string
		return ret
	}
	return *o.LicenseEmailDomain
}

// GetLicenseEmailDomainOk returns a tuple with the LicenseEmailDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLicense) GetLicenseEmailDomainOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseEmailDomain) {
		return nil, false
	}
	return o.LicenseEmailDomain, true
}

// HasLicenseEmailDomain returns a boolean if a field has been set.
func (o *ModelLicense) HasLicenseEmailDomain() bool {
	if o != nil && !IsNil(o.LicenseEmailDomain) {
		return true
	}

	return false
}

// SetLicenseEmailDomain gets a reference to the given string and assigns it to the LicenseEmailDomain field.
func (o *ModelLicense) SetLicenseEmailDomain(v string) {
	o.LicenseEmailDomain = &v
}

// GetLicenseType returns the LicenseType field value if set, zero value otherwise.
func (o *ModelLicense) GetLicenseType() string {
	if o == nil || IsNil(o.LicenseType) {
		var ret string
		return ret
	}
	return *o.LicenseType
}

// GetLicenseTypeOk returns a tuple with the LicenseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLicense) GetLicenseTypeOk() (*string, bool) {
	if o == nil || IsNil(o.LicenseType) {
		return nil, false
	}
	return o.LicenseType, true
}

// HasLicenseType returns a boolean if a field has been set.
func (o *ModelLicense) HasLicenseType() bool {
	if o != nil && !IsNil(o.LicenseType) {
		return true
	}

	return false
}

// SetLicenseType gets a reference to the given string and assigns it to the LicenseType field.
func (o *ModelLicense) SetLicenseType(v string) {
	o.LicenseType = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ModelLicense) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLicense) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ModelLicense) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ModelLicense) SetMessage(v string) {
	o.Message = &v
}

// GetNoOfCloudAccounts returns the NoOfCloudAccounts field value if set, zero value otherwise.
func (o *ModelLicense) GetNoOfCloudAccounts() int32 {
	if o == nil || IsNil(o.NoOfCloudAccounts) {
		var ret int32
		return ret
	}
	return *o.NoOfCloudAccounts
}

// GetNoOfCloudAccountsOk returns a tuple with the NoOfCloudAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLicense) GetNoOfCloudAccountsOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfCloudAccounts) {
		return nil, false
	}
	return o.NoOfCloudAccounts, true
}

// HasNoOfCloudAccounts returns a boolean if a field has been set.
func (o *ModelLicense) HasNoOfCloudAccounts() bool {
	if o != nil && !IsNil(o.NoOfCloudAccounts) {
		return true
	}

	return false
}

// SetNoOfCloudAccounts gets a reference to the given int32 and assigns it to the NoOfCloudAccounts field.
func (o *ModelLicense) SetNoOfCloudAccounts(v int32) {
	o.NoOfCloudAccounts = &v
}

// GetNoOfHosts returns the NoOfHosts field value if set, zero value otherwise.
func (o *ModelLicense) GetNoOfHosts() int32 {
	if o == nil || IsNil(o.NoOfHosts) {
		var ret int32
		return ret
	}
	return *o.NoOfHosts
}

// GetNoOfHostsOk returns a tuple with the NoOfHosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLicense) GetNoOfHostsOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfHosts) {
		return nil, false
	}
	return o.NoOfHosts, true
}

// HasNoOfHosts returns a boolean if a field has been set.
func (o *ModelLicense) HasNoOfHosts() bool {
	if o != nil && !IsNil(o.NoOfHosts) {
		return true
	}

	return false
}

// SetNoOfHosts gets a reference to the given int32 and assigns it to the NoOfHosts field.
func (o *ModelLicense) SetNoOfHosts(v int32) {
	o.NoOfHosts = &v
}

// GetNoOfImagesInRegistry returns the NoOfImagesInRegistry field value if set, zero value otherwise.
func (o *ModelLicense) GetNoOfImagesInRegistry() int32 {
	if o == nil || IsNil(o.NoOfImagesInRegistry) {
		var ret int32
		return ret
	}
	return *o.NoOfImagesInRegistry
}

// GetNoOfImagesInRegistryOk returns a tuple with the NoOfImagesInRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLicense) GetNoOfImagesInRegistryOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfImagesInRegistry) {
		return nil, false
	}
	return o.NoOfImagesInRegistry, true
}

// HasNoOfImagesInRegistry returns a boolean if a field has been set.
func (o *ModelLicense) HasNoOfImagesInRegistry() bool {
	if o != nil && !IsNil(o.NoOfImagesInRegistry) {
		return true
	}

	return false
}

// SetNoOfImagesInRegistry gets a reference to the given int32 and assigns it to the NoOfImagesInRegistry field.
func (o *ModelLicense) SetNoOfImagesInRegistry(v int32) {
	o.NoOfImagesInRegistry = &v
}

// GetNoOfRegistries returns the NoOfRegistries field value if set, zero value otherwise.
func (o *ModelLicense) GetNoOfRegistries() int32 {
	if o == nil || IsNil(o.NoOfRegistries) {
		var ret int32
		return ret
	}
	return *o.NoOfRegistries
}

// GetNoOfRegistriesOk returns a tuple with the NoOfRegistries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLicense) GetNoOfRegistriesOk() (*int32, bool) {
	if o == nil || IsNil(o.NoOfRegistries) {
		return nil, false
	}
	return o.NoOfRegistries, true
}

// HasNoOfRegistries returns a boolean if a field has been set.
func (o *ModelLicense) HasNoOfRegistries() bool {
	if o != nil && !IsNil(o.NoOfRegistries) {
		return true
	}

	return false
}

// SetNoOfRegistries gets a reference to the given int32 and assigns it to the NoOfRegistries field.
func (o *ModelLicense) SetNoOfRegistries(v int32) {
	o.NoOfRegistries = &v
}

// GetNotificationThresholdPercentage returns the NotificationThresholdPercentage field value if set, zero value otherwise.
func (o *ModelLicense) GetNotificationThresholdPercentage() int32 {
	if o == nil || IsNil(o.NotificationThresholdPercentage) {
		var ret int32
		return ret
	}
	return *o.NotificationThresholdPercentage
}

// GetNotificationThresholdPercentageOk returns a tuple with the NotificationThresholdPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLicense) GetNotificationThresholdPercentageOk() (*int32, bool) {
	if o == nil || IsNil(o.NotificationThresholdPercentage) {
		return nil, false
	}
	return o.NotificationThresholdPercentage, true
}

// HasNotificationThresholdPercentage returns a boolean if a field has been set.
func (o *ModelLicense) HasNotificationThresholdPercentage() bool {
	if o != nil && !IsNil(o.NotificationThresholdPercentage) {
		return true
	}

	return false
}

// SetNotificationThresholdPercentage gets a reference to the given int32 and assigns it to the NotificationThresholdPercentage field.
func (o *ModelLicense) SetNotificationThresholdPercentage(v int32) {
	o.NotificationThresholdPercentage = &v
}

// GetNotificationThresholdUpdatedAt returns the NotificationThresholdUpdatedAt field value if set, zero value otherwise.
func (o *ModelLicense) GetNotificationThresholdUpdatedAt() int32 {
	if o == nil || IsNil(o.NotificationThresholdUpdatedAt) {
		var ret int32
		return ret
	}
	return *o.NotificationThresholdUpdatedAt
}

// GetNotificationThresholdUpdatedAtOk returns a tuple with the NotificationThresholdUpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLicense) GetNotificationThresholdUpdatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.NotificationThresholdUpdatedAt) {
		return nil, false
	}
	return o.NotificationThresholdUpdatedAt, true
}

// HasNotificationThresholdUpdatedAt returns a boolean if a field has been set.
func (o *ModelLicense) HasNotificationThresholdUpdatedAt() bool {
	if o != nil && !IsNil(o.NotificationThresholdUpdatedAt) {
		return true
	}

	return false
}

// SetNotificationThresholdUpdatedAt gets a reference to the given int32 and assigns it to the NotificationThresholdUpdatedAt field.
func (o *ModelLicense) SetNotificationThresholdUpdatedAt(v int32) {
	o.NotificationThresholdUpdatedAt = &v
}

// GetRegistryCredentials returns the RegistryCredentials field value if set, zero value otherwise.
func (o *ModelLicense) GetRegistryCredentials() ModelRegistryCredentials {
	if o == nil || IsNil(o.RegistryCredentials) {
		var ret ModelRegistryCredentials
		return ret
	}
	return *o.RegistryCredentials
}

// GetRegistryCredentialsOk returns a tuple with the RegistryCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLicense) GetRegistryCredentialsOk() (*ModelRegistryCredentials, bool) {
	if o == nil || IsNil(o.RegistryCredentials) {
		return nil, false
	}
	return o.RegistryCredentials, true
}

// HasRegistryCredentials returns a boolean if a field has been set.
func (o *ModelLicense) HasRegistryCredentials() bool {
	if o != nil && !IsNil(o.RegistryCredentials) {
		return true
	}

	return false
}

// SetRegistryCredentials gets a reference to the given ModelRegistryCredentials and assigns it to the RegistryCredentials field.
func (o *ModelLicense) SetRegistryCredentials(v ModelRegistryCredentials) {
	o.RegistryCredentials = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ModelLicense) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelLicense) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ModelLicense) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *ModelLicense) SetStartDate(v string) {
	o.StartDate = &v
}

func (o ModelLicense) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelLicense) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentHosts) {
		toSerialize["current_hosts"] = o.CurrentHosts
	}
	if !IsNil(o.DeepfenceSupportEmail) {
		toSerialize["deepfence_support_email"] = o.DeepfenceSupportEmail
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	if !IsNil(o.IsActive) {
		toSerialize["is_active"] = o.IsActive
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.LicenseEmail) {
		toSerialize["license_email"] = o.LicenseEmail
	}
	if !IsNil(o.LicenseEmailDomain) {
		toSerialize["license_email_domain"] = o.LicenseEmailDomain
	}
	if !IsNil(o.LicenseType) {
		toSerialize["license_type"] = o.LicenseType
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.NoOfCloudAccounts) {
		toSerialize["no_of_cloud_accounts"] = o.NoOfCloudAccounts
	}
	if !IsNil(o.NoOfHosts) {
		toSerialize["no_of_hosts"] = o.NoOfHosts
	}
	if !IsNil(o.NoOfImagesInRegistry) {
		toSerialize["no_of_images_in_registry"] = o.NoOfImagesInRegistry
	}
	if !IsNil(o.NoOfRegistries) {
		toSerialize["no_of_registries"] = o.NoOfRegistries
	}
	if !IsNil(o.NotificationThresholdPercentage) {
		toSerialize["notification_threshold_percentage"] = o.NotificationThresholdPercentage
	}
	if !IsNil(o.NotificationThresholdUpdatedAt) {
		toSerialize["notification_threshold_updated_at"] = o.NotificationThresholdUpdatedAt
	}
	if !IsNil(o.RegistryCredentials) {
		toSerialize["registry_credentials"] = o.RegistryCredentials
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	return toSerialize, nil
}

type NullableModelLicense struct {
	value *ModelLicense
	isSet bool
}

func (v NullableModelLicense) Get() *ModelLicense {
	return v.value
}

func (v *NullableModelLicense) Set(val *ModelLicense) {
	v.value = val
	v.isSet = true
}

func (v NullableModelLicense) IsSet() bool {
	return v.isSet
}

func (v *NullableModelLicense) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelLicense(val *ModelLicense) *NullableModelLicense {
	return &NullableModelLicense{value: val, isSet: true}
}

func (v NullableModelLicense) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelLicense) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


