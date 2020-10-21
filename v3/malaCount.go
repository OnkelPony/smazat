package arithmetic

import "errors"

// Counts number of malas to achieve given repetitions
func MalaCounts(n int) (int, error) {
	if n%108 != 0 {
		return -1, errors.New("Repetitions not divisible by 108!")
	}
	return n / 108, nil
}
