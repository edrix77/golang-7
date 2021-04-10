package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	type names struct {
		fname string
		lname string
	}
	var text names
	var a string

	fmt.Println("My file is (test.txt)")
	fmt.Print("Enter filename = ")
	fmt.Scan(&a)

	data, _ := ioutil.ReadFile(a)

	var datax string = fmt.Sprintf("%s", data)
	// x := make([]string, 0)
	x := strings.Split(datax, " ")
	text.fname = x[0]
	text.lname = x[1]

	println(text.fname, text.lname)

}

// inputbuffer := bufio.NewScanner(os.Stdin)

// for true {
// 	fmt.Print("First Name : ")
// 	inputbuffer.Scan()
// 	text.fname = inputbuffer.Text()
// 	if len(text.fname) <= 20 {
// 		break
// 	} else {
// 		println("The string size is exceeded, max 20 letters")
// 	}
// }

// for true {
// 	fmt.Print("Last Name : ")
// 	inputbuffer.Scan()
// 	text.lname = inputbuffer.Text()
// 	if len(text.lname) <= 20 {
// 		break
// 	} else {
// 		println("The string size is exceeded, max 20 letters")
// 	}
// }

// println(text.fname, " + ", text.lname)
