package selection

import (
	"github.com/charmbracelet/bubbles/list"
	"os/user"
	"path"
	"strings"
)

var Filter = list.DefaultFilter

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
