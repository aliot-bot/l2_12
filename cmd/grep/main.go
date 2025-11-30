package main

import (
	"fmt"
	"os"

	"l2_12/internal/flags"
	"l2_12/internal/input"
	"l2_12/internal/matcher"
	"l2_12/internal/output"
	"l2_12/internal/search"
)

func main() {
	cfg, err := flags.ParseFlags(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	lines, err := input.ReadInput(cfg.Filepath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	var m matcher.Matcher
	if cfg.Fixed {
		m = matcher.NewFixedMatcher(cfg.Pattern, cfg.IgnoreCase)
	} else {
		m, err = matcher.NewRegexMatcher(cfg.Pattern, cfg.IgnoreCase)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:", err)
			os.Exit(1)
		}
	}

	matches := search.FindMatches(lines, m, cfg)
	out := output.BuildOutput(lines, matches, cfg)
	formatted := output.FormatOutput(out, cfg)

	for _, line := range formatted {
		fmt.Println(line)
	}
}
