package vo

import (
	"sort"
	"strings"

	"github.com/unknwon/com"
)

var TwoToneSphereRuleMap = map[int]map[int]int{
	6: {1: 1, 0: 2},
	5: {1: 3, 0: 4},
	4: {1: 4, 0: 5},
	3: {1: 5, 0: 0},
	2: {1: 6, 0: 0},
	1: {1: 6, 0: 0},
	0: {1: 6, 0: 0},
}

var TwoToneSpherePriceMap = map[int]int{
	1: 5000000,
	2: 200000,
	3: 3000,
	4: 200,
	5: 10,
	6: 5,
}

type TwoToneSphereCalculator struct {
}

func (t TwoToneSphereCalculator) Calculate(number, winNumber string) (level, prize int) {
	splitNumber := strings.Split(number, "|")
	if len(splitNumber) != 2 {
		return level, prize
	}

	splitWinNumber := strings.Split(winNumber, "|")
	if len(splitWinNumber) != 2 {
		return level, prize
	}
	red, blue := t.helper(number)
	redWin, blueWin := t.helper(winNumber)

	redCount, blueCount := t.helperTriple(red, redWin), t.helperTriple(blue, blueWin)
	return t.helperQuadra(len(redCount), len(blueCount))
}

func (t TwoToneSphereCalculator) helper(number string) (red, blue []int) {
	splitNumber := strings.Split(number, "|")
	if len(splitNumber) != 2 {
		return red, blue
	}
	red = t.helperAgain(splitNumber[0])
	blue = t.helperAgain(splitNumber[1])
	return red, blue
}

func (t TwoToneSphereCalculator) helperAgain(number string) (out []int) {
	numbers := strings.Split(number, ",")
	numberMap := make(map[string]bool, len(numbers))

	for _, v := range numbers {
		if _, ok := numberMap[v]; ok {
			continue
		}
		if com.StrTo(v).MustInt() == 0 {
			continue
		}
		out = append(out, com.StrTo(v).MustInt())
	}
	sort.Ints(out)
	return out
}

func (t TwoToneSphereCalculator) helperTriple(numsOne, numsTwo []int) (out []int) {
	m := make(map[int]bool)
	for _, v := range numsOne {
		m[v] = true
	}
	for _, v := range numsTwo {
		if m[v] {
			out = append(out, v)
		}
	}
	return out
}

func (t TwoToneSphereCalculator) helperQuadra(redCount, blueCount int) (level, price int) {
	if redCount > 6 {
		redCount = 6
	}
	if redCount < 0 {
		redCount = 0
	}
	if blueCount < 0 {
		blueCount = 0
	}
	if blueCount > 1 {
		redCount = 1
	}
	level = TwoToneSphereRuleMap[redCount][blueCount]
	return level, TwoToneSpherePriceMap[level]
}
func init() {
	CalculatorMap[TwoToneSphere] = TwoToneSphereCalculator{}
}
