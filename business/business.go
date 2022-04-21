package business

import "zerohash/models"

var productResponseMap map[string]models.VwapPackage

func Init(keys []string, queueSize int) {
	productResponseMap = make(map[string]models.VwapPackage)

	if len(keys) > 0 {
		for _, key := range keys {
			var vwap models.VwapPackage
			vwap.Initialize(queueSize)
			productResponseMap[key] = vwap
		}
	}
}
