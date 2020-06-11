package vo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoToneSphereCalculator_Calculate(t *testing.T) {
	tests := []struct {
		desc      string
		number    string
		winNumber string
		level     int
		price     int
	}{
		{desc: "一等奖", number: "1,2,3,4,5,6|2", winNumber: "1,2,3,4,5,6|2", level: 1, price: TwoToneSpherePriceMap[1]},
		{desc: "二等奖", number: "01,2,3,4,5,6|2", winNumber: "1,02,3,4,5,6|1", level: 2, price: TwoToneSpherePriceMap[2]},
		{desc: "三等奖", number: "1,2,3,4,5,8|2", winNumber: "1,2,3,4,5,6|2", level: 3, price: TwoToneSpherePriceMap[3]},
		{desc: "四等奖", number: "1,2,3,4,15,8|02", winNumber: "1,2,3,4,5,6|2", level: 4, price: TwoToneSpherePriceMap[4]},
		{desc: "四等奖", number: "1,2,3,4,5,8,9|12", winNumber: "1,2,3,4,5,6|2", level: 4, price: TwoToneSpherePriceMap[4]},
		{desc: "五等奖", number: "1,2,3,4,8,9,20|2", winNumber: "1,2,3,4,5,6|05", level: 5, price: TwoToneSpherePriceMap[5]},
		{desc: "五等奖", number: "1,2,3,14,15,16,17|2", winNumber: "1,2,3,4,5,6|2", level: 5, price: TwoToneSpherePriceMap[5]},
		{desc: "六等奖", number: "1,2,13,14,15,16,17|2", winNumber: "1,2,3,4,5,6|2", level: 6, price: TwoToneSpherePriceMap[6]},
		{desc: "六等奖", number: "1,12,13,14,15,16,17|2", winNumber: "1,2,3,4,5,6|2", level: 6, price: TwoToneSpherePriceMap[6]},
		{desc: "六等奖", number: "11,12,13,14,15,16,17|2", winNumber: "1,2,3,4,5,6|2", level: 6, price: TwoToneSpherePriceMap[6]},
	}
	calculator := TwoToneSphereCalculator{}
	for _, tt := range tests {
		level, price := calculator.Calculate(tt.number, tt.winNumber)
		assert.Equal(t, tt.level, level, "%s: unexpected level. number:%s winNumber:%s", tt.desc, tt.number, tt.winNumber)
		assert.Equal(t, tt.price, price, "%s: unexpected price. number:%s winNumber:%s", tt.desc, tt.number, tt.winNumber)
	}
}
