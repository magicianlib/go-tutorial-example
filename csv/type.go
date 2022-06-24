package csv

import "time"

type Order struct {
	OrderId        uint64    `csv:"订单号"`
	PlaceOrderTime time.Time `csv:"下单时间"`
	PayableAmount  uint64    `csv:"应付金额"`
	Currency       string    `csv:"币种"`
}
