package presentation

import (
	"io"
	"log"
	"os"
	"testing"
	"time"
)

func TestUsingDefer(t *testing.T) {
	f, err := os.Open("testdata/sample.json")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		log.Println("Closing the file")
		if err:=f.Close(); err != nil {
			log.Println("Closing file failed...")
		}
	}()

	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
	Something()
	log.Println("Something is finished")
}

func TestUsingDefer2(t *testing.T) {
	f, cleanup, err := getResource("testdata/sample.json")

	if err != nil {
		t.Fatal("Something bad happened")
	}

	defer func() {
		cleanup()
	}()

	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}

func getResource(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		file.Close()
	}, nil
}


func Something() {
	defer func() {
		// proses ini hanya dimiliki oleh something
		log.Println("Something cleaning up")
	}()
	time.Sleep(1 * time.Second)
}