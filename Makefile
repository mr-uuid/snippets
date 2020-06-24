to-string:
	find . -name "stringer.go" -exec go run {} \;

parsing-strings:
	find . -name "parsing-strings.go" -exec go run {} \;

adding-methods:
	find . -name "adding-methods.go" -exec go run {} \;
	find . -name "adding-methods-bad-1-alternative.go" -exec go run {} \;
	find . -name "adding-methods-bad-1.go" -exec go run {} \; || true
	find . -name "adding-methods-bad-2.go" -exec go run {} \; || true

making-ids:
	go run code-snippets/making-ids.go
