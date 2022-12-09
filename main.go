package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

var filter = flag.String("filter", "", "Initial selection filter.")

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 1 {
		bail("No paths given.")
	}

	var items []list.Item
	for _, arg := range flag.Args() {
		items = append(items, item(arg))
	}
	items = filterItems(items, *filter)

	if len(items) == 0 {
		bail("No matches found.")
	} else if len(items) == 1 {
		fmt.Print(items[0])
		os.Exit(0)
	}

	p := tea.NewProgram(newModel(items), tea.WithAltScreen())

	var m model
	if tmp, err := p.Run(); err != nil {
		bail(err)
	} else {
		m = tmp.(model)
	}

	if val := m.choice; val != "" {
		fmt.Print(m.choice)
	} else {
		os.Exit(1)
	}
}

func bail(msg any) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: selector [options] [paths ...]")
	fmt.Fprintln(os.Stderr, "Options:")
	flag.PrintDefaults()
	os.Exit(2)
}

func filterItems(items []list.Item, filter string) []list.Item {
	if filter == "" {
		return items
	}

	var filterValues []string
	for _, item := range items {
		filterValues = append(filterValues, item.FilterValue())
	}

	var filtered []list.Item
	for _, rank := range list.DefaultFilter(filter, filterValues) {
		filtered = append(filtered, items[rank.Index])
	}
	return filtered
}

// -----------------------------------------------------------------------------

type item string

func (i item) FilterValue() string {
	return i.Title()
}

func (i item) Title() string {
	return path.Base(string(i))
}

func (i item) Description() string {
	desc := path.Dir(string(i))

	if u, err := user.Current(); err == nil {
		desc = strings.Replace(desc, u.HomeDir, "~", 1)
	}

	return desc
}

func (i item) String() string {
	return string(i)
}

// -----------------------------------------------------------------------------

// model is the main model for the application user interface,
// containing the program state.
type model struct {
	list     list.Model
	choice   string
	quitting bool
}

func newModel(items []list.Item) model {
	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Select directory"
	l.SetStatusBarItemName("Directory", "Directories")
	return model{list: l}
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m model) Init() tea.Cmd {
	return nil
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		m.list.SetHeight(msg.Height)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit

		default:
			if n, err := strconv.Atoi(msg.String()); err == nil {
				n = (n + 9) % 10
				if n < len(m.list.Items()) {
					m.list.Select(n)
				}
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m model) View() string {
	if m.quitting {
		return ""
	}
	return m.list.View()
}
