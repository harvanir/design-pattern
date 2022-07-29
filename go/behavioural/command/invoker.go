package command

// Invoker implementation
type serviceAvailabilityInvoker struct {
	onIbft   Command
	offIbft  Command
	onTopup  Command
	offTopup Command
	onKyc    Command
	offKyc   Command
}

func (sa *serviceAvailabilityInvoker) On() {
	sa.onIbft.Execute()
	sa.onTopup.Execute()
	sa.onKyc.Execute()
}

func (sa *serviceAvailabilityInvoker) Off() {
	sa.offIbft.Execute()
	sa.offTopup.Execute()
	sa.offKyc.Execute()
}

func (sa *serviceAvailabilityInvoker) OnIbft() {
	sa.onIbft.Execute()
}

func (sa *serviceAvailabilityInvoker) OnTopup() {
	sa.onTopup.Execute()
}

func (sa *serviceAvailabilityInvoker) OnKyc() {
	sa.onKyc.Execute()
}

func (sa *serviceAvailabilityInvoker) OffIbft() {
	sa.offIbft.Execute()
}

func (sa *serviceAvailabilityInvoker) OffTopup() {
	sa.offTopup.Execute()
}

func (sa *serviceAvailabilityInvoker) OffKyc() {
	sa.offKyc.Execute()
}
func (sa *serviceAvailabilityInvoker) Undo() {
	// pop queue
	// execute Undo()
}
