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

// checks if the SearchApp200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchApp200Response{}

// SearchApp200Response struct for SearchApp200Response
type SearchApp200Response struct {
	Code *int32 `json:"code,omitempty"`
	Data *RequestRegisterStruct `json:"data,omitempty"`
	Msg *string `json:"msg,omitempty"`
	TraceId *string `json:"trace_id,omitempty"`
}

// NewSearchApp200Response instantiates a new SearchApp200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchApp200Response() *SearchApp200Response {
	this := SearchApp200Response{}
	return &this
}

// NewSearchApp200ResponseWithDefaults instantiates a new SearchApp200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchApp200ResponseWithDefaults() *SearchApp200Response {
	this := SearchApp200Response{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *SearchApp200Response) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchApp200Response) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *SearchApp200Response) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *SearchApp200Response) SetCode(v int32) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SearchApp200Response) GetData() RequestRegisterStruct {
	if o == nil || IsNil(o.Data) {
		var ret RequestRegisterStruct
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchApp200Response) GetDataOk() (*RequestRegisterStruct, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SearchApp200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given RequestRegisterStruct and assigns it to the Data field.
func (o *SearchApp200Response) SetData(v RequestRegisterStruct) {
	o.Data = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *SearchApp200Response) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchApp200Response) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *SearchApp200Response) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *SearchApp200Response) SetMsg(v string) {
	o.Msg = &v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *SearchApp200Response) GetTraceId() string {
	if o == nil || IsNil(o.TraceId) {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchApp200Response) GetTraceIdOk() (*string, bool) {
	if o == nil || IsNil(o.TraceId) {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *SearchApp200Response) HasTraceId() bool {
	if o != nil && !IsNil(o.TraceId) {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *SearchApp200Response) SetTraceId(v string) {
	o.TraceId = &v
}

func (o SearchApp200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchApp200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !IsNil(o.TraceId) {
		toSerialize["trace_id"] = o.TraceId
	}
	return toSerialize, nil
}

type NullableSearchApp200Response struct {
	value *SearchApp200Response
	isSet bool
}

func (v NullableSearchApp200Response) Get() *SearchApp200Response {
	return v.value
}

func (v *NullableSearchApp200Response) Set(val *SearchApp200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchApp200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchApp200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchApp200Response(val *SearchApp200Response) *NullableSearchApp200Response {
	return &NullableSearchApp200Response{value: val, isSet: true}
}

func (v NullableSearchApp200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchApp200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


