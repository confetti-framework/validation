module github.com/lanvard/validation

go 1.15

require (
	github.com/lanvard/contract v0.0.0
	github.com/lanvard/errors v0.10.0
	github.com/lanvard/support v0.1.0
	github.com/stretchr/testify v1.6.1
)

replace (
	github.com/lanvard/contract v0.0.0 => ../contract
	github.com/lanvard/support v0.1.0 => ../support
	github.com/lanvard/errors v0.10.0 => ../errors
)
