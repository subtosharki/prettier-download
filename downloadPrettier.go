package main

import (
	"os"
	"os/exec"
)

func DownloadPrettier() {
	println("Downloading Prettier")
	packageManager := ""
	_, err := os.Stat("yarn.lock")
	if err == nil {
		packageManager = "yarn"
	}
	_, err = os.Stat("package-lock.json")
	if err == nil {
		packageManager = "npm"
	}
	_, err = os.Stat("pnpm-lock.yaml")
	if err == nil {
		packageManager = "pnpm"
	}
	if packageManager == "" {
		println("No package manager found, skipping prettier download")
		return
	}
	_, err = os.Stat("package.json")
	if err != nil {
		println("package.json not found, skipping prettier download")
		return
	}
	if packageManager == "yarn" {
		err = exec.Command("yarn", "add", "prettier", "--dev").Run()
		if err != nil {
			panic(err)
		}
	} else if packageManager == "npm" {
		err = exec.Command("npm", "install", "prettier", "--dev").Run()
		if err != nil {
			panic(err)
		}
	} else if packageManager == "pnpm" {
		err = exec.Command("pnpm", "install", "prettier", "--dev").Run()
		if err != nil {
			panic(err)
		}
	}
	println("Prettier Downloaded")
}
