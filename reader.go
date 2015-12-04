package main

//go:generate protoc --python_out=. data.proto
//go:generate protoc --go_out=. data.proto
import (
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"
)

func main() {
	
	// Note that we do not handle any errors for the sake of brevity
	d := Demo{}
	f, _ := os.Open("tst.bin")
	fi, _ := f.Stat()
	
	// We create a buffer which is big enough to hold the entire message
	b := make([]byte,fi.Size())

	f.Read(b)
	
	proto.Unmarshal(b, &d)
	fmt.Println("* Go reading from file")
	
	// Note the explicit pointer dereference, as the fields are pointers to a pointers
	fmt.Printf("\tDemo.A: %d, Demo.B: %d, Demo.C: %d\n",*d.A,*d.B,*d.C)
}
