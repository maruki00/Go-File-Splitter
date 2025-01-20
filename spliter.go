package main

import (
	"io"
	"os"
	"strconv"

	"github.com/pkg/profile"
)

func readByte(r io.Reader, buffer []byte) error {
	// var buffer [1]byte
	_, err := r.Read(buffer[:])
	return err
}
func writeTofile(data []byte, name string, chunk int) error {
	file, err := os.Create("spl/" + name + strconv.Itoa(chunk))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}

func main() {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	// defer profile.Start(profile.BlockProfile, profile.ProfilePath(".")).Stop()
	d := make([]byte, 0)
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()
	os.MkdirAll("spl", 0777)
	reader := file //bufio.NewReader(file)
	r := make([]byte, 1)
	lines := 1
	chunk := 1
	for {
		err := readByte(reader, r)
		if err == io.EOF {
			break
		}

		d = append(d, r[0])
		if r[0] == '\n' {
			lines++
		}

		if lines >= 1000 {
			writeTofile(d, "chunk", chunk)
			chunk++
			lines = 1
		}
	}
	if len(d) != 0 {
		writeTofile(d, "chunk", chunk)
	}
}
