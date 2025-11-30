BINARY=mygrep

SRC=./cmd/grep

TEST=test/test_grep.sh

build:
	go build -o $(BINARY) $(SRC)

clean:
	rm -f $(BINARY)
	rm -f *.out

test: build
	bash $(TEST)

all: build test
