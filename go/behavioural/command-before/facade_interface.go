package command_before

// ServiceToggleFacade - Facade
type ServiceToggleFacade interface {
	On()
	Off()
	OnIbft()
	OnTopup()
	OnKyc()
	OffIbft()
	OffTopup()
	OffKyc()
}
