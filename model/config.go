package model

type Config struct {
	Before int
	After  int

	CountOnly   bool // -c
	IgnoreCase  bool // -i
	Invert      bool // -v
	Fixed       bool // -F
	ShowLineNum bool // -n

	Pattern  string
	Filepath string
}
