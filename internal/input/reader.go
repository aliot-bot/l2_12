package input

import (
	"bufio"
	"os"
)

func ReadInput(filepath string) ([]string, error) {
	if filepath == "" {
		return readFromStdin()
	}
	return readFromFile(filepath)
}

func readFromStdin() ([]string, error) {
	var lines []string
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines, sc.Err()
}

func readFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	return lines, sc.Err()
}
