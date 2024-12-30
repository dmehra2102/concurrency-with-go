package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	main()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	fmt.Print(output)
	os.Stdout = stdOut

	if !strings.Contains(output, "509600") {
		t.Error("Wrong bank balance")
	}
}
