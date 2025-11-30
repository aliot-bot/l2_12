package matcher

import "strings"

type FixedMatcher struct {
	pattern    string
	ignoreCase bool
}

func NewFixedMatcher(pattern string, ignoreCase bool) Matcher {
	if ignoreCase {
		pattern = strings.ToLower(pattern)
	}
	return &FixedMatcher{
		pattern:    pattern,
		ignoreCase: ignoreCase,
	}
}

func (fm *FixedMatcher) Match(line string) bool {
	if fm.ignoreCase {
		line = strings.ToLower(line)
	}
	return strings.Contains(line, fm.pattern)
}
