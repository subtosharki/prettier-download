package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"os"
	"strings"
)

func AddScriptToJSON() {
	println("Adding script to package.json")
	file, err := os.Open("package.json")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	formatField := `"format": "prettier --write \"src/**/*.ts\""`
	for i, line := range lines {
		if strings.Contains(line, `"scripts": {}`) {
			copy(lines[i+1:], lines[i:])
			lines = append(lines, "}")
			lines[i] = `"scripts": {`
			lines[i+1] = formatField + "},"
			break
		} else if strings.Contains(line, `"scripts": {`) {
			for _, line := range lines {
				if strings.Contains(line, formatField) {
					println("Script already exists, skipping")
					return
				}
			}
			if !strings.Contains(lines[i+1], "}") {
				lines[i+1] = formatField
			} else {
				copy(lines[i+1:], lines[i:])
				lines[i+1] = formatField
				if lines[i+2] != "}" {
					lines = append(lines, "}")
				}
			}
			break
		} else if line == "}" {
			lines = append(lines, formatField, "}")
			lines[i] = `,"scripts": {`
			break
		}
	}
	finalJSON := strings.Join(lines, ""+"")
	finalJSONBuffer := bytes.Buffer{}
	err = json.Indent(&finalJSONBuffer, []byte(finalJSON), "", "    ")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("package.json", []byte(finalJSONBuffer.String()), 0644)
	if err != nil {
		panic(err)
	}
	println("Script added to package.json")

}
