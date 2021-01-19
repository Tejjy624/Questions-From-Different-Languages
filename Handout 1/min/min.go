package min

// Min returns the minimum value in the arr,
// and 0 if arr is nil.
func Min(arr []int) int {
	// TODO: implement this function.
	if len(arr) == 0 {
		return 0
	} else{
		temp := arr[0]
		for _, value := range arr{
			if (value < temp) {
				temp = value
			}
		}
		return temp
	}
	//return 9876
}
