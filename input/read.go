package input

import (
	"bufio"
	"io"
	"log"
	"os"
)

// MaxBufferSize 4KB is plenty to allocate initially
// append() will grow the array dynamically as needed
const MaxBufferSize = 1024 * 4

func Read(reader *bufio.Reader) (error, []byte) {
	buf := make([]byte, 0, MaxBufferSize)
	for {
		tmp := make([]byte, 0, MaxBufferSize)
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
	err, buf := Read(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(buf)
}

func readFromStdin() string {
	r := bufio.NewReader(os.Stdin)
	err, buf := Read(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(buf)
}

func ReadInput() string {
	if len(os.Args) > 1 {
		return readInputFile(os.Args[1])
	}
	return readFromStdin()
}
