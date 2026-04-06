package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/GrazianoJoa/jasl/scan"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: jasl [source]")
		os.Exit(1)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}

func runFile(src string) {
	data, err := os.ReadFile(src)
	if err != nil {
		panic("ERROR")
	}
	run(data)
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for ;; {
		fmt.Print("jasl> ")
		text, _ := reader.ReadBytes('\n')
		if strings.TrimSpace(string(text)) == "quit" {
			break
		}
		run(text)
	}
}

func run(src []byte) {
	sc := scan.NewScanner(src)
	l := sc.ScanTokens()

	for _, v := range l {
		fmt.Println(v.String())
	}
}
