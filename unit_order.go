package main

type unitOrder struct {
	orderCode int
	x, y      int
}

func (o *unitOrder) getCoords() (int, int) {
	return o.x, o.y
}

const (
	ORDER_NONE = iota
	ORDER_MOVE
	ORDER_SETTLE
	ORDER_HARVEST
	ORDER_EXPLORE
	ORDERS_COUNT
)

func getNameOfOrder(code int) string {
	switch code {
	case ORDER_NONE:
		return "Wait"
	case ORDER_MOVE:
		return "Move"
	case ORDER_HARVEST:
		return "Harvest"
	case ORDER_SETTLE:
		return "Settle"
	case ORDER_EXPLORE:
		return "Explore"
	}
	panic("No order desc for code!")
}
