package main

import (
	"github.com/bcpitutor/tiki/tlogger"
)

func init() {
	tlogger.InitiliazeTikiLogger()
}

func main() {
	tlogger.Log.Info("Hello World!")
}
