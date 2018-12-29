package main

import (
	csremotelib "github.com/jcteng/cryptosocksgo/pkg/csremotelib"

	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	csremotelib.Start()
}
