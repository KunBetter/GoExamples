// Defer
package main

/*
Defer is used to ensure that a function call
is performed later in a program has execution,
usually for purposes of cleanup.
defer is often used where e.g. ensure and
finally would be used in other languages.
*/

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("creating.")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

func Defer_() {
	/*
		Suppose we wanted to create a file,
		write to it, and then close when we are done.
		Here is how we could do that with defer.

		Immediately after getting a file object with createFile,
		we defer the closing of that file with closeFile.
		This will be executed at the end of
		the enclosing function (main),
		after writeFile has finished.
	*/
	f := createFile("defer.txt")
	defer closeFile(f)
	writeFile(f)
}
