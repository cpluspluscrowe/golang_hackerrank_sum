package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func simpleArraySum(ar []string) int32 {
	var sum int32 = 0
	var cnt int = 0
	for _, numberString := range ar {
		cnt += 1
		number, err := strconv.Atoi(numberString)
		checkError(err)
		sum += int32(number)
	}
	fmt.Println(cnt)
	return sum
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 20000)
	reader.ReadLine()
	arTemp := strings.Split(readLine(reader), " ")
	result := simpleArraySum(arTemp)
	fmt.Println(result)
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
