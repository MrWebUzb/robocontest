package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader *bufio.Reader = bufio.NewReader(os.Stdin)
	writer *bufio.Writer = bufio.NewWriter(os.Stdout)
)

func printf(f string, args ...interface{}) { fmt.Fprintf(writer, f, args...) }
func scanf(f string, args ...interface{})  { fmt.Fscanf(reader, f, args...) }

func main() {
	defer writer.Flush()

	var n uint64

	scanf("%d\n", &n)

	printf("%d\n", n/2+1)
}
