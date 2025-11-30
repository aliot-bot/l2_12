package output

import (
	"sort"

	"l2_12/model"
)

func BuildOutput(lines []string, matches []model.Match, cfg model.Config) []model.OutputLine {
	if cfg.CountOnly {
		out := make([]model.OutputLine, 0, len(matches))
		for _, m := range matches {
			out = append(out, model.OutputLine{
				LineNum: m.LineNum,
				Text:    m.Text,
				IsMatch: true,
			})
		}
		return out
	}

	type rangeGroup struct {
		start int
		end   int
	}

	var groups []rangeGroup
	addContext := !cfg.Invert

	// Формируем группы строк для вывода
	for _, m := range matches {
		start := m.LineNum
		end := m.LineNum
		if addContext {
			start -= cfg.Before
			end += cfg.After
			if start < 1 {
				start = 1
			}
			if end > len(lines) {
				end = len(lines)
			}
		}
		groups = append(groups, rangeGroup{start, end})
	}

	// Объединяем пересекающиеся группы
	sort.Slice(groups, func(i, j int) bool { return groups[i].start < groups[j].start })
	merged := []rangeGroup{}
	for _, g := range groups {
		if len(merged) == 0 {
			merged = append(merged, g)
		} else {
			last := &merged[len(merged)-1]
			if g.start <= last.end+1 {
				if g.end > last.end {
					last.end = g.end
				}
			} else {
				merged = append(merged, g)
			}
		}
	}

	// Строим OutputLine с правильным "--"
	var out []model.OutputLine
	for i, g := range merged {
		if i != 0 && addContext && g.start > merged[i-1].end+1 {
			out = append(out, model.OutputLine{LineNum: 0, Text: "--", IsMatch: false})
		}
		for idx := g.start; idx <= g.end; idx++ {
			isMatch := false
			for _, m := range matches {
				if m.LineNum == idx {
					isMatch = true
					break
				}
			}
			out = append(out, model.OutputLine{
				LineNum: idx,
				Text:    lines[idx-1],
				IsMatch: isMatch,
			})
		}
	}

	// Для -v: выводим только строки без совпадений
	if cfg.Invert {
		out = []model.OutputLine{}
		for _, m := range matches {
			out = append(out, model.OutputLine{
				LineNum: m.LineNum,
				Text:    m.Text,
				IsMatch: true,
			})
		}
	}

	return out
}
