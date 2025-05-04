package main

import (
	"fmt"
	"log"
	"os"

	"go.creack.net/bistro-matic/parser"
)

func main() {
	if len(os.Args) != 4 {
		log.Fatalf(
			"Usage: %s base ops\"()+=-*/%%\" exp_len\nExample:\n"+
				"\techo '1+2*3/(4-5)' | %s 0123456789 '()+-*/%%' 11",
			os.Args[0], os.Args[0])
	}
	checkBase(os.Args[1])
	checkOps(os.Args[2])
	size, err := parser.ParseNumberBase(os.Args[3], "0123456789")
	if err != nil {
		log.Fatalf("invalid size: %s\n", err)
	}
	expr := getExpr(size)
	result, err := evalExpr(os.Args[1], os.Args[2], expr, size)
	if err != nil {
		log.Fatalf("could not eval: %s\n", err)
	}
	fmt.Print(result)
}

func checkBase(base string) {
	if len(base) < 2 {
		log.Fatal("Bad base\n")
	}
}

func getExpr(size int) string {
	if size <= 0 {
		log.Printf("Bad expr len\n")
		os.Exit(2)
	}
	expr := make([]byte, size)
	n, err := os.Stdin.Read(expr)
	if err != nil {
		log.Printf("could not read: %s\n", err)
		os.Exit(4)
	}
	if n != size {
		log.Printf("could not read\n")
		os.Exit(4)
	}
	return string(expr)
}

func checkOps(ops string) {
	if len(ops) != 7 {
		log.Printf("Bad ops\n")
		os.Exit(5)
	}
}
