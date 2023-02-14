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

// checks if the LedgerAccountSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LedgerAccountSource{}

// LedgerAccountSource struct for LedgerAccountSource
type LedgerAccountSource struct {
	Id string `json:"id"`
	Ledger string `json:"ledger"`
}

// NewLedgerAccountSource instantiates a new LedgerAccountSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLedgerAccountSource(id string, ledger string) *LedgerAccountSource {
	this := LedgerAccountSource{}
	this.Id = id
	this.Ledger = ledger
	return &this
}

// NewLedgerAccountSourceWithDefaults instantiates a new LedgerAccountSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLedgerAccountSourceWithDefaults() *LedgerAccountSource {
	this := LedgerAccountSource{}
	return &this
}

// GetId returns the Id field value
func (o *LedgerAccountSource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LedgerAccountSource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LedgerAccountSource) SetId(v string) {
	o.Id = v
}

// GetLedger returns the Ledger field value
func (o *LedgerAccountSource) GetLedger() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ledger
}

// GetLedgerOk returns a tuple with the Ledger field value
// and a boolean to check if the value has been set.
func (o *LedgerAccountSource) GetLedgerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ledger, true
}

// SetLedger sets field value
func (o *LedgerAccountSource) SetLedger(v string) {
	o.Ledger = v
}

func (o LedgerAccountSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LedgerAccountSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["ledger"] = o.Ledger
	return toSerialize, nil
}

type NullableLedgerAccountSource struct {
	value *LedgerAccountSource
	isSet bool
}

func (v NullableLedgerAccountSource) Get() *LedgerAccountSource {
	return v.value
}

func (v *NullableLedgerAccountSource) Set(val *LedgerAccountSource) {
	v.value = val
	v.isSet = true
}

func (v NullableLedgerAccountSource) IsSet() bool {
	return v.isSet
}

func (v *NullableLedgerAccountSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLedgerAccountSource(val *LedgerAccountSource) *NullableLedgerAccountSource {
	return &NullableLedgerAccountSource{value: val, isSet: true}
}

func (v NullableLedgerAccountSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLedgerAccountSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

