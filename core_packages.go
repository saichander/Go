package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"hash/crc32"
)

func main() {
	//strings package
	fmt.Println(
		strings.Contains("netsetter", "set"),
		strings.Count("levied", "e"),
		strings.HasPrefix("test", "te"), //true
		strings.HasSuffix("sdfd", "fd"), //true
		strings.ToUpper("upper"),
		strings.ToLower("lower"),
		strings.Join([]string{"a", "b"}, "-"), //"a-b"
	)

	//files and folders
	file, err := os.Open("test.txt")
	if err != nil {
		return
	}
	defer file.Close()

	stat, err := file.Stat() //gets file size
	if err != nil {
		return
	}
	fmt.Println("Size of file", stat.Size())

	//read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)

	str := string(bs)
	fmt.Println(str, err)

	//create a file
	file1, err := os.Create("test1.txt")
	if err != nil {
		return
	}
	defer file1.Close()
	file1.WriteString("test")
	bs1, err := ioutil.ReadFile("test1.txt")
	fmt.Println("str1", string(bs1))

	//errors
	fmt.Println(errors.New("test eror message"))

	//cryptography
	h := crc32.NewIEEE()
	h.Write([]byte("test"))
	fmt.Println(h)
	v := h.Sum32()
	fmt.Println(v)
}
