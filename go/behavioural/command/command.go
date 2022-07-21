package command

import "fmt"

type (
	// Command - Command
	Command interface {
		Execute()
		Undo()
	}
	command struct {
		R Receiver
	}
)

func (c command) constructRequest(mode string, service string) string {
	return fmt.Sprintf("{\"mode\": \"%v\", \"service\": \"%v\"}", mode, service)
}

// ibft - start
type ibftOnCommand command

func (c ibftOnCommand) Execute() {
	c.R.call("http://firebase.com", command(c).constructRequest("on", "ibft"))
}

func (c ibftOnCommand) Undo() {
	c.R.call("http://firebase.com", command(c).constructRequest("off", "ibft"))
}

type ibftOffCommand command

func (c ibftOffCommand) Execute() {
	c.R.call("http://firebase.com", command(c).constructRequest("off", "ibft"))
}

func (c ibftOffCommand) Undo() {
	c.R.call("http://firebase.com", command(c).constructRequest("on", "ibft"))
} // ibft - end

// topup - start
type topupOnCommand command

func (c topupOnCommand) Execute() {
	c.R.call("http://firebase.com", command(c).constructRequest("on", "topup"))
}

func (c topupOnCommand) Undo() {
	c.R.call("http://firebase.com", command(c).constructRequest("off", "topup"))
}

type topupOffCommand command

func (c topupOffCommand) Execute() {
	c.R.call("http://firebase.com", command(c).constructRequest("off", "topup"))
}

func (c topupOffCommand) Undo() {
	c.R.call("http://firebase.com", command(c).constructRequest("on", "topup"))
} // topup - end

// kyc - start
type kycOnCommand command

func (c kycOnCommand) Execute() {
	c.R.call("http://firebase.com", command(c).constructRequest("on", "kyc"))
}

func (c kycOnCommand) Undo() {
	c.R.call("http://firebase.com", command(c).constructRequest("off", "kyc"))
}

type kycOffCommand command

func (c kycOffCommand) Execute() {
	c.R.call("http://firebase.com", command(c).constructRequest("off", "kyc"))
}

func (c kycOffCommand) Undo() {
	c.R.call("http://firebase.com", command(c).constructRequest("on", "kyc"))
} // kyc - end
