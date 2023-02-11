package main

type producible interface {
	getCosts() (int, int)
	getProductionTypeRequired() int
	getName() string
}

type productionAbility struct {
	prodCode  int
	prodPower int
}

const (
	PRODUCTION_BUILDING = iota
	PRODUCTION_INFANTRY
	PRODUCTION_GROUND_MECH
)
