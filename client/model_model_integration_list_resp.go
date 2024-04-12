/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.2.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ModelIntegrationListResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelIntegrationListResp{}

// ModelIntegrationListResp struct for ModelIntegrationListResp
type ModelIntegrationListResp struct {
	Config map[string]interface{} `json:"config,omitempty"`
	Filters *ModelIntegrationFilters `json:"filters,omitempty"`
	Id *int32 `json:"id,omitempty"`
	IntegrationType *string `json:"integration_type,omitempty"`
	LastErrorMsg *string `json:"last_error_msg,omitempty"`
	NotificationType *string `json:"notification_type,omitempty"`
}

// NewModelIntegrationListResp instantiates a new ModelIntegrationListResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelIntegrationListResp() *ModelIntegrationListResp {
	this := ModelIntegrationListResp{}
	return &this
}

// NewModelIntegrationListRespWithDefaults instantiates a new ModelIntegrationListResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelIntegrationListRespWithDefaults() *ModelIntegrationListResp {
	this := ModelIntegrationListResp{}
	return &this
}

// GetConfig returns the Config field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelIntegrationListResp) GetConfig() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelIntegrationListResp) GetConfigOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Config) {
		return map[string]interface{}{}, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *ModelIntegrationListResp) HasConfig() bool {
	if o != nil && !IsNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *ModelIntegrationListResp) SetConfig(v map[string]interface{}) {
	o.Config = v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ModelIntegrationListResp) GetFilters() ModelIntegrationFilters {
	if o == nil || IsNil(o.Filters) {
		var ret ModelIntegrationFilters
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelIntegrationListResp) GetFiltersOk() (*ModelIntegrationFilters, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ModelIntegrationListResp) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given ModelIntegrationFilters and assigns it to the Filters field.
func (o *ModelIntegrationListResp) SetFilters(v ModelIntegrationFilters) {
	o.Filters = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelIntegrationListResp) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelIntegrationListResp) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelIntegrationListResp) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelIntegrationListResp) SetId(v int32) {
	o.Id = &v
}

// GetIntegrationType returns the IntegrationType field value if set, zero value otherwise.
func (o *ModelIntegrationListResp) GetIntegrationType() string {
	if o == nil || IsNil(o.IntegrationType) {
		var ret string
		return ret
	}
	return *o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelIntegrationListResp) GetIntegrationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IntegrationType) {
		return nil, false
	}
	return o.IntegrationType, true
}

// HasIntegrationType returns a boolean if a field has been set.
func (o *ModelIntegrationListResp) HasIntegrationType() bool {
	if o != nil && !IsNil(o.IntegrationType) {
		return true
	}

	return false
}

// SetIntegrationType gets a reference to the given string and assigns it to the IntegrationType field.
func (o *ModelIntegrationListResp) SetIntegrationType(v string) {
	o.IntegrationType = &v
}

// GetLastErrorMsg returns the LastErrorMsg field value if set, zero value otherwise.
func (o *ModelIntegrationListResp) GetLastErrorMsg() string {
	if o == nil || IsNil(o.LastErrorMsg) {
		var ret string
		return ret
	}
	return *o.LastErrorMsg
}

// GetLastErrorMsgOk returns a tuple with the LastErrorMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelIntegrationListResp) GetLastErrorMsgOk() (*string, bool) {
	if o == nil || IsNil(o.LastErrorMsg) {
		return nil, false
	}
	return o.LastErrorMsg, true
}

// HasLastErrorMsg returns a boolean if a field has been set.
func (o *ModelIntegrationListResp) HasLastErrorMsg() bool {
	if o != nil && !IsNil(o.LastErrorMsg) {
		return true
	}

	return false
}

// SetLastErrorMsg gets a reference to the given string and assigns it to the LastErrorMsg field.
func (o *ModelIntegrationListResp) SetLastErrorMsg(v string) {
	o.LastErrorMsg = &v
}

// GetNotificationType returns the NotificationType field value if set, zero value otherwise.
func (o *ModelIntegrationListResp) GetNotificationType() string {
	if o == nil || IsNil(o.NotificationType) {
		var ret string
		return ret
	}
	return *o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelIntegrationListResp) GetNotificationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationType) {
		return nil, false
	}
	return o.NotificationType, true
}

// HasNotificationType returns a boolean if a field has been set.
func (o *ModelIntegrationListResp) HasNotificationType() bool {
	if o != nil && !IsNil(o.NotificationType) {
		return true
	}

	return false
}

// SetNotificationType gets a reference to the given string and assigns it to the NotificationType field.
func (o *ModelIntegrationListResp) SetNotificationType(v string) {
	o.NotificationType = &v
}

func (o ModelIntegrationListResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelIntegrationListResp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Config != nil {
		toSerialize["config"] = o.Config
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IntegrationType) {
		toSerialize["integration_type"] = o.IntegrationType
	}
	if !IsNil(o.LastErrorMsg) {
		toSerialize["last_error_msg"] = o.LastErrorMsg
	}
	if !IsNil(o.NotificationType) {
		toSerialize["notification_type"] = o.NotificationType
	}
	return toSerialize, nil
}

type NullableModelIntegrationListResp struct {
	value *ModelIntegrationListResp
	isSet bool
}

func (v NullableModelIntegrationListResp) Get() *ModelIntegrationListResp {
	return v.value
}

func (v *NullableModelIntegrationListResp) Set(val *ModelIntegrationListResp) {
	v.value = val
	v.isSet = true
}

func (v NullableModelIntegrationListResp) IsSet() bool {
	return v.isSet
}

func (v *NullableModelIntegrationListResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelIntegrationListResp(val *ModelIntegrationListResp) *NullableModelIntegrationListResp {
	return &NullableModelIntegrationListResp{value: val, isSet: true}
}

func (v NullableModelIntegrationListResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelIntegrationListResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


