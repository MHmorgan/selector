package main

import (
	"flag"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/mhmorgan/selector/selection"
	"log"
	"os"
)

var filter = flag.String("filter", "", "Initial selection filter.")

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 1 {
		os.Exit(1)
	}

	model := selection.New(flag.Args(), *filter)

	if ranks := selection.Filter(*filter, model.FilterValues()); len(ranks) == 1 {
		fmt.Print(model.GetValue(0))
		os.Exit(0)
	}

	p := tea.NewProgram(model, tea.WithAltScreen())

	var sl selection.Model
	if tmp, err := p.Run(); err != nil {
		log.Fatalln(err)
	} else {
		sl = tmp.(selection.Model)
	}

	if val := sl.Choice(); val != "" {
		fmt.Print(sl.Choice())
	} else {
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: selector [options] [paths ...]")
	fmt.Fprintln(os.Stderr, "Options:")
	flag.PrintDefaults()
	os.Exit(2)
}
