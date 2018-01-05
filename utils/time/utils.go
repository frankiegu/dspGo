package common

func GetBitOneCount(input uint64) uint8 {
	var count uint8
	for input != 0 {
		count += 1
		input = input & (input - 1)
	}

	return count
}
