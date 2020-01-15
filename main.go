package main

import (
	"fmt"
	"os"
	"strings"
)

func convertRange(qualysRange string) string {
	octets := strings.Split(qualysRange, ".")
	endOctet := octets[len(octets)-1:][0]
	return strings.SplitAfter(qualysRange, "-")[0] + endOctet
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("qualysconvert [Qualys Network Range]")
		os.Exit(-1)
	}

	fmt.Println(convertRange(os.Args[1]))
}
