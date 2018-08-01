package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/objx"
)

func TestParseJSONString(t *testing.T) {
	stringTest := `{
		"testKey": "testValue"
	 }`
	key := "testKey"
	jsonData, _ := objx.FromJSON(stringTest)
	expectedValue := jsonData.Get("testKey").Data().(string)
	actualValue, err := parseJSON(jsonData, key)
	assert.Equal(t, expectedValue, actualValue, "should both be strings")
	assert.NoError(t, err, "should have no error")
}

func TestParseJSONBool(t *testing.T) {
	boolTest := `{
		"testKey": true
	 }`
	key := "testKey"
	jsonData, _ := objx.FromJSON(boolTest)
	expectedValue := 1
	actualValue, err := parseJSON(jsonData, key)
	assert.Equal(t, expectedValue, actualValue, "should both be 1")
	assert.NoError(t, err, "should have no error")
}
func TestParseJSONFloat64(t *testing.T) {
	floatTest := `{
		"testKey": 42.42
	 }`
	key := "testKey"
	jsonData, _ := objx.FromJSON(floatTest)
	expectedValue := jsonData.Get("testKey").Data().(float64)
	actualValue, err := parseJSON(jsonData, key)
	assert.Equal(t, expectedValue, actualValue, "should both be float64")
	assert.NoError(t, err, "should have no error")
}
func TestParseJSONInt(t *testing.T) {
	intTest := `{
		"testKey": 42
	 }`
	key := "testKey"
	jsonData, _ := objx.FromJSON(intTest)
	expectedValue := jsonData.Get("testKey").Data().(int)
	actualValue, err := parseJSON(jsonData, key)
	assert.Equal(t, expectedValue, actualValue, "should both be integers")
	assert.NoError(t, err, "should have no error")
}

func TestParseJSONNil(t *testing.T) {
	nilTest := `{
		"testKey": "testValue"
	 }`
	key := "badKey"
	jsonData, _ := objx.FromJSON(nilTest)
	actualValue, err := parseJSON(jsonData, key)
	assert.Nil(t, actualValue, "value should be nil")
	assert.Error(t, err, "should have an error")
}

func TestConvertBoolToIntTrue(t *testing.T) {
	actual := convertBoolToInt(true)
	expected := 1
	assert.Equal(t, expected, actual, "should both equal 1")
}

func TestConvertBoolToIntFalse(t *testing.T) {
	actual := convertBoolToInt(false)
	expected := 0
	assert.Equal(t, expected, actual, "should both equal 0")
}
