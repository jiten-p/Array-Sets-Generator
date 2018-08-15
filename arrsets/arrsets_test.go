package arrsets

import (
	"reflect"
	"testing"
)

func TestArraySetsGeneratorSuccessForInt(t *testing.T) {
	arr := []interface{}{1, 2, 3}
	arr_expected := [][]interface{}{
		{2, 1, 3},
		{2, 3, 1},
		{3, 2, 1},
		{3, 1, 2},
		{1, 3, 2},
		{1, 2, 3},
	}

	//This is success case
	arr_actual := ArraySetsGenerator(arr)
	if !reflect.DeepEqual(arr_actual, arr_expected) {
		t.Fatal("expected and actual results are not same")
	}
}

func TestArraySetsGeneratorFailForInt(t *testing.T) {
	// This is a failure case
	arr := []interface{}{1, 2, 3}
	arr_expected_fail := [][]interface{}{
		{2, 1, 3},
		{2, 3, 1},
		{3, 2, 1},
		{3, 1, 2},
		{1, 3, 2},
		{1, 2, 2},
	}
	arr_actual := ArraySetsGenerator(arr)
	if reflect.DeepEqual(arr_actual, arr_expected_fail) {
		t.Fatal("expected and actual results should not be same")
	}
}

func TestArraySetsGeneratorNilSuccess(t *testing.T) {
	// This is nil case
	var arr_nil []interface{}

	arr_actual_nul := ArraySetsGenerator(arr_nil)

	if arr_actual_nul != nil {
		t.Fatal("expected result is nil")
	}

}

func TestArraySetsGeneratorSigleSuccess(t *testing.T) {
	arr_single := []interface{}{1}
	arr_expected_single := [][]interface{}{{1}}

	//This is success case
	arr_actual_single := ArraySetsGenerator(arr_single)
	if !reflect.DeepEqual(arr_expected_single, arr_actual_single) {
		t.Fatal("expected and actual results are not same")
	}

}

func TestArraySetsGeneratorSuccessForString(t *testing.T) {
	arr := []interface{}{"A", "B", "C"}
	arr_expected := [][]interface{}{
		{"B", "A", "C"},
		{"B", "C", "A"},
		{"C", "B", "A"},
		{"C", "A", "B"},
		{"A", "C", "B"},
		{"A", "B", "C"},
	}
	//This is success case
	arr_actual := ArraySetsGenerator(arr)
	if !reflect.DeepEqual(arr_actual, arr_expected) {
		t.Fatal("expected and actual results are not same")
	}

}
func TestArraySetsGeneratorFailForString(t *testing.T) {
	arr := []interface{}{"A", "B", "C"}
	arr_expected := [][]interface{}{
		{"B", "A", "C"},
		{"B", "C", "A"},
		{"C", "B", "A"},
		{"C", "A", "B"},
		{"A", "C", "C"},
	}
	//This is success case
	arr_actual := ArraySetsGenerator(arr)
	if reflect.DeepEqual(arr_actual, arr_expected) {
		t.Fatal("expected and actual results should not be same")
	}
}
