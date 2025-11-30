package matcher

type Matcher interface {
	Match(line string) bool
}
