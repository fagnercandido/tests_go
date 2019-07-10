package main

func main() {

}

//SumBigger function to get the bigger value
func SumBigger(alice []int32, bob []int32) [2]int32 {
	var aIndex int32
	var bIndex int32 = 1

	result := [...]int32{0, 0}

	for index := range alice {
		if alice[index] > bob[index] {
			result[aIndex] = result[aIndex] + 1
		}
		if bob[index] > alice[index] {
			result[bIndex] = result[bIndex] + 1
		}
	}

	return result
}
