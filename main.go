package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

func must(err error) {
  if err != nil {
    panic(err)
  }
}

func must2[T any](val T, err error) T {
  must(err)
  return val
}

func usage() {
  fmt.Fprintln(os.Stdout, "Usage: cutnstitch DELIMITER FORMAT")
  os.Exit(1)
}

func main() {
  flag.Parse()
  if flag.NArg() != 2 {
    usage()
  }
  
  delimiter := flag.Arg(0)
  format := must2(template.New("").Parse(flag.Arg(1)))
  input := strings.Split(string(must2(io.ReadAll(os.Stdin))), delimiter)
  must(format.Execute(os.Stdout, input))
}
