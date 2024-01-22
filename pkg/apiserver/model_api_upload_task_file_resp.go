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

// checks if the ApiUploadTaskFileResp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiUploadTaskFileResp{}

// ApiUploadTaskFileResp struct for ApiUploadTaskFileResp
type ApiUploadTaskFileResp struct {
	Code *int32 `json:"code,omitempty"`
	Data *ResponseUploadTaskResp `json:"data,omitempty"`
	Msg *string `json:"msg,omitempty"`
	TraceId *string `json:"trace_id,omitempty"`
}

// NewApiUploadTaskFileResp instantiates a new ApiUploadTaskFileResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiUploadTaskFileResp() *ApiUploadTaskFileResp {
	this := ApiUploadTaskFileResp{}
	return &this
}

// NewApiUploadTaskFileRespWithDefaults instantiates a new ApiUploadTaskFileResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiUploadTaskFileRespWithDefaults() *ApiUploadTaskFileResp {
	this := ApiUploadTaskFileResp{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiUploadTaskFileResp) GetCode() int32 {
	if o == nil || IsNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUploadTaskFileResp) GetCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiUploadTaskFileResp) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *ApiUploadTaskFileResp) SetCode(v int32) {
	o.Code = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ApiUploadTaskFileResp) GetData() ResponseUploadTaskResp {
	if o == nil || IsNil(o.Data) {
		var ret ResponseUploadTaskResp
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUploadTaskFileResp) GetDataOk() (*ResponseUploadTaskResp, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ApiUploadTaskFileResp) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ResponseUploadTaskResp and assigns it to the Data field.
func (o *ApiUploadTaskFileResp) SetData(v ResponseUploadTaskResp) {
	o.Data = &v
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *ApiUploadTaskFileResp) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUploadTaskFileResp) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *ApiUploadTaskFileResp) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *ApiUploadTaskFileResp) SetMsg(v string) {
	o.Msg = &v
}

// GetTraceId returns the TraceId field value if set, zero value otherwise.
func (o *ApiUploadTaskFileResp) GetTraceId() string {
	if o == nil || IsNil(o.TraceId) {
		var ret string
		return ret
	}
	return *o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUploadTaskFileResp) GetTraceIdOk() (*string, bool) {
	if o == nil || IsNil(o.TraceId) {
		return nil, false
	}
	return o.TraceId, true
}

// HasTraceId returns a boolean if a field has been set.
func (o *ApiUploadTaskFileResp) HasTraceId() bool {
	if o != nil && !IsNil(o.TraceId) {
		return true
	}

	return false
}

// SetTraceId gets a reference to the given string and assigns it to the TraceId field.
func (o *ApiUploadTaskFileResp) SetTraceId(v string) {
	o.TraceId = &v
}

func (o ApiUploadTaskFileResp) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiUploadTaskFileResp) ToMap() (map[string]interface{}, error) {
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

type NullableApiUploadTaskFileResp struct {
	value *ApiUploadTaskFileResp
	isSet bool
}

func (v NullableApiUploadTaskFileResp) Get() *ApiUploadTaskFileResp {
	return v.value
}

func (v *NullableApiUploadTaskFileResp) Set(val *ApiUploadTaskFileResp) {
	v.value = val
	v.isSet = true
}

func (v NullableApiUploadTaskFileResp) IsSet() bool {
	return v.isSet
}

func (v *NullableApiUploadTaskFileResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiUploadTaskFileResp(val *ApiUploadTaskFileResp) *NullableApiUploadTaskFileResp {
	return &NullableApiUploadTaskFileResp{value: val, isSet: true}
}

func (v NullableApiUploadTaskFileResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiUploadTaskFileResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


