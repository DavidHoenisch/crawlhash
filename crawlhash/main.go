package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Out struct {
	Hash      string `json:"hash"`
	Path      string `json:"path"`
	Generated string `json:"generated"`
}

var entries []Out

func getDatetime() string {
	currentTime := time.Now()

	return currentTime.Format(time.UnixDate)

}

func jsonWriter() {
	jsonData, err := json.MarshalIndent(entries, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("log.json", jsonData, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func entryHandler(path string, hash, generated string) {

	entry := Out{
		Generated: generated,
		Path:      path,
		Hash:      hash,
	}

	entries = append(entries, entry)

}

func hashFile(filep string) {
	// filep: path to the file to hash
	f, err := os.Open(filep)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		print(err)
	}

	fmt.Printf("%s	%x	%s\n", getDatetime(), h.Sum(nil), filep)
	entryHandler(filep, hex.EncodeToString(h.Sum(nil)), getDatetime())
}

func handler() filepath.WalkFunc {
	return func(filep string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}

		pathInfo, err := os.Stat(filep)
		if err != nil {
			return nil
		}

		if pathInfo.IsDir() == false {

			hashFile(filep)
		}

		return nil
	}
}

func main() {

	if len(os.Args) > 2 {
		fmt.Println("Missing path to scan")
		os.Exit(1)
	}

	rootPath := os.Args[1]

	err := filepath.Walk(rootPath, handler())
	if err != nil {
		fmt.Printf("the follow error occure %s ", err)
	}
	jsonWriter()
}
