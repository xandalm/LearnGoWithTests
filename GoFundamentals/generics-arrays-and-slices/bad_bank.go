package genericsarraysandslices

type Transaction struct {
	From, To string
	Sum      float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	adjusteBalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == name {
			currentBalance -= t.Sum
		}
		if t.To == name {
			currentBalance += t.Sum
		}
		return currentBalance
	}
	return Reduce(transactions, adjusteBalance, 0.0)
}
