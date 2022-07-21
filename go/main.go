package main

import (
	"github.com/harvanir/design-pattern/behavioural/command"
	command_before "github.com/harvanir/design-pattern/behavioural/command-before"
	"github.com/harvanir/design-pattern/behavioural/common"
	"log"
)

func main() {
	commandBefore()
	log.Println("################################################")
	commandAfter()
}

func commandBefore() {
	log.Println("setup commandBefore...")
	command_before.Setup()

	log.Println("[before command] handle On")
	command_before.Handle(common.On)

	log.Println()
	log.Println("[before command] handle Off")
	command_before.Handle(common.Off)

	log.Println()
	log.Println("[before command] handle OnIbft")
	command_before.Handle(common.OnIbft)

	log.Println()
	log.Println("[before command] handle OffIbft")
	command_before.Handle(common.OffIbft)

	log.Println()
	log.Println("[before command] handle OnTopup")
	command_before.Handle(common.OnTopup)

	log.Println()
	log.Println("[before command] handle OffTopup")
	command_before.Handle(common.OffTopup)

	log.Println()
	log.Println("[before command] handle OnKyc")
	command_before.Handle(common.OnKyc)

	log.Println()
	log.Println("[before command] handle OffKyc")
	command_before.Handle(common.OffKyc)
}

func commandAfter() {
	log.Println("setup commandAfter...")
	command.Setup()

	log.Println("[command] handle On")
	command.Handle(common.On)

	log.Println()
	log.Println("[command] handle Off")
	command.Handle(common.Off)

	log.Println()
	log.Println("[command] handle OnIbft")
	command.Handle(common.OnIbft)

	log.Println()
	log.Println("[command] handle OffIbft")
	command.Handle(common.OffIbft)

	log.Println()
	log.Println("[command] handle OnTopup")
	command.Handle(common.OnTopup)

	log.Println()
	log.Println("[command] handle OffTopup")
	command.Handle(common.OffTopup)

	log.Println()
	log.Println("[command] handle OnKyc")
	command.Handle(common.OnKyc)

	log.Println()
	log.Println("[before command] handle OffKyc")
	command.Handle(common.OffKyc)
}
