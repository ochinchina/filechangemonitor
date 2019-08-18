package filechangemonitor

import (
	"testing"
)

func TestMatchAllFile(t *testing.T) {
	matcher := MatchAllFile{}
	if !matcher.Match("anyFile.txt") {
		t.Fatal()
	}

	if !matcher.Match("test.dat") {
		t.Fatal()
	}
}

func TestPatternFileMatcher(t *testing.T) {
	matcher := NewPatternFileMatcher(`*.cfg`)

	if !matcher.Match("/test/test.cfg") {
		t.Fatal()
	}

	if matcher.Match("/test/test.txt") {
		t.Fatal()
	}
}

func TestRegexFileMatcher(t *testing.T) {
	matcher, err := NewRegexFileMatcher(`^[a-z]+\[[0-9]+\]$`)
	if err != nil {
		t.Fatal()
	}

	if !matcher.Match("adam[23]") {
		t.Fatal()
	}

	if matcher.Match("snakey") {
		t.Fatal()
	}
}

func TestAnyFileMatcher(t *testing.T) {
	matcher1, err := NewRegexFileMatcher(`^[a-z]+\[[0-9]+\]$`)
	if err != nil {
		t.Fatal()
	}
	matcher2, err := NewRegexFileMatcher(`^[A-Z]+$`)
	if err != nil {
		t.Fatal()
	}

	matcher := NewAnyFileMatcher(matcher1, matcher2)
	if !matcher.Match("ABC") {
		t.Fatal()
	}
	if !matcher.Match("adam[23]") {
		t.Fatal()
	}
}

func TestAllFileMatcher(t *testing.T) {
	matcher1, err := NewRegexFileMatcher(`^[a-z]+\[[0-9]+\]$`)
	if err != nil {
		t.Fatal()
	}
	matcher2, err := NewRegexFileMatcher(`^[e-z].+$`)
	if err != nil {
		t.Fatal()
	}

	matcher := NewAllFileMatcher(matcher1, matcher2)

	if matcher.Match("adam[23]") {
		t.Fatal()
	}

	if !matcher.Match("fifo[123]") {
		t.Fatal()
	}

}
