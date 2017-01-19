package main

import (
	"strings"
	"bufio"
	"fmt"
	"reflect"
	//"io"
)

func main()  {

	//io.Reader()

	s := strings.NewReader("ABCDEFG")

	fmt.Printf(reflect.TypeOf(s).Kind().String())

	br := bufio.NewReader(s)

	fmt.Printf(reflect.TypeOf(br).Kind().String())

	b, _ := br.Peek(5)

	fmt.Printf("%s\n", b)
}
