package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// CLI application
// 1. Read the input file from the command line (e.g `go run main.go input.txt` or cat input.txt | go run main.go)

func readInputFile(filename string) {
	// Read the file
}

func main() {

	nBytes, nChunks := int64(0), int64(0)
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*1024)

	for {

		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]

		if n == 0 {

			if err == nil {
				continue
			}

			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		nChunks++
		nBytes += int64(len(buf))

		fmt.Println(string(buf))

		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}

	fmt.Println("Bytes:", nBytes, "Chunks:", nChunks)
}
