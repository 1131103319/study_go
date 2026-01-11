package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("Creating", p)
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}
func writeFile(f *os.File) {
	fmt.Println("Writing", f.Name())
	fmt.Fprintln(f, "data")
}
func closeFile(f *os.File) {
	fmt.Println("Closing", f.Name())
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error closing file: %v", err)
		os.Exit(1)
	}
}
func main() {
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}
