/*
 * CAMAO.domains
 *
 * <h1>Additional info</h1> <h2><a href=\"/nexus-info\">Nexus Info</h2></a>  <h2><a href=\"/markdown/OPENAPIGATEWAY-USAGE.MD\">OPENAPIGATEWAY-USAGE.MD</h2></a> <h2><a href=\"/markdown/CHANGELOG.md\">CHANGELOG.md</h2></a> <h2><a href=\"/markdown/version.md\">version.md</h2></a>
 *
 * API version: v1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package stackit

import (
	"encoding/json"
)

// CreateRecordRequestModel struct for CreateRecordRequestModel
type CreateRecordRequestModel struct {
	Provider    string         `json:"provider"`
	Environment NullableString `json:"environment,omitempty"`
	Zone        string         `json:"zone"`
	Name        string         `json:"name"`
	Class       *RecordClass   `json:"class,omitempty"`
	Type        RecordType     `json:"type"`
	Data        string         `json:"data"`
	Ttl         *int32         `json:"ttl,omitempty"`
}

// NewCreateRecordRequestModel instantiates a new CreateRecordRequestModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRecordRequestModel(provider string, zone string, name string, type_ RecordType, data string) *CreateRecordRequestModel {
	this := CreateRecordRequestModel{}
	this.Provider = provider
	this.Zone = zone
	this.Name = name
	this.Type = type_
	this.Data = data
	return &this
}

// NewCreateRecordRequestModelWithDefaults instantiates a new CreateRecordRequestModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRecordRequestModelWithDefaults() *CreateRecordRequestModel {
	this := CreateRecordRequestModel{}
	return &this
}

// GetProvider returns the Provider field value
func (o *CreateRecordRequestModel) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *CreateRecordRequestModel) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *CreateRecordRequestModel) SetProvider(v string) {
	o.Provider = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateRecordRequestModel) GetEnvironment() string {
	if o == nil || o.Environment.Get() == nil {
		var ret string
		return ret
	}
	return *o.Environment.Get()
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRecordRequestModel) GetEnvironmentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Environment.Get(), o.Environment.IsSet()
}

// HasEnvironment returns a boolean if a field has been set.
func (o *CreateRecordRequestModel) HasEnvironment() bool {
	if o != nil && o.Environment.IsSet() {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given NullableString and assigns it to the Environment field.
func (o *CreateRecordRequestModel) SetEnvironment(v string) {
	o.Environment.Set(&v)
}

// SetEnvironmentNil sets the value for Environment to be an explicit nil
func (o *CreateRecordRequestModel) SetEnvironmentNil() {
	o.Environment.Set(nil)
}

// UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
func (o *CreateRecordRequestModel) UnsetEnvironment() {
	o.Environment.Unset()
}

// GetZone returns the Zone field value
func (o *CreateRecordRequestModel) GetZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Zone
}

// GetZoneOk returns a tuple with the Zone field value
// and a boolean to check if the value has been set.
func (o *CreateRecordRequestModel) GetZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Zone, true
}

// SetZone sets field value
func (o *CreateRecordRequestModel) SetZone(v string) {
	o.Zone = v
}

// GetName returns the Name field value
func (o *CreateRecordRequestModel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateRecordRequestModel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRecordRequestModel) SetName(v string) {
	o.Name = v
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *CreateRecordRequestModel) GetClass() RecordClass {
	if o == nil || o.Class == nil {
		var ret RecordClass
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecordRequestModel) GetClassOk() (*RecordClass, bool) {
	if o == nil || o.Class == nil {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *CreateRecordRequestModel) HasClass() bool {
	if o != nil && o.Class != nil {
		return true
	}

	return false
}

// SetClass gets a reference to the given RecordClass and assigns it to the Class field.
func (o *CreateRecordRequestModel) SetClass(v RecordClass) {
	o.Class = &v
}

// GetType returns the Type field value
func (o *CreateRecordRequestModel) GetType() RecordType {
	if o == nil {
		var ret RecordType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateRecordRequestModel) GetTypeOk() (*RecordType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateRecordRequestModel) SetType(v RecordType) {
	o.Type = v
}

// GetData returns the Data field value
func (o *CreateRecordRequestModel) GetData() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateRecordRequestModel) GetDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *CreateRecordRequestModel) SetData(v string) {
	o.Data = v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *CreateRecordRequestModel) GetTtl() int32 {
	if o == nil || o.Ttl == nil {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateRecordRequestModel) GetTtlOk() (*int32, bool) {
	if o == nil || o.Ttl == nil {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *CreateRecordRequestModel) HasTtl() bool {
	if o != nil && o.Ttl != nil {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *CreateRecordRequestModel) SetTtl(v int32) {
	o.Ttl = &v
}

func (o CreateRecordRequestModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if o.Environment.IsSet() {
		toSerialize["environment"] = o.Environment.Get()
	}
	if true {
		toSerialize["zone"] = o.Zone
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Class != nil {
		toSerialize["class"] = o.Class
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if o.Ttl != nil {
		toSerialize["ttl"] = o.Ttl
	}
	return json.Marshal(toSerialize)
}

type NullableCreateRecordRequestModel struct {
	value *CreateRecordRequestModel
	isSet bool
}

func (v NullableCreateRecordRequestModel) Get() *CreateRecordRequestModel {
	return v.value
}

func (v *NullableCreateRecordRequestModel) Set(val *CreateRecordRequestModel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRecordRequestModel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRecordRequestModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRecordRequestModel(val *CreateRecordRequestModel) *NullableCreateRecordRequestModel {
	return &NullableCreateRecordRequestModel{value: val, isSet: true}
}

func (v NullableCreateRecordRequestModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRecordRequestModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
