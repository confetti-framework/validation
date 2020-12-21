module github.com/confetti-framework/validation

go 1.15

require (
	github.com/confetti-framework/contract v0.0.0
	github.com/confetti-framework/errors v0.10.0
	github.com/confetti-framework/support v0.1.0
	github.com/confetti-framework/syslog v0.0.0-20201116213126-66df7a6162c3
	github.com/spf13/cast v1.3.1
	github.com/stretchr/testify v1.6.1
	github.com/uniplaces/carbon v0.1.6
	github.com/vigneshuvi/GoDateFormat v0.0.0-20190923034126-379ee8a8c45f
)

replace (
	github.com/confetti-framework/contract v0.0.0 => ../contract
	github.com/confetti-framework/errors v0.10.0 => ../errors
	github.com/confetti-framework/support v0.1.0 => ../support
	github.com/confetti-framework/syslog v0.0.0-20201116213126-66df7a6162c3 => ../syslog
)
