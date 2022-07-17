package export

import (
	"encoding/json"
	"time"

	"itknown.io/design/strategy"
)

type OrderExportService struct {
}

type OrderExportServiceCondition struct {
	MerchantId          uint64    `json:"merchantId"`
	PlaceOrderStartTime time.Time `json:"placeOrderStartTime"`
	PlaceOrderEndTime   time.Time `json:"placeOrderEndTime"`
}

func (s *OrderExportService) QueryData(bizCond string) (error, interface{}) {

	var condition = &OrderExportServiceCondition{}

	if err := json.Unmarshal([]byte(bizCond), condition); err != nil {
		return err, nil
	}

	return nil, "订单数据"
}

func (s *OrderExportService) GetBizType() strategy.BizType {
	return strategy.BizOrder
}
