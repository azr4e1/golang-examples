package firstset_test

import (
	firstset "cryptochallenges/firstSet"
	"testing"
)

func TestHummingDistance(t *testing.T) {
	t.Parallel()
	input1 := []byte("this is a test")
	input2 := []byte("wokka wokka!!!")

	got := firstset.EditDistance(input1, input2)
	want := 37

	if got != want {
		t.Errorf("want %d, got %d", want, got)
	}
}
