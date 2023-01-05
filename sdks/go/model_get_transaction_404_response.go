/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: SDK_VERSION
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// GetTransaction404Response struct for GetTransaction404Response
type GetTransaction404Response struct {
	ErrorCode string `json:"error_code"`
	ErrorMessage *string `json:"error_message,omitempty"`
}

// NewGetTransaction404Response instantiates a new GetTransaction404Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTransaction404Response(errorCode string) *GetTransaction404Response {
	this := GetTransaction404Response{}
	this.ErrorCode = errorCode
	return &this
}

// NewGetTransaction404ResponseWithDefaults instantiates a new GetTransaction404Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTransaction404ResponseWithDefaults() *GetTransaction404Response {
	this := GetTransaction404Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *GetTransaction404Response) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *GetTransaction404Response) GetErrorCodeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *GetTransaction404Response) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *GetTransaction404Response) GetErrorMessage() string {
	if o == nil || isNil(o.ErrorMessage) {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTransaction404Response) GetErrorMessageOk() (*string, bool) {
	if o == nil || isNil(o.ErrorMessage) {
    return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *GetTransaction404Response) HasErrorMessage() bool {
	if o != nil && !isNil(o.ErrorMessage) {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *GetTransaction404Response) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o GetTransaction404Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error_code"] = o.ErrorCode
	}
	if !isNil(o.ErrorMessage) {
		toSerialize["error_message"] = o.ErrorMessage
	}
	return json.Marshal(toSerialize)
}

type NullableGetTransaction404Response struct {
	value *GetTransaction404Response
	isSet bool
}

func (v NullableGetTransaction404Response) Get() *GetTransaction404Response {
	return v.value
}

func (v *NullableGetTransaction404Response) Set(val *GetTransaction404Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTransaction404Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTransaction404Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTransaction404Response(val *GetTransaction404Response) *NullableGetTransaction404Response {
	return &NullableGetTransaction404Response{value: val, isSet: true}
}

func (v NullableGetTransaction404Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTransaction404Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


