package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aiialzy/wasmer/binary"
)

func main() {
	dumpFlag := flag.Bool("d", false, "dump")
	flag.Parse()
	if flag.NArg() != 1 {
		fmt.Println("Usage: wasmgo [-d] filename")
		os.Exit(1)
	}

	module, err := binary.DecodeFile(flag.Args()[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *dumpFlag {
		dump(module)
	}
}
