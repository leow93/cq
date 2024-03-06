package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

const MAX_BUFFER_SIZE = 1024 * 4 // 4KB is plenty for now

func read(reader *bufio.Reader) (error, []byte) {
	buf := make([]byte, 0, MAX_BUFFER_SIZE)
	for {
		tmp := make([]byte, 0, MAX_BUFFER_SIZE)
		n, err := reader.Read(tmp[:cap(tmp)])
		buf = append(buf, tmp[:n]...)

		if n == 0 {
			if err == nil {
				continue
			}

			if err == io.EOF {
				break
			}

			log.Fatal(err)
		}

		if err != nil && err != io.EOF {
			return err, nil
		}
	}
	return nil, buf
}

func readInputFile(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	err, buf := read(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(buf)
}

func readFromStdin() string {
	r := bufio.NewReader(os.Stdin)
	err, buf := read(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(buf)
}

func readInput() string {
	if len(os.Args) > 1 {
		return readInputFile(os.Args[1])
	}
	return readFromStdin()
}

func main() {
	data := readInput()
	fmt.Print(data)
}
