package main

import (
	"os"
)

func main() {
	_, err := os.Stat(".prettierrc")
	if err == nil {
		println(".prettierrc found, skipping download")
	} else {
		DownloadFile()
	}

	_, err = os.Stat("package.json")
	if err != nil {
		println("package.json not found, skipping script setup")
	} else {
		AddScriptToJSON()
	}
}
