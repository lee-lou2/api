package convert

import (
	"reflect"
	"testing"
)

func TestStringToMapNormal(t *testing.T) {
	params := `{"name": "John", "age": 30}`
	expected := map[string]interface{}{
		"name": "John",
		"age":  float64(30),
	}
	result, err := StringToMap(params)
	if err != nil {
		t.Errorf("StringToMap(%v) error = %v, expected nil", params, err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("StringToMap(%v) = %v, expected %v", params, result, expected)
	}
}

func TestStringToMapEmpty(t *testing.T) {
	params := ``
	expected := map[string]interface{}{}
	result, err := StringToMap(params)
	if err != nil {
		t.Errorf("StringToMap(%v) error = %v, expected nil", params, err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("StringToMap(%v) = %v, expected %v", params, result, expected)
	}
}

func TestStringToMapInvalidJSON(t *testing.T) {
	params := `{"name": "John", "age": }`
	expected := map[string]interface{}{}
	result, err := StringToMap(params)
	if err == nil {
		t.Errorf("StringToMap(%v) error = nil, expected error", params)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("StringToMap(%v) = %v, expected %v", params, result, expected)
	}
}

func TestStringToMapNested(t *testing.T) {
	params := `{"name": {"first": "John", "last": "Doe"}, "age": 30}`
	expected := map[string]interface{}{
		"name": map[string]interface{}{
			"first": "John",
			"last":  "Doe",
		},
		"age": float64(30),
	}
	result, err := StringToMap(params)
	if err != nil {
		t.Errorf("StringToMap(%v) error = %v, expected nil", params, err)
	}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("StringToMap(%v) = %v, expected %v", params, result, expected)
	}
}
