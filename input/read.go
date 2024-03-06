package input

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// 4KB is plenty to allocate initially
// append() will grow the array dynamically as needed
const MAX_BUFFER_SIZE = 1024 * 4

func Read(reader *bufio.Reader) (error, []byte) {
	buf := make([]byte, 0, MAX_BUFFER_SIZE)
	for {
		tmp := make([]byte, 0, MAX_BUFFER_SIZE)
		fmt.Println(len(buf), cap(buf))
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
	fmt.Println(len(buf), cap(buf))
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
