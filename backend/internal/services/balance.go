package services

import "math"

type Balance struct {
	UserID uint    `json:"user_id"`
	Amount float64 `json:"amount"`
}

type Transfer struct {
	From   uint    `json:"from"`
	To     uint    `json:"to"`
	Amount float64 `json:"amount"`
}

// CalculateBalances computes net balances from expenses.
// Positive = owed money (others owe you), Negative = owes money (you owe others).
func CalculateBalances(expenses []ExpenseData) []Balance {
	balanceMap := make(map[uint]float64)

	for _, exp := range expenses {
		balanceMap[exp.PaidBy] += exp.Amount
		for _, split := range exp.Splits {
			balanceMap[split.UserID] -= split.Amount
		}
	}

	var balances []Balance
	for userID, amount := range balanceMap {
		balances = append(balances, Balance{
			UserID: userID,
			Amount: math.Round(amount*100) / 100,
		})
	}

	return balances
}

// SimplifyDebts uses a greedy algorithm to minimize the number of transfers.
func SimplifyDebts(balances []Balance) []Transfer {
	var debtors, creditors []Balance

	for _, b := range balances {
		if b.Amount < -0.01 {
			debtors = append(debtors, Balance{UserID: b.UserID, Amount: -b.Amount})
		} else if b.Amount > 0.01 {
			creditors = append(creditors, b)
		}
	}

	var transfers []Transfer
	i, j := 0, 0

	for i < len(debtors) && j < len(creditors) {
		amount := math.Min(debtors[i].Amount, creditors[j].Amount)
		amount = math.Round(amount*100) / 100

		if amount > 0.01 {
			transfers = append(transfers, Transfer{
				From:   debtors[i].UserID,
				To:     creditors[j].UserID,
				Amount: amount,
			})
		}

		debtors[i].Amount -= amount
		creditors[j].Amount -= amount

		if debtors[i].Amount < 0.01 {
			i++
		}
		if creditors[j].Amount < 0.01 {
			j++
		}
	}

	return transfers
}

type ExpenseData struct {
	PaidBy uint
	Amount float64
	Splits []SplitData
}

type SplitData struct {
	UserID uint
	Amount float64
}
