package models

type VwapPackage struct {
	Size           int
	Start          int
	End            int
	Queue          []CustomChannelResp
	NewTick        CustomChannelResp
	OldTick        CustomChannelResp
	VolumeSum      float64
	PriceVolumeSum float64
}
