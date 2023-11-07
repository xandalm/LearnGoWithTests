package genericsarraysandslices

type Transaction struct {
	From, To string
	Sum      float64
}

func BalanceFor(transaction []Transaction, name string) float64 {
	var balance float64
	for _, t := range transaction {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
	}
	return balance
}
