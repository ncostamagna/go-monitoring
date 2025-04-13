package bootstrap

import (
	"log"
	"os"
	"strconv"

	"github.com/ncostamagna/go-logger-hub/loghub"
)

func NewLogger() loghub.Logger {

	strTrace := os.Getenv("TRACE_LEVEL")
	trace, err := strconv.Atoi(strTrace)
	if err != nil {
		log.Printf("Error parsing TRACE_LEVEL: %v", err)
	}
	return loghub.New(
		loghub.NewNativeLogger(log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile), trace),
	)
}
