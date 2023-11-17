package main

import (
	"bufio"
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
  reader := bufio.NewReader(os.Stdin)

  for {
    line, err := reader.ReadString('\n')
    if err != nil && err != io.EOF {
        panic(err)
    }
    cut := strings.Split(line, delimiter)
    must(format.Execute(os.Stdout, cut))
    if err == io.EOF {
      break
    }
  }
}
