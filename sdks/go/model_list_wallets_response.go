/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: develop
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
)

// checks if the ListWalletsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListWalletsResponse{}

// ListWalletsResponse struct for ListWalletsResponse
type ListWalletsResponse struct {
	Cursor ListWalletsResponseCursor `json:"cursor"`
}

// NewListWalletsResponse instantiates a new ListWalletsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListWalletsResponse(cursor ListWalletsResponseCursor) *ListWalletsResponse {
	this := ListWalletsResponse{}
	this.Cursor = cursor
	return &this
}

// NewListWalletsResponseWithDefaults instantiates a new ListWalletsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListWalletsResponseWithDefaults() *ListWalletsResponse {
	this := ListWalletsResponse{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *ListWalletsResponse) GetCursor() ListWalletsResponseCursor {
	if o == nil {
		var ret ListWalletsResponseCursor
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *ListWalletsResponse) GetCursorOk() (*ListWalletsResponseCursor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *ListWalletsResponse) SetCursor(v ListWalletsResponseCursor) {
	o.Cursor = v
}

func (o ListWalletsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListWalletsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cursor"] = o.Cursor
	return toSerialize, nil
}

type NullableListWalletsResponse struct {
	value *ListWalletsResponse
	isSet bool
}

func (v NullableListWalletsResponse) Get() *ListWalletsResponse {
	return v.value
}

func (v *NullableListWalletsResponse) Set(val *ListWalletsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListWalletsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListWalletsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListWalletsResponse(val *ListWalletsResponse) *NullableListWalletsResponse {
	return &NullableListWalletsResponse{value: val, isSet: true}
}

func (v NullableListWalletsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListWalletsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


