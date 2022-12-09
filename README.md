Selector
========

Simple terminal application which lets the user select
between its input arguments and prints out the selected
choice.

The usage is simple and intuitive: use the arrow keys to
move between values; press enter to select a value;
type `/` to start filtering.

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

Startup filtering
-----------------

Parameter `-filter` applies a filter to the list of choices
at startup.  
This is useful when the list of choices is long and the user
wants to filter it down to a few choices before selecting
one.

If the filter matches exactly one choice, it is
automatically selected and printed.  
This allows implementing a `goto` shell function which will
automatically switch to a directory if the sufficient filter
is provided by the user. With the function below,
calling `goto foo` will automatically switch to
the `$HOME/Documents/foodir` directory.

```Zsh
#!/bin/zsh

function goto {
	mypaths=(
		$HOME/{Documents,Downloads}
		$HOME/Documents/foodir
	)
	local DIR=$(selector -filter "$*" ${=mypaths})
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
go install github.com/mhmorgan/selector
```
