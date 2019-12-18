package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	lorem "github.com/drhodes/golorem"
	bytesize "github.com/inhies/go-bytesize"
)

const (
	defaultSizeStr = "32kb"
)

var (
	size = flag.String("size", defaultSizeStr, "desired size of output")
	seed = flag.Int64("seed", 0, "use specific random seed")
)

func main() {
	flag.Parse()
	desiredBytes, err := bytesize.Parse(*size)
	if err != nil {
		log.Fatal(err)
	}

	if *seed == 0 {
		*seed = time.Now().UTC().UnixNano()
	}
	rand.Seed(*seed)

	fmt.Fprintf(os.Stderr, "Generating >= %d bytes of lorem ipsum with seed %d\n", desiredBytes, *seed)

	f := bufio.NewWriter(os.Stdout)
	defer f.Flush()
	for bs := uint64(0); bs < uint64(desiredBytes); {
		bW, err := f.WriteString(lorem.Paragraph(1, 7) + "\n\n")
		bs += uint64(bW)
		if err != nil {
			log.Fatal(err)
		}
	}
}
