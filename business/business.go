package business

import "zerohash/models"

var productResponseMap map[string]models.VwapPackage

func Init(keys []string, queueSize int) {
	productResponseMap = make(map[string]models.VwapPackage)

	if len(keys) > 0 {
		for _, key := range keys {
			var vwap models.VwapPackage
			vwap.Size = queueSize
			vwap.Start = 0
			vwap.End = 0
			vwap.Queue = make([]models.CustomChannelResp, 0)
			productResponseMap[key] = vwap
		}
	}
}
