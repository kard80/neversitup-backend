package oddnumber

func OddNumber(input []int) int {
	reserver := map[int]int{}

	for _, v := range input {
		if reserver[v] == 0 {
			reserver[v] = 1
		} else {
			reserver[v]++
		}
	}

	for k, v := range reserver {
		if v%2 != 0 {
			return k
		}
	}

	panic("No odd number found")
}
