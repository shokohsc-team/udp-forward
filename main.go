package main

import "os"
import "fmt"

func main() {
	// Forward(src, dst). It's asynchronous.
	src := fmt.Sprintf("%s:%s", os.Getenv("SRC_ADDR"),os.Getenv("SRC_PORT"))
	dest := fmt.Sprintf("%s:%s", os.Getenv("DEST_ADDR"),os.Getenv("DEST_PORT"))
	forwarder, err := Forward(src, dest, DefaultTimeout)
	if err != nil {
		panic(err)
	}

	// Do something...
	for true {
		continue
	}

	// Stop the forwarder
	forwarder.Close()
}
