package main

import (
	"fmt"
	"os"
	"os/signal"
	"plugin"
	"syscall"
)

type DirectHandler interface {
	Handle() error
}

func main() {
	if len(os.Args) < 2 {
		panic("input plugin")
	}

	mod := os.Args[1]

	// load plugin
	plug, err := plugin.Open(mod)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// look up a symbol (an exported function or variable)
	// in this case, variable Handler
	symGreeter, err := plug.Lookup("Handler")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Assert that loaded symbol is of a desired type
	// in this case interface type DirectHandler (defined above)
	var caller DirectHandler
	caller, ok := symGreeter.(DirectHandler)
	if !ok {
		fmt.Println("unexpected type from module symbol")
		os.Exit(1)
	}

	// start the plugin
	fmt.Println("start service")
	go func() {
		err := caller.Handle()
		if err != nil {
			panic(err)
		}
	}()

	// listen signal
	sigch := make(chan os.Signal, 1)
	signal.Notify(sigch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for sig := range sigch {
		fmt.Println("recv sig: %s", sig)
		os.Exit(0)
	}
}
