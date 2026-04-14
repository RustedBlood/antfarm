package domain

import "strconv"

type Amount int64

type Order struct {
	OrderID     int64
	Amount      Amount
	Description string
}

func (a Amount) ToString() string {
	return strconv.FormatInt(int64(a), 10)
}

func AmountFromString(a string) (Amount, error) {
	a_int, err := strconv.ParseInt(a, 10, 64)

	if err != nil {
		return 0, err
	}

	return Amount(a_int), nil
}
