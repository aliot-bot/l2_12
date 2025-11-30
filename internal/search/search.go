package search

import (
	"l2_12/internal/matcher"
	"l2_12/model"
)

func FindMatches(lines []string, m matcher.Matcher, cfg model.Config) []model.Match {
	var out []model.Match

	for i, line := range lines {
		match := m.Match(line)
		if cfg.Invert {
			match = !match
		}
		if match {
			out = append(out, model.Match{
				LineNum: i + 1,
				Text:    line,
			})
		}
	}
	return out
}
