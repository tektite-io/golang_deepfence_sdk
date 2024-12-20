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
)

// checks if the ModelExportReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelExportReport{}

// ModelExportReport struct for ModelExportReport
type ModelExportReport struct {
	CreatedAt *int32 `json:"created_at,omitempty"`
	Filters *string `json:"filters,omitempty"`
	FromTimestamp *int32 `json:"from_timestamp,omitempty"`
	ReportId *string `json:"report_id,omitempty"`
	Status *string `json:"status,omitempty"`
	StatusMessage *string `json:"status_message,omitempty"`
	StoragePath *string `json:"storage_path,omitempty"`
	ToTimestamp *int32 `json:"to_timestamp,omitempty"`
	Type *string `json:"type,omitempty"`
	UpdatedAt *int32 `json:"updated_at,omitempty"`
	Url *string `json:"url,omitempty"`
}

// NewModelExportReport instantiates a new ModelExportReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelExportReport() *ModelExportReport {
	this := ModelExportReport{}
	return &this
}

// NewModelExportReportWithDefaults instantiates a new ModelExportReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelExportReportWithDefaults() *ModelExportReport {
	this := ModelExportReport{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelExportReport) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelExportReport) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelExportReport) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *ModelExportReport) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *ModelExportReport) GetFilters() string {
	if o == nil || IsNil(o.Filters) {
		var ret string
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelExportReport) GetFiltersOk() (*string, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *ModelExportReport) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given string and assigns it to the Filters field.
func (o *ModelExportReport) SetFilters(v string) {
	o.Filters = &v
}

// GetFromTimestamp returns the FromTimestamp field value if set, zero value otherwise.
func (o *ModelExportReport) GetFromTimestamp() int32 {
	if o == nil || IsNil(o.FromTimestamp) {
		var ret int32
		return ret
	}
	return *o.FromTimestamp
}

// GetFromTimestampOk returns a tuple with the FromTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelExportReport) GetFromTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.FromTimestamp) {
		return nil, false
	}
	return o.FromTimestamp, true
}

// HasFromTimestamp returns a boolean if a field has been set.
func (o *ModelExportReport) HasFromTimestamp() bool {
	if o != nil && !IsNil(o.FromTimestamp) {
		return true
	}

	return false
}

// SetFromTimestamp gets a reference to the given int32 and assigns it to the FromTimestamp field.
func (o *ModelExportReport) SetFromTimestamp(v int32) {
	o.FromTimestamp = &v
}

// GetReportId returns the ReportId field value if set, zero value otherwise.
func (o *ModelExportReport) GetReportId() string {
	if o == nil || IsNil(o.ReportId) {
		var ret string
		return ret
	}
	return *o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelExportReport) GetReportIdOk() (*string, bool) {
	if o == nil || IsNil(o.ReportId) {
		return nil, false
	}
	return o.ReportId, true
}

// HasReportId returns a boolean if a field has been set.
func (o *ModelExportReport) HasReportId() bool {
	if o != nil && !IsNil(o.ReportId) {
		return true
	}

	return false
}

// SetReportId gets a reference to the given string and assigns it to the ReportId field.
func (o *ModelExportReport) SetReportId(v string) {
	o.ReportId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ModelExportReport) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelExportReport) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ModelExportReport) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ModelExportReport) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *ModelExportReport) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelExportReport) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *ModelExportReport) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *ModelExportReport) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetStoragePath returns the StoragePath field value if set, zero value otherwise.
func (o *ModelExportReport) GetStoragePath() string {
	if o == nil || IsNil(o.StoragePath) {
		var ret string
		return ret
	}
	return *o.StoragePath
}

// GetStoragePathOk returns a tuple with the StoragePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelExportReport) GetStoragePathOk() (*string, bool) {
	if o == nil || IsNil(o.StoragePath) {
		return nil, false
	}
	return o.StoragePath, true
}

// HasStoragePath returns a boolean if a field has been set.
func (o *ModelExportReport) HasStoragePath() bool {
	if o != nil && !IsNil(o.StoragePath) {
		return true
	}

	return false
}

// SetStoragePath gets a reference to the given string and assigns it to the StoragePath field.
func (o *ModelExportReport) SetStoragePath(v string) {
	o.StoragePath = &v
}

// GetToTimestamp returns the ToTimestamp field value if set, zero value otherwise.
func (o *ModelExportReport) GetToTimestamp() int32 {
	if o == nil || IsNil(o.ToTimestamp) {
		var ret int32
		return ret
	}
	return *o.ToTimestamp
}

// GetToTimestampOk returns a tuple with the ToTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelExportReport) GetToTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.ToTimestamp) {
		return nil, false
	}
	return o.ToTimestamp, true
}

// HasToTimestamp returns a boolean if a field has been set.
func (o *ModelExportReport) HasToTimestamp() bool {
	if o != nil && !IsNil(o.ToTimestamp) {
		return true
	}

	return false
}

// SetToTimestamp gets a reference to the given int32 and assigns it to the ToTimestamp field.
func (o *ModelExportReport) SetToTimestamp(v int32) {
	o.ToTimestamp = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ModelExportReport) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelExportReport) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ModelExportReport) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ModelExportReport) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ModelExportReport) GetUpdatedAt() int32 {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret int32
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelExportReport) GetUpdatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ModelExportReport) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int32 and assigns it to the UpdatedAt field.
func (o *ModelExportReport) SetUpdatedAt(v int32) {
	o.UpdatedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ModelExportReport) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelExportReport) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ModelExportReport) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *ModelExportReport) SetUrl(v string) {
	o.Url = &v
}

func (o ModelExportReport) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelExportReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.FromTimestamp) {
		toSerialize["from_timestamp"] = o.FromTimestamp
	}
	if !IsNil(o.ReportId) {
		toSerialize["report_id"] = o.ReportId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.StatusMessage) {
		toSerialize["status_message"] = o.StatusMessage
	}
	if !IsNil(o.StoragePath) {
		toSerialize["storage_path"] = o.StoragePath
	}
	if !IsNil(o.ToTimestamp) {
		toSerialize["to_timestamp"] = o.ToTimestamp
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableModelExportReport struct {
	value *ModelExportReport
	isSet bool
}

func (v NullableModelExportReport) Get() *ModelExportReport {
	return v.value
}

func (v *NullableModelExportReport) Set(val *ModelExportReport) {
	v.value = val
	v.isSet = true
}

func (v NullableModelExportReport) IsSet() bool {
	return v.isSet
}

func (v *NullableModelExportReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelExportReport(val *ModelExportReport) *NullableModelExportReport {
	return &NullableModelExportReport{value: val, isSet: true}
}

func (v NullableModelExportReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelExportReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


