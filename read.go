package tgkorean

import (
	"bytes"
	"encoding/csv"
	"io"
	"log"
	"os"
)

func ReadCsv(name string) [][]string {
	s, err := Read(name)
	if err != nil {
		log.Fatal(err)
	}
	return ReadBytesCsv(s)
}

func ReadBytesCsv(s []byte) [][]string {
	r := csv.NewReader(bytes.NewReader(s))
	r.Comma = rune([]byte(";")[0])

	var res [][]string

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, record)
	}
	return res
}

func Read(name string) ([]byte, error) {
	file, err := os.Open(name)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	// defer the closing of our csvFile so that we can parse it later on
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("close file error %s", err)
		}
	}()
	byteArray, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return byteArray, nil
}
