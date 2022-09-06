//Ini Executable = package yg langsung dieksekusi.
package main

import (
	"fmt"
	"testGo1/go2"
)

func main() {
	hasil := go2.Tambah()
	kurang := go2.Kurang(5, 2)
	fmt.Println("hello world", hasil, "Pengurangan = ", kurang)
}
