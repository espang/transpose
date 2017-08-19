package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	var (
		from  = flag.String("source", "", "Source filename")
		to    = flag.String("target", "", "Target filename")
		force = flag.Bool("force", false, "Overwrite target")
	)

	flag.Parse()

	log.Printf("transpose file '%s' into '%s'", *from, *to)

	if !fileExists(*from) {
		log.Fatalf("source file '%s' not found", *from)
	}

	input, err := os.Open(*from)
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	if fileExists(*to) {
		if !*force {
			log.Fatalf("target file '%s' exists. Use -force to overwrite existing file", *to)
		}
	}

	output, err := os.Create(*to)
	if err != nil {
		log.Fatal(err)
	}

	err = transposeCsv(input, output)
	if err != nil {
		_ = output.Close()
		log.Fatal(err)
	}

	if err = output.Close(); err != nil {
		log.Fatal("closing output file: ", err)
	}

	log.Print("transpose finished")
}
