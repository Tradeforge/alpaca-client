package util

type Stringer interface {
	~string
}

func JoinTickers[T Stringer](tickers ...T) string {
	var str string
	for i, ticker := range tickers {
		if i != 0 {
			str += ","
		}
		str += string(ticker)
	}
	return str
}
