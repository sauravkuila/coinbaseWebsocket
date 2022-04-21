package models

import (
	"strconv"
	"time"
)

type ChannelRequest struct {
	Type     string    `json:"type"`
	Channels []Channel `json:"channels"`
}

type Channel struct {
	Name       string   `json:"name"`
	ProductIds []string `json:"product_ids"`
}

type ChannelResponse struct {
	Type         string    `json:"type"`
	TradeID      int       `json:"trade_id"`
	MakerOrderID string    `json:"maker_order_id"`
	TakerOrderID string    `json:"taker_order_id"`
	Side         string    `json:"side"`
	Size         string    `json:"size"`
	Price        string    `json:"price"`
	ProductID    string    `json:"product_id"`
	Sequence     int64     `json:"sequence"`
	Time         time.Time `json:"time"`
}

func (obj *ChannelResponse) Convert() CustomChannelResp {
	var resp CustomChannelResp
	resp.Type = obj.Type
	resp.Side = obj.Side
	resp.ProductID = obj.ProductID
	resp.Sequence = obj.Sequence
	resp.Time = obj.Time
	resp.Size, _ = strconv.ParseFloat(obj.Size, 64)
	resp.Price, _ = strconv.ParseFloat(obj.Price, 64)
	return resp
}

type CustomChannelResp struct {
	Type      string
	Side      string
	Size      float64
	Price     float64
	ProductID string
	Sequence  int64
	Time      time.Time
}
