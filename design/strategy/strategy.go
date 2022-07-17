package strategy

type BizType uint8

const (
	BizOrder BizType = iota
	BizSettlement
)

type ExportService interface {
	QueryData(condition string) (error, interface{})

	GetBizType() BizType
}

type ShapeType uint8

const (
	Square ShapeType = iota
	Circle
)

type Blueprint struct {
	Biz string
}

type Draw interface {
	Support(ShapeType) bool

	Shape(*Blueprint)
}
