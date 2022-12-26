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
		println(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			println(err)
		}
	}(resp.Body)
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		println(err)
	}
	err = os.WriteFile(".prettierrc", content, 0644)
	if err != nil {
		println(err)
	}

	println("Downloaded")
}
