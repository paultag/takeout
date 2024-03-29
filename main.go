package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/iancoleman/strcase"
)

func main() {
	pkg := os.Args[1]

	fmt.Printf(`// THIS FILE WAS AUTOGENERATED BY TAKEOUT; DO NOT EDIT
// THIS FILE BY HAND, PLEASE REBUILD USING TAKEOUT AND NOT BY HAND.
//

`)
	fmt.Printf("package %s\n", pkg)

	for _, filePath := range os.Args[2:] {
		perLine := 10
		name := strcase.ToLowerCamel(strings.Replace(path.Base(filePath), ".", "_", -1))

		bytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Printf("\nvar (\n")
		fmt.Printf("	%s []byte = []byte{\n", name)

		for i := 0; i < len(bytes); i = i + perLine {
			if i+perLine > len(bytes) {
				perLine = len(bytes) % perLine
			}

			fmt.Printf("		")
			for i, b := range bytes[i : i+perLine] {
				if i != 0 {
					fmt.Printf(" ")
				}
				fmt.Printf("%#x,", b)
			}
			fmt.Printf("\n")
		}

		fmt.Printf("	}\n")
		fmt.Printf("	%sLen uint = %d\n", name, len(bytes))
		fmt.Printf(")\n")
	}
}
