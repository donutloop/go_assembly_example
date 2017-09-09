package math_test

import (
	"testing"
	"github.com/donutloop/go_assembly_example/math"
)

func TestAdd(t *testing.T) {
	sum := math.Add(9,9)
	expectedValue := int64(18)
	if sum != expectedValue {
		t.Errorf("Sum of addtion is bad (Actual: %s, Expected: %s)", sum, expectedValue)
	}
}