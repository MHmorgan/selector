package selection

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"strconv"
)

type newItems []list.Item

// Model is the main model for the application user interface,
// containing the program state.
type Model struct {
	list     list.Model
	choice   string
	quitting bool
}

func New(choices []string, filter string) Model {
	var items []list.Item
	for _, s := range choices {
		items = append(items, item(s))
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Select directory"
	l.Filter = Filter
	l.SetStatusBarItemName("Directory", "Directories")
	l.FilterInput.SetValue(filter)

	return Model{list: l}
}

// Choice returns the string of the currently selected item.
func (m Model) Choice() string {
	return m.choice
}

// FilterValues returns a list of strings with all the
// filter values of all items in this model.
func (m Model) FilterValues() []string {
	values := make([]string, len(m.list.Items()))
	for i, itm := range m.list.Items() {
		values[i] = itm.FilterValue()
	}
	return values
}

// GetValue returns the string value of the item with
// index i.
func (m Model) GetValue(i int) string {
	itm, ok := m.list.Items()[i].(item)
	if !ok {
		return ""
	}
	return string(itm)
}

// Init is the first function that will be called. It returns an optional
// initial command. To not perform an initial command return nil.
func (m Model) Init() tea.Cmd {
	filter := m.list.FilterValue()
	if filter == "" {
		return nil
	}

	var (
		items  []list.Item
		values = m.FilterValues()
		old    = m.list.Items()
	)

	for _, rank := range Filter(filter, values) {
		items = append(items, old[rank.Index])
	}

	return func() tea.Msg {
		return newItems(items)
	}
}

// Update is called when a message is received. Use it to inspect messages
// and, in response, update the model and/or send a command.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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

	case newItems:
		m.list.SetItems(msg)
		m.list.FilterInput.SetValue("")
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View renders the program's UI, which is just a string. The view is
// rendered after every Update.
func (m Model) View() string {
	if m.quitting {
		return ""
	}
	return m.list.View()
}
