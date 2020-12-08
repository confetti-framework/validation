module github.com/lanvard/validation

go 1.15

require (
	github.com/lanvard/contract v0.0.0
	github.com/lanvard/errors v0.10.0
	github.com/lanvard/support v0.1.0
	github.com/lanvard/syslog v0.0.0-20201116213126-66df7a6162c3
	github.com/stretchr/testify v1.6.1
	github.com/uniplaces/carbon v0.1.6
	github.com/vigneshuvi/GoDateFormat v0.0.0-20190923034126-379ee8a8c45f
)

replace (
	github.com/lanvard/contract v0.0.0 => ../contract
	github.com/lanvard/errors v0.10.0 => ../errors
	github.com/lanvard/support v0.1.0 => ../support
	github.com/lanvard/syslog v0.0.0-20201116213126-66df7a6162c3 => ../syslog
)
