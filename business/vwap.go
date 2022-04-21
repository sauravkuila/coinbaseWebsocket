package business

import (
	"log"
	"zerohash/models"
)

func CalculateVWAP(tick models.ChannelResponse) float64 {
	if tick.Type == "match" {
		vwapStruct := productResponseMap[tick.ProductID]
		customTick := tick.Convert()
		vwapStruct.NewTick = customTick

		if vwapStruct.End == vwapStruct.Size {
			vwapStruct.OldTick = vwapStruct.Queue[vwapStruct.Start]
			vwapStruct.Queue[vwapStruct.Start] = vwapStruct.NewTick
			vwapStruct.Start += 1
			if vwapStruct.Start == vwapStruct.End {
				vwapStruct.Start = 0
			}
		} else {
			vwapStruct.Queue = append(vwapStruct.Queue, customTick)
			vwapStruct.End += 1
		}
		vwap := printVWAP(&vwapStruct)
		productResponseMap[tick.ProductID] = vwapStruct
		return vwap
	}
	return -1
}

func printVWAP(data *models.VwapPackage) float64 {
	data.PriceVolumeSum = data.PriceVolumeSum + data.NewTick.Price*data.NewTick.Size - data.OldTick.Price*data.OldTick.Size
	data.VolumeSum = data.VolumeSum + data.NewTick.Size - data.OldTick.Size
	actualVWAP := data.PriceVolumeSum / data.VolumeSum
	log.Printf("%s VWAP %s = %v", data.NewTick.Time.Format("2006-01-02 15:04:05"), data.NewTick.ProductID, actualVWAP)
	return actualVWAP
}
