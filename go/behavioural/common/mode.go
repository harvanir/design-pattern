package common

type Mode = int

const (
	On Mode = iota
	Off
	OnIbft
	OffIbft
	OnTopup
	OffTopup
	OnKyc
	OffKyc
)
