package command_before

type (
	// ServiceToggleFacade - Facade
	ServiceToggleFacade interface {
		On()
		Off()
		OnIbft()
		OnTopup()
		OnKyc()
		OffIbft()
		OffTopup()
		OffKyc()
	}

	// HttpClient - external
	HttpClient interface {
		Call(url string, requestBody interface{})
	}
)
