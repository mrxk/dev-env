#!/bin/bash

echo '### dev-env'
echo
echo '```text'
./dev-env --help
echo '```'
echo
./dev-env --help | grep -E "^  [[:alpha:]]+ " | cut -f3 -d' ' | while read cmd
do
	echo "### dev-env $cmd"
	echo
	echo '```text'
	./dev-env $cmd --help
	echo '```'
	echo
done
