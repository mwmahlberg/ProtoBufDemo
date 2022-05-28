[![Stability: Unsupported](https://masterminds.github.io/stability/unsupported.svg)](https://masterminds.github.io/stability/unsupported.html)

## Protocol Buffer Demo

This demo was written as part of an answer to the question [How to read packed binary data in Go?][so:q] on Stackoverflow. 

## Requirements

* [The protocoll buffer C++ Distribution][gh:google-protobuf]
* [Protocol Buffers for Python][gh:py-protobuf]
* [Protocol Buffers for Go][gh:go-protobuf]
* A [make][wp:make] implementation (optional)


## How to run

After properly installing the requirements, you can run the project by issuing

    $ make run
    
on the commandline in the project directory. Alternatively, you can run

    go generate
    go build
    python writer.py
    ./ProtoBufDemo



[so:q]: http://stackoverflow.com/questions/34078427/how-to-read-packed-binary-data-in-go
[gh:google-protobuf]: https://github.com/google/protobuf
[gh:py-protobuf]: https://github.com/google/protobuf/tree/master/python
[gh:go-protobuf]: https://github.com/golang/protobuf
[wp:make]: https://en.wikipedia.org/wiki/Make_(software)
