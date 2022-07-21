package command_before

import (
	"github.com/harvanir/design-pattern/behavioural/common"
	"log"
)

var stf ServiceToggleFacade

func Setup() {
	stf = &serviceToggleFacade{client: newHttpClient()}
}

// Handle - delivery/controller/entry-point layer
func Handle(request common.Mode) {
	switch request {
	case common.On:
		stf.On()
		break
	case common.Off:
		stf.Off()
		break
	case common.OnIbft:
		stf.OnIbft()
		break
	case common.OffIbft:
		stf.OffIbft()
		break
	case common.OnTopup:
		stf.OnTopup()
		break
	case common.OffTopup:
		stf.OffTopup()
		break
	case common.OnKyc:
		stf.OnKyc()
		break
	case common.OffKyc:
		stf.OffKyc()
		break
	default:
		log.Fatal("unimplemented yet")
	}
}
