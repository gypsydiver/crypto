package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/gypsydiver/caesar/pkg/caesar"
)

// cli flags
var (
	n   = flag.Int("n", 1, "shift, use sign to specify left(-) or right(+)")
	dec = flag.Bool("d", false, "decrypt, uses opposite sign of shift")
	b   = flag.Bool("b", false, "bruteforce")
)

func main() {
	flag.Parse()

	rdr := bufio.NewReader(os.Stdin)
	str, err := rdr.ReadString('\n')
	if err != nil {
		// TODO: lift
		panic(err)
	}

	shift := *n
	if *dec {
		shift = -shift
	}

	if *b {
		caesar.Bruteforce(str)
		os.Exit(0)
	}

	fmt.Print(caesar.Apply(str, shift))
}
