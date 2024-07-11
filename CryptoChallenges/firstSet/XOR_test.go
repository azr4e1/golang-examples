package firstset_test

import (
	"cryptochallenges/firstSet"
	"testing"
)

func TestEditDistance(t *testing.T) {
	t.Parallel()
	input1 := "this is a test"
	input2 := "wokka wokka!!!"
	want := 37
	got := firstset.EditDistance([]byte(input1), []byte(input2))
	if want != got {
		t.Errorf("got %d, want %d", want, got)
	}
}
