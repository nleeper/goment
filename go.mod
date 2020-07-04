module github.com/nleeper/goment

go 1.14

require (
	github.com/nleeper/goment/internal/constants v1.0.0
	github.com/nleeper/goment/internal/regexps v1.0.0
	github.com/stretchr/testify v1.6.1
)

replace github.com/nleeper/goment/internal/regexps v1.0.0 => ./internal/regexps

replace github.com/nleeper/goment/internal/constants v1.0.0 => ./internal/constants
