Selector
========

Simple terminal application which lets the user select between its input arguments and prints out the the selected choice.  

The usage is simple and intuitive: use the arrow keys to move between values; press enter to select a value;
type anything to filter the values; press backspace to remove the last filter character, or delete to clear the entire filter text.
Whitespaces in the filter text are ignored and treated as separators between substrings, all of which must match a value for it
to pass the filter.

Intended to be used in shell functions and aliases, such as:

```Zsh
alias goto='cd $(selector $HOME/*)'

# or

function goto {
	local DIR=$(selector $HOME/*)/$1
	cd $DIR
	ls
}
```

Automatic selection with startup filter
---------------------------------------

Parameter `-f`/`--filter` sets the startup value of the selector filter.  
Parameter `-a`/`--auto` enable automatic selection at startup: if there's only one value to choose from after the startup filters have been applied, choose this value automatically and return.

This allows implementing a `goto` shell function which will automatically switch to a directory if the sufficient filter is provided by the user.
With the function below, calling `goto foo` will automatically switch to the `$HOME/Documents/foodir` directory.

```Zsh
#!/bin/zsh

function goto {
	mypaths=(
		$HOME/{Documents,Downloads}
		$HOME/Documents/foodir
	)
	local DIR=$(selector ${=mypaths} -af "$*")
	# Don't try to change directory if selector didn't return a value
	[[ -n "$DIR" ]] || return
	echo $DIR
	cd $DIR
	ls
}
```

Installation
------------

```
cargo install selector
```
