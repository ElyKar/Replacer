package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// This is a script, so it's okay
var replacement = strings.NewReplacer("toto", "titi", "dog", "cat", "Vivaldi", "Bach")
var visited = make([]string, 128)
var isInterestingFile = func(fileContent string) bool {
	return strings.Contains(fileContent, "dog")
}

func replaceFile(filename string) {
	content, _ := ioutil.ReadFile(filename)
	stringContent := string(content)
	if isInterestingFile(stringContent) {
		fmt.Println("Visited ", filename)
		stringFields := strings.FieldsFunc(stringContent, func(c rune) bool { return c == '\n' })
		for i, line := range stringFields {
			stringFields[i] = replacement.Replace(line)
		}
		newString := strings.Join(stringFields, "\n")
		_ = ioutil.WriteFile(filename, []byte(newString), os.ModePerm)
	}
}
func getAllFiles(root string) {
	_ = filepath.Walk(root, visit)
}
func visit(path string, info os.FileInfo, err error) error {
	if !info.IsDir() {
		fmt.Println("Add ", path)
		visited = append(visited, path)
	}
	return nil
}
func main() {
	flag.Parse()
	getAllFiles(flag.Arg(0))
	for _, filename := range visited {
		replaceFile(filename)
	}
}
