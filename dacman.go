package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	var outstr []byte
	var i int = 0
	files, _ := ioutil.ReadDir("./data")
	for _, file := range files {
		fname := file.Name()
		ext := strings.ToLower(fname[len(fname)-3:])
		if ext == "dat" || ext == "ccl" {
			outstr = []byte("file=" + fname + "\nusepasswords=n\nruleset=ms\n")
			ioutil.WriteFile("./sets/" + fname + "-MS.dac", outstr, 0644)
			outstr = []byte("file=" + fname + "\nusepasswords=n\nruleset=lynx\n")
			ioutil.WriteFile("./sets/" + fname + "-Lynx.dac", outstr, 0644)
			i++
		}
	}
	fmt.Printf("Generated DAC files for %d level sets\n", i)
}
