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

func (obj *VwapPackage) Initialize(size int) {
	obj.Size = size
	obj.Start = 0
	obj.End = 0
	obj.Queue = make([]CustomChannelResp, 0)
}
