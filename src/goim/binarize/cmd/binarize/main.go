package main

import (
	"flag"
	"goim/binarize"
	"log"
)

var sourcePath = flag.String("source", "", "the path to the source file")
var destinationPath = flag.String("destination", "", "the path to th destination file")

var binarizeImport = flag.String("import_binarize", "", "import path in go to binarize package")
var serializeFuncName = flag.String("serialize_fname", "", "the name of serialization function")
var deserializeFuncName = flag.String("deserialize_fname", "", "the name of deserialization function")

func main() {
	defer func() {
		// Throughout this package panic used to raise errors because performance impact of it's usage in this
		// specific case is acceptable
		r := recover()
		if r != nil {
			log.Fatalf("%v", r)
		}
	}()

	flag.Parse()
	binarize.Generate(*sourcePath, *destinationPath, *binarizeImport, *serializeFuncName, *deserializeFuncName)
}
