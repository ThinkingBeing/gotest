package main

import (
	"fmt"

	"github.com/arnauddri/algorithms/algorithms/sorting/bubble-sort"
	//"github.com/gotest/bubble"
	"github.com/gotest/stringutil"
)

//Name name
type Name struct {
	first string
	last  string
}

func main() {

	msg := "Hello, 世界! Hello,"
	fmt.Printf(msg + "\n")
	rmsg := stringutil.Reverse(msg)
	fmt.Printf(rmsg + "\n")

	myName := Name{"高飞", "李"}
	fmt.Println(myName.last + " " + myName.first)

	np := myName
	np.first = "世民"
	fmt.Println(myName.first)

	mn := [2]string{"李", "高飞"}
	fmt.Println(mn[0], mn[1])

	mnSlice := []string{"li", "gao", "fei"}
	fmt.Println(mnSlice)

	for i := 0; i < len(mnSlice); i++ {
		fmt.Println(mnSlice[i])
	}

	fmt.Println(mnSlice[1:3])
	fmt.Println(stringutil.WordCount(msg))

	f := stringutil.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d,", f())
	}

	fmt.Println("")

	arr := []int{4, 2, 23, 1, 43, 13, 56}
	bubble.Sort(arr)
	fmt.Println(arr)
}
