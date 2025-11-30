package flags

import (
	"errors"
	"strconv"

	"l2_12/model"
)

func ParseFlags(args []string) (model.Config, error) {
	var cfg model.Config

	i := 0
	for i < len(args) {
		arg := args[i]

		switch arg {
		case "-A":
			i++
			if i >= len(args) {
				return cfg, errors.New("-A requires number")
			}
			n, err := strconv.Atoi(args[i])
			if err != nil {
				return cfg, err
			}
			cfg.After = n

		case "-B":
			i++
			if i >= len(args) {
				return cfg, errors.New("-B requires number")
			}
			n, err := strconv.Atoi(args[i])
			if err != nil {
				return cfg, err
			}
			cfg.Before = n

		case "-C":
			i++
			if i >= len(args) {
				return cfg, errors.New("-C requires number")
			}
			n, err := strconv.Atoi(args[i])
			if err != nil {
				return cfg, err
			}
			cfg.Before = n
			cfg.After = n

		case "-c":
			cfg.CountOnly = true
		case "-i":
			cfg.IgnoreCase = true
		case "-v":
			cfg.Invert = true
		case "-F":
			cfg.Fixed = true
		case "-n":
			cfg.ShowLineNum = true

		default:
			if cfg.Pattern == "" {
				cfg.Pattern = arg
			} else if cfg.Filepath == "" {
				cfg.Filepath = arg
			}
		}

		i++
	}

	if cfg.Pattern == "" {
		return cfg, errors.New("pattern is required")
	}

	return cfg, nil
}
