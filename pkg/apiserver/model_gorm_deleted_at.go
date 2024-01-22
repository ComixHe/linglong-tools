/*
linglong仓库

玲珑仓库接口

API version: 1.0.0
Contact: wurongjie@deepin.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiserver

import (
	"encoding/json"
)

// checks if the GormDeletedAt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GormDeletedAt{}

// GormDeletedAt struct for GormDeletedAt
type GormDeletedAt struct {
	Time *string `json:"time,omitempty"`
	// Valid is true if Time is not NULL
	Valid *bool `json:"valid,omitempty"`
}

// NewGormDeletedAt instantiates a new GormDeletedAt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGormDeletedAt() *GormDeletedAt {
	this := GormDeletedAt{}
	return &this
}

// NewGormDeletedAtWithDefaults instantiates a new GormDeletedAt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGormDeletedAtWithDefaults() *GormDeletedAt {
	this := GormDeletedAt{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *GormDeletedAt) GetTime() string {
	if o == nil || IsNil(o.Time) {
		var ret string
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GormDeletedAt) GetTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *GormDeletedAt) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given string and assigns it to the Time field.
func (o *GormDeletedAt) SetTime(v string) {
	o.Time = &v
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *GormDeletedAt) GetValid() bool {
	if o == nil || IsNil(o.Valid) {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GormDeletedAt) GetValidOk() (*bool, bool) {
	if o == nil || IsNil(o.Valid) {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *GormDeletedAt) HasValid() bool {
	if o != nil && !IsNil(o.Valid) {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *GormDeletedAt) SetValid(v bool) {
	o.Valid = &v
}

func (o GormDeletedAt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GormDeletedAt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.Valid) {
		toSerialize["valid"] = o.Valid
	}
	return toSerialize, nil
}

type NullableGormDeletedAt struct {
	value *GormDeletedAt
	isSet bool
}

func (v NullableGormDeletedAt) Get() *GormDeletedAt {
	return v.value
}

func (v *NullableGormDeletedAt) Set(val *GormDeletedAt) {
	v.value = val
	v.isSet = true
}

func (v NullableGormDeletedAt) IsSet() bool {
	return v.isSet
}

func (v *NullableGormDeletedAt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGormDeletedAt(val *GormDeletedAt) *NullableGormDeletedAt {
	return &NullableGormDeletedAt{value: val, isSet: true}
}

func (v NullableGormDeletedAt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGormDeletedAt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


