package utils // Change to utils package

func FindAvg(scores []int) int {
	total := FindTotal(scores)
	return total / len(scores)
}

func FindTotal(scores []int) int {
	total := 0
	for _, score := range scores {
		total += score
	}
	return total
}

func FindMinMax(scores []int) (int, int) {

	if len(scores) == 0 {
		return 0, 0 // Return 0 if scores is empty
	}

	min := scores[0]
	max := scores[0]
	for _, score := range scores {
		if score < min {
			min = score
		}
		if score > max {
			max = score
		}
	}
	return min, max
}

// FindStudentScoreGreaterThanAvg filters students whose scores are greater than the average.
func FindStudentScoreGreaterThanAvg(scores map[string]int, avg int) map[string]int {
	result := make(map[string]int)

	if len(scores) == 0 {
		return result
	}

	for key, value := range scores {
		if value > avg {
			result[key] = value
		}
	}

	return result
}

// FindMinMaxAvg returns the minimum, maximum, and average of the scores
// Improved performance to handle 3 times iteration (min, max, total)
func FindMinMaxAvg(scores []int) (int, int, int) {

	// Check if scores is empty
	if len(scores) == 0 {
		return 0, 0, 0 // Return 0 if scores is empty
	}

	min := scores[0]
	max := scores[0]
	total := 0

	for _, score := range scores {
		if score < min {
			min = score
		}

		if score > max {
			max = score
		}
		total += score
	}

	avg := total / len(scores)
	return min, max, avg
}
