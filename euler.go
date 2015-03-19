package main

import "fmt"

func get_table() [][15]int {
	return [][15]int{
		{75},
		{95, 64},
		{17, 47, 82},
		{18, 35, 87, 10},
		{20, 04, 82, 47, 65},
		{19, 01, 23, 75, 03, 34},
		{88, 02, 77, 73, 07, 63, 67},
		{99, 65, 04, 28, 06, 16, 70, 92},
		{41, 41, 26, 56, 83, 40, 80, 70, 33},
		{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
		{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
		{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
		{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
		{63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
		{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23},
	}

}

func MinInt(x, y int) (r int) {
	if x < y {
		return x
	}
	return y
}

func MaxInt(x, y int) (r int) {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println("This is a test.")
	table := get_table()
	fmt.Println("%v", table)
	for i := len(table) - 1; i > 0; i-- {
		// For each level in the array, we know how many elements should be
		// there, and we need to decrease it by one, taking the max between
		// each element
		for j := 0; j < i; j++ {
			table[i][j] = MaxInt(table[i][j], table[i][j+1])
		}
		// then, we need to add the current row to the next row, and then
		// continue

		for j := 0; j < i; j++ {
			table[i-1][j] = table[i-1][j] + table[i][j]
		}
		fmt.Println(table[i-1])
	}
}
