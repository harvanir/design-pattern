package command

import (
	"github.com/harvanir/design-pattern/behavioural/common"
	"log"
)

func Setup() {
	r := newHttpClient()
	sai = &serviceAvailabilityInvoker{
		onIbft:   &ibftOnCommand{R: r},
		offIbft:  &ibftOffCommand{R: r},
		onTopup:  &topupOnCommand{R: r},
		offTopup: &topupOffCommand{R: r},
		onKyc:    &kycOnCommand{R: r},
		offKyc:   &kycOffCommand{R: r},
	}
}

// Handle - delivery/controller/entry-point layer
func Handle(request common.Mode) {
	switch request {
	case common.On:
		sai.On()
		break
	case common.Off:
		sai.Off()
		break
	case common.OnIbft:
		sai.OnIbft()
		break
	case common.OffIbft:
		sai.OffIbft()
		break
	case common.OnTopup:
		sai.OnTopup()
		break
	case common.OffTopup:
		sai.OffTopup()
		break
	case common.OnKyc:
		sai.OnKyc()
		break
	case common.OffKyc:
		sai.OffKyc()
		break
	default:
		log.Fatal("unimplemented yet")
	}
}
