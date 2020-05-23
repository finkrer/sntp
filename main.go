package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
)

const usage = `SNTP server
Usage:
  sntp [-d delay] [-p port]

Options:
  -d --delay <delay>  Increase time returned by this number, in seconds [default: 0].
  -p --port <port>	  Port to listen on [default: 123].
  -h --help    		  Show this message.`

var delay int

func main() {
	args, _ := docopt.ParseDoc(usage)
	d, err1 := args.Int("--delay")
	port, err2 := args.Int("--port")
	if err1 != nil || err2 != nil {
		fmt.Println("Option value should be a number")
		return
	}

	delay = d
	run(delay, port)
}
