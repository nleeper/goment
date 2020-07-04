module github.com/nleeper/goment

go 1.14

require (
	github.com/nleeper/goment/internal/constants v0.0.0-20200704205958-1d76f1925454
	github.com/nleeper/goment/internal/regexps v0.0.0-20200704205958-1d76f1925454
	github.com/stretchr/testify v1.6.1
	internal/regexps v0.0.0
)

replace internal/constants => ./internal/constants

replace internal/regexps => ./internal/regexps
