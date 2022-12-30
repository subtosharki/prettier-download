# prettier-download
A go program that downloads my `.prettierrc` file to my current directory, then adds my format script to my `package.json` file. 

To read the data from the `package.json` file, I write my own parsing system (which was a huge pain) which is used in the `addToScriptJSON.go` file on line 30.

This is for personal use but can easily be customized to your needs by doing the following:

1. Fork this repo
2. Change the `downloadUrl` variable in `downloadFile.go` to your own source of a `.prettierrc` file.
3. If you want to modify the script that is added to the package.json file, change the `formatField` variable in `addToScriptJSON.go` to your own script.
4. Use `make` to add the program to your `$GOPATH/bin` directory, **or**  `make binary` and add it to your `$PATH` manually.
5. Run `prettier-download` in your project directory to download the `.prettierrc` file and add the script to your `package.json` file.