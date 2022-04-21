package business

import (
	"fmt"
	"log"
	"zerohash/models"
)

func CalculateVWAP(tick models.ChannelResponse) {
	if tick.Type == "match" {
		vwapStruct := productResponseMap[tick.ProductID]
		customTick := tick.Convert()
		vwapStruct.NewTick = customTick
		fmt.Println(vwapStruct.Queue)

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
			// log.Printf("ProductID: %s | Price: %v | Size: %v | Date: %s", customTick.ProductID, customTick.Price, customTick.Size, customTick.Time.GoString())
		}
		printVWAP(&vwapStruct)
		productResponseMap[tick.ProductID] = vwapStruct
	}
}

func printVWAP(data *models.VwapPackage) {
	data.PriceVolumeSum = data.PriceVolumeSum + data.NewTick.Price*data.NewTick.Size - data.OldTick.Price*data.OldTick.Size
	data.VolumeSum = data.VolumeSum + data.NewTick.Size - data.OldTick.Size
	log.Printf("%s VMAC %s = %v", data.NewTick.Time.Format("2006-01-02 15:04:05"), data.NewTick.ProductID, data.PriceVolumeSum/data.VolumeSum)
}
