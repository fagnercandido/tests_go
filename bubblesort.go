package main

import "fmt"

func main() {

	var items = [6]int{3, 4, 1, 2, 6, 5}

	for aux := 0; aux < len(items); aux++ {
		for cont := 1; cont < len(items)-1; cont++ {
			if items[cont-1] > items[cont] {
				var temp = items[cont-1]
				items[cont-1] = items[cont]
				items[cont] = temp
			}
		}
	}
	for aux := 0; aux < len(items); aux++ {
		fmt.Println(items[aux])
	}

}
