/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.2.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"time"
)

// checks if the PostgresqlDbScheduler type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostgresqlDbScheduler{}

// PostgresqlDbScheduler struct for PostgresqlDbScheduler
type PostgresqlDbScheduler struct {
	Action *string `json:"action,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	CronExpr *string `json:"cron_expr,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *int32 `json:"id,omitempty"`
	IsEnabled *bool `json:"is_enabled,omitempty"`
	IsSystem *bool `json:"is_system,omitempty"`
	LastRanAt map[string]interface{} `json:"last_ran_at,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewPostgresqlDbScheduler instantiates a new PostgresqlDbScheduler object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostgresqlDbScheduler() *PostgresqlDbScheduler {
	this := PostgresqlDbScheduler{}
	return &this
}

// NewPostgresqlDbSchedulerWithDefaults instantiates a new PostgresqlDbScheduler object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostgresqlDbSchedulerWithDefaults() *PostgresqlDbScheduler {
	this := PostgresqlDbScheduler{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *PostgresqlDbScheduler) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDbScheduler) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *PostgresqlDbScheduler) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *PostgresqlDbScheduler) SetAction(v string) {
	o.Action = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PostgresqlDbScheduler) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDbScheduler) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PostgresqlDbScheduler) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PostgresqlDbScheduler) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCronExpr returns the CronExpr field value if set, zero value otherwise.
func (o *PostgresqlDbScheduler) GetCronExpr() string {
	if o == nil || IsNil(o.CronExpr) {
		var ret string
		return ret
	}
	return *o.CronExpr
}

// GetCronExprOk returns a tuple with the CronExpr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDbScheduler) GetCronExprOk() (*string, bool) {
	if o == nil || IsNil(o.CronExpr) {
		return nil, false
	}
	return o.CronExpr, true
}

// HasCronExpr returns a boolean if a field has been set.
func (o *PostgresqlDbScheduler) HasCronExpr() bool {
	if o != nil && !IsNil(o.CronExpr) {
		return true
	}

	return false
}

// SetCronExpr gets a reference to the given string and assigns it to the CronExpr field.
func (o *PostgresqlDbScheduler) SetCronExpr(v string) {
	o.CronExpr = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PostgresqlDbScheduler) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDbScheduler) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PostgresqlDbScheduler) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PostgresqlDbScheduler) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PostgresqlDbScheduler) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDbScheduler) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PostgresqlDbScheduler) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *PostgresqlDbScheduler) SetId(v int32) {
	o.Id = &v
}

// GetIsEnabled returns the IsEnabled field value if set, zero value otherwise.
func (o *PostgresqlDbScheduler) GetIsEnabled() bool {
	if o == nil || IsNil(o.IsEnabled) {
		var ret bool
		return ret
	}
	return *o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDbScheduler) GetIsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsEnabled) {
		return nil, false
	}
	return o.IsEnabled, true
}

// HasIsEnabled returns a boolean if a field has been set.
func (o *PostgresqlDbScheduler) HasIsEnabled() bool {
	if o != nil && !IsNil(o.IsEnabled) {
		return true
	}

	return false
}

// SetIsEnabled gets a reference to the given bool and assigns it to the IsEnabled field.
func (o *PostgresqlDbScheduler) SetIsEnabled(v bool) {
	o.IsEnabled = &v
}

// GetIsSystem returns the IsSystem field value if set, zero value otherwise.
func (o *PostgresqlDbScheduler) GetIsSystem() bool {
	if o == nil || IsNil(o.IsSystem) {
		var ret bool
		return ret
	}
	return *o.IsSystem
}

// GetIsSystemOk returns a tuple with the IsSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDbScheduler) GetIsSystemOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSystem) {
		return nil, false
	}
	return o.IsSystem, true
}

// HasIsSystem returns a boolean if a field has been set.
func (o *PostgresqlDbScheduler) HasIsSystem() bool {
	if o != nil && !IsNil(o.IsSystem) {
		return true
	}

	return false
}

// SetIsSystem gets a reference to the given bool and assigns it to the IsSystem field.
func (o *PostgresqlDbScheduler) SetIsSystem(v bool) {
	o.IsSystem = &v
}

// GetLastRanAt returns the LastRanAt field value if set, zero value otherwise.
func (o *PostgresqlDbScheduler) GetLastRanAt() map[string]interface{} {
	if o == nil || IsNil(o.LastRanAt) {
		var ret map[string]interface{}
		return ret
	}
	return o.LastRanAt
}

// GetLastRanAtOk returns a tuple with the LastRanAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDbScheduler) GetLastRanAtOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.LastRanAt) {
		return map[string]interface{}{}, false
	}
	return o.LastRanAt, true
}

// HasLastRanAt returns a boolean if a field has been set.
func (o *PostgresqlDbScheduler) HasLastRanAt() bool {
	if o != nil && !IsNil(o.LastRanAt) {
		return true
	}

	return false
}

// SetLastRanAt gets a reference to the given map[string]interface{} and assigns it to the LastRanAt field.
func (o *PostgresqlDbScheduler) SetLastRanAt(v map[string]interface{}) {
	o.LastRanAt = v
}

// GetPayload returns the Payload field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PostgresqlDbScheduler) GetPayload() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PostgresqlDbScheduler) GetPayloadOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return &o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *PostgresqlDbScheduler) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given interface{} and assigns it to the Payload field.
func (o *PostgresqlDbScheduler) SetPayload(v interface{}) {
	o.Payload = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PostgresqlDbScheduler) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDbScheduler) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PostgresqlDbScheduler) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PostgresqlDbScheduler) SetStatus(v string) {
	o.Status = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PostgresqlDbScheduler) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostgresqlDbScheduler) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PostgresqlDbScheduler) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PostgresqlDbScheduler) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o PostgresqlDbScheduler) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostgresqlDbScheduler) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.CronExpr) {
		toSerialize["cron_expr"] = o.CronExpr
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IsEnabled) {
		toSerialize["is_enabled"] = o.IsEnabled
	}
	if !IsNil(o.IsSystem) {
		toSerialize["is_system"] = o.IsSystem
	}
	if !IsNil(o.LastRanAt) {
		toSerialize["last_ran_at"] = o.LastRanAt
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullablePostgresqlDbScheduler struct {
	value *PostgresqlDbScheduler
	isSet bool
}

func (v NullablePostgresqlDbScheduler) Get() *PostgresqlDbScheduler {
	return v.value
}

func (v *NullablePostgresqlDbScheduler) Set(val *PostgresqlDbScheduler) {
	v.value = val
	v.isSet = true
}

func (v NullablePostgresqlDbScheduler) IsSet() bool {
	return v.isSet
}

func (v *NullablePostgresqlDbScheduler) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostgresqlDbScheduler(val *PostgresqlDbScheduler) *NullablePostgresqlDbScheduler {
	return &NullablePostgresqlDbScheduler{value: val, isSet: true}
}

func (v NullablePostgresqlDbScheduler) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostgresqlDbScheduler) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


