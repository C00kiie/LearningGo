package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func wrapper_logic(data int) int {
	rest := 100 + data
	return rest
}
func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			return
		}

		fmt.Print("-> ", string(netData))

		results, err := strconv.ParseInt(string(netData)[:1], 10, 32)
		results_with_logic_done := wrapper_logic(int(results))
		if err != nil {
			fmt.Println(string(netData))
			panic(err)
		}
		fmt.Fprintf(c, strconv.Itoa(results_with_logic_done)+"\n")

	}
}
