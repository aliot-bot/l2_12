package output

import (
	"fmt"
	"strconv"

	"l2_12/model"
)

func FormatOutput(out []model.OutputLine, cfg model.Config) []string {
	if cfg.CountOnly {
		return []string{strconv.Itoa(countMatches(out))}
	}

	var res []string
	for _, line := range out {
		if line.Text == "--" {
			res = append(res, "--")
			continue
		}
		if cfg.ShowLineNum {
			res = append(res, fmt.Sprintf("%d:%s", line.LineNum, line.Text))
		} else {
			res = append(res, line.Text)
		}
	}
	return res
}

func countMatches(out []model.OutputLine) int {
	n := 0
	for _, line := range out {
		if line.IsMatch {
			n++
		}
	}
	return n
}
