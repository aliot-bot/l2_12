package matcher

import (
	"regexp"
)

type RegexMatcher struct {
	r *regexp.Regexp
}

func NewRegexMatcher(pattern string, ignoreCase bool) (Matcher, error) {
	if ignoreCase {
		pattern = "(?i)" + pattern
	}
	r, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &RegexMatcher{r: r}, nil
}

func (rm *RegexMatcher) Match(line string) bool {
	return rm.r.MatchString(line)
}
