package main

import "golds/oops"

func main() {
	creature := oops.Creature{"Anuj", true}
	creature.Dump()
	runner := oops.Runner(creature)

	flyingCreature := oops.FlyingCreature{creature, false}
	flyingCreature.Dump()

	runner.Run()
}
