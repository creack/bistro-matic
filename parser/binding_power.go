package parser

type bindingPower int

const (
	bpDefault bindingPower = iota

	bpAdditive
	bpMultiplicative
)
