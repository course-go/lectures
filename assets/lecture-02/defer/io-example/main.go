package main

import (
	"fmt"
	"io"
	"os"
)

// START OMIT

func main() {
	testCopyFile("08_defer_practical_usage.go", "new.go")
	testCopyFile("not_exists", "new.go")
	testCopyFile("08_defer_practical_usage.go", "")
	testCopyFile("08_defer_practical_usage.go", "/dev/full")
	testCopyFile("/dev/null", "new2.go")
}

func copyFile(srcName, dstName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}

	defer src.Close()
	dst, err := os.Create(dstName)
	if err != nil {
		return
	}

	defer dst.Close()
	return io.Copy(dst, src)
}

// END OMIT

func testCopyFile(srcName, dstName string) {
	copied, err := copyFile(srcName, dstName)
	if err != nil {
		fmt.Printf("failed running copyFile('%s', '%s')\n", srcName, dstName)
	} else {
		fmt.Printf("copyFile('%s', '%s') copied %d bytes\n", srcName, dstName, copied)
	}
}
