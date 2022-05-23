package main

import "fmt"

var block = "package"

func main() {
	var a uint32 = 1
	var b uint32 = 2
	fmt.Println(a - b)
	fmt.Println(2^32)

}

func ttt(aa *[2]int) {
	aa[0] = 1212
}

func addStrings(num1 string, num2 string) string {
	ln := max(len(num1), len(num2))
	byteDance := make([]byte, ln+1)
	var i, j = len(num1) - 1, len(num2) - 1
	for i >= 0 || j >= 0 {
		if i >= 0 && j >= 0 {
			sum := num1[i] + num2[j] - 2*'0' + byteDance[max(i, j)+1]
			byteDance[max(i, j)] = sum / 10
			byteDance[max(i, j)+1] = sum%10 + '0'
			i--
			j--
			continue
		}
		if i >= 0 {
			sum := num1[i] - '0' + byteDance[i+1]
			byteDance[i] = sum / 10
			byteDance[i+1] = sum%10 + '0'
			i--
			continue
		}
		if j >= 0 {
			sum := num2[j] - '0' + byteDance[j+1]
			byteDance[j] = sum / 10
			byteDance[j+1] = sum%10 + '0'
			j--
		}
	}
	if byteDance[0] == 0 {
		return string(byteDance[1:])
	}
	byteDance[0] += '0'
	return string(byteDance)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	ln1, ln2 := len(num1), len(num2)
	ln := ln1 + ln2
	res := make([]byte, ln)
	for i := ln1 - 1; i >= 0; i-- {
		n1 := num1[i] - '0'
		for j := ln2 - 1; j >= 0; j-- {
			n2 := num2[j] - '0'
			sum := n1*n2 + res[i+j+1]
			res[i+j+1] = sum % 10
			res[i+j] += sum / 10
		}
	}
	if res[0] == 0 {
		res = res[1:]
	}
	for i, _ := range res {
		res[i] += '0'
	}
	return string(res)
}

type test struct {
	InnerTest innerTest
}

type innerTest struct {
	addr string
}
