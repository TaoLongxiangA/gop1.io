package main

import (
	"fmt"
	"sort"
)

func main() {
	//s := "hello"
	////s1 := make([]byte, 5, 10)
	//slice := []byte{'h', 'e', 'l', 'l', 'o'}
	//fmt.Println(s[:3], reflect.TypeOf(s[:3]))
	//fmt.Println(string(slice[:3]), reflect.TypeOf(slice[:3]))
	//
	//for i := range s {
	//	fmt.Println(i)
	//}

	//var w io.Write
	//fmt.Println(reflect.TypeOf(w))
	//w = os.Stdout
	//fmt.Println(reflect.TypeOf(w))
	//w = new(bytes.Buffer)
	//fmt.Println(reflect.TypeOf(w))
	//var rwc io.ReadWriteCloser

	//var w io.Write = os.Stdout
	//f, ok := w.(*os.File)
	//fmt.Println("f:", reflect.TypeOf(f), "ok:", ok)

	//s := make(map[int]int)
	//fmt.Println(s)
	//s[0]++
	//s[0]++
	//s[1]++
	//fmt.Println(s)
	//bytes := []byte{'b', 'b', 'q'}
	//fmt.Printf("%s",bytes)
	//fmt.Printf("%4d",123)

	////fmt.Println(runtime.NumCPU())
	//var a io.Writer
	//a = os.Stdout
	//_, _ = a.Write([]byte("hello"))
	//var s interface{} = time.Now()
	//var s interface{} = []byte{'b','t','t'}
	//fmt.Println(s==s)
	strings := []string{"dog", "cat", "tiger"}
	sort.Strings(strings)
	fmt.Println(strings)
}
