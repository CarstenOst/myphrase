package myphrase

import "testing"

type PhraseResult struct {
	i        int
	expected string
}

var phraseResult = []PhraseResult{
	{0, "I can eat glass and it doesn't hurt me."},
	{1, "Don't communicate by sharing memory, share memory by communicating."},
	{2, "99 bottles of beer on the wall, 99 bottles of beer, ..."},
	{3, "If a program is too slow, it must have a loop."},
}

func TestPithySay(t *testing.T) {
	for i, test := range phraseResult {
		if result := PithySay(i); result != test.expected {
			t.Errorf("PithySay(i) = %q, want %q", result, phraseResult)
		}
	}
}
