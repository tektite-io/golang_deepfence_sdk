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

// checks if the IngestersMetaRules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestersMetaRules{}

// IngestersMetaRules struct for IngestersMetaRules
type IngestersMetaRules struct {
	Author *string `json:"author,omitempty"`
	Date *string `json:"date,omitempty"`
	Description *string `json:"description,omitempty"`
	FileSeverity *string `json:"file_severity,omitempty"`
	Filetype *string `json:"filetype,omitempty"`
	Info *string `json:"info,omitempty"`
	Reference *string `json:"reference,omitempty"`
	RuleId *string `json:"rule_id,omitempty"`
	RuleName *string `json:"rule_name,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewIngestersMetaRules instantiates a new IngestersMetaRules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestersMetaRules() *IngestersMetaRules {
	this := IngestersMetaRules{}
	return &this
}

// NewIngestersMetaRulesWithDefaults instantiates a new IngestersMetaRules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestersMetaRulesWithDefaults() *IngestersMetaRules {
	this := IngestersMetaRules{}
	return &this
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *IngestersMetaRules) GetAuthor() string {
	if o == nil || IsNil(o.Author) {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersMetaRules) GetAuthorOk() (*string, bool) {
	if o == nil || IsNil(o.Author) {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *IngestersMetaRules) HasAuthor() bool {
	if o != nil && !IsNil(o.Author) {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *IngestersMetaRules) SetAuthor(v string) {
	o.Author = &v
}

// GetDate returns the Date field value if set, zero value otherwise.
func (o *IngestersMetaRules) GetDate() string {
	if o == nil || IsNil(o.Date) {
		var ret string
		return ret
	}
	return *o.Date
}

// GetDateOk returns a tuple with the Date field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersMetaRules) GetDateOk() (*string, bool) {
	if o == nil || IsNil(o.Date) {
		return nil, false
	}
	return o.Date, true
}

// HasDate returns a boolean if a field has been set.
func (o *IngestersMetaRules) HasDate() bool {
	if o != nil && !IsNil(o.Date) {
		return true
	}

	return false
}

// SetDate gets a reference to the given string and assigns it to the Date field.
func (o *IngestersMetaRules) SetDate(v string) {
	o.Date = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IngestersMetaRules) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersMetaRules) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IngestersMetaRules) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IngestersMetaRules) SetDescription(v string) {
	o.Description = &v
}

// GetFileSeverity returns the FileSeverity field value if set, zero value otherwise.
func (o *IngestersMetaRules) GetFileSeverity() string {
	if o == nil || IsNil(o.FileSeverity) {
		var ret string
		return ret
	}
	return *o.FileSeverity
}

// GetFileSeverityOk returns a tuple with the FileSeverity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersMetaRules) GetFileSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.FileSeverity) {
		return nil, false
	}
	return o.FileSeverity, true
}

// HasFileSeverity returns a boolean if a field has been set.
func (o *IngestersMetaRules) HasFileSeverity() bool {
	if o != nil && !IsNil(o.FileSeverity) {
		return true
	}

	return false
}

// SetFileSeverity gets a reference to the given string and assigns it to the FileSeverity field.
func (o *IngestersMetaRules) SetFileSeverity(v string) {
	o.FileSeverity = &v
}

// GetFiletype returns the Filetype field value if set, zero value otherwise.
func (o *IngestersMetaRules) GetFiletype() string {
	if o == nil || IsNil(o.Filetype) {
		var ret string
		return ret
	}
	return *o.Filetype
}

// GetFiletypeOk returns a tuple with the Filetype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersMetaRules) GetFiletypeOk() (*string, bool) {
	if o == nil || IsNil(o.Filetype) {
		return nil, false
	}
	return o.Filetype, true
}

// HasFiletype returns a boolean if a field has been set.
func (o *IngestersMetaRules) HasFiletype() bool {
	if o != nil && !IsNil(o.Filetype) {
		return true
	}

	return false
}

// SetFiletype gets a reference to the given string and assigns it to the Filetype field.
func (o *IngestersMetaRules) SetFiletype(v string) {
	o.Filetype = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *IngestersMetaRules) GetInfo() string {
	if o == nil || IsNil(o.Info) {
		var ret string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersMetaRules) GetInfoOk() (*string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *IngestersMetaRules) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given string and assigns it to the Info field.
func (o *IngestersMetaRules) SetInfo(v string) {
	o.Info = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *IngestersMetaRules) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersMetaRules) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *IngestersMetaRules) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *IngestersMetaRules) SetReference(v string) {
	o.Reference = &v
}

// GetRuleId returns the RuleId field value if set, zero value otherwise.
func (o *IngestersMetaRules) GetRuleId() string {
	if o == nil || IsNil(o.RuleId) {
		var ret string
		return ret
	}
	return *o.RuleId
}

// GetRuleIdOk returns a tuple with the RuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersMetaRules) GetRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.RuleId) {
		return nil, false
	}
	return o.RuleId, true
}

// HasRuleId returns a boolean if a field has been set.
func (o *IngestersMetaRules) HasRuleId() bool {
	if o != nil && !IsNil(o.RuleId) {
		return true
	}

	return false
}

// SetRuleId gets a reference to the given string and assigns it to the RuleId field.
func (o *IngestersMetaRules) SetRuleId(v string) {
	o.RuleId = &v
}

// GetRuleName returns the RuleName field value if set, zero value otherwise.
func (o *IngestersMetaRules) GetRuleName() string {
	if o == nil || IsNil(o.RuleName) {
		var ret string
		return ret
	}
	return *o.RuleName
}

// GetRuleNameOk returns a tuple with the RuleName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersMetaRules) GetRuleNameOk() (*string, bool) {
	if o == nil || IsNil(o.RuleName) {
		return nil, false
	}
	return o.RuleName, true
}

// HasRuleName returns a boolean if a field has been set.
func (o *IngestersMetaRules) HasRuleName() bool {
	if o != nil && !IsNil(o.RuleName) {
		return true
	}

	return false
}

// SetRuleName gets a reference to the given string and assigns it to the RuleName field.
func (o *IngestersMetaRules) SetRuleName(v string) {
	o.RuleName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *IngestersMetaRules) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersMetaRules) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *IngestersMetaRules) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *IngestersMetaRules) SetVersion(v string) {
	o.Version = &v
}

func (o IngestersMetaRules) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestersMetaRules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Author) {
		toSerialize["author"] = o.Author
	}
	if !IsNil(o.Date) {
		toSerialize["date"] = o.Date
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.FileSeverity) {
		toSerialize["file_severity"] = o.FileSeverity
	}
	if !IsNil(o.Filetype) {
		toSerialize["filetype"] = o.Filetype
	}
	if !IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.RuleId) {
		toSerialize["rule_id"] = o.RuleId
	}
	if !IsNil(o.RuleName) {
		toSerialize["rule_name"] = o.RuleName
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableIngestersMetaRules struct {
	value *IngestersMetaRules
	isSet bool
}

func (v NullableIngestersMetaRules) Get() *IngestersMetaRules {
	return v.value
}

func (v *NullableIngestersMetaRules) Set(val *IngestersMetaRules) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestersMetaRules) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestersMetaRules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestersMetaRules(val *IngestersMetaRules) *NullableIngestersMetaRules {
	return &NullableIngestersMetaRules{value: val, isSet: true}
}

func (v NullableIngestersMetaRules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestersMetaRules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


