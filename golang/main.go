package main

import(
		"fmt"
		"unsafe"
      )

const(
		s = "hello world"
		m = len(s)
		n = unsafe.Sizeof(m)
     )

func main() {
	fmt.Println(s,m,n)
}
