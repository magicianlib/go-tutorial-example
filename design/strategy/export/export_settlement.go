package export

import (
	"encoding/json"

	"itknown.io/design/strategy"
)

type SettlementExportService struct {
}

type SettlementExportServiceCondition struct {
	MerchantId       uint64 `json:"merchantId"`
	SettlementStatus uint8  `json:"settlementStatus"`
}

func (s *SettlementExportService) QueryData(bizCond string) (error, interface{}) {

	var condition = &SettlementExportServiceCondition{}

	if err := json.Unmarshal([]byte(bizCond), condition); err != nil {
		return err, nil
	}

	return nil, "结算数据"
}

func (s *SettlementExportService) GetBizType() strategy.BizType {
	return strategy.BizSettlement
}
