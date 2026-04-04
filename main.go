package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/GrazianoJoa/jasl/scan"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadBytes('\n')

	sc := scan.NewScanner(text)
	l := sc.ScanTokens()
	
	for _, v := range l {
		fmt.Println(v.String())
	}
}
