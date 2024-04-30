package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scan int
var tuble [][]int

func main() {

	rider := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Ведите число:")
		rider.Scan()
		read, _ := strconv.Atoi(rider.Text())
		scan = read
		start()
	}
}

func start() {

	for i := 1; i <= scan; i++ {
		rows := []int{}
		for v := 1; v <= scan; v++ {
			rows = append(rows, i*v)
		}
		tuble = append(tuble, rows)
	}
	render()
}

func render() {
	str := ""
	for a := 0; a < scan; a++ {
		for i, _ := range tuble {
			num := tuble[i][a]
			b := strconv.Itoa(num)
			str += b + " "
			if num < 10 {
				str += "  "
			}
			if num < 100 && num > 9 {
				str += " "
			}

		}
		fmt.Println(str)
		str = ""
	}
	tuble = tuble[:0][:0]
}
