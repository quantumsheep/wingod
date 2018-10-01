package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

var programs = map[string]string{
	"Background changer":                                "brain bc [url|path]",
	"Activate boring things (Narator, Contrast, etc..)": "brain abt [url|path]",
}

func main() {
	numargs := len(os.Args)

	if numargs == 1 {
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.TabIndent|tabwriter.Debug)

		for k, v := range programs {
			fmt.Fprintln(w, k+"\t  "+v)
		}

		w.Flush()
	}
}
