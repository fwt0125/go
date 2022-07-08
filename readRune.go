package main
import (
	"bufio"
	"bytes"
	"fmt"
)
func main() {
	data := []byte("语言中文网")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	ch, size, err := r.ReadRune()
	fmt.Println(string(ch), size, err)

	slice()
	sliceDelim()
	buff()
	peek()

}

func slice(){
	data := []byte("C语言中文网, slice End语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadSlice(delim)
	fmt.Println(string(line), err)
	line, err = r.ReadSlice(delim)
	fmt.Println(string(line), err)
	line, err = r.ReadSlice(delim)
	fmt.Println(string(line), err)
}

func sliceDelim(){
	data := []byte("sliceDelim 语言中文网, sliceDelim End语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var delim byte = ','
	line, err := r.ReadString(delim)
	fmt.Println(line, err)
}

func buff()  {
	data := []byte("Buff 语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	var buf [14]byte
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err)
	rn := r.Buffered()
	fmt.Println("rb: ",rn)
}

func peek()  {
	data := []byte("Go语言入门教程")
	rd := bytes.NewReader(data)
	r := bufio.NewReader(rd)
	bl, err := r.Peek(5)
	fmt.Println(string(bl), err)
}