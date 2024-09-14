package utility_test

import (
	"testing"

	"github.com/fitm-elite/elebs/packages/utility"
)

func TestCostDivided(t *testing.T) {
	t.Parallel()

	t.Run("TestCostDivided", func(t *testing.T) {
		cost := 100.0
		n := 10
		expected := 10.0

		result := utility.CostDivider(cost, n)

		if result != expected {
			t.Errorf("Expected %f but got %f", expected, result)
		}
	})
}
