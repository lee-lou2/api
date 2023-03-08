package utils

import (
	"testing"
)

func TestSum(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	expected := 15.0
	result := Sum(data)
	if result != expected {
		t.Errorf("Sum(%v) = %v, expected %v", data, result, expected)
	}
}

func TestPredictNext(t *testing.T) {
	data := []float64{1, 2, 3, 4, 5}
	expected := 6.0
	result := PredictNext(data)
	if result != expected {
		t.Errorf("PredictNext(%v) = %v, expected %v", data, result, expected)
	}
}

func TestPredictNextSingleValue(t *testing.T) {
	data := []float64{1}
	expected := 1.0
	result := PredictNext(data)
	if result != expected {
		t.Errorf("PredictNext(%v) = %v, expected %v", data, result, expected)
	}
}

func TestPredictNextNegativeValue(t *testing.T) {
	data := []float64{1, 2, 3, -4, 5}
	expected := 6.0
	result := PredictNext(data)
	if result != expected {
		t.Errorf("PredictNext(%v) = %v, expected %v", data, result, expected)
	}
}
func TestPredictNextNonIntegerValue(t *testing.T) {
	data := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	expected := 6.6
	result := PredictNext(data)
	if result != expected {
		t.Errorf("PredictNext(%v) = %v, expected %v", data, result, expected)
	}
}
