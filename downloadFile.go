package main

import (
	"github.com/go-resty/resty/v2"
	"os"
)

func DownloadFile() {
	println("Downloading .prettierrc")
	downloadUrl := "https://raw.githubusercontent.com/subtosharki/my-prettier-file/main/.prettierrc"
	client := resty.New()
	content, err := client.R().Get(downloadUrl)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(".prettierrc", []byte(content.String()), 0644)
	if err != nil {
		panic(err)
	}
	println(".prettierrc Downloaded")
}
