#!/bin/bash

set -e

GO_BUILD="./mygrep"
go build -o $GO_BUILD ./cmd/grep

TEST_FILE="testfile.txt"
cat > $TEST_FILE <<EOL
apple
banana
Apple
orange
banana apple
APPLE pie
grape
EOL

declare -a TESTS=(
"1|plain search|\"apple\"|"
"2|ignore case|-i \"apple\"|"
"3|invert|-v \"banana\"|"
"4|count|-c \"banana\"|"
"5|after context|-A 1 \"banana\"|"
"6|before context|-B 1 \"banana\"|"
"7|context|-C 1 \"banana\"|"
"8|fixed string|-F \"banana\"|"
"9|line numbers|-n \"apple\"|"
"10|combo|-i -v -n -C 1 \"apple\"|"
)

echo "=== RUNNING TESTS ==="
for t in "${TESTS[@]}"; do
	IFS="|" read -r id name flags <<< "$t"
	echo "Test $id: $name"
	
	$GO_BUILD $flags $TEST_FILE > mygrep.out 2>/dev/null
	grep $flags $TEST_FILE > grep.out 2>/dev/null || true

	if diff -u grep.out mygrep.out >/dev/null; then
		echo "  PASS ✅"
	else
		echo "  FAIL ❌"
		echo "---- Expected (grep) ----"
		cat grep.out
		echo "---- Got (mygrep) ----"
		cat mygrep.out
	fi
done

rm -f $GO_BUILD $TEST_FILE grep.out mygrep.out
echo "=== ALL TESTS FINISHED ==="
