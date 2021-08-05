package slices

func InSlices(slices []string, target string) bool  {
	for _, v := range slices {
		if v == target {
			return true
		}
	}

	return false
}

func SearchSlices(slices []string, target string) int  {
	for k, v := range slices {
		if v == target {
			return k
		}
	}

	return -1
}

func Remove(slices []string, i int) []string  {
	slices[len(slices)-1], slices[i] = slices[i], slices[len(slices)-1]
	return slices[:len(slices)-1]
}


func SearchSlicesUint64(slices []uint64, target uint64) int  {
	for k, v := range slices {
		if v == target {
			return k
		}
	}

	return -1
}

func RemoveUint64(slices []uint64, i int) []uint64  {
	slices[len(slices)-1], slices[i] = slices[i], slices[len(slices)-1]
	return slices[:len(slices)-1]
}

