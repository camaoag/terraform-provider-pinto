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
	"fmt"
)

// RecordClass Resource record classes  as defined in  <see href=\"https://tools.ietf.org/html/rfc1035#section-3.2.4\">rfc1035</see>
type RecordClass string

// List of RecordClass
const (
	IN RecordClass = "IN"
	CS RecordClass = "CS"
	CH RecordClass = "CH"
	HS RecordClass = "HS"
)

func (v *RecordClass) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RecordClass(value)
	for _, existing := range []RecordClass{"IN", "CS", "CH", "HS"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RecordClass", value)
}

// Ptr returns reference to RecordClass value
func (v RecordClass) Ptr() *RecordClass {
	return &v
}

type NullableRecordClass struct {
	value *RecordClass
	isSet bool
}

func (v NullableRecordClass) Get() *RecordClass {
	return v.value
}

func (v *NullableRecordClass) Set(val *RecordClass) {
	v.value = val
	v.isSet = true
}

func (v NullableRecordClass) IsSet() bool {
	return v.isSet
}

func (v *NullableRecordClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecordClass(val *RecordClass) *NullableRecordClass {
	return &NullableRecordClass{value: val, isSet: true}
}

func (v NullableRecordClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecordClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
