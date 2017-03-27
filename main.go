package main

import (
	"log"
	"os"

	"fmt"

	"github.com/crissilvaeng/pcv-brute-force/arguments"
	"github.com/crissilvaeng/pcv-brute-force/graph"
)

func main() {
	args, err := arguments.New(os.Args)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	raw, err := args.Input()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	route, err := graph.New(raw)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	paths, err := route.Paths()
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	for index, element := range paths {
		fmt.Printf("%d -> %v -> %d\n", index, element, route.Cost(element))
	}
}
