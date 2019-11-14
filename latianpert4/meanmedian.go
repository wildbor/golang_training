package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	var commands string
	var array1 []string
	var array2 []string
	array3 := make([]int, 4, 4)

	fmt.Println("masukan command: ")
	fmt.Scan(commands)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	commands = scanner.Text()
	array1 = strings.Split(commands, ",")
	fmt.Println(len(array1))
	for i := 0; i < len(array1); i++ {
		array2 = strings.Split(strings.Trim(array1[i], " "), " ")
		switch array2[0] {
		case "insert":
			posisi, _ := strconv.Atoi(array2[1])
			value, _ := strconv.Atoi(array2[2])
			array3[posisi-1] = value
			fmt.Println(array3)
		case "remove":
			posisi, _ := strconv.Atoi(array2[1])
			array3[posisi-1] = array3[len(array3)-1]
			array3[len(array3)-1] = 0
			array3 = array3[:len(array3)-1]
			jumlah := len(array3) - 1
			array3 = array3[:jumlah]
			fmt.Println(array3)
		case "append":
			value, _ := strconv.Atoi(array2[1])
			fmt.Println(value)
			array3 = append(array3, value)
			fmt.Println(array3)
		case "sort":
			sort.Ints(array3)
			fmt.Println(array3)
		case "pop":
			array3[len(array3)-1] = 0
			array3 = array3[:len(array3)-1]
			fmt.Println(array3)
		case "reverse":
			sort.Ints(array3)
			sort.Slice(array3, func(i, j int) bool {
				return array3[i] > array3[j]
			})
			fmt.Println(array3)
		}

	}

}
