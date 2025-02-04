/*
Device Type-7 APIs

APIs to get information about the HPE Alletra Storage MP X10000 system

API version: 0.1.0
Contact: object-svc-team@hpe.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BucketTagsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BucketTagsInner{}

// BucketTagsInner struct for BucketTagsInner
type BucketTagsInner struct {
	// Key of the Tag
	Key string `json:"key"`
	// Value for that Key
	Value *string `json:"value,omitempty"`
}

type _BucketTagsInner BucketTagsInner

// NewBucketTagsInner instantiates a new BucketTagsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucketTagsInner(key string) *BucketTagsInner {
	this := BucketTagsInner{}
	this.Key = key
	var value string = ""
	this.Value = &value
	return &this
}

// NewBucketTagsInnerWithDefaults instantiates a new BucketTagsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketTagsInnerWithDefaults() *BucketTagsInner {
	this := BucketTagsInner{}
	var value string = ""
	this.Value = &value
	return &this
}

// GetKey returns the Key field value
func (o *BucketTagsInner) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *BucketTagsInner) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *BucketTagsInner) SetKey(v string) {
	o.Key = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *BucketTagsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketTagsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *BucketTagsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *BucketTagsInner) SetValue(v string) {
	o.Value = &v
}

func (o BucketTagsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BucketTagsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

func (o *BucketTagsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"key",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBucketTagsInner := _BucketTagsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBucketTagsInner)

	if err != nil {
		return err
	}

	*o = BucketTagsInner(varBucketTagsInner)

	return err
}

type NullableBucketTagsInner struct {
	value *BucketTagsInner
	isSet bool
}

func (v NullableBucketTagsInner) Get() *BucketTagsInner {
	return v.value
}

func (v *NullableBucketTagsInner) Set(val *BucketTagsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketTagsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketTagsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketTagsInner(val *BucketTagsInner) *NullableBucketTagsInner {
	return &NullableBucketTagsInner{value: val, isSet: true}
}

func (v NullableBucketTagsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketTagsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


