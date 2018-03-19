package main

import (
	m "fmt"
	. "math"

	. "./packone"
	. "./packtwo"

	"./packthree"
	"./fmt"

	uuid "github.com/satori/go.uuid"
)

var x packthree.Internal
var y fmt.Apple

func main() {
	m.Println(Exp2(6))      // 64
	Add()                   //SomeThings
	Reduce()                //SomeThingsElse
	packthree.Exponential() //SomeThingsOnSomeThings

	fmt.HelloWorld()

	u1 := uuid.Must(uuid.NewV4())
	m.Printf("UUIDv4: %s\n", u1) //UUIDv4: 1b214112-d698-4443-a35d-ee3dcd8355e9
}
