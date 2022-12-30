package main

import (
	"bytes"
	"github.com/buger/jsonparser"
	"github.com/tidwall/pretty"
	"os"
)

func AddScriptToJSON() {
	println("Adding script to package.json")
	file, err := os.ReadFile("package.json")
	if err != nil {
		panic(err)
	}
	formatScript := `"prettier --write \"src/**/*.ts\""`
	scripts, _, _, err := jsonparser.Get(file, "scripts")
	if err != nil {
		println(`"scripts" field not found in package.json, skipping`)
		return
	}
	if bytes.Contains(scripts, []byte(`"format"`)) {
		println(`"format" script already exists, skipping`)
		return
	}
	file, err = jsonparser.Set(file, []byte(formatScript), "scripts", "format")
	if err != nil {
		panic(err)
	}
	file = pretty.Pretty(file)
	err = os.WriteFile("package.json", file, 0644)
	if err != nil {
		panic(err)
	}
	println("Script added to package.json")
}
