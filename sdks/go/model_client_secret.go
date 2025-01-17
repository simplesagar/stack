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

// checks if the ClientSecret type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClientSecret{}

// ClientSecret struct for ClientSecret
type ClientSecret struct {
	LastDigits string `json:"lastDigits"`
	Name string `json:"name"`
	Id string `json:"id"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// NewClientSecret instantiates a new ClientSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientSecret(lastDigits string, name string, id string) *ClientSecret {
	this := ClientSecret{}
	this.LastDigits = lastDigits
	this.Name = name
	this.Id = id
	return &this
}

// NewClientSecretWithDefaults instantiates a new ClientSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientSecretWithDefaults() *ClientSecret {
	this := ClientSecret{}
	return &this
}

// GetLastDigits returns the LastDigits field value
func (o *ClientSecret) GetLastDigits() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastDigits
}

// GetLastDigitsOk returns a tuple with the LastDigits field value
// and a boolean to check if the value has been set.
func (o *ClientSecret) GetLastDigitsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastDigits, true
}

// SetLastDigits sets field value
func (o *ClientSecret) SetLastDigits(v string) {
	o.LastDigits = v
}

// GetName returns the Name field value
func (o *ClientSecret) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClientSecret) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClientSecret) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *ClientSecret) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ClientSecret) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ClientSecret) SetId(v string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ClientSecret) GetMetadata() map[string]interface{} {
	if o == nil || IsNil(o.Metadata) {
		var ret map[string]interface{}
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientSecret) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Metadata) {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ClientSecret) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]interface{} and assigns it to the Metadata field.
func (o *ClientSecret) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

func (o ClientSecret) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClientSecret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lastDigits"] = o.LastDigits
	toSerialize["name"] = o.Name
	toSerialize["id"] = o.Id
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

type NullableClientSecret struct {
	value *ClientSecret
	isSet bool
}

func (v NullableClientSecret) Get() *ClientSecret {
	return v.value
}

func (v *NullableClientSecret) Set(val *ClientSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableClientSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableClientSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientSecret(val *ClientSecret) *NullableClientSecret {
	return &NullableClientSecret{value: val, isSet: true}
}

func (v NullableClientSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


