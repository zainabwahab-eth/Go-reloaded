package operations

func Remove(checkIndex int, isdouble bool, inputSlice []string) []string {
	if isdouble {
		if checkIndex <= 0 {
			return inputSlice
		}
		return append(inputSlice[:checkIndex-1], inputSlice[checkIndex+1:]...)
	}

	return append(inputSlice[:checkIndex], inputSlice[checkIndex+1:]...)
}
