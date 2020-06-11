package vo

type Calculator interface {
	Calculate(number, winNumber string) (level, prize int)
}

var CalculatorMap = make(map[TicketType]Calculator)

func GetCalculator(ticketType TicketType) Calculator {
	out, ok := CalculatorMap[ticketType]
	if !ok {
		return CalculatorMap[TwoToneSphere]
	}
	return out
}
