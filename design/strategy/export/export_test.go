package export

import (
	"encoding/json"
	"itknown.io/design/strategy"
	"testing"
	"time"
)

func TestExportService(t *testing.T) {

	var orderCondition = OrderExportServiceCondition{
		MerchantId:          5000247160,
		PlaceOrderStartTime: time.Now().Add(-12 * time.Hour),
		PlaceOrderEndTime:   time.Now(),
	}

	condition, _ := json.Marshal(&orderCondition)
	t.Log(condition)

	var services []strategy.ExportService
	services = append(services, &OrderExportService{})
	services = append(services, &SettlementExportService{})

	for _, srv := range services {

		if srv.GetBizType() == strategy.BizOrder {
			_, data := srv.QueryData(string(condition))
			t.Log(data)
		}
	}

}
