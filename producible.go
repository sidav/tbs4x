package main

type producible interface {
	getProductionCost() int
	getMoneyCost() int
	getProductionTypeRequired() int
	getName() string
}

type productionAbility struct {
	prodCode  int
	prodPower int
}

const (
	PRODUCTION_NONE = iota
	PRODUCTION_BUILDING
	PRODUCTION_INFANTRY
	PRODUCTION_GROUND_MECH
)

func getProductionTypeString(ptype int) string {
	switch ptype {
	case PRODUCTION_INFANTRY:
		return "training"
	case PRODUCTION_BUILDING:
		return "building"
	case PRODUCTION_GROUND_MECH:
		return "assembling"
	default:
		return "ERROR: NO NAME FOR PROD"
	}
}
