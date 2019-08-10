package filechangemonitor

import (
	"regexp"
)

type FileMatcher interface {
	// check if the file matches the expected
	Match(path string) bool
}

type MatchAllFile struct {
}

func (maf MatchAllFile) Match(path string) bool {
	return true
}

type RegexFileMatcher struct {
	pattern *regexp.Regexp
}

func NewRegexFileMatcher(expr string) (*RegexFileMatcher, error) {
	r, err := regexp.Compile(expr)

	if err != nil {
		return nil, err
	}
	return &RegexFileMatcher{pattern: r}, nil
}

func (rfm *RegexFileMatcher) Match(path string) bool {
	return rfm.pattern.MatchString(path)
}

// match any of matcher
type AnyFileMatcher struct {
	fileMatchers []FileMatcher
}

func NewAnyFileMatcher(matchers ...FileMatcher) *AnyFileMatcher {
	r := &AnyFileMatcher{fileMatchers: make([]FileMatcher, 0)}
	for _, matcher := range matchers {
		r.fileMatchers = append(r.fileMatchers, matcher)
	}
	return r
}

func (afm *AnyFileMatcher) Match(path string) bool {
	for _, matcher := range afm.fileMatchers {
		if matcher.Match(path) {
			return true
		}
	}
	return false
}

type AllFileMatcher struct {
	fileMatchers []FileMatcher
}

func NewAllFileMatcher(matchers ...FileMatcher) *AllFileMatcher {
	r := &AllFileMatcher{fileMatchers: make([]FileMatcher, 0)}
	for _, matcher := range matchers {
		r.fileMatchers = append(r.fileMatchers, matcher)
	}
	return r
}

func (afm *AllFileMatcher) Match(path string) bool {
	for _, matcher := range afm.fileMatchers {
		if !matcher.Match(path) {
			return false
		}
	}
	return true
}
