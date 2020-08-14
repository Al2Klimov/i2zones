package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"i2zones/lib/base"
	"io"
	"os"
)

type object struct {
	Name       string `json:"name"`
	Properties struct {
		Parent string `json:"parent"`
	} `json:"properties"`
	Type string `json:"type"`
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)

	fmt.Fprintln(out, "digraph {")
	fmt.Fprintln(out, "  rankdir = LR;")

	for {
		ns, errRN := base.ReadNetStringFromStream(in, -1)
		if errRN == io.EOF {
			break
		}

		var obj object
		if errUJ := json.Unmarshal(ns, &obj); errUJ != nil {
			fmt.Fprintln(os.Stderr, errUJ.Error())
			os.Exit(1)
		}

		if obj.Type == "Zone" && obj.Properties.Parent != "" {
			fmt.Fprintf(out, "  \"%s\" -> \"%s\";\n", obj.Properties.Parent, obj.Name)
		}
	}

	fmt.Fprintln(out, "}")
	out.Flush()
}
