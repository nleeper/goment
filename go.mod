module github.com/nleeper/goment

go 1.14

require (
	github.com/nleeper/goment/internal/constants v0.0.0-00010101000000-000000000000
	github.com/nleeper/goment/internal/regexps v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.6.1
)

replace github.com/nleeper/goment/internal/regexps => ./internal/regexps

replace github.com/nleeper/goment/internal/constants => ./internal/constants
