package main

type unitOrder struct {
	orderCode int
	x, y      int
}

const (
	ORDER_NONE = iota
	ORDER_MOVE
	ORDER_SETTLE
	ORDER_HARVEST
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
	}
	panic("No order desc for code!")
}
