package services

import "math"

// SplitEqual divides the total amount equally among participants.
// Remainder cents are distributed to the first participants.
func SplitEqual(totalAmount float64, participantCount int) []float64 {
	if participantCount <= 0 {
		return nil
	}

	baseAmount := math.Floor(totalAmount*100/float64(participantCount)) / 100
	remainder := totalAmount - baseAmount*float64(participantCount)
	remainderCents := int(math.Round(remainder * 100))

	splits := make([]float64, participantCount)
	for i := range splits {
		splits[i] = baseAmount
		if i < remainderCents {
			splits[i] += 0.01
		}
	}

	return splits
}

// ValidateCustomSplit checks that custom split amounts sum to the total.
func ValidateCustomSplit(totalAmount float64, splits []float64) bool {
	var sum float64
	for _, s := range splits {
		sum += s
	}
	return math.Abs(sum-totalAmount) < 0.01
}
