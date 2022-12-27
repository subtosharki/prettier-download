package main

import (
	"io"
	"net/http"
	"os"
)

var DownloadUrl = "https://raw.githubusercontent.com/subtosharki/my-prettier-file/main/.prettierrc"

func main() {
	resp, err := http.Get(DownloadUrl)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(".prettierrc", content, 0644)
	if err != nil {
		panic(err)
	}

	println("Downloaded")
}
