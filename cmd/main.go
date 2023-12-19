package main

import (
	"os"

	"github.com/taylormonacelli/couplefly"
)

func main() {
	code := couplefly.Execute()
	os.Exit(code)
}
