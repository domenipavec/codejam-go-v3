package main

import "fmt"

func testCase(input *Input, testCaseNumber int) {
	fmt.Printf("Case #%d: ", testCaseNumber)

	N := input.Int()
	fmt.Println(N)
}
