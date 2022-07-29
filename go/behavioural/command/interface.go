package command

type (
	// Invoker - Invoker
	Invoker interface {
		On()
		OnIbft()
		OnTopup()
		OnKyc()
		Off()
		OffIbft()
		OffTopup()
		OffKyc()
		Undo()
	}

	// Command - Command
	Command interface {
		Execute()
		Undo()
	}

	// Receiver - external code
	Receiver interface {
		call(url string, requestBody interface{})
	}
)
