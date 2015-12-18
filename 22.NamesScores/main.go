/*
Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/
package main

import (
	// "fmt"
	"io/ioutil"
	// "reflect"
	"sort"
	"strings"
)

func main() {
	fileBuf, err := ioutil.ReadFile("p022_names.txt")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%v", reflect.TypeOf(fileBuf))
	fileStr := string(fileBuf)
	arr := strings.Split(fileStr, ",")
	for i := range arr {
		arr[i] = strings.Split(arr[i], "\"")[1]
	}

	sort.Strings(arr)

	result := uint64(0)
	for i := range arr {
		result += uint64(i+1) * uint64(alphabeticValue(arr[i]))
	}
	println(result)
}

func alphabeticValue(in string) int {
	sum := 0
	for i := range in {
		sum += int(in[i]) - int('A') + 1
	}
	return sum
}
