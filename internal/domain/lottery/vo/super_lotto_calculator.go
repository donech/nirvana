package vo

import "strings"

func init() {
	CalculatorMap[SuperLotto] = SuperLottoCalculator{}
}

type rule struct {
	red   int
	blue  int
	level int
	price int
}

var SuperLottoRuleMap = []rule{
	{red: 5, blue: 2, level: 1, price: 5000000},
	{red: 5, blue: 1, level: 2, price: 5000000},
	{red: 5, blue: 0, level: 3, price: 10000},
	{red: 4, blue: 2, level: 4, price: 3000},
	{red: 4, blue: 1, level: 5, price: 300},
	{red: 3, blue: 2, level: 6, price: 200},
	{red: 4, blue: 0, level: 7, price: 100},
	{red: 3, blue: 1, level: 8, price: 15},
	{red: 3, blue: 0, level: 9, price: 5},
	{red: 1, blue: 2, level: 9, price: 5},
	{red: 2, blue: 1, level: 9, price: 5},
	{red: 0, blue: 2, level: 9, price: 5},
}

type SuperLottoCalculator struct{}

func (c SuperLottoCalculator) Calculate(number, winNumber string) (level, prize int) {
	splitNumber := strings.Split(number, "|")
	if len(splitNumber) != 2 {
		return 0, 0
	}
	splitWinNumber := strings.Split(winNumber, "|")
	if len(splitWinNumber) != 2 {
		return 0, 0
	}
	redCount := c.calculateCount(splitNumber[0], splitWinNumber[0])
	blueCount := c.calculateCount(splitNumber[1], splitWinNumber[1])
	return c.calculateLevel(redCount, blueCount)
}

func (SuperLottoCalculator) calculateCount(number, winNumber string) int {
	numbers := strings.Split(number, ",")
	winNumbers := strings.Split(winNumber, ",")
	winNumbersMap := make(map[string]bool)
	for _, v := range winNumbers {
		winNumbersMap[v] = true
	}
	count := 0
	for _, v := range numbers {
		if winNumbersMap[v] {
			count++
		}
	}
	return count
}

func (SuperLottoCalculator) calculateLevel(redCount, blueCount int) (int, int) {
	for _, v := range SuperLottoRuleMap {
		if redCount == v.red && blueCount == v.blue {
			return v.level, v.price
		}
	}
	return 0, 0
}
