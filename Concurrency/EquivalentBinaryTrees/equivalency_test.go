package equivalency_test

import (
	"equivalency"
	"errors"
	"testing"
)

func getRange(k int) ([]int, error) {
	if k < 0 {
		return nil, errors.New("Cannot build negative length sequence")
	}
	result := make([]int, k)
	for i := 0; i < k; i++ {
		result[i] = i + 1
	}
	return result, nil
}

func TestWalk_CanCorrectlyReadTreeValues(t *testing.T) {
	t.Parallel()
}
