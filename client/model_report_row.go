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

// checks if the ReportRow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportRow{}

// ReportRow struct for ReportRow
type ReportRow struct {
	Entries map[string]string `json:"entries,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewReportRow instantiates a new ReportRow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportRow() *ReportRow {
	this := ReportRow{}
	return &this
}

// NewReportRowWithDefaults instantiates a new ReportRow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportRowWithDefaults() *ReportRow {
	this := ReportRow{}
	return &this
}

// GetEntries returns the Entries field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ReportRow) GetEntries() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}
	return o.Entries
}

// GetEntriesOk returns a tuple with the Entries field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportRow) GetEntriesOk() (*map[string]string, bool) {
	if o == nil || isNil(o.Entries) {
		return nil, false
	}
	return &o.Entries, true
}

// HasEntries returns a boolean if a field has been set.
func (o *ReportRow) HasEntries() bool {
	if o != nil && isNil(o.Entries) {
		return true
	}

	return false
}

// SetEntries gets a reference to the given map[string]string and assigns it to the Entries field.
func (o *ReportRow) SetEntries(v map[string]string) {
	o.Entries = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReportRow) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportRow) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReportRow) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ReportRow) SetId(v string) {
	o.Id = &v
}

func (o ReportRow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportRow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Entries != nil {
		toSerialize["entries"] = o.Entries
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableReportRow struct {
	value *ReportRow
	isSet bool
}

func (v NullableReportRow) Get() *ReportRow {
	return v.value
}

func (v *NullableReportRow) Set(val *ReportRow) {
	v.value = val
	v.isSet = true
}

func (v NullableReportRow) IsSet() bool {
	return v.isSet
}

func (v *NullableReportRow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportRow(val *ReportRow) *NullableReportRow {
	return &NullableReportRow{value: val, isSet: true}
}

func (v NullableReportRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportRow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


