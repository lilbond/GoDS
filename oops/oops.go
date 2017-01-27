package oops

import (
	"fmt"
)

//Creature represents a creature, it could be real or imaginary like a comic creature
type Creature struct {
	Name string
	Real bool
}

//FlyingCreature represents a creature who can fly
type FlyingCreature struct {
	Creature
	HighFlying bool
}

func (c Creature) Dump() {
	fmt.Printf("Name: %s, Real: %t\n", c.Name, c.Real)
}

func (f FlyingCreature) Dump() {
	fmt.Printf("Name: %s, Real: %t, HighFlying: %t\n", f.Name, f.Real, f.HighFlying)
}

//Runner implements run
type Runner interface {
	Run()
}

func (c Creature) Run() {
	fmt.Printf("Name: %s is running\n", c.Name)
}
