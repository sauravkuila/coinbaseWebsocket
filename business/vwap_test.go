package business

import (
	"testing"
	"time"
	"zerohash/models"

	"github.com/stretchr/testify/assert"
)

func TestCalculateVWAP(t *testing.T) {
	var tick models.ChannelResponse
	tick.Type = "heartbeat"

	val := CalculateVWAP(tick)
	assert.Equal(t, val, float64(-1))

	var m models.VwapPackage
	m.Initialize(2)
	productResponseMap = make(map[string]models.VwapPackage)
	productResponseMap["ETH-USD"] = m

	var tick1 models.ChannelResponse
	tick1.Type = "match"
	tick1.ProductID = "ETH-USD"
	tick1.Price = "3075.65"
	tick1.Size = "0.06158154"
	tick1.Side = "sell"
	tick1.Time = time.Now().Add(-5 * time.Second)

	val = CalculateVWAP(tick1)
	assert.Equal(t, val, 3075.65)

	var tick2 models.ChannelResponse
	tick2.Type = "match"
	tick2.ProductID = "ETH-USD"
	tick2.Price = "3060"
	tick2.Size = "0.00238"
	tick2.Side = "buy"
	tick2.Time = time.Now().Add(-2 * time.Second)

	CalculateVWAP(tick2)
	tick3 := tick2
	tick3.Size = "0.1124"
	val = CalculateVWAP(tick3)
	assert.Equal(t, val, 3060.0000000000005)

	tick4 := tick3
	tick4.Size = "0.00056"
	val = CalculateVWAP(tick3)
	assert.Equal(t, val, 3060.0000000000005)
}
