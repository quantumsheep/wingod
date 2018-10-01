package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/reujab/wallpaper"
)

func loading() {
	fmt.Print("Loading")

	for true {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(".")
	}
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func main() {
	path := ""

	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		path = getCurrentDirectory() + "/wallpaper.jpg"
	}

	if strings.HasPrefix(path, "http") {
		go loading()

		wallpaper.SetFromURL(path)
	} else {
		wallpaper.SetFromFile(path)
	}

	fmt.Println("\nDone!")

	os.Exit(0)
}
